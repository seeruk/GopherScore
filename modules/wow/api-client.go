package wow

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const urlTemplate = "https://%s.api.battle.net/wow/%s"

type ApiClient struct {
	ApiKey string
	Locale string
}

// NewApiClient returns a new World of Warcraft API client.
func NewApiClient(apiKey string) *ApiClient {
	return &ApiClient{
		ApiKey: apiKey,
		Locale: "en_GB",
	}
}

// SetLocale sets the locale for future requests via this API client.
func (c *ApiClient) SetLocale(locale string) *ApiClient {
	c.Locale = locale
	return c
}

// doRequest provides a base for making requests to the World of Warcraft API.
func doRequest(region string, uri string, model interface{}) error {
	resp, err := http.Get(fmt.Sprintf(urlTemplate, region, uri))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&model)
}
