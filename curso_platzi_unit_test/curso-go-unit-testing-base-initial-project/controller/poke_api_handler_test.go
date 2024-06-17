package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	c := require.New(t)

	pokemon, err := GetPokemonFromApi("1")
	c.NoError(err)

	body, err := ioutil.ReadFile("poke_api_readed.json")

	if err != nil {
		c.Fail("Error reading file")
	}

	var expected models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &expected)
	c.NoError(err)

	c.Equal(pokemon, expected)
}

func TestGetPokemonFromPokeApiSuccessWithMock(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "balbasaur"

	body, err := ioutil.ReadFile("poke_api_response.json")
	c.NoError(err)

	// Register a mock HTTP responder for the PokeAPI
	httpmock.RegisterResponder(
		"GET", // Method type
		fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id),  // URL pattern to match
		httpmock.NewStringResponder(http.StatusOK, string(body)), // Responder that returns a status code and response body
	)

	pokemon, err := GetPokemonFromApi(id)
	c.NoError(err)

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal(body, &expected)
	c.NoError(err)

	c.Equal(pokemon, expected)
}

func TestGetPokemonFromPokeApiInternalServerError(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "balbasaur"

	body, err := ioutil.ReadFile("poke_api_response.json")
	c.NoError(err)

	// Register a mock HTTP responder for the PokeAPI
	httpmock.RegisterResponder(
		"GET", // Method type
		fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id),                   // URL pattern to match
		httpmock.NewStringResponder(http.StatusInternalServerError, string(body)), // Responder that returns a status code and response body
	)

	_, err = GetPokemonFromApi(id)
	c.NotNil(err)
	c.Equal(err, ErrPokeApiFailure)
}
