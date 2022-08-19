package game

import (
	"math/rand"
)

type Suggestions map[string]string

func GenSuggest() Suggestions {
	easyWordlist := []string{"carrot", "terminal", "search", "selection"}

	suggest := make(Suggestions)

	for {

		randomNumber := rand.Intn(len(easyWordlist))
		suggest[easyWordlist[randomNumber]] = easyWordlist[randomNumber]
		if len(suggest) == 3 {
			break
		}
	}
	return suggest
}
