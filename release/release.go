package release

import (
	"net/http"
	"fmt"
	"encoding/json"
	"regexp"
	"github.com/pkg/errors"
)

type Release struct {
	Name    string  `json:"name"`
	Version string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name  string `json:"name"`
	Type  string `json:"content_type"`
	Count int    `json:"download_count"`
	Size  int    `json:"size"`
	URL   string `json:"browser_download_url"`
}

func Latest(owner, repo string) (*Release, error) {
	var release Release
	release.Assets = []Asset{}
	r, e := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo))
	if e != nil {
		return nil, e
	}
	return &release, json.NewDecoder(r.Body).Decode(&release)
}

func (r *Release) Asset(matcher regexp.Regexp) (*Asset, error) {
	for _, a := range r.Assets {
		if matcher.MatchString(a.Name) {
			return &a, nil
		}
	}
	return nil, errors.New("asset not found")
}