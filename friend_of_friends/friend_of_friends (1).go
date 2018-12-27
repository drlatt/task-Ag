package main

import (
	"fmt"
)

type person struct {
	name       string
	friends    []string
	friendsMap map[string][]string
	matchedMap map[string][]string
	// tempMap    map[string][]string
	tempSlice []string
}

func (p *person) friendOfFriends() {

	tempMap := make(map[string][]string)
	for k, v := range p.friendsMap {
		// fmt.Println(k, v)
		for _, y := range v {
			// fmt.Println(y)

			// if friend's name exists in friendlist, then that is a friend of friend
			if x, ok := p.friendsMap[y]; ok {
				// fmt.Printf("%s and %s are friends \n", k, y)
				// x = append(x, y)
				// fmt.Println(len(p.matchedMap))
				fmt.Println(k, x)
				p.tempSlice = x
				tempMap[k] = x
				// }
			}
			// fmt.Println("slice", p.tempSlice)

		}
		fmt.Println("slice", p.tempSlice)

	}

	fmt.Println(tempMap)
	// fmt.Println(p.tempSlice)
	return
}

func main() {
	p := person{friendsMap: map[string][]string{
		"Lateef": []string{"Kenneth", "Ayo", "Anthonia"},
		"Ayo":    []string{"Lateef", "Abubakar", "Deji"},
		"Deji":   []string{"Abdullahi", "Aanu", "Ayo"},
	}, matchedMap: map[string][]string{
		"Lateef": []string{},
		"Ayo":    []string{},
		"Deji":   []string{},
	},
	}

	p.friendOfFriends()

}
