package auth

import (
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	hasspwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hasspwd), nil
}
func Comparepasswords(hash, plain string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil

}
