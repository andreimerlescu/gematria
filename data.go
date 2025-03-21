package gematria

var (
	// eightsCodes uses the numerology theory of 3+5=8 then repeating 8 to 192 for A-Z and Simplify(Z)=3
	eightsCodes = ResultMap{
		"A": 3, "B": 5, "C": 8, "D": 16, "E": 32, "F": 40,
		"G": 48, "H": 54, "I": 63, "J": 69, "K": 77, "L": 80,
		"M": 88, "N": 96, "O": 104, "P": 112, "Q": 120, "R": 128,
		"S": 136, "T": 144, "U": 152, "V": 160, "W": 168, "X": 176,
		"Y": 184, "Z": 192,
		// Lowercase equivalents
		"a": 3, "b": 5, "c": 8, "d": 16, "e": 32, "f": 40,
		"g": 48, "h": 54, "i": 63, "j": 69, "k": 77, "l": 80,
		"m": 88, "n": 96, "o": 104, "p": 112, "q": 120, "r": 128,
		"s": 136, "t": 144, "u": 152, "v": 160, "w": 168, "x": 176,
		"y": 184, "z": 192,
	}
	// majesticCodes uses the numerology theory of 3+6=9 for A+B=C to Z=78 and Simplify(Z) = 6
	majesticCodes = ResultMap{
		"A": 3, "B": 6, "C": 9, "D": 12, "E": 15, "F": 18,
		"G": 21, "H": 24, "I": 27, "J": 30, "K": 33, "L": 36,
		"M": 39, "N": 42, "O": 45, "P": 47, "Q": 51, "R": 54,
		"S": 57, "T": 60, "U": 63, "V": 66, "W": 69, "X": 72,
		"Y": 75, "Z": 78,
		// Lowercase equivalents
		"a": 3, "b": 6, "c": 9, "d": 12, "e": 15, "f": 18,
		"g": 21, "h": 24, "i": 27, "j": 30, "k": 33, "l": 36,
		"m": 39, "n": 42, "o": 45, "p": 47, "q": 51, "r": 54,
		"s": 57, "t": 60, "u": 63, "v": 66, "w": 69, "x": 72,
		"y": 75, "z": 78,
	}
	// mysteryCodes uses the numerology of Saint Andrei - Life is a mystery - Everyone must stand alone and Simplify(Z)=6
	mysteryCodes = ResultMap{
		"A": 369, "B": 3, "C": 144, "D": 6, "E": 17, "F": 9,
		"G": 22, "H": 222, "I": 333, "J": 300, "K": 666, "L": 600,
		"M": 963, "N": 900, "O": 45, "P": 47, "Q": 1776, "R": 639,
		"S": 1618, "T": 60, "U": 999, "V": 434, "W": 99, "X": 88,
		"Y": 66, "Z": 33,
		// Lowercase equivalents
		"a": 369, "b": 3, "c": 144, "d": 6, "e": 17, "f": 9,
		"g": 22, "h": 222, "i": 333, "j": 300, "k": 666, "l": 600,
		"m": 963, "n": 900, "o": 45, "p": 47, "q": 1776, "r": 639,
		"s": 1618, "t": 60, "u": 999, "v": 434, "w": 99, "x": 88,
		"y": 66, "z": 33,
	}
	// jewishCodes uses the numerology of Gematria settled over thousands of years of human history and Simplify(Z)=5
	jewishCodes = ResultMap{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6,
		"G": 7, "H": 8, "I": 9, "J": 600, "K": 10, "L": 20,
		"M": 30, "N": 40, "O": 50, "P": 60, "Q": 70, "R": 80,
		"S": 90, "T": 100, "U": 200, "V": 700, "W": 900, "X": 300,
		"Y": 400, "Z": 500,
		// Lowercase equivalents
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
		"g": 7, "h": 8, "i": 9, "j": 600, "k": 10, "l": 20,
		"m": 30, "n": 40, "o": 50, "p": 60, "q": 70, "r": 80,
		"s": 90, "t": 100, "u": 200, "v": 700, "w": 900, "x": 300,
		"y": 400, "z": 500,
	}
	// englishCodes uses the numerology of Gematria settled over thousands of years of human history and Simplify(Z)=6
	englishCodes = ResultMap{
		"A": 6, "B": 12, "C": 18, "D": 24, "E": 30, "F": 36,
		"G": 42, "H": 48, "I": 54, "J": 60, "K": 66, "L": 72,
		"M": 78, "N": 84, "O": 90, "P": 96, "Q": 102, "R": 108,
		"S": 114, "T": 120, "U": 126, "V": 132, "W": 138, "X": 144,
		"Y": 150, "Z": 156,
		// Lowercase equivalents
		"a": 6, "b": 12, "c": 18, "d": 24, "e": 30, "f": 36,
		"g": 42, "h": 48, "i": 54, "j": 60, "k": 66, "l": 72,
		"m": 78, "n": 84, "o": 90, "p": 96, "q": 102, "r": 108,
		"s": 114, "t": 120, "u": 126, "v": 132, "w": 138, "x": 144,
		"y": 150, "z": 156,
	}
	// simpleCodes uses the numerology of Gematria to establish A=1, B=2 and Z=26 and Simplify(Z)=7
	simpleCodes = ResultMap{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6,
		"G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12,
		"M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18,
		"S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24,
		"Y": 25, "Z": 26,
		// Lowercase equivalents
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
		"g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12,
		"m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18,
		"s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24,
		"y": 25, "z": 26,
	}
	// and for the record Simplify(Z) = 5 + Simplify(Z) = 7 === 5+7 = 12 = 1+2 = 3 therefore all codes 369
)

// ResultMap stores the character string and the Gematria uint64 value
type ResultMap map[string]uint64

// Gematria stores the original string and the various Gematria uint64 values
type Gematria struct {
	original string
	Jewish   uint64 `json:"jewish" yaml:"Jewish"`
	English  uint64 `json:"english" yaml:"English"`
	Simple   uint64 `json:"simple" yaml:"Simple"`
	Mystery  uint64 `json:"mystery" yaml:"Mystery"`
	Majestic uint64 `json:"majestic" yaml:"Majestic"`
	Eights   uint64 `json:"eights" yaml:"Eights"`
}
