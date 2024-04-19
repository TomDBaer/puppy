package puppy

import "github.com/TomDBaer/dog"

func Bark() string {
	return "Bork!"
}

func Barks() string {
	return "Bork! Bork! Bork!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
