package main

import (
	"fmt"
	"./animals"
)

func main() {
	dogs := []animals.AnimalInterface {
		animals.Dog{
			Animal: animals.Animal {
				Name: "Fido",
				Age: 2,
			},
			BarkStrength: 8,
		},
		animals.Dog{
			Animal: animals.Animal {
				Name: "Rover",
				Age: 4,
			},
			BarkStrength: 5,
		},
		animals.Cat{
			Animal: animals.Animal{
				Name: "Snowball",
				Age: 3,
			},
			MeowStrength: 2,
		},
		animals.Cat{
			Animal: animals.Animal{
				Name: "Whiskers",
				Age: 3,
			},
			MeowStrength: 4,
		},
	}

	for _, value := range dogs {
		fmt.Printf("%s says %s \n", value.GetName(), value.Talk())
	}
}
