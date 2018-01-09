package dl

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"sync"
	"time"
)

const expireTime = time.Duration(5 * time.Minute)

type Cache struct {
	lock  sync.RWMutex
	owner string
	cache map[string]*Result
}

type Result struct {
	rule regexp.Regexp
	age  time.Time
	url  *string
}

func NewCache(c Config) Cache {
	var cache Cache
	cache.owner = c.Owner
	cache.cache = make(map[string]*Result)

	for repo, rules := range c.Repos {
		for _, rule := range rules {
			lookup := fmt.Sprint(repo, ".", rule.ID)
			reg, err := regexp.Compile(rule.Regex)
			if err == nil {
				cache.cache[lookup] = &Result{
					rule: *reg,
					age:  time.Now(),
				}
			} else {
				fmt.Println("Err regex: ", err)
			}
		}
	}

	return cache
}

func (c *Cache) Get(repo, id string) (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	lookup := fmt.Sprint(repo, ".", id)
	if res, ok := c.cache[lookup]; ok {
		if res.url == nil || time.Since(res.age) > expireTime {
			err := update(c, repo)
			if err != nil {
				return "", err
			}
		}

		if res.url != nil {
			return *res.url, nil
		}
	}

	return "", errors.New("not found: " + lookup)
}

func update(c *Cache, repo string) (error) {
	latest, err := Latest(c.owner, repo)
	if err != nil {
		return err
	}

	match := fmt.Sprint(repo, ".")
	for path, res := range c.cache {
		if !strings.HasPrefix(path, match) {
			continue
		}

		ass, err := latest.Asset(res.rule)
		if err != nil {
			fmt.Println("Err match asset: ", err)
			continue
		}

		res.url = &ass.URL
		res.age = time.Now()
	}

	return nil
}
