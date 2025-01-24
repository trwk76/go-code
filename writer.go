package code

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// WriteFile is a short hand for:
// - ensuring that the given path directory exists.
// - creating an empty file at the given path.
// - writing the content of the file using the given WriteFunc.
func WriteFile(path string, tabString string, f WriteFunc) error {
	os.MkdirAll(filepath.Dir(path), 0777)

	fil, err := os.Create(path)
	if err != nil {
		return err
	}

	w := NewWriter(fil, tabString)

	defer func() {
		w.Flush()
		fil.Close()
	}()

	f(&w)
	return nil
}

// WriteString is a short hand for rendering the content of a WriteFunc as a string.
func WriteString(tabString string, f WriteFunc) string {
	buf := strings.Builder{}
	w := NewWriter(&buf, tabString)

	f(&w)

	w.Flush()
	return buf.String()
}

// NewWriter creates a new Writer that writes to the given io.Writer using the specified tabString for indentation.
func NewWriter(w io.Writer, tabString string) Writer {
	if tabString == "" {
		tabString = "\t"
	}

	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}

	return Writer{
		w:   bw,
		ts:  tabString,
		nl:  true,
		ind: 0,
	}
}

// Flush will flush the underlying bufio.Writer.
// Must be called before closing the underlying io.Writer if you want output to be complete.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) Write(p []byte) (int, error) {
	res := 0

	for idx, line := range bytes.Split(p, []byte{'\n'}) {
		if idx > 0 {
			if err := w.Newline(); err != nil {
				return res, err
			}
		}

		if len(line) > 0 {
			w.ensureIndented()
		}

		if done, err := w.w.Write(line); err != nil {
			return res, err
		} else {
			res += done
		}
	}

	return res, nil
}

func (w *Writer) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}

func (w *Writer) WriteByte(b byte) error {
	if b == '\n' {
		w.nl = true
	} else {
		w.ensureIndented()
	}

	return w.w.WriteByte(b)
}

func (w *Writer) Newline() error {
	return w.WriteByte('\n')
}

func (w *Writer) Space() error {
	return w.WriteByte(' ')
}

func (w *Writer) Indent(f WriteFunc) {
	w.ind++
	f(w)
	w.ind--
}

func (w *Writer) Table(rows ...TableRow) {
	colw := make([]int, 0)
	maxw := 0

	for _, row := range rows {
		for idx, col := range row.Columns {
			l := len(col)

			if idx < len(colw) {
				if l > colw[idx] {
					colw[idx] = l
				}
			} else {
				colw = append(colw, l)
			}

			if l > maxw {
				maxw = l
			}
		}
	}

	pad := strings.Repeat(" ", maxw)

	for _, row := range rows {
		if row.Prefix != "" {
			if _, err := w.WriteString(row.Prefix); err != nil {
				panic(err)
			}

			if !w.nl {
				// Make sure we write the row bulk on a newline
				if err := w.Newline(); err != nil {
					panic(err)
				}
			}
		}

		for idx, col := range row.Columns {
			if _, err := w.WriteString(col); err != nil {
				panic(err)
			}

			if idx < len(row.Columns)-1 {
				if _, err := w.WriteString(pad[:colw[idx]-len(col)]); err != nil {
					panic(err)
				}

				if err := w.WriteByte(' '); err != nil {
					panic(err)
				}
			}
		}

		if err := w.Newline(); err != nil {
			panic(err)
		}
	}
}

type (
	Writer struct {
		w   *bufio.Writer
		ts  string
		nl  bool
		ind uint16
	}

	TableRow struct {
		Prefix  string
		Columns []string
	}

	WriteFunc func(w *Writer)
)

func (w *Writer) ensureIndented() {
	if !w.nl {
		return
	}

	for range w.ind {
		if _, err := w.w.WriteString(w.ts); err != nil {
			panic(err)
		}
	}

	w.nl = false
}
