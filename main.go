package main

import (
	bcrypt "github.com/memcachier/bcrypt"
)

func main() {
	_, _ := bcrypt.Crypt("A", "B")
}
