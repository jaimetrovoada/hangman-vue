package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameList(t *testing.T) {
	gamelist := NewGameList()

	game1 := NewGame(gamelist.NewRoomId())
	game2 := NewGame(gamelist.NewRoomId())
	gamelist.AddGame(game1)
	gamelist.AddGame(game2)
	playerList1 := game1.playerList
	playerList2 := game2.playerList

	player1 := game1.NewPlayer()
	player2 := game1.NewPlayer()
	playerList1 = append(playerList1, *player1, *player2)

	g2player1 := game2.NewPlayer()
	playerList2 = append(playerList2, *g2player1)

	want := "game list has 2 games, game 1 has 2 players, game 2 has 1 player"
	got := fmt.Sprintf("game list has %d games, game 1 has %d players, game 2 has %d player", gamelist.GameListLength(), len(playerList1), len(playerList2))

	// fmt.Printf("gl len: %d", gamelist.GameListLength())
	// fmt.Printf("game 1 id: %d", game1.roomId)

	if ok := assert.Equal(t, want, got); !ok {
		t.Errorf("want:\t%s\ngot:\t%s", want, got)
	}

}

func TestGame(t *testing.T) {
	type TestCaseInput struct {
		word       string
		guess      []string
		numOfGuess int
	}
	type TestCase struct {
		want  string
		input TestCaseInput
	}

	// guess charCOunt ok
	test1 := TestCase{"0 0 false", TestCaseInput{"torrent", []string{"e", "b", "c", "d", "e", "f", "g", "h", "i", "j", "e"}, 10}}
	testCases := []TestCase{
		test1,
	}

	gameList := NewGameList()
	for _, test := range testCases {
		game := NewGame(gameList.NewRoomId())
		game.SetWord(test.input.word)
		var charCount int
		var ok bool
		for _, guess := range test.input.guess {
			charCount, ok = game.MakeGuess(guess)
		}
		guessCount := game.counter
		// charCount, ok := game.MakeGuess("t")
		// count, charIndex, ok := Guess("torrent", "t")

		want := test.want
		got := fmt.Sprintf("%d %d %v", guessCount, charCount, ok)
		if ok := assert.Equal(t, want, got); !ok {
			t.Errorf("want: {%s}, got: {%s}", want, got)
		}

	}
}

func TestGamePlayers(t *testing.T) {
	gameList := NewGameList()
	game := NewGame(gameList.NewRoomId())
	playerList := game.playerList

	player1 := game.NewPlayer()
	player2 := game.NewPlayer()
	playerList = append(playerList, *player1, *player2)

	want := "2 players joined"
	got := fmt.Sprintf("%d players joined", len(playerList))

	if ok := assert.Equal(t, want, got); !ok {
		t.Errorf("want:\t%s\ngot:\t%s", want, got)
	}
}
