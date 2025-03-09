package gematria

import (
	"log"
	"strconv"
	"strings"
)

type ComboType string

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
