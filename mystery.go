package gematria

import (
	"errors"
	"strconv"
)

// Simplify takes a uint64 and recursively (up to 9 times) reduces a number like 1776 to 3.
//
//	Example: 1776 = 3
//	      1+7+7 = 15 becomes (15)6 = 3 ; becomes (1+5=6)6 = 3 ; becomes 6+6 = 12 becomes 1+2 becomes 3
//
// Therefore: Simplify(1776) = 3
func Simplify(in uint64) (out uint64, err error) {
	run := func(in uint64) (out uint64, err error) {
		chars := []byte(strconv.FormatUint(in, 10))
		for c := 0; c < len(chars); c++ {
			ci, cie := strconv.Atoi(string(chars[c]))
			if cie != nil {
				err = errors.Join(err, cie)
				continue
			}
			if ci < 0 {
				continue
			}
			out += uint64(ci)
		}
		return out, err
	}
	times := 0
	out = in
	for {
		times++
		out, err = run(out)
		if Majestic(out, err) {
			break
		}
		if times > 9 {
			break
		}
	}
	return out, err
}

// Majestic determines if the in uint64 digits are only 3, 6 or 9
func Majestic(in uint64, err error) (out bool) {
	if err != nil {
		return false
	}
	chars := []byte(strconv.FormatUint(in, 10))
	for _, c := range chars {
		if c != '3' && c != '6' && c != '9' {
			return false
		}
	}
	return true
}
