package roleValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/inventapp_v2/validators"
)

var (
	isValidName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateName(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("name only can contain letters or spaces")
	}
	return nil
}
