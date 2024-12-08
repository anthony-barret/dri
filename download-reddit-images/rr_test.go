package downloadredditimages

import (
	"reflect"
	"testing"
)

func TestGetLinks(t *testing.T) {
	config := Config{
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
	redditResponse := RedditResponse{
		Data: struct {
			Children []struct {
				Data struct {
					URL    string `json:"url"`
					Over18 bool   `json:"over_18"`
				} `json:"data"`
			} `json:"children"`
		}{
			Children: []struct {
				Data struct {
					URL    string `json:"url"`
					Over18 bool   `json:"over_18"`
				} `json:"data"`
			}{
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_jpg_nover.jpg",
						Over18: false,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_jpeg_nover.jpeg",
						Over18: false,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_png_nover.png",
						Over18: false,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_aaa_nover.aaa",
						Over18: false,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_jpg_over.jpg",
						Over18: true,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_jpeg_over.jpeg",
						Over18: true,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_png_over.png",
						Over18: true,
					},
				},
				{
					Data: struct {
						URL    string `json:"url"`
						Over18 bool   `json:"over_18"`
					}{
						URL:    "https://redd.it/test_aaa_over.aaa",
						Over18: true,
					},
				},
			},
		},
	}
	t.Run("NSFW Allowed", func(t *testing.T) {
		config.Config.Over18 = true
		want := []string{
			"https://redd.it/test_jpg_nover.jpg",
			"https://redd.it/test_jpeg_nover.jpeg",
			"https://redd.it/test_png_nover.png",
			"https://redd.it/test_jpg_over.jpg",
			"https://redd.it/test_jpeg_over.jpeg",
			"https://redd.it/test_png_over.png",
		}
		got := redditResponse.GetLinks(config)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Slices are not equal. Expected %v, got %v", want, got)
		}
	})
	t.Run("NSFW Not Allowed", func(t *testing.T) {
		config.Config.Over18 = false
		want := []string{
			"https://redd.it/test_jpg_nover.jpg",
			"https://redd.it/test_jpeg_nover.jpeg",
			"https://redd.it/test_png_nover.png",
		}
		got := redditResponse.GetLinks(config)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Slices are not equal. Expected %v, got %v", want, got)
		}
	})
}
