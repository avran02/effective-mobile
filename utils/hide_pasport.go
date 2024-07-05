package utils

import "strings"

const numSymbolsToShow = 6

func HidePassportNumber(passportNumber string) string {
	if len(passportNumber) < numSymbolsToShow {
		return passportNumber
	}

	passportLength := len(passportNumber)

	return passportNumber[0:4] + strings.Repeat("*", passportLength-numSymbolsToShow) + passportNumber[passportLength-2:]
}
