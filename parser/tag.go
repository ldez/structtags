// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"strconv"
)

// Filler is the interface implemented by types that can fill elements from a struct tag.
type Filler[T any] interface {
	Data() T
	Fill(key, value string) error
}

// Tag parses a struct tag.
//
// Based on https://github.com/golang/go/blob/411c250d64304033181c46413a6e9381e8fe9b82/src/reflect/type.go#L1030-L1108
func Tag[T any](tag string, filler Filler[T]) (T, error) {
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}

		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}

		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			var zero T

			return zero, fmt.Errorf("syntax error in tag %q", tag)
		}

		name := tag[:i]
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}

			i++
		}

		if i >= len(tag) {
			var zero T

			return zero, fmt.Errorf("syntax error in tag %q", tag)
		}

		qvalue := tag[:i+1]
		tag = tag[i+1:]

		value, err := strconv.Unquote(qvalue)
		if err != nil {
			var zero T

			return zero, fmt.Errorf("syntax error in tag %q", tag)
		}

		err = filler.Fill(name, value)
		if err != nil {
			var zero T

			return zero, err
		}
	}

	return filler.Data(), nil
}
