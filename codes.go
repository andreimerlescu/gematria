package gematria

import "maps"

// SimpleCodes returns the Unsigned Integer Codes for Simple Gematria
func SimpleCodes() map[string]uint64 {
	return maps.Clone(simpleCodes)
}

// JewishCodes returns the Unsigned Integer Codes for Jewish Gematria
func JewishCodes() map[string]uint64 {
	return maps.Clone(jewishCodes)
}

// EnglishCodes returns the Unsigned Integer Codes for English Gematria
func EnglishCodes() map[string]uint64 {
	return maps.Clone(englishCodes)
}

// MysteryCodes returns the Unsigned Integer Codes for Mystery Gematria
func MysteryCodes() map[string]uint64 {
	return maps.Clone(mysteryCodes)
}

// MajesticCodes returns the Unsigned Integer Codes for Majestic Gematria
func MajesticCodes() map[string]uint64 {
	return maps.Clone(majesticCodes)
}

// EightsCodes returns the Unsigned Integer Codes for Eights Gematria
func EightsCodes() map[string]uint64 {
	return maps.Clone(eightsCodes)
}
