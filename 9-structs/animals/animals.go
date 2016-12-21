package animals

import "strings"

type Animal struct {
	Name string
	Age int
}

type Cat struct {
	Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

type AnimalInterface interface {
	GetName() string
	Talk() string
}

func (cat Cat) GetName() string {
	return cat.Name
}

func (cat Cat) Talk() string {
	return says("Meow!", cat.MeowStrength)
}

func (dog Dog) GetName() string {
	return dog.Name
}

func (dog Dog) Talk() string {
	return says("Woof!", dog.BarkStrength)
}

func says(says string, strength int) (string) {
	if strength >= 7 {
		return strings.ToUpper(says)
	} else if strength <= 3 {
		return strings.ToLower(says)
	}
	return strings.Title(says)
}

