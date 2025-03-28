package validation

import (
	"content_validator/internal/data"
	"reflect"
	"sort"
	"testing"
)

func TestValidateNameNotEmpty(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name          string
		input         []data.GameGenre
		expectedValid bool
	}{
		{
			name:          "all valid names",
			input:         []data.GameGenre{{Name: "action", AltNames: nil}, {Name: "rpg", AltNames: nil}},
			expectedValid: true,
		},
		{
			name:          "single empty name",
			input:         []data.GameGenre{{Name: "", AltNames: nil}},
			expectedValid: false,
		},
		{
			name: "multiple empty names",
			input: []data.GameGenre{
				{Name: "", AltNames: nil},
				{Name: "action", AltNames: nil},
				{Name: "", AltNames: nil},
			},
			expectedValid: false,
		},
		{
			name:          "empty input slice",
			input:         []data.GameGenre{},
			expectedValid: true,
		},
		{
			name: "mixed valid and empty",
			input: []data.GameGenre{
				{Name: "strategy", AltNames: nil},
				{Name: "", AltNames: nil},
				{Name: "puzzle", AltNames: nil},
			},
			expectedValid: false,
		},
		{
			name:          "multiple whitespace",
			input:         []data.GameGenre{{Name: "   ", AltNames: nil}},
			expectedValid: false,
		},
		{
			name:          "unicode empty check",
			input:         []data.GameGenre{{Name: "", AltNames: nil}, {Name: "アクション", AltNames: nil}},
			expectedValid: false,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			valid := ValidateNameNotEmpty(test.input)

			if valid != test.expectedValid {
				runner.Errorf("Validity mismatch: got %v, want %v", valid, test.expectedValid)
			}
		})
	}
}

