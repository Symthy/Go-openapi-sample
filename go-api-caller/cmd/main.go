package main

import (
	"fmt"

	"github.com/golang-practices/go-api-caller/api"
)

func main() {

	poke, err := api.GetPokemon("306")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(poke)
}
