// Package domain check the domain is valid.
package domain

import "strings"

// IsValid checks the domain whether is legal by RFC 1035, RFC 3696
func IsValid(s string) bool {
	l := len(s)
	if l == 0 || l > 254 || l == 254 && s[l-1] != '.' {
		return false
	}

	n := 0
	notNumber := false
	lastByte := byte('.')

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			notNumber = true
			n++
		case '0' <= c && c <= '9':
			n++
		case c == '-':
			if lastByte == '.' {
				return false
			}
			n++
			notNumber = true
		case c == '.':
			if lastByte == '.' || lastByte == '-' {
				return false
			}
			if n > 63 || n == 0 {
				return false
			}
			n = 0
		default:
			return false
		}

		lastByte = c
	}

	if lastByte == '-' || n > 63 {
		return false
	}

	return notNumber
}

// LastIndexDot returns the subdomain of the last instance of '.' in domain
func LastIndexDot(domain string, ndot int) (string, bool) {
	pos := len(domain)
	for i := 0; i < ndot; i++ {
		pos = strings.LastIndexByte(domain[:pos], '.')
		if pos == -1 {
			return "", false
		}
	}

	n := strings.LastIndexByte(domain[:pos], '.')
	if n > -1 {
		return domain[n+1 : pos], true
	}

	return domain[:pos], true
}
