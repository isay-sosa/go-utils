package main

import (
	ar "github.com/OscarSwanros/go-utils/arraylist"
	"log"
)

func main() {
	al := ar.New()

	al.Add("Oscar")
	log.Println(al)

	log.Printf("%v", al.Remove("iscar"))
}
