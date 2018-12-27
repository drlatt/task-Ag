package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Friends []*Person
}

// print friend of friend's names
func (p *Person) printFriendOfFriends() {
	fmt.Printf("%s's FOF: ", p.Name)

	for _, f := range p.FriendOfFriends() {
		fmt.Printf("%s,", f.Name)
	}

	fmt.Println()
}

func (p *Person) FriendOfFriends() []*Person {
	friends := make(map[string]struct{})
	friendsOfFriends := make([]*Person, 0)

	friends[p.Name] = struct{}{}
	for _, f := range p.Friends {
		friends[f.Name] = struct{}{}
	}

	for _, f := range p.Friends {
		for _, ff := range f.Friends {
			// ignore individual's friends
			if _, isfriend := friends[ff.Name]; !isfriend {
				friendsOfFriends = append(friendsOfFriends, ff)
			}
		}
	}

	return friendsOfFriends
}

func addFriends(person *Person, friends []*Person) {
	for _, f := range friends {
		person.Friends = append(person.Friends, f)
	}
}

func addAndGet(person *Person, persons []*Person) (*Person, []*Person) {
	return person, append(persons, person)
}

func main() {
	persons := make([]*Person, 0)
	lateef, persons := addAndGet(&Person{Name: "Lateef", Friends: []*Person{}}, persons)
	ayo, persons := addAndGet(&Person{Name: "Ayo", Friends: []*Person{}}, persons)
	deji, persons := addAndGet(&Person{Name: "Deji", Friends: []*Person{}}, persons)
	kenneth, persons := addAndGet(&Person{Name: "Kenneth", Friends: []*Person{}}, persons)
	anthonia, persons := addAndGet(&Person{Name: "Anthonia", Friends: []*Person{}}, persons)
	abubakar, persons := addAndGet(&Person{Name: "Abubakar", Friends: []*Person{}}, persons)
	abdullahi, persons := addAndGet(&Person{Name: "Abdullahi", Friends: []*Person{}}, persons)
	aanu, persons := addAndGet(&Person{Name: "Aanu", Friends: []*Person{}}, persons)

	addFriends(lateef, []*Person{kenneth, ayo, anthonia})
	addFriends(ayo, []*Person{lateef, abubakar, deji})
	addFriends(deji, []*Person{abdullahi, aanu, ayo, lateef})

	for _, person := range persons {
		person.printFriendOfFriends()
	}
}
