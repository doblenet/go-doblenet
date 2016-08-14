package cliapp

import (
	"fmt"
)

// Someone who has contributed to a project.
type Author struct {
        Name  string
        Email string
}

// Implement the Stringer interface
func (a Author) String() string {
	mail := ""  
	if "" != a.Email {
		mail = fmt.Sprintf(" <%s>", a.Email)
	}
	return a.Name+mail
}
