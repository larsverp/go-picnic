package picnic

import "encoding/json"

type Cart struct {
	Items []Product `json:"items"`
	Price float64   `json:"price"`
}

type getCartResponse struct {
	Items []Cart `json:"items"`
}

func (cl Client) GetCart() (*Cart, error) {
	cartRaw, err := cl.get("/cart")
	if err != nil {
		return nil, err
	}
	return toCart(cartRaw)

}

func toCart(cartRaw []byte) (*Cart, error) {
	var response getCartResponse

	err := json.Unmarshal(cartRaw, &response)
	if err != nil {
		return nil, err
	}
	return &response.Items[0], nil
}
