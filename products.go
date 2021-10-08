package picnic

import (
	"encoding/json"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Image_id string `json:"image_id"`
}

type SearchResponse struct {
	Picnictype string    `json:"type"`
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Items      []Product `json:"items"`
}

func (cl Client) SearchProducts(searchQuery string) ([]Product, error) {
	endpoint := "/search?search_term=" + searchQuery
	productsRaw, err := cl.Get(endpoint)
	if err != nil {
		return nil, err
	}

	return toProducts(productsRaw)
}

func toProducts(productsRaw []byte) ([]Product, error) {
	var searchResponse []SearchResponse

	err2 := json.Unmarshal(productsRaw, &searchResponse)
	if err2 != nil {
		return nil, err2
	}
	return searchResponse[0].Items, nil
}
