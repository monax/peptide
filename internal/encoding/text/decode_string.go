// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"bytes"
	"fmt"
	"github.com/monax/peptide/internal/strs"
	"io"
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// Stripped down prototext decoder - we only need it for decoding default bytes value

// parseString parses a string value enclosed in " or '.
func UnmarshalString(str string) (string, error) {
	in := []byte(str)
	if len(in) == 0 {
		return "", io.ErrUnexpectedEOF
	}
	quote := in[0]
	in = in[1:]
	i := indexNeedEscapeInBytes(in)
	in, out := in[i:], in[:i:i] // set cap to prevent mutations
	for len(in) > 0 {
		switch r, n := utf8.DecodeRune(in); {
		case r == utf8.RuneError && n == 1:
			return "", fmt.Errorf("invalid UTF-8 detected")
		case r == 0 || r == '\n':
			return "", fmt.Errorf("invalid character %q in string", r)
		case r == rune(quote):
			in = consume(in, len(in)-len(in)+1)
			return string(out), nil
		case r == '\\':
			if len(in) < 2 {
				return "", io.ErrUnexpectedEOF
			}
			switch r := in[1]; r {
			case '"', '\'', '\\', '?':
				in, out = in[2:], append(out, r)
			case 'a':
				in, out = in[2:], append(out, '\a')
			case 'b':
				in, out = in[2:], append(out, '\b')
			case 'n':
				in, out = in[2:], append(out, '\n')
			case 'r':
				in, out = in[2:], append(out, '\r')
			case 't':
				in, out = in[2:], append(out, '\t')
			case 'v':
				in, out = in[2:], append(out, '\v')
			case 'f':
				in, out = in[2:], append(out, '\f')
			case '0', '1', '2', '3', '4', '5', '6', '7':
				// One, two, or three octal characters.
				n := len(in[1:]) - len(bytes.TrimLeft(in[1:], "01234567"))
				if n > 3 {
					n = 3
				}
				v, err := strconv.ParseUint(string(in[1:1+n]), 8, 8)
				if err != nil {
					return "", fmt.Errorf("invalid octal escape code %q in string", in[:1+n])
				}
				in, out = in[1+n:], append(out, byte(v))
			case 'x':
				// One or two hexadecimal characters.
				n := len(in[2:]) - len(bytes.TrimLeft(in[2:], "0123456789abcdefABCDEF"))
				if n > 2 {
					n = 2
				}
				v, err := strconv.ParseUint(string(in[2:2+n]), 16, 8)
				if err != nil {
					return "", fmt.Errorf("invalid hex escape code %q in string", in[:2+n])
				}
				in, out = in[2+n:], append(out, byte(v))
			case 'u', 'U':
				// Four or eight hexadecimal characters
				n := 6
				if r == 'U' {
					n = 10
				}
				if len(in) < n {
					return "", io.ErrUnexpectedEOF
				}
				v, err := strconv.ParseUint(string(in[2:n]), 16, 32)
				if utf8.MaxRune < v || err != nil {
					return "", fmt.Errorf("invalid Unicode escape code %q in string", in[:n])
				}
				in = in[n:]

				r := rune(v)
				if utf16.IsSurrogate(r) {
					if len(in) < 6 {
						return "", io.ErrUnexpectedEOF
					}
					v, err := strconv.ParseUint(string(in[2:6]), 16, 16)
					r = utf16.DecodeRune(r, rune(v))
					if in[0] != '\\' || in[1] != 'u' || r == unicode.ReplacementChar || err != nil {
						return "", fmt.Errorf("invalid Unicode escape code %q in string", in[:6])
					}
					in = in[6:]
				}
				out = append(out, string(r)...)
			default:
				return "", fmt.Errorf("invalid escape code %q in string", in[:2])
			}
		default:
			i := indexNeedEscapeInBytes(in[n:])
			in, out = in[n+i:], append(out, in[:n+i]...)
		}
	}
	return "", io.ErrUnexpectedEOF
}

// indexNeedEscapeInString returns the index of the character that needs
// escaping. If no characters need escaping, this returns the input length.
func indexNeedEscapeInBytes(b []byte) int { return indexNeedEscapeInString(strs.UnsafeString(b)) }

// indexNeedEscapeInString returns the index of the character that needs
// escaping. If no characters need escaping, this returns the input length.
func indexNeedEscapeInString(s string) int {
	for i := 0; i < len(s); i++ {
		if c := s[i]; c < ' ' || c == '"' || c == '\'' || c == '\\' || c >= 0x7f {
			return i
		}
	}
	return len(s)
}

// consume consumes n bytes of input and any subsequent whitespace or comments.
func consume(b []byte, n int) []byte {
	b = b[n:]
	for len(b) > 0 {
		switch b[0] {
		case ' ', '\n', '\r', '\t':
			b = b[1:]
		case '#':
			if i := bytes.IndexByte(b, '\n'); i >= 0 {
				b = b[i+len("\n"):]
			} else {
				b = nil
			}
		default:
			return b
		}
	}
	return b
}
