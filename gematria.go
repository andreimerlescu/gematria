package gematria

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// allowedInput is used to strip characters from an input target that Gematria cannot calculate
var allowedInput = regexp.MustCompile(`[^a-zA-Z\d.\s]`)

// FromString is an alias to NewGematria that suppresses the error if one is present
func FromString(in string) (out Gematria) {
	out, _ = NewGematria(in)
	return
}

// NewGematria calculates the value of data and returns Gematria and an error
func NewGematria(data string) (Gematria, error) {
	data = strings.TrimSpace(data)
	data = allowedInput.ReplaceAllString(data, "")
	dataBytes := []byte(data)
	var letters []Gematria
	for i := 0; i < len(dataBytes); i++ {
		wat := strings.ToUpper(string(dataBytes[i]))
		if len(wat) == 1 && wat != "" && wat != " " {
			j, e, s, my, maj, eights, err := characterScores(wat)
			if err != nil {
				return Gematria{}, err
			}
			letters = append(letters, Gematria{
				Jewish:   j,
				English:  e,
				Simple:   s,
				Mystery:  my,
				Majestic: maj,
				Eights:   eights,
			})
		}
	}
	var jf, ef, sf, myf, majf, eigf uint64
	for _, gs := range letters {
		jf += gs.Jewish
		ef += gs.English
		sf += gs.Simple
		myf += gs.Mystery
		majf += gs.Majestic
		eigf += gs.Eights
	}
	return Gematria{original: data, Jewish: jf, English: ef, Simple: sf, Mystery: myf, Majestic: majf, Eights: eigf}, nil
}

// characterScores uses the data maps to reference the score of the incoming character
func characterScores(str string) (uint64, uint64, uint64, uint64, uint64, uint64, error) {
	if len(str) != 1 {
		return 0, 0, 0, 0, 0, 0, fmt.Errorf("string must only contain 1 character")
	}
	j, e, s, my, maj, eights := jewishCodes[str], englishCodes[str], simpleCodes[str], mysteryCodes[str], majesticCodes[str], eightsCodes[str]
	return j, e, s, my, maj, eights, nil
}

// String formats the output of the Gematria Struct
func (s Gematria) String() string {
	o := strings.Builder{}
	o.WriteString("Q: " + s.original + "\n")
	o.WriteString("\tJewish = " + strconv.FormatUint(s.Jewish, 10) + " \n")
	o.WriteString("\tEnglish = " + strconv.FormatUint(s.English, 10) + " \n")
	o.WriteString("\tSimple = " + strconv.FormatUint(s.Simple, 10) + " \n")
	o.WriteString("\tMajestic = " + strconv.FormatUint(s.Majestic, 10) + " \n")
	o.WriteString("\tMystery = " + strconv.FormatUint(s.Mystery, 10) + " \n")
	o.WriteString("\tEights  = " + strconv.FormatUint(s.Eights, 10) + " \n\n")
	return o.String()
}

// SIMPLIFICATION

// Simplify uses the SimplifiedEnglish and other methods to modify the Gematria and return a new Gematria
func (s Gematria) Simplify() Gematria {
	s.English = s.SimplifiedEnglish()
	s.Jewish = s.SimplifiedJewish()
	s.Simple = s.SimplifiedSimple()
	s.Mystery = s.SimplifiedMystery()
	s.Majestic = s.SimplifiedMajestic()
	s.Eights = s.SimplifiedEights()
	return s
}

// SimplifiedSimple discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedSimple() (i uint64) {
	i, _ = Simplify(s.Simple)
	return
}

// SimplifiedEnglish discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedEnglish() (i uint64) {
	i, _ = Simplify(s.English)
	return
}

// SimplifiedJewish discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedJewish() (i uint64) {
	i, _ = Simplify(s.Jewish)
	return
}

// SimplifiedMystery discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedMystery() (i uint64) {
	i, _ = Simplify(s.Mystery)
	return
}

// SimplifiedMajestic discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedMajestic() (i uint64) {
	i, _ = Simplify(s.Majestic)
	return
}

// SimplifiedEights discards the error from Simplify and returns the uint64 if it can be simplified,
// otherwise 0 is returned
func (s Gematria) SimplifiedEights() (i uint64) {
	i, _ = Simplify(s.Eights)
	return
}

// Combinations

// SimpleEnglish uses the Gematria.Simple and Gematria.English to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) SimpleEnglish() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Simple, 10))
	sb.WriteString(strconv.FormatUint(s.English, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// SimpleJewishEnglish uses the Gematria.Simple Gematria.Jewish and Gematria.English to merge the two
// integers into a new string, then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) SimpleJewishEnglish() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Simple, 10))
	sb.WriteString(strconv.FormatUint(s.Jewish, 10))
	sb.WriteString(strconv.FormatUint(s.English, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i

}

// SimpleJewish uses the Gematria.Simple and Gematria.Jewish to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) SimpleJewish() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Simple, 10))
	sb.WriteString(strconv.FormatUint(s.Jewish, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// MajesticEights uses the Gematria.Majestic and Gematria.Eights to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) MajesticEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Majestic, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// MajesticMystery uses the Gematria.Majestic and Gematria.Mystery to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) MajesticMystery() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Majestic, 10))
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// MajesticMysteryEights uses the Gematria.Majestic Gematria.Mystery and Gematria.Eights to merge the two integers
// into a new string, then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) MajesticMysteryEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Majestic, 10))
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// SimpleEights uses the Gematria.Simple and Gematria.Eights to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) SimpleEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Simple, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// SimpleMystery uses the Gematria.Simple and Gematria.Mystery to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) SimpleMystery() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Simple, 10))
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// EnglishMystery uses the Gematria.English and Gematria.Mystery to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) EnglishMystery() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.English, 10))
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// EnglishEights uses the Gematria.English and Gematria.Eights to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) EnglishEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.English, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// MysteryJewish uses the Gematria.Mystery and Gematria.Jewish to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) MysteryJewish() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	sb.WriteString(strconv.FormatUint(s.Jewish, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// MysteryEights uses the Gematria.Mystery and Gematria.Eights to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) MysteryEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Mystery, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// JewishEights uses the Gematria.Jewish and Gematria.Eights to merge the two integers into a new string,
// then return the parsed uint64 of the result. If an error occurs, 0 is returned.
func (s Gematria) JewishEights() uint64 {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatUint(s.Jewish, 10))
	sb.WriteString(strconv.FormatUint(s.Eights, 10))
	i, _ := strconv.ParseUint(sb.String(), 10, 64)
	return i
}

// ReadCombo accepts multiple ComboType and runs MajesticGematria on those pairs and returns the simplified uint64
func (s Gematria) ReadCombo(ct ...ComboType) uint64 {
	return MajesticGematria(s, ct...)
}

// JSON compiles the Gematria result into a formatted JSON string
func (s Gematria) JSON() string {
	b, err := json.MarshalIndent(s, ``, `   `)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
