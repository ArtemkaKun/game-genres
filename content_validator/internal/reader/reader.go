package reader

import (
	"content_validator/internal/data"
	"errors"

	"encoding/json"
	"fmt"
	"os"
)

var errNoGameGenresFound = errors.New("no game genres found in JSON")

// ReadGameGenresFromJSON reads and parses game genres from a JSON file.
//
// Parameters:
//
//	jsonFilePath: The path to the JSON file containing game genre data
//
// Returns:
//
//	[]data.GameGenre: A slice of GameGenre objects parsed from the JSON file
//	error: An error if the file cannot be read or if the JSON structure is invalid
//
// Examples:
//
//	genres, err := ReadGameGenresFromJSON("game_genres.json")
//
//	if err != nil {
//	    log.Fatalf("Failed to read game genres: %v", err)
//	}
//
// Errors:
//
//   - Returns "error reading file: [underlying error]" if the file cannot be read
//   - Returns "invalid structure: [underlying error]" if the JSON cannot be parsed into GameGenre objects
//
// Note:
//
//	The function expects the JSON file to contain an array of objects that can be
//	unmarshaled into the data.GameGenre struct. Make sure the JSON structure
//	matches the GameGenre definition.
func ReadGameGenresFromJSON(jsonFilePath string) ([]data.GameGenre, error) {
	content, err := os.ReadFile(jsonFilePath)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var gameGenres []data.GameGenre

	err = json.Unmarshal(content, &gameGenres)

	if err != nil {
		return nil, fmt.Errorf("invalid structure: %w", err)
	}

	if len(gameGenres) == 0 {
		return nil, errNoGameGenresFound
	}

	return gameGenres, nil
}
