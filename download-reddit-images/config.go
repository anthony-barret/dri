package downloadredditimages

import (
	"log"
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Config struct {
		Directory string `yaml:"directory"`
		Limit     int    `yaml:"limit"`
	} `yaml:"config"`
	SubReddits []string `yaml:"subreddits"`
}

func ParseConfig(configFile string) Config {
	var config Config
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Cannot read the configuration file", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln("Cannot parse the configuration file", err)
	}
	links := make([]string, 0)
	for _, sub := range config.SubReddits {
		link := "https://www.reddit.com/r/" + sub + "/new.json?limit=" + strconv.Itoa(config.Config.Limit)
		links = append(links, link)
	}
	config.SubReddits = links
	return config
}
