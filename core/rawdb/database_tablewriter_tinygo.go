package rawdb

import "io"

type Table struct{}

func NewTableWriter(w io.Writer) *Table     { return &Table{} }
func (t *Table) SetHeader(h []string)       {}
func (t *Table) SetFooter(f []string)       {}
func (t *Table) Append(r []string)          {}
func (t *Table) AppendBulk(rows [][]string) {}
func (t *Table) Render()                    {}
