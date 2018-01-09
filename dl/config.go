package dl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const filename = "config.json"

type Config struct {
	Owner string
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
		config.Owner = "Conquest-Reforged"
		config.Repos = make(map[string][]Rule)
		config.Repos = map[string][]Rule{
			"launcher": {
				{ID: "windows", Regex: "^.*?\\.exe"},
				{ID: "other", Regex: "^.*?\\.jar"},
			},
		}

		data, err := json.MarshalIndent(&config, "", "  ")
		if err == nil {
			err = ioutil.WriteFile(filename, data, os.ModePerm)
		}
		if err != nil {
			fmt.Println(err)
		}

		return config
	}
	json.Unmarshal(data, &config)
	return config
}
