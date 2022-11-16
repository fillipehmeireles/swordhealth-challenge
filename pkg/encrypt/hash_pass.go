package encrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(hashPass)

}

func CheckPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
