// Package validation defined to store validation-oriented functions
package validation

import "net/mail"

func ValidMail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
