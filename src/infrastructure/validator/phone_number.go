package validator

import "regexp"

func ValidateUzPhoneNumber(phoneNumber string) bool {
	regex := `^998\d{9}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(phoneNumber)
}
