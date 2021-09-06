package isValid

import (
	"log"
	"regexp"
)

//will return true if the provided email is valid, false if not,
// and err if there is a problem with the regex
func Email(email string) (bool, error) {
	emailReg, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		log.Printf("email validation regexp didn't compile")
		return false, err
	}
	if !emailReg.MatchString(email) {
		return false, nil
	}

	return true, nil
}

//will return true if the provided gsm is valid, false if not,
// and err if there is a problem with the regex
func GSM(gsm string) (bool, error) {
	phoneReg, err := regexp.Compile("^((080)|(081)|(070)|(090))\\d{8}$")
	if err != nil {
		log.Printf("phone validation regexp didn't compile")
		return false, err
	}
	if !phoneReg.MatchString(gsm) {
		return false, nil
	}

	return true, nil
}
