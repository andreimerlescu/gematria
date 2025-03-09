package gematria

import (
	"errors"
	"strconv"
)

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

func Majestic(in uint64, err error) (out bool) {
	if err != nil {
		return false
	}
	chars := []byte(strconv.FormatUint(uint64(in), 10))
	for _, c := range chars {
		if c != '3' && c != '6' && c != '9' {
			return false
		}
	}
	return true
}
