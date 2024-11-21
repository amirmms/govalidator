package govalidator

import (
	"regexp"
	"strings"
)

const (
	// Email represents rule name which will be used to find the default error message.
	Email = "email"
	// EmailMsg is the default error message format for fields with Email validation rule.
	EmailMsg = "%s is not valid"

	minEmailValidLen = 3
	maxEmailValidLen = 320
)

// EmailRegex is the default pattern to validate email field by RFC 5322 rule.
var EmailRegex = regexp.MustCompile(`^([a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)])$`)

// Email checks the value under validation must match the EmailRegex regular expression.
//
// Example:
//
//	v := validator.New()
//	v.Email("john.doe@example.com", "email", "email address is not valid.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) Email(s, field, msg string) Validator {
	if len(s) < minEmailValidLen || len(s) > maxEmailValidLen || !strings.Contains(s, "@") {
		v.check(false, field, v.msg(Email, msg, field))
		return v
	}

	v.check(EmailRegex.MatchString(s), field, v.msg(Email, msg, field))

	return v
}
