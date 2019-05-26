package main

import "fmt"

type Person struct {
	name    string
	friends []string
	persons []*Person
}

func (p *Person) friendOfFriends() {
	fmt.Printf("%s's 2nd level network includes ", p.name)
}

func addFriends(p *Person, persons []*Person) {
	for _, b := range persons {
		p.persons = append(p.persons, b)
	}
}

func main() {

	friendsList := map[string][]string{
		"ken":    {"lateef", "pamela", "afolabi", "ken"},
		"lateef": {"antho", "pamela", "chike", "ken"},
		"chike":  {"toun", "ife", "lateef"},
	}
}
