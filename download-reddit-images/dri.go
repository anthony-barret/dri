package downloadredditimages

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseConfig(configFile string) []string {
	// Read config file
	fmt.Println(configFile)
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Cannot read the configuration file")
	}
	lines := strings.Split(string(data), "\n")
	// Extract subreddits
	links := make([]string, 0)
	for _, line := range lines {
		if trimmedLine := strings.TrimSpace(line); trimmedLine != "" {
			link := "https://www.reddit.com/r/" + trimmedLine + "/new.json?limit=10"
			links = append(links, link)
		}

	}
	// Return list of subreddits
	return links
}
