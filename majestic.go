package gematria

import (
	"log"
	"strconv"
	"strings"
)

// ComboType is a string for the name of the flavor of Gematria you're displaying as a uint64
type ComboType string

// define the combination types for Gematria
const (
	CT_S      ComboType = "simple"
	CT_J      ComboType = "jewish"
	CT_E      ComboType = "english"
	CT_MY     ComboType = "mystery"
	CT_MJ     ComboType = "majestic"
	CT_EI     ComboType = "eights"
	CT_SJ     ComboType = "simple jewish"
	CT_SE     ComboType = "simple english"
	CT_SMY    ComboType = "simple mystery"
	CT_SMJ    ComboType = "simple majestic"
	CT_SEI    ComboType = "simple eights"
	CT_JE     ComboType = "jewish english"
	CT_JMY    ComboType = "jewish mystery"
	CT_JMJ    ComboType = "jewish majestic"
	CT_JEI    ComboType = "jewish eights"
	CT_MYMJ   ComboType = "mystery majestic"
	CT_MYEI   ComboType = "mystery eights"
	CT_EIMJ   ComboType = "eights majestic"
	CT_SJE    ComboType = "simple jewish english"
	CT_SJMY   ComboType = "simple jewish mystery"
	CT_SJMJ   ComboType = "simple jewish majestic"
	CT_SJEI   ComboType = "simple jewish majestic"
	CT_SEMY   ComboType = "simple english mystery"
	CT_SEMJ   ComboType = "simple english majestic"
	CT_SEEI   ComboType = "simple english eights"
	CT_JEMY   ComboType = "jewish english mystery"
	CT_JEMJ   ComboType = "jewish english majestic"
	CT_JEEI   ComboType = "jewish english eights"
	CT_EMYMJ  ComboType = "english mystery majestic"
	CT_EMYEI  ComboType = "english mystery eights"
	CT_EMJEI  ComboType = "english majestic eights"
	CT_MJMYEI ComboType = "majestic mystery eights"
)

// ComboTypes returns a slice of ComboType that has all supported combinations by the package
func ComboTypes() []ComboType {
	return []ComboType{
		CT_S, CT_E, CT_J,
		CT_MY, CT_MJ, CT_EI, CT_SJ,
		CT_SE, CT_SMY, CT_SMJ, CT_SEI, CT_JE,
		CT_JMY, CT_JMJ, CT_JEI, CT_MYMJ, CT_MYEI, CT_EIMJ,
		CT_SJE, CT_SJMY, CT_SJMJ, CT_SJEI, CT_SEMY, CT_SEMJ,
		CT_SEEI, CT_JEMY, CT_JEMJ, CT_JEEI, CT_EMYMJ, CT_EMYEI,
		CT_EMJEI, CT_MJMYEI,
	}
}

// Types returns a slice of the Gematria keys as strings
func Types() []string {
	return []string{
		"simple", "jewish", "english",
		"majestic", "mystery", "eights",
	}
}

// Combine accepts multiple uint64 values and merges them like the Gematria.Combine method does
func Combine(in ...uint64) (uint64, error) {
	sb := strings.Builder{}
	for _, i := range in {
		sb.WriteString(strconv.FormatUint(i, 10))
	}
	return strconv.ParseUint(sb.String(), 10, 64)
}

// Combine reads multiple ComboType fields and concatenates the digits then converts to uint64
// otherwise 0 is returned
func (s Gematria) Combine(fields ...ComboType) (uint64, error) {
	sb := strings.Builder{}
	for _, field := range fields {
		switch field {
		case CT_S:
			sb.WriteString(strconv.FormatUint(s.Simple, 10))
		case CT_E:
			sb.WriteString(strconv.FormatUint(s.English, 10))
		case CT_J:
			sb.WriteString(strconv.FormatUint(s.Jewish, 10))
		case CT_MY:
			sb.WriteString(strconv.FormatUint(s.Mystery, 10))
		case CT_MJ:
			sb.WriteString(strconv.FormatUint(s.Majestic, 10))
		case CT_EI:
			sb.WriteString(strconv.FormatUint(s.Eights, 10))
		}
	}
	return strconv.ParseUint(sb.String(), 10, 64)
}

// SimplifyWithError accepts a uint64, error combo and performs Simplify on it or returns 0 for an error
func SimplifyWithError(in uint64, err error) (out uint64, e error) {
	if err != nil {
		return 0, err
	}
	out, e = Simplify(in)
	if e != nil {
		return 0, e
	}
	return out, nil
}

// MajesticGematria - TODO - Update this function later to be more majestic than just an (out uint64)
func MajesticGematria(in Gematria, choices ...ComboType) (out uint64) {
	for _, choice := range choices {
		c := strings.Clone(string(choice))
		if en, err := Simplify(in.English); err == nil {
			c = strings.ReplaceAll(c, "english", strconv.FormatUint(en, 10))
		}
		if je, err := Simplify(in.Jewish); err == nil {
			c = strings.ReplaceAll(c, "jewish", strconv.FormatUint(je, 10))
		}
		if si, err := Simplify(in.Simple); err == nil {
			c = strings.ReplaceAll(c, "simple", strconv.FormatUint(si, 10))
		}
		if my, err := Simplify(in.Mystery); err == nil {
			c = strings.ReplaceAll(c, "mystery", strconv.FormatUint(my, 10))
		}
		if ma, err := Simplify(in.Majestic); err == nil {
			c = strings.ReplaceAll(c, "majestic", strconv.FormatUint(ma, 10))
		}
		if ei, err := Simplify(in.Eights); err == nil {
			c = strings.ReplaceAll(c, "eights", strconv.FormatUint(ei, 10))
		}
		nums := strings.Split(c, " ")
		var sum uint64
		var err error
		for _, num := range nums {
			pre := sum
			sum, err = strconv.ParseUint(num, 10, 64)
			if err != nil {
				log.Println(err)
				break
			}
			sum = sum + pre
		}
		c = strings.ReplaceAll(c, " ", "")
		ci, cie := strconv.ParseUint(c, 10, 64)
		if cie != nil {
			log.Println(cie)
			continue
		}
		out, err = Simplify(ci)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
