package game

import (
	"math/rand"
	"strings"
)

type Player struct {
	id int
}

func (game *Game) NewPlayer() *Player {
	if len(game.playerList) == 0 {

		return &Player{
			id: 1,
		}
	}
	return &Player{
		id: 2,
	}
}

type GameList map[int]Game

func NewGameList() *GameList {
	return &GameList{}
}
func (gamelist *GameList) AddGame(game *Game) {
	gamelistPtr := *gamelist
	gamePtr := *game
	gamelistPtr[gamePtr.roomId] = gamePtr
}
func (gamelist *GameList) GameListLength() int {
	return len(*gamelist)
}
func (gamelist *GameList) NewRoomId() int {
	var roomId int
	for {

		roomId = rand.Intn(100)
		gamelistPtr := *gamelist
		if _, ok := gamelistPtr[roomId]; !ok {
			break
		}
	}
	return roomId
}

type Game struct {
	roomId       int
	counter      int
	selectedWord string
	guess        string
	playerList   []struct{ id int }
}

func NewGame(roomId int) *Game {

	game := &Game{
		roomId:       roomId,
		counter:      0,
		selectedWord: "",
		guess:        "",
		playerList:   make([]struct{ id int }, 0, 2),
	}
	return game
}

func (game *Game) IncrCounter() {
	game.counter++
}

func (game *Game) SetWord(selWord string) {
	game.selectedWord = selWord
}
func (game *Game) MakeGuess(guess string) (int, bool) {
	// game.guess = guess

	if game.counter >= 10 {
		game.counter = 0
		return 0, false
	}
	ok := strings.Contains(game.selectedWord, guess)
	count := strings.Count(game.selectedWord, guess)
	// charIndex := strings.Index(selectedWord, userGuess)

	game.IncrCounter()
	return count, ok
}
