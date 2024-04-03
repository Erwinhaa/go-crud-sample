package tools

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	} else if err != nil {
		panic(err)
	}
	return nil
}
