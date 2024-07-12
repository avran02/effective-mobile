package validation

import "regexp"

func PassportNumber(passportNumber string) error {
	re := regexp.MustCompile(`^\d{4} \d{6}$`)
	isValid := re.MatchString(passportNumber)

	if !isValid {
		return ErrInvalidPassportNumber
	}

	return nil
}
