package download

import (
	"encoding/json"
	"io/ioutil"
)

const filename = "config.json"

type Config struct {
	Repos map[string][]Rule
}

type Rule struct {
	ID    string `json:"id"`
	Regex string `json:"regex"`
}

func LoadConfig() (Config) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		config.Repos["launcher"] = []Rule{
			{ID: "windows", Regex: "^.*?\\.exe"},
			{ID: "other", Regex: "^.*?\\.jar"},
		}
		json.MarshalIndent(&config, "", "  ")
		return config
	}
	json.Unmarshal(data, &config)
	return config
}
