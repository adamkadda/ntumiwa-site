package pagedata

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/adamkadda/ntumiwa-site/shared/cache"
)

/*
	I've decided to avoid using an interface for the Pages
	field as it introduced a lot of unncessary complexity.

	Instead, I'm leaving it as a wrapper that allows it to
	still serve its core functions:

	1. Access the caches' values
	2. Update the caches
	3. Check the caches' health
*/

type PageData interface {
	GetHomeData() (*HomeData, error)
}

type PageCache struct {
	apiURL string
	client *http.Client
	home   *cache.Cache[HomeData]
}

func NewPageCache(apiURL string, client *http.Client) (*PageCache, error) {
	pc := &PageCache{
		apiURL: apiURL,
		client: client,
		home:   &cache.Cache[HomeData]{},
	}

	homeData, err := pc.GetHomeData()
	if err != nil {
		return nil, err
	}
	pc.home.Set(homeData)

	return pc, nil
}

func (pc *PageCache) GetHomeData() (*HomeData, error) {
	resp, err := pc.client.Get(pc.apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &HomeData{}
	err = json.Unmarshal(body, data)

	return data, nil
}

// low priority feature
func (p *PageCache) checkHealth() error {

	/*
		Call Get() on each field/cache, check for nil.
		Fetch relevant data from API if Get() returns nil.
		Update data in field/cache with Set().

		Ideally a goroutine periodically calls checkHealth().
	*/

	return nil
}
