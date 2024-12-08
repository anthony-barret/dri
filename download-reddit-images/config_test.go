package downloadredditimages

import (
	"os"
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	t.Run("Valid configuration", func(t *testing.T) {
		content := `
config:
  directory: img
  limit: 10
  sort_by: hot
  over_18: no

subreddits:
  - memes
  - programminghumour
`
		file, err := os.CreateTemp("", "valid_config.yaml")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name())

		_, err = file.Write([]byte(content))
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
		file.Close()

		config, err := ParseConfig(file.Name())
		if err != nil {
			t.Fatalf("ParseConfig returned an error: %v", err)
		}

		want := Config{
			Config: struct {
				Directory string `yaml:"directory"`
				Limit     int    `yaml:"limit"`
				SortBy    SortBy `yaml:"sort_by"`
				Over18    bool   `yaml:"over_18"`
			}{
				Directory: "img",
				Limit:     10,
				SortBy:    "hot",
				Over18:    false,
			},
			SubReddits: []string{
				"https://www.reddit.com/r/memes/hot.json?limit=10",
				"https://www.reddit.com/r/programminghumour/hot.json?limit=10",
			},
		}

		if !reflect.DeepEqual(want, config) {
			t.Errorf("Config are not equal. Expected %v, got %v", want, config)
		}
	})

	t.Run("Nonexistent file", func(t *testing.T) {
		config, err := ParseConfig("nonexistent_file.yaml")
		if !reflect.DeepEqual(Config{}, config) {
			t.Errorf("Expected config to be empty but got %v", config)
		}
		if err == nil {
			t.Errorf("Expected error for nonexistent file, got nil")
		}
	})

	t.Run("Malformed file", func(t *testing.T) {
		content := `abcdefghijklmnopqrstuvwxyz`
		file, err := os.CreateTemp("", "malformed_config.yaml")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name())

		_, err = file.Write([]byte(content))
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
		file.Close()

		config, err := ParseConfig(file.Name())
		if !reflect.DeepEqual(Config{}, config) {
			t.Errorf("Expected config to be empty but got %v", config)
		}
		if err == nil {
			t.Fatalf("Expected error for nonexistent file, got nil")
		}
	})

	t.Run("Empty file", func(t *testing.T) {
		content := ``
		file, err := os.CreateTemp("", "empty_config.yaml")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name())

		_, err = file.Write([]byte(content))
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
		file.Close()

		want := Config{
			Config: struct {
				Directory string `yaml:"directory"`
				Limit     int    `yaml:"limit"`
				SortBy    SortBy `yaml:"sort_by"`
				Over18    bool   `yaml:"over_18"`
			}{
				Directory: "",
				Limit:     0,
				SortBy:    "",
				Over18:    false,
			},
			SubReddits: []string{},
		}

		config, err := ParseConfig(file.Name())
		if !reflect.DeepEqual(want, config) {
			t.Errorf("Config are not equal. Expected %v, got %v", want, config)
		}
		if err != nil {
			t.Fatalf("ParseConfig returned an error: %v", err)
		}
	})

	t.Run("Invalid sort_by", func(t *testing.T) {
		content := `
config:
  directory: img
  limit: 10
  sort_by: aaa
  over_18: no

subreddits:
  - memes
  - programminghumour
`
		file, err := os.CreateTemp("", "invalid_config_sort_by.yaml")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name())

		_, err = file.Write([]byte(content))
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
		file.Close()

		config, err := ParseConfig(file.Name())
		if !reflect.DeepEqual(Config{}, config) {
			t.Errorf("Expected config to be empty but got %v", config)
		}
		if err == nil {
			t.Fatalf("Expected error for invalid sort_by, got nil")
		}
	})
}
