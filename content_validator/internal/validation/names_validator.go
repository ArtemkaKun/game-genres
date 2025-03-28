package validation

import (
	"content_validator/internal/data"
	"slices"
	"strings"
)

// ValidateNameNotEmpty checks if all game genres in the provided slice have non-empty names.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all genres have non-empty names, false otherwise
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "Action"},
//	    {Name: "Adventure"},
//	}
//
//	ValidateNameNotEmpty(validGenres)  // returns true
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "Action"},
//	    {Name: ""},
//	}
//
//	ValidateNameNotEmpty(invalidGenres)  // returns false
//
// Note:
//
//	If the input slice is empty, the function returns true.
func ValidateNameNotEmpty(genres []data.GameGenre) bool {
	for _, genre := range genres {
		if strings.TrimSpace(genre.Name) == "" {
			return false
		}
	}

	return true
}

// ValidateAltNamesNotEmpty checks if all alternative names for all game genres are non-empty.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all alternative names are non-empty, false otherwise
//	[]string: A slice containing the names of genres with empty alternative names, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", "Fighting"}},
//	    {Name: "Adventure", AltNames: []string{"Adv"}},
//	}
//
//	valid, _ := ValidateAltNamesNotEmpty(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", ""}},
//	    {Name: "Adventure", AltNames: []string{"Adv"}},
//	}
//
//	valid, invalid := ValidateAltNamesNotEmpty(invalidGenres)  // returns false, []string{"Action"}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", ""}},
//	    {Name: "Adventure", AltNames: []string{"Adv", "  "}},
//	}
//
//	valid, invalid := ValidateAltNamesNotEmpty(multipleInvalidGenres)  // returns false, []string{"Action", "Adventure"}
//
// Note:
//
//	A genre name appears only once in the returned slice even if it has multiple empty alternative names.
//	Genres with no alternative names (empty slice) are considered valid.
func ValidateAltNamesNotEmpty(genres []data.GameGenre) (bool, []string) {
	var invalidGenreNames []string

	for _, genre := range genres {
		for _, altName := range genre.AltNames {
			if strings.TrimSpace(altName) == "" {
				if !slices.Contains(invalidGenreNames, genre.Name) {
					invalidGenreNames = append(invalidGenreNames, genre.Name)
				}
			}
		}
	}

	if len(invalidGenreNames) == 0 {
		return true, nil
	}

	return false, invalidGenreNames
}

// ValidateNameTrimmed checks if all game genre names are properly trimmed of leading and trailing whitespace.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all genre names are properly trimmed, false otherwise
//	[]string: A slice containing the untrimmed genre names, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "Action"},
//	    {Name: "Adventure"},
//	}
//
//	valid, _ := ValidateNameTrimmed(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "Action"},
//	    {Name: "Adventure "},
//	}
//
//	valid, invalid := ValidateNameTrimmed(invalidGenres)  // returns false, []string{"Adventure "}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: " Action"},
//	    {Name: "Adventure "},
//	}
//
//	valid, invalid := ValidateNameTrimmed(multipleInvalidGenres)  // returns false, []string{" Action", "Adventure "}
//
// Note:
//
//	The function only checks for leading and trailing whitespace, not whitespace within the name.
func ValidateNameTrimmed(genres []data.GameGenre) (bool, []string) {
	var invalidNames []string

	for _, genre := range genres {
		if genre.Name != strings.TrimSpace(genre.Name) {
			invalidNames = append(invalidNames, genre.Name)
		}
	}

	if len(invalidNames) == 0 {
		return true, nil
	}

	return false, invalidNames
}

// ValidateAltNamesTrimmed checks if all alternative names for all game genres are properly trimmed
// of leading and trailing whitespace.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all alternative names are properly trimmed, false otherwise
//	[]string: A slice containing the names of genres with untrimmed alternative names, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", "Fighting"}},
//	    {Name: "Adventure", AltNames: []string{"Adv"}},
//	}
//
//	valid, _ := ValidateAltNamesTrimmed(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", "Fighting "}},
//	    {Name: "Adventure", AltNames: []string{"Adv"}},
//	}
//
//	valid, invalid := ValidateAltNamesTrimmed(invalidGenres)  // returns false, []string{"Action"}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: "Action", AltNames: []string{"Act", " Fighting"}},
//	    {Name: "Adventure", AltNames: []string{"Adv", "  RPG "}},
//	}
//
//	valid, invalid := ValidateAltNamesTrimmed(multipleInvalidGenres)  // returns false, []string{"Action", "Adventure"}
//
// Note:
//
//	A genre name appears only once in the returned slice even if it has multiple untrimmed alternative names.
//	The function returns the genre names (not the alternative names) that contain untrimmed alternative names.
//	Genres with no alternative names (empty slice) are considered valid.
func ValidateAltNamesTrimmed(genres []data.GameGenre) (bool, []string) {
	var invalidGenreNames []string

	for _, genre := range genres {
		for _, altName := range genre.AltNames {
			if altName != strings.TrimSpace(altName) {
				if !slices.Contains(invalidGenreNames, genre.Name) {
					invalidGenreNames = append(invalidGenreNames, genre.Name)
				}
			}
		}
	}

	if len(invalidGenreNames) == 0 {
		return true, nil
	}

	return false, invalidGenreNames
}

