package main

import (
	"content_validator/internal/data"
	"content_validator/internal/reader"
	"content_validator/internal/validation"
	"log"
	"os"
)

const expectedNumberOfArguments = 2

func main() {
	if len(os.Args) != expectedNumberOfArguments {
		log.Fatalf("Usage: %s <path-to-json-file>", os.Args[0])
	}

	filePath := os.Args[1]

	gameGenres, err := reader.ReadGameGenresFromJSON(filePath)

	if err != nil {
		log.Fatalf("Failed to read game genres: %v", err)
	}

	validateNamesNotEmpty(gameGenres)
	validateNamesTrimmed(gameGenres)
	validateNamesCase(gameGenres)
	validateNamesUnique(gameGenres)
	validateNamesCollision(gameGenres)
}

func validateNamesNotEmpty(gameGenres []data.GameGenre) {
	if !validation.ValidateNameNotEmpty(gameGenres) {
		log.Fatal("There are game genres with empty names")
	}

	result, invalidEntities := validation.ValidateAltNamesNotEmpty(gameGenres)

	if !result {
		log.Println("There are game genres with empty alternative names:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}
}

func validateNamesTrimmed(gameGenres []data.GameGenre) {
	result, invalidEntities := validation.ValidateNameTrimmed(gameGenres)

	if !result {
		log.Println("There are game genres with leading or trailing whitespace in their names:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}

	result, invalidEntities = validation.ValidateAltNamesTrimmed(gameGenres)

	if !result {
		log.Println("There are game genres with leading or trailing whitespace in their alternative names:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}
}

func validateNamesCase(gameGenres []data.GameGenre) {
	result, invalidEntities := validation.ValidateNameCase(gameGenres)

	if !result {
		log.Println("There are game genres with names that are not in lowercase:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}

	result, invalidEntities = validation.ValidateAltNamesCase(gameGenres)

	if !result {
		log.Println("There are game genres with alternative names that are not in lowercase:")

		for _, altName := range invalidEntities {
			log.Println(altName)
		}

		os.Exit(1)
	}
}

func validateNamesUnique(gameGenres []data.GameGenre) {
	result, invalidEntities := validation.ValidateNameUnique(gameGenres)

	if !result {
		log.Println("There are game genres with duplicate names:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}

	result, invalidEntities = validation.ValidateAltNamesUnique(gameGenres)

	if !result {
		log.Println("There are game genres with duplicate alternative names:")

		for _, genre := range invalidEntities {
			log.Println(genre)
		}

		os.Exit(1)
	}
}

func validateNamesCollision(gameGenres []data.GameGenre) {
	result, nameCollisions := validation.ValidateGenreNameNoCollisionsWithAltNames(gameGenres)

	if !result {
		log.Println("There are game genres with names that are also alternative names:")

		for _, collision := range nameCollisions {
			log.Printf("%s - %s", collision.CollidingGenreName, collision.GenreWithCollidingAltName)
		}

		os.Exit(1)
	}

	result, altNameCollisions := validation.ValidateCollidingAltNames(gameGenres)

	if !result {
		log.Println("There are game genres with alternative names that are also names:")

		for _, collision := range altNameCollisions {
			log.Printf("%s: %s - %s", collision.AltName, collision.CollidingGenreName, collision.GenreWithCollidingAltName)
		}

		os.Exit(1)
	}
}
