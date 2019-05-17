package data

import (
	"fmt"
	"time"
)

// Person has an ID and Name
type Person struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
}

// People is a list of Person
type People []Person

// Friendship is a map of Person.ID to []Person.ID
type Friendship map[int][]int

var (
	dbPeople = People{
		Person{1, "Frank Ocean", time.Date(1987, time.August, 28, 0, 0, 0, 0, time.UTC)},
		Person{2, "Paul Gascoigne", time.Date(1967, time.May, 27, 0, 0, 0, 0, time.UTC)},
		Person{3, "Uwe Seeler", time.Date(1936, time.November, 5, 0, 0, 0, 0, time.UTC)},
		Person{4, "Marko Marin", time.Date(1989, time.March, 13, 0, 0, 0, 0, time.UTC)},
	}

	dbFriends = Friendship{
		1: []int{2, 3, 4},
		2: []int{1, 4},
		3: []int{1},
		4: []int{2, 1},
	}
)

// All people returns the list of all Person
func All() People {
	return dbPeople
}

// PersonByID looks up a Person
func PersonByID(id int) (Person, error) {
	for _, p := range dbPeople {
		if p.ID == id {
			return p, nil
		}
	}

	return Person{}, fmt.Errorf("Cannot find person with ID: %d", id)
}

// FriendsByID looks up the list of friends for a Person
func FriendsByID(id int) (People, error) {
	list := People{}

	if refs, found := dbFriends[id]; found {
		for _, friendID := range refs {
			friend, err := PersonByID(friendID)

			if err != nil {
				return nil, err
			}

			list = append(list, friend)
		}
	}

	return list, nil
}
