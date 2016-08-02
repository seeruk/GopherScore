package wow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/SeerUK/GopherScore/modules/wow/models"
)

// ApiClient provides easy access to specific World of Warcraft API endpoints.
type ApiClient struct {
	// A Battle.net API key.
	ApiKey string
	// A locale string (e.g. en_GB).
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

// FindCharacter attempts to find a given character.
func (c *ApiClient) FindCharacter(region string, realm string, name string) (*models.Character, error) {
	character := &models.Character{}
	uri := fmt.Sprintf(
		"wow/character/%s/%s",
		realm,
		name,
	)

	query := make(map[string]string)
	query["fields"] = "titles,items,professions,progression"

	err := c.doRequest(region, uri, query, character)

	if err != nil {
		return character, err
	}

	return character, nil
}

// doRequest provides a base for making requests to the World of Warcraft API.
func (c *ApiClient) doRequest(region string, path string, requestQuery map[string]string, model interface{}) error {
	url := url.URL{}
	url.Scheme = "https"
	url.Host = fmt.Sprintf("%s.api.battle.net", strings.ToLower(region))
	url.Path = path

	query := url.Query()
	query.Set("locale", c.Locale)
	query.Set("apikey", c.ApiKey)

	for k := range requestQuery {
		query.Set(k, requestQuery[k])
	}

	url.RawQuery = query.Encode()

	resp, err := http.Get(url.String())

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// This could be improved!
	if resp.StatusCode != 200 {
		return errors.New("Unable to handle request.")
	}

	return json.NewDecoder(resp.Body).Decode(&model)
}
