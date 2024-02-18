package jooble

import "fmt"

var (
	Err4xx = fmt.Errorf("client-side error; did we sent something wrong? check the audit response")
	Err5xx = fmt.Errorf("server-side error; try again later? check the audit response")
)