// ValidateNameCase checks if all game genre names are in lowercase.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all genre names are in lowercase, false otherwise
//	[]string: A slice containing the genre names that are not in lowercase, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "action"},
//	    {Name: "adventure"},
//	}
//
//	valid, _ := ValidateNameCase(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "action"},
//	    {Name: "Adventure"},
//	}
//
//	valid, invalid := ValidateNameCase(invalidGenres)  // returns false, []string{"Adventure"}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: "Action"},
//	    {Name: "Adventure"},
//	}
//
//	valid, invalid := ValidateNameCase(multipleInvalidGenres)  // returns false, []string{"Action", "Adventure"}
//
// Note:
//
//	Names with numbers, symbols, or spaces are considered valid as long as any letters are lowercase.
func ValidateNameCase(genres []data.GameGenre) (bool, []string) {
	var invalidNames []string

	for _, genre := range genres {
		if genre.Name != strings.ToLower(genre.Name) {
			invalidNames = append(invalidNames, genre.Name)
		}
	}

	if len(invalidNames) == 0 {
		return true, nil
	}

	return false, invalidNames
}

// ValidateAltNamesCase checks if all alternative names for all game genres are in lowercase.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all alternative names are in lowercase, false otherwise
//	[]string: A slice containing the alternative names that are not in lowercase, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv"}},
//	}
//
//	valid, _ := ValidateAltNamesCase(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "Fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv"}},
//	}
//
//	valid, invalid := ValidateAltNamesCase(invalidGenres)  // returns false, []string{"Fighting"}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"Act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"Adv", "RPG"}},
//	}
//
//	valid, invalid := ValidateAltNamesCase(multipleInvalidGenres)  // returns false, []string{"Act", "Adv", "RPG"}
//
// Note:
//
//	Alternative names with numbers, symbols, or spaces are considered valid as long as any letters are lowercase.
//	Genres with no alternative names (empty slice) are considered valid.
func ValidateAltNamesCase(genres []data.GameGenre) (bool, []string) {
	var invalidAltNames []string

	for _, genre := range genres {
		for _, altName := range genre.AltNames {
			if altName != strings.ToLower(altName) {
				invalidAltNames = append(invalidAltNames, altName)
			}
		}
	}

	if len(invalidAltNames) == 0 {
		return true, nil
	}

	return false, invalidAltNames
}

// ValidateNameUnique checks if there are any duplicate genre names in the provided slice.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all genre names are unique, false if duplicates exist
//	[]string: A slice containing the duplicate genre names, or nil if none found
//
// Examples:
//
//	uniqueGenres := []data.GameGenre{
//	    {Name: "action"},
//	    {Name: "adventure"},
//	    {Name: "rpg"},
//	}
//
//	valid, _ := ValidateNameUnique(uniqueGenres)  // returns true, nil
//
//	duplicateGenres := []data.GameGenre{
//	    {Name: "action"},
//	    {Name: "adventure"},
//	    {Name: "action"},  // duplicate name
//	}
//
//	valid, dupes := ValidateNameUnique(duplicateGenres)  // returns false, []string{"action"}
//
//	multipleGenres := []data.GameGenre{
//	    {Name: "action"},
//	    {Name: "action"},  // duplicate
//	    {Name: "rpg"},
//	    {Name: "rpg"},     // duplicate
//	}
//
//	valid, dupes := ValidateNameUnique(multipleGenres)  // returns false, []string{"action", "rpg"}
//
// Note:
//
//	The function only checks for exact string matches and is case-sensitive.
//	The order of duplicate names in the returned slice is not guaranteed.
func ValidateNameUnique(genres []data.GameGenre) (bool, []string) {
	nameCounts := make(map[string]int)

	for _, category := range genres {
		nameCounts[category.Name]++
	}

	var duplicates []string

	for name, count := range nameCounts {
		if count > 1 {
			duplicates = append(duplicates, name)
		}
	}

	if len(duplicates) > 0 {
		return false, duplicates
	}

	return true, nil
}

