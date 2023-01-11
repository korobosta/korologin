// hash.go
package korologin

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        panic(err.Error())
    }
    return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}