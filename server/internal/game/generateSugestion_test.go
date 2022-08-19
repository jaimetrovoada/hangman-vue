package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenSuggest(t *testing.T) {

	suggestion := GenSuggest()
	// fmt.Printf("suugestions %v", suggestion)

	assert.Equal(t, 3, len(suggestion))

}
