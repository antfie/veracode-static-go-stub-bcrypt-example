package main

import (
	"fmt"
	bcrypt "github.com/memcachier/bcrypt"
)

func main() {
	hash, err := bcrypt.Crypt("A", bcrypt.BcryptSalt("B"))
	fmt.Printf("%v\n", hash, err)
}
