package errorformat

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Username sudah ada")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Password Salah")
	}
	return errors.New("Detail salah")
}
