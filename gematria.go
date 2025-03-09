package gematria

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func characterScores(str string) (uint64, uint64, uint64, uint64, uint64, uint64, error) {
	if len(str) != 1 {
		return 0, 0, 0, 0, 0, 0, fmt.Errorf("string must only contain 1 character")
	}
	j, e, s, my, maj, eights := jewishCodes[str], englishCodes[str], simpleCodes[str], mysteryCodes[str], majesticCodes[str], eightsCodes[str]
	return j, e, s, my, maj, eights, nil
}

func FromString(in string) (out Gematria) {
	var err error
	out, err = NewGematria(in)
	if err != nil {
		log.Println(err)
	}
	return
}

func NewGematria(data string) (Gematria, error) {
	re, err := regexp.Compile(`[^a-zA-Z\d.\s]`)
	if err != nil {
		return Gematria{}, fmt.Errorf("failed to compile gematria regex: %w", err)
	}
	data = re.ReplaceAllString(data, "")
	data = strings.TrimLeft(data, "")
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
	return Gematria{Jewish: jf, English: ef, Simple: sf, Mystery: myf, Majestic: majf, Eights: eigf}, nil
}

func (s Gematria) String() string {
	output := "\t"
	output += fmt.Sprintf("%s = %d \t", "Jewish", s.Jewish)
	output += fmt.Sprintf("%s = %d \t", "English", s.English)
	output += fmt.Sprintf("%s = %d \t", "Simple", s.Simple)
	output += fmt.Sprintf("%s = %d \t", "Mystery", s.Mystery)
	output += fmt.Sprintf("%s = %d \t", "Majestic", s.Majestic)
	output += fmt.Sprintf("%s = %d \t", "Eights", s.Eights)
	return output
}
