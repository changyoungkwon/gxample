package igntserver

import (
	"encoding/json"
	"net/http"

	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/models"
)

// IgntRepo represents repo
type IgntRepo struct {
	uri string
}

// NewIgntRepo generate igntserver in msa architecture
func NewIgntRepo() *IgntRepo {
	return &IgntRepo{
		uri: igserverURI,
	}
}

// List get name matching ignts from igserver
func (i *IgntRepo) list(name string) ([]*ignt, error) {
	// generate uri
	req, _ := http.NewRequest("GET", i.uri, nil)
	q := req.URL.Query()
	q.Add("searchWord", name)
	req.URL.RawQuery = q.Encode()

	// get response
	client := http.Client{
		Timeout: timeoutGetSingle,
	}
	res, err := client.Get(req.URL.String())
	if err != nil {
		logging.Errorf("error getting data from ignt server, %v", err)
		return nil, err
	}

	// parse response
	var igntRes igntListResponse
	err = json.NewDecoder(res.Body).Decode(&igntRes)
	if err != nil {
		logging.Errorf("error parsing response from ignt server, %v", err)
		return nil, err
	}
	return igntRes.Embedded.Ingredients, nil
}

// Get get ignt from igserver
func (i *IgntRepo) listAll() ([]*ignt, error) {
	// get response
	client := http.Client{
		Timeout: timeoutGetSingle,
	}
	res, err := client.Get(i.uri)
	if err != nil {
		logging.Errorf("error getting data from ignt server, %v", err)
		return nil, err
	}
	// parse response
	var igntRes igntListResponse
	err = json.NewDecoder(res.Body).Decode(&igntRes)
	if err != nil {
		logging.Errorf("error parsing response from ignt server, %v", err)
		return nil, err
	}
	return igntRes.Embedded.Ingredients, nil
}

func toModel(i ignt) *models.Ingredient {
	return &models.Ingredient{}
}
