package util

import "fmt"

type MhallError struct {
	Where		string
	line		int
	ErrorMsg	string
}

func (m *MhallError) Error() string {
	return fmt.Sprintf("@ %s %d --> %s", m.Where, m.line, m.ErrorMsg)
}

func New(where string, line int,  errormsg string) *MhallError {
	return &MhallError{where, line, errormsg}
}