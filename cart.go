package picnic

import (
	"encoding/json"
)

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

func (cl Client) AddToCart(productId string, amount int) error {
	payload, err := json.Marshal(map[string]interface{}{
		"product_id": productId,
		"count":      amount,
	})
	if err != nil {
		return err
	}
	_, err2 := cl.post("/cart/add_product", payload)
	if err2 != nil {
		return err2
	}
	return nil
}

func (cl Client) RemoveFromCart(productId string, amount int) error {
	payload, err := json.Marshal(map[string]interface{}{
		"product_id": productId,
		"count":      amount,
	})
	if err != nil {
		return err
	}
	_, err2 := cl.post("/cart/remove_product", payload)
	if err2 != nil {
		return err2
	}
	return nil
}

func (cl Client) ClearCart() error {
	_, err := cl.post("/cart/clear", nil)
	if err != nil {
		return err
	}
	return nil
}

func toCart(cartRaw []byte) (*Cart, error) {
	var response getCartResponse

	err := json.Unmarshal(cartRaw, &response)
	if err != nil {
		return nil, err
	}
	for i, item := range response.Items[0].Items {
		response.Items[0].Items[i].Amount = item.Decorators[0].Quantity
	}
	return &response.Items[0], nil
}
