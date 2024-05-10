package main

import "fmt"

type ResponseBuilder struct {
	Response Response
}

func NewBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response: Response{},
	}
}

type Response struct {
	User       User
	AccountBal Account
	Address    Address
}

type User struct {
	Email string
}

type Account struct {
	Balance float64
}

type Address struct {
	City string
}

func (rb *ResponseBuilder) SetEmail(email string) *ResponseBuilder {
	rb.Response.User.Email = email
	return rb
}

func (rb *ResponseBuilder) SetBalance(bal float64) *ResponseBuilder {
	rb.Response.AccountBal.Balance = bal
	return rb
}

func (rb *ResponseBuilder) SetCity(city string) *ResponseBuilder {
	rb.Response.Address.City = city
	return rb
}

func (rb *ResponseBuilder) Build() Response {
	return rb.Response
}

func main() {
	rb := NewBuilder().
		SetEmail("saikumar@gmail.com").
		SetBalance(20000).
		SetCity("Vizianagaram")
	fmt.Printf("response %+v\n", rb.Build())
}
