package ics

import "strings"

func trimField(field, s string) string {
	return strings.TrimRight(strings.TrimPrefix(s, field), "\r\n")
}

/** These codes are forked form "github.com/PuloV/ics-golang" package. (licensed under MIT) */
