package main

import (
	"fmt"

	"github.com/larsverp/go-picnic"
)

func main() {

	//////////////////////////////////////////////////////////////////////////////////////////////////////////
	// The _ symbol means "ignore this value" the second value in all client functions is a potention error.
	// In this example we will ignore them, but in production you should catch and handle them.
	//////////////////////////////////////////////////////////////////////////////////////////////////////////

	client, _ := picnic.NewClient(picnic.NewUser("username", "password"))
	// From now on we can use client to call all our functions. The calls will be authenticated from the client we just made

	// For example this will print the current cart from the client
	cart, _ := client.GetCart()
	fmt.Println(cart)

	// Let's search for bananas and print their Picnic ID
	bananas, _ := client.SearchProducts("Bananen")
	fmt.Println(bananas[0].Id)

	// You know what, I acctual would love to order some bananas! Let's add 1 pack off bananas:
	client.AddToCart(bananas[0].Id, 1) // This returns a posible error, in production you should handle them.

	// Let's view our cart again
	cart, _ = client.GetCart()
	fmt.Println(cart)

	// I now remember I don't have any money. Let's clear the cart and forget about it
	client.ClearCart() // This returns a posible error, in production you should handle them.

	// There's more to dicover, but I'll let you figure that our on your own.

}
