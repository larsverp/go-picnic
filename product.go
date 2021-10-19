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

type searchProductsResponse struct {
	Picnictype string    `json:"type"`
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Items      []Product `json:"items"`
}

type getProductByIdResponse struct {
	ProductDetails Product `json:"product_details"`
}

func (product Product) Add(cl Client, amount int) error {
	err := cl.AddToCart(product.Id, amount)
	if err != nil {
		return err
	}
	return nil
}
func (product Product) Remove(cl Client, amount int) error {
	err := cl.RemoveFromCart(product.Id, amount)
	if err != nil {
		return err
	}
	return nil
}

func (cl Client) SearchProducts(searchQuery string) ([]Product, error) {
	endpoint := "/search?search_term=" + searchQuery
	productsRaw, err := cl.get(endpoint)
	if err != nil {
		return nil, err
	}

	return toProducts(productsRaw)
}

func (cl Client) GetProductById(productId string) (*Product, error) {
	endpoint := "/product/" + productId
	productRaw, err := cl.get(endpoint)
	if err != nil {
		return nil, err
	}

	return toProduct(productRaw)
}

func toProduct(productRaw []byte) (*Product, error) {
	var response getProductByIdResponse

	err := json.Unmarshal(productRaw, &response)
	if err != nil {
		return nil, err
	}
	return &response.ProductDetails, nil
}

func toProducts(productsRaw []byte) ([]Product, error) {
	var response []searchProductsResponse

	err := json.Unmarshal(productsRaw, &response)
	if err != nil {
		return nil, err
	}
	return response[0].Items, nil
}
