package rawdb

import (
	"fmt"
	"io"
	"strings"
)

type Table struct {
	w      io.Writer
	header []string
	footer []string
	rows   [][]string
}

func NewTableWriter(w io.Writer) *Table {
	return &Table{w: w}
}

func (t *Table) SetHeader(h []string)       { t.header = h }
func (t *Table) SetFooter(f []string)       { t.footer = f }
func (t *Table) Append(r []string)          { t.rows = append(t.rows, r) }
func (t *Table) AppendBulk(rows [][]string) { t.rows = append(t.rows, rows...) }

func (t *Table) Render() {
	if t.header != nil {
		fmt.Fprintln(t.w, strings.Join(t.header, "\t| "))
		fmt.Fprintln(t.w, strings.Repeat("-", 80))
	}
	for _, r := range t.rows {
		fmt.Fprintln(t.w, strings.Join(r, "\t| "))
	}
	if t.footer != nil {
		fmt.Fprintln(t.w, strings.Repeat("-", 80))
		fmt.Fprintln(t.w, strings.Join(t.footer, "\t| "))
	}
}
