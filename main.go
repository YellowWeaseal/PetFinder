package main

import "C"
import (
	"PetProject/PetFinder"
	"fmt"
	"log"
	"time"
)

func main() {
	PetFinder.Access()
	petFinderClient, err := PetFinder.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}
	myAnimal, err := PetFinder.Client.GetAnimal(*petFinderClient, 142)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myAnimal.ID, myAnimal.Species, myAnimal.Breeds)
}
