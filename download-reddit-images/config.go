package downloadredditimages

import (
	"fmt"
	"os"
	"strconv"

	yaml "gopkg.in/yaml.v3"
)

type SortBy string

func (sb *SortBy) UnmarshalYAML(value *yaml.Node) error {
	var allowedValues = []string{"hot", "new", "top", "rising"}
	var rawValue string
	if err := value.Decode(&rawValue); err != nil {
		return err
	}
	for _, allowed := range allowedValues {
		if rawValue == allowed {
			*sb = SortBy(rawValue)
			return nil
		}
	}
	return fmt.Errorf("invalid value %q. The allowed values are: %v", rawValue, allowedValues)
}

type Config struct {
	Config struct {
		Directory string `yaml:"directory"`
		Limit     int    `yaml:"limit"`
		SortBy    SortBy `yaml:"sort_by"`
		Over18    bool   `yaml:"over_18"`
	} `yaml:"config"`
	SubReddits []string `yaml:"subreddits"`
}

func ParseConfig(configFile string) (Config, error) {
	var config Config
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	links := make([]string, 0)
	for _, sub := range config.SubReddits {
		link := "https://www.reddit.com/r/" + sub + "/" + string(config.Config.SortBy) + ".json?limit=" + strconv.Itoa(config.Config.Limit)
		links = append(links, link)
	}
	config.SubReddits = links
	return config, nil
}
