package downloadredditimages

import "strings"

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				URL    string `json:"url"`
				Over18 bool   `json:"over_18"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func (rr RedditResponse) GetLinks(config Config) []string {
	var links []string
	for _, child := range rr.Data.Children {
		if !config.Config.Over18 && child.Data.Over18 {
			continue
		}
		if strings.HasSuffix(child.Data.URL, ".jpg") || strings.HasSuffix(child.Data.URL, ".jpeg") || strings.HasSuffix(child.Data.URL, ".png") {
			links = append(links, string(child.Data.URL))
		}
	}
	return links
}
