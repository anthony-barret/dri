package downloadredditimages

import "strings"

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func (rr RedditResponse) GetLinks() []string {
	var links []string
	for _, child := range rr.Data.Children {
		if strings.HasSuffix(child.Data.URL, ".jpg") || strings.HasSuffix(child.Data.URL, ".png") {
			links = append(links, string(child.Data.URL))
		}
	}
	return links
}
