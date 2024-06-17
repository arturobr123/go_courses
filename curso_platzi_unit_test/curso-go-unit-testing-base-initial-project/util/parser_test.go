package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)
	c.NotNil(body)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)
	c.NotNil(response)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)
	c.NotNil(parsedPokemon)

	//////////////////////////////

	body_api, err := ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)
	c.NotNil(body_api)

	var expectedPokemon models.Pokemon
	err = json.Unmarshal(body_api, &expectedPokemon)
	c.NoError(err)
	c.NotNil(expectedPokemon)

	c.Equal(expectedPokemon, parsedPokemon)
}

// go test ./util
// go test ./util -run=TestParsePokemonSuccess

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)
	c.NotNil(body)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)
	c.NotNil(response)

	response.PokemonType = []models.PokemonType{}

	_, err2 := ParsePokemon(response)
	c.Error(err2)
	fmt.Println(err2.Error())
	c.Equal(ErrNotFoundPokemonType.Error(), err2.Error())

}