// ValidateAltNamesUnique checks if all alternative names within each game genre are unique.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if all genres have unique alternative names, false otherwise
//	[]string: A slice containing the names of genres with duplicate alternative names, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv", "quest"}},
//	}
//
//	valid, _ := ValidateAltNamesUnique(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting", "act"}},  // duplicate "act"
//	    {Name: "adventure", AltNames: []string{"adv", "quest"}},
//	}
//
//	valid, invalid := ValidateAltNamesUnique(invalidGenres)  // returns false, []string{"action"}
//
//	multipleInvalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "act"}},  // duplicate
//	    {Name: "adventure", AltNames: []string{"adv", "quest", "adv"}},  // duplicate
//	}
//
//	valid, invalid := ValidateAltNamesUnique(multipleInvalidGenres)  // returns false, []string{"action", "adventure"}
//
// Note:
//
//	The function returns the names of genres that contain duplicate alternative names, not the duplicate names
//	themselves.
//	The function checks for duplicates within each genre's alternative names, not across different genres.
//	Genres with no alternative names or only one alternative name are always considered valid.
func ValidateAltNamesUnique(genres []data.GameGenre) (bool, []string) {
	var invalidGenres []string

	for _, genre := range genres {
		seen := make(map[string]bool)
		hasDuplicates := false

		for _, altName := range genre.AltNames {
			if seen[altName] {
				hasDuplicates = true

				break
			}

			seen[altName] = true
		}

		if hasDuplicates {
			invalidGenres = append(invalidGenres, genre.Name)
		}
	}

	if len(invalidGenres) == 0 {
		return true, nil
	}

	return false, invalidGenres
}

type GenreWithCollidingAltName struct {
	CollidingGenreName        string
	GenreWithCollidingAltName string
}

// ValidateGenreNameNoCollisionsWithAltNames checks if any genre name appears as an alternative name in any other genre.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if no collisions exist, false otherwise
//	[]GenreWithCollidingAltName: A slice containing details about each collision, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv", "quest"}},
//	}
//
//	valid, _ := ValidateGenreNameNoCollisionsWithAltNames(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv", "action"}},  // "action" appears as alt name
//	}
//
//	valid, collisions := ValidateGenreNameNoCollisionsWithAltNames(invalidGenres)
//	// returns false, [{CollidingGenreName: "action", GenreWithCollidingAltName: "adventure"}]
//
//	multipleCollisions := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"rpg", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"action", "quest"}},
//	    {Name: "rpg", AltNames: []string{"role-playing", "adventure"}},
//	}
//
//	valid, collisions := ValidateGenreNameNoCollisionsWithAltNames(multipleCollisions)
//	// returns false with multiple collision entries
//
// Note:
//
//	The function checks for exact string matches and is case-sensitive.
//	A genre name can appear in multiple collisions if it appears as an alternative name in multiple genres.
func ValidateGenreNameNoCollisionsWithAltNames(genres []data.GameGenre) (bool, []GenreWithCollidingAltName) {
	var collisions []GenreWithCollidingAltName

	for _, genre := range genres {
		for _, otherGenre := range genres {
			if slices.Contains(otherGenre.AltNames, genre.Name) {
				collisions = append(collisions, GenreWithCollidingAltName{
					CollidingGenreName:        genre.Name,
					GenreWithCollidingAltName: otherGenre.Name,
				})
			}
		}
	}

	if len(collisions) == 0 {
		return true, nil
	}

	return false, collisions
}

type AltNameCollision struct {
	AltName                   string
	CollidingGenreName        string
	GenreWithCollidingAltName string
}

// ValidateCollidingAltNames checks if any genres have the same alternative names.
//
// Parameters:
//
//	genres: A slice of data.GameGenre objects to validate
//
// Returns:
//
//	bool: true if no collisions exist, false otherwise
//	[]AltNameCollision: A slice containing details about each collision, or nil if none found
//
// Examples:
//
//	validGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv", "quest"}},
//	}
//
//	valid, _ := ValidateCollidingAltNames(validGenres)  // returns true, nil
//
//	invalidGenres := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "fighting"}},
//	    {Name: "adventure", AltNames: []string{"adv", "fighting"}},  // "fighting" appears in both
//	}
//
//	valid, collisions := ValidateCollidingAltNames(invalidGenres)
//	// returns false, [{AltName: "fighting", CollidingGenreName: "action", GenreWithCollidingAltName: "adventure"}]
//
//	multipleCollisions := []data.GameGenre{
//	    {Name: "action", AltNames: []string{"act", "game"}},
//	    {Name: "adventure", AltNames: []string{"act", "quest"}},
//	    {Name: "rpg", AltNames: []string{"game", "role-playing"}},
//	}
//
//	valid, collisions := ValidateCollidingAltNames(multipleCollisions)
//	// returns false with multiple collision entries
//
// Note:
//
//	The function checks for exact string matches and is case-sensitive.
//	The function only reports collisions between different genres (not within the same genre).
func ValidateCollidingAltNames(genres []data.GameGenre) (bool, []AltNameCollision) {
	var collisions []AltNameCollision

	for _, genre := range genres {
		for _, altName := range genre.AltNames {
			for _, otherGenre := range genres {
				if genre.Name == otherGenre.Name {
					continue
				}

				if slices.Contains(otherGenre.AltNames, altName) {
					collisions = append(collisions, AltNameCollision{
						AltName:                   altName,
						CollidingGenreName:        genre.Name,
						GenreWithCollidingAltName: otherGenre.Name,
					})
				}
			}
		}
	}

	if len(collisions) == 0 {
		return true, nil
	}

	return false, collisions
}
