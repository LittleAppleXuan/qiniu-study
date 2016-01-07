package main

import (
	"./stringutils"
	"log"
)

// test stringutils
func main() {

	log.Println(stringutils.IsEmail("contact@idrmfly.com"))
	log.Println(stringutils.IsEmail("contact.liw@idrmfly.com"))
	log.Println(stringutils.IsEmail("contact-liw@idrmfly.com"))
	log.Println(stringutils.IsEmail("contact-liwidrmfly.com"))
	log.Println(stringutils.IsEmail("contact-liw@idrmflycom"))

}
