package download

import (
	"regexp"
	"time"
	"fmt"
)

type Cache struct {
	cache map[string]Result
}

type Result struct {
	rule regexp.Regexp
	age  time.Time
	url  *string
}

func NewCache(c Config) Cache {
	var cache Cache
	cache.cache = make(map[string]Result)
	for repo, rules := range c.Repos {
		for _, rule := range rules {
			lookup := fmt.Sprint(repo, rule.ID)
			reg, err := regexp.Compile(rule.Regex)
			if err == nil {
				cache.cache[lookup] = Result{
					rule: *reg,
					age: time.Now(),
				}
			}
		}
	}
	return cache
}