func TestValidateAltNamesNotEmpty(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantInvalid []string
	}{
		{
			name:        "empty genres slice",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "genre with empty AltNames slice",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "single invalid alt name (empty string)",
			genres: []data.GameGenre{
				{Name: "Action", AltNames: []string{""}},
			},
			wantValid:   false,
			wantInvalid: []string{"Action"},
		},
		{
			name: "single invalid alt name (whitespace)",
			genres: []data.GameGenre{
				{Name: "Adventure", AltNames: []string{"   "}},
			},
			wantValid:   false,
			wantInvalid: []string{"Adventure"},
		},
		{
			name: "multiple invalid alt names in one genre",
			genres: []data.GameGenre{
				{Name: "Strategy", AltNames: []string{"", "  ", "   "}},
			},
			wantValid:   false,
			wantInvalid: []string{"Strategy"},
		},
		{
			name: "mixed valid and invalid alt names",
			genres: []data.GameGenre{
				{Name: "Puzzle", AltNames: []string{"BrainTeaser", ""}},
			},
			wantValid:   false,
			wantInvalid: []string{"Puzzle"},
		},
		{
			name: "multiple genres with invalid entries",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"", "RolePlaying"}},
				{Name: "FPS", AltNames: []string{"Shooter", "   "}},
			},
			wantValid:   false,
			wantInvalid: []string{"RPG", "FPS"},
		},
		{
			name: "genre with empty name and invalid alt",
			genres: []data.GameGenre{
				{Name: "", AltNames: []string{""}},
			},
			wantValid:   false,
			wantInvalid: []string{""},
		},
		{
			name: "all valid alt names",
			genres: []data.GameGenre{
				{Name: "Racing", AltNames: []string{"Driving", "Cars"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotInvalid := ValidateAltNamesNotEmpty(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("got valid %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotInvalid, test.wantInvalid) {
				runner.Errorf("got invalid %v, want %v", gotInvalid, test.wantInvalid)
			}
		})
	}
}

func TestValidateNameTrimmed(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantInvalid []string
	}{
		{
			name:        "empty slice",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "valid names",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
				{Name: "Action-Adventure", AltNames: nil},
				{Name: "Real-Time Strategy", AltNames: nil},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "leading whitespace",
			genres: []data.GameGenre{
				{Name: " RPG", AltNames: nil},
				{Name: "  Simulation", AltNames: nil},
			},
			wantValid:   false,
			wantInvalid: []string{" RPG", "  Simulation"},
		},
		{
			name: "trailing whitespace",
			genres: []data.GameGenre{
				{Name: "FPS ", AltNames: nil},
				{Name: "Racing  ", AltNames: nil},
			},
			wantValid:   false,
			wantInvalid: []string{"FPS ", "Racing  "},
		},
		{
			name: "mixed whitespace types",
			genres: []data.GameGenre{
				{Name: "\tTabbed", AltNames: nil},
				{Name: "Newline\n", AltNames: nil},
				{Name: " \t\nAll Three\n\t ", AltNames: nil},
			},
			wantValid:   false,
			wantInvalid: []string{"\tTabbed", "Newline\n", " \t\nAll Three\n\t "},
		},
		{
			name: "whitespace-only names",
			genres: []data.GameGenre{
				{Name: " ", AltNames: nil},
				{Name: "\t", AltNames: nil},
				{Name: "\n", AltNames: nil},
			},
			wantValid:   false,
			wantInvalid: []string{" ", "\t", "\n"},
		},
		{
			name: "mixed valid and invalid",
			genres: []data.GameGenre{
				{Name: "Valid", AltNames: nil},
				{Name: " Invalid", AltNames: nil},
				{Name: "AlsoValid", AltNames: nil},
				{Name: "Invalid ", AltNames: nil},
			},
			wantValid:   false,
			wantInvalid: []string{" Invalid", "Invalid "},
		},
		{
			name: "empty name",
			genres: []data.GameGenre{
				{Name: "", AltNames: nil},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotInvalid := ValidateNameTrimmed(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("valid mismatch: got %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotInvalid, test.wantInvalid) {
				runner.Errorf("invalid names mismatch: got %v, want %v", gotInvalid, test.wantInvalid)
			}
		})
	}
}

func TestValidateAltNamesTrimmed(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantInvalid []string
	}{
		{
			name:        "empty input",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "all valid altNames",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"RolePlaying", "CRPG"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "leading/trailing whitespace",
			genres: []data.GameGenre{
				{Name: "Action", AltNames: []string{" Action ", "Arcade"}},
			},
			wantValid:   false,
			wantInvalid: []string{"Action"},
		},
		{
			name: "multiple invalid in one genre",
			genres: []data.GameGenre{
				{Name: "FPS", AltNames: []string{" FPS ", "Shooter ", "Gun "}},
			},
			wantValid:   false,
			wantInvalid: []string{"FPS"},
		},
		{
			name: "mixed valid/invalid across genres",
			genres: []data.GameGenre{
				{Name: "RTS", AltNames: []string{"RealTime "}},
				{Name: "MMO", AltNames: []string{"MassMultiplayer"}},
				{Name: "RPG", AltNames: []string{" RPG", "RolePlay"}},
			},
			wantValid:   false,
			wantInvalid: []string{"RTS", "RPG"},
		},
		{
			name: "special whitespace characters",
			genres: []data.GameGenre{
				{Name: "Tab", AltNames: []string{"\tIndented"}},
				{Name: "Newline", AltNames: []string{"Line\n"}},
			},
			wantValid:   false,
			wantInvalid: []string{"Tab", "Newline"},
		},
		{
			name: "whitespace-only altName",
			genres: []data.GameGenre{
				{Name: "Empty", AltNames: []string{"  ", "\t\n"}},
			},
			wantValid:   false,
			wantInvalid: []string{"Empty"},
		},
		{
			name: "empty genre name with invalid alt",
			genres: []data.GameGenre{
				{Name: "", AltNames: []string{" Invalid "}},
			},
			wantValid:   false,
			wantInvalid: []string{""},
		},
		{
			name: "multiple validation errors",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{" Valid", "AlsoValid"}},
				{Name: "B", AltNames: []string{"Perfect"}},
				{Name: "C", AltNames: []string{" Problem ", "Issue"}},
			},
			wantValid:   false,
			wantInvalid: []string{"A", "C"},
		},
		{
			name: "valid empty altName",
			genres: []data.GameGenre{
				{Name: "Strategy", AltNames: []string{""}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotInvalid := ValidateAltNamesTrimmed(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("valid mismatch: got %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotInvalid, test.wantInvalid) {
				runner.Errorf("invalid list mismatch: got %v, want %v", gotInvalid, test.wantInvalid)
			}
		})
	}
}

func TestValidateNameCase(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name          string
		input         []data.GameGenre
		expectedValid bool
		expectedNames []string
	}{
		{
			name:          "all lowercase",
			input:         []data.GameGenre{{Name: "action", AltNames: nil}, {Name: "rpg", AltNames: nil}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "single uppercase",
			input:         []data.GameGenre{{Name: "Action", AltNames: nil}},
			expectedValid: false,
			expectedNames: []string{"Action"},
		},
		{
			name: "mixed casing",
			input: []data.GameGenre{
				{Name: "action", AltNames: nil},
				{Name: "RPG", AltNames: nil},
				{Name: "FPS", AltNames: nil},
			},
			expectedValid: false,
			expectedNames: []string{"RPG", "FPS"},
		},
		{
			name:          "empty input",
			input:         []data.GameGenre{},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "unicode characters",
			input:         []data.GameGenre{{Name: "アクション", AltNames: nil}, {Name: "アクション", AltNames: nil}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "numbers and symbols",
			input:         []data.GameGenre{{Name: "game-2", AltNames: nil}, {Name: "mod!", AltNames: nil}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "case in middle",
			input:         []data.GameGenre{{Name: "actionGame", AltNames: nil}},
			expectedValid: false,
			expectedNames: []string{"actionGame"},
		},
		{
			name:          "empty name",
			input:         []data.GameGenre{{Name: "", AltNames: nil}},
			expectedValid: true,
			expectedNames: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			valid, names := ValidateNameCase(test.input)

			if valid != test.expectedValid {
				runner.Errorf("valid mismatch: got %v, want %v", valid, test.expectedValid)
			}

			if len(names) != len(test.expectedNames) {
				runner.Fatalf("names length mismatch: got %d, want %d", len(names), len(test.expectedNames))
			}

			for nameIndex := range names {
				if names[nameIndex] != test.expectedNames[nameIndex] {
					runner.Errorf("name mismatch at index %d: got %q, want %q", nameIndex, names[nameIndex],
						test.expectedNames[nameIndex])
				}
			}
		})
	}
}

func TestValidateAltNamesCase(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name          string
		input         []data.GameGenre
		expectedValid bool
		expectedNames []string
	}{
		{
			name:          "all lowercase altnames",
			input:         []data.GameGenre{{Name: "", AltNames: []string{"action", "rpg"}}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name: "mixed case in multiple genres",
			input: []data.GameGenre{
				{Name: "", AltNames: []string{"Action", "platform"}},
				{Name: "", AltNames: []string{"RPG", "Strategy"}},
			},
			expectedValid: false,
			expectedNames: []string{"Action", "RPG", "Strategy"},
		},
		{
			name:          "case in middle of word",
			input:         []data.GameGenre{{Name: "", AltNames: []string{"actionGame"}}},
			expectedValid: false,
			expectedNames: []string{"actionGame"},
		},
		{
			name:          "empty altnames list",
			input:         []data.GameGenre{{Name: "", AltNames: []string{}}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "empty string altname",
			input:         []data.GameGenre{{Name: "", AltNames: []string{""}}},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name: "unicode characters",
			input: []data.GameGenre{
				{Name: "", AltNames: []string{"ÄCTION", "ßpecial"}},
			},
			expectedValid: false,
			expectedNames: []string{"ÄCTION"},
		},
		{
			name: "special characters and numbers",
			input: []data.GameGenre{
				{Name: "", AltNames: []string{"mod!", "game2"}},
			},
			expectedValid: true,
			expectedNames: nil,
		},
		{
			name:          "multiple duplicates",
			input:         []data.GameGenre{{Name: "", AltNames: []string{"Action", "Action"}}},
			expectedValid: false,
			expectedNames: []string{"Action", "Action"},
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			valid, names := ValidateAltNamesCase(test.input)

			if valid != test.expectedValid {
				runner.Errorf("validity mismatch: got %v, want %v", valid, test.expectedValid)
			}

			if !reflect.DeepEqual(names, test.expectedNames) {
				runner.Errorf("invalid names mismatch:\ngot:  %v\nwant: %v", names, test.expectedNames)
			}
		})
	}
}

func TestValidateNameUnique(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name           string
		gameGenres     []data.GameGenre
		wantValid      bool
		wantDuplicates []string
	}{
		{
			name:           "empty slice",
			gameGenres:     []data.GameGenre{},
			wantValid:      true,
			wantDuplicates: nil,
		},
		{
			name: "single genre",
			gameGenres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
			},
			wantValid:      true,
			wantDuplicates: nil,
		},
		{
			name: "two genres with same name",
			gameGenres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
				{Name: "RPG", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{"RPG"},
		},
		{
			name: "three genres with same name",
			gameGenres: []data.GameGenre{
				{Name: "Action", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Action", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{"Action"},
		},
		{
			name: "case sensitive names",
			gameGenres: []data.GameGenre{
				{Name: "rpg", AltNames: nil},
				{Name: "RPG", AltNames: nil},
			},
			wantValid:      true,
			wantDuplicates: nil,
		},
		{
			name: "multiple duplicate groups",
			gameGenres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
				{Name: "RPG", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Simulation", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{"Action", "RPG"},
		},
		{
			name: "mixed duplicates and unique",
			gameGenres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
				{Name: "RPG", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Strategy", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{"RPG"},
		},
		{
			name: "empty name duplicates",
			gameGenres: []data.GameGenre{
				{Name: "", AltNames: nil},
				{Name: "", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{""},
		},
		{
			name: "whitespace names",
			gameGenres: []data.GameGenre{
				{Name: " RPG ", AltNames: nil},
				{Name: "RPG", AltNames: nil},
			},
			wantValid:      true,
			wantDuplicates: nil,
		},
		{
			name: "multiple duplicates with non-duplicates",
			gameGenres: []data.GameGenre{
				{Name: "RPG", AltNames: nil},
				{Name: "RPG", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Action", AltNames: nil},
				{Name: "Simulation", AltNames: nil},
			},
			wantValid:      false,
			wantDuplicates: []string{"Action", "RPG"},
		},
		{
			name: "single empty name",
			gameGenres: []data.GameGenre{
				{Name: "", AltNames: nil},
			},
			wantValid:      true,
			wantDuplicates: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotDuplicates := ValidateNameUnique(test.gameGenres)

			if gotValid != test.wantValid {
				runner.Errorf("got valid %v, want %v", gotValid, test.wantValid)
			}

			sort.Strings(gotDuplicates)

			sortedWant := append([]string(nil), test.wantDuplicates...)

			sort.Strings(sortedWant)

			if !reflect.DeepEqual(gotDuplicates, sortedWant) {
				runner.Errorf("got duplicates %v, want %v", gotDuplicates, test.wantDuplicates)
			}
		})
	}
}

func TestValidateAltNamesUnique(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantInvalid []string
	}{
		{
			name:        "empty input",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "no duplicates",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"CRPG", "RolePlaying"}},
				{Name: "FPS", AltNames: []string{"Shooter", "FPS"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "single duplicate in one genre",
			genres: []data.GameGenre{
				{Name: "RTS", AltNames: []string{"Strategy", "Strategy", "RTS"}},
			},
			wantValid:   false,
			wantInvalid: []string{"RTS"},
		},
		{
			name: "case sensitivity no duplicates",
			genres: []data.GameGenre{
				{Name: "Action", AltNames: []string{"action", "ACTION"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "whitespace differences no duplicates",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"RPG ", " RPG", "RPG"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "multiple duplicates in one genre",
			genres: []data.GameGenre{
				{Name: "GenreA", AltNames: []string{"A", "A", "B", "B"}},
			},
			wantValid:   false,
			wantInvalid: []string{"GenreA"},
		},
		{
			name: "multiple genres with duplicates",
			genres: []data.GameGenre{
				{Name: "Genre1", AltNames: []string{"X", "X"}},
				{Name: "Genre2", AltNames: []string{"Y", "Y", "Z"}},
			},
			wantValid:   false,
			wantInvalid: []string{"Genre1", "Genre2"},
		},
		{
			name: "empty altNames list",
			genres: []data.GameGenre{
				{Name: "Empty", AltNames: []string{}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "single altName",
			genres: []data.GameGenre{
				{Name: "Single", AltNames: []string{"Only"}},
			},
			wantValid:   true,
			wantInvalid: nil,
		},
		{
			name: "mixed valid and invalid genres",
			genres: []data.GameGenre{
				{Name: "Valid", AltNames: []string{"A", "B"}},
				{Name: "Invalid", AltNames: []string{"C", "C"}},
				{Name: "Valid2", AltNames: []string{"D", "E"}},
			},
			wantValid:   false,
			wantInvalid: []string{"Invalid"},
		},
		{
			name: "duplicates with empty genre name",
			genres: []data.GameGenre{
				{Name: "", AltNames: []string{"X", "X"}},
			},
			wantValid:   false,
			wantInvalid: []string{""},
		},
		{
			name: "whitespace only duplicates",
			genres: []data.GameGenre{
				{Name: "Space", AltNames: []string{"  ", "  ", " "}},
			},
			wantValid:   false,
			wantInvalid: []string{"Space"},
		},
		{
			name: "all altNames duplicated",
			genres: []data.GameGenre{
				{Name: "AllDuplicates", AltNames: []string{"A", "A", "A"}},
			},
			wantValid:   false,
			wantInvalid: []string{"AllDuplicates"},
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotInvalid := ValidateAltNamesUnique(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("validity mismatch: got %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotInvalid, test.wantInvalid) {
				runner.Errorf("invalid list mismatch: got %v, want %v", gotInvalid, test.wantInvalid)
			}
		})
	}
}

func TestValidateGenreNameNoCollisionsWithAltNames(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantCollide []GenreWithCollidingAltName
	}{
		{
			name:        "empty input",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "no collisions",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"CRPG"}},
				{Name: "FPS", AltNames: []string{"Shooter"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "self collision",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"RPG"}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"RPG", "RPG"},
			},
		},
		{
			name: "cross-genre collision",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"CRPG"}},
				{Name: "CRPG", AltNames: []string{}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"CRPG", "RPG"},
			},
		},
		{
			name: "multiple collisions",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"B"}},
				{Name: "B", AltNames: []string{"C"}},
				{Name: "C", AltNames: []string{"A"}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"A", "C"},
				{"B", "A"},
				{"C", "B"},
			},
		},
		{
			name: "case sensitivity",
			genres: []data.GameGenre{
				{Name: "Rpg", AltNames: []string{"RPG"}},
				{Name: "rpg", AltNames: []string{"Rpg"}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"Rpg", "rpg"},
			},
		},
		{
			name: "whitespace differences",
			genres: []data.GameGenre{
				{Name: "Action", AltNames: []string{" Action "}},
				{Name: "Action ", AltNames: []string{"Action"}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"Action", "Action "},
			},
		},
		{
			name: "empty name collision",
			genres: []data.GameGenre{
				{Name: "", AltNames: []string{""}},
				{Name: "Empty", AltNames: []string{""}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"", ""},
				{"", "Empty"},
			},
		},
		{
			name: "multiple alt name sources",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"X"}},
				{Name: "B", AltNames: []string{"X"}},
				{Name: "X", AltNames: []string{}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"X", "A"},
				{"X", "B"},
			},
		},
		{
			name: "complex collision chain",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"B", "C"}},
				{Name: "B", AltNames: []string{"D"}},
				{Name: "D", AltNames: []string{"A"}},
			},
			wantValid: false,
			wantCollide: []GenreWithCollidingAltName{
				{"A", "D"},
				{"B", "A"},
				{"D", "B"},
			},
		},
		{
			name: "duplicate alt names in same genre",
			genres: []data.GameGenre{
				{Name: "RPG", AltNames: []string{"CRPG", "CRPG"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotCollide := ValidateGenreNameNoCollisionsWithAltNames(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("Validity mismatch: got %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotCollide, test.wantCollide) {
				runner.Errorf("Collisions mismatch:\nGot: %+v\nWant: %+v", gotCollide, test.wantCollide)
			}
		})
	}
}

func TestValidateCollidingAltNames(testRunner *testing.T) {
	testRunner.Parallel()

	tests := []struct {
		name        string
		genres      []data.GameGenre
		wantValid   bool
		wantCollide []AltNameCollision
	}{
		{
			name:        "empty input",
			genres:      []data.GameGenre{},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "no collisions",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"a1"}},
				{Name: "B", AltNames: []string{"b1"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "self collision with duplicate alt names",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"x", "x"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "cross-genre collision",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"x"}},
				{Name: "B", AltNames: []string{"x"}},
			},
			wantValid: false,
			wantCollide: []AltNameCollision{
				{"x", "A", "B"},
				{"x", "B", "A"},
			},
		},
		{
			name: "case sensitivity",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"RPG"}},
				{Name: "B", AltNames: []string{"rpg"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "whitespace differences",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{" RPG "}},
				{Name: "B", AltNames: []string{"RPG"}},
			},
			wantValid:   true,
			wantCollide: nil,
		},
		{
			name: "empty alt names",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{""}},
				{Name: "B", AltNames: []string{""}},
			},
			wantValid: false,
			wantCollide: []AltNameCollision{
				{"", "A", "B"},
				{"", "B", "A"},
			},
		},
		{
			name: "multiple genre collision",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"x"}},
				{Name: "B", AltNames: []string{"x"}},
				{Name: "C", AltNames: []string{"x"}},
			},
			wantValid: false,
			wantCollide: []AltNameCollision{
				{"x", "A", "B"},
				{"x", "A", "C"},
				{"x", "B", "A"},
				{"x", "B", "C"},
				{"x", "C", "A"},
				{"x", "C", "B"},
			},
		},
		{
			name: "mixed collisions",
			genres: []data.GameGenre{
				{Name: "A", AltNames: []string{"x", "y"}},
				{Name: "B", AltNames: []string{"x", "z"}},
			},
			wantValid: false,
			wantCollide: []AltNameCollision{
				{"x", "A", "B"},
				{"x", "B", "A"},
			},
		},
	}

	for _, test := range tests {
		testRunner.Run(test.name, func(runner *testing.T) {
			runner.Parallel()

			gotValid, gotCollide := ValidateCollidingAltNames(test.genres)

			if gotValid != test.wantValid {
				runner.Errorf("Validity mismatch: got %v, want %v", gotValid, test.wantValid)
			}

			if !reflect.DeepEqual(gotCollide, test.wantCollide) {
				runner.Errorf("Collision mismatch:\nGot: %+v\nWant: %+v", gotCollide, test.wantCollide)
			}
		})
	}
}
