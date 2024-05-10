// Package classification awesome.
//
// Documentation of our awesome API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8001
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- basic
//
// swagger:meta
package docs

import (
	_ "github.com/go-swagger/go-swagger"
)

// swagger:route GET /api/v1/car GetCars idOfGetCarsEndpoint
// Get cars by filter
//
// responses:
//  200: GetCarsResponse
// 	400: ErrorResponse
//	500: ErrorResponse

// swagger:response GetCarsResponse
type GetCarsResponse struct {
	// in:body
	Body Body
}

type Body struct {
	Id     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
	Owner  struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic,omitempty"`
	} `json:"owner"`
	Page int `json:"page,omitempty"`
}

// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	Error ErrorParams
}

type ErrorParams struct {
	Error string `json:"error"`
}

// swagger:parameters idOfGetCarsEndpoint
type GetCarsParams struct {
	//in:query
	RegNum string
	//in:query
	Mark string
	//in:query
	Model string
	//in:query
	Year int
	//in:query
	Name string
	//in:query
	Surname string
	//in:query
	Patronymic string
}

// swagger:route POST /api/v1/car CreateCars idOfCreateCarsEndpoint
// Create cars
//
// Responses:
//
//  201: CreateCarReponse
//  400: ErrorResponse
//  500: ErrorResponse

// swagger:parameters idOfCreateCarsEndpoint
type CreateCarsParams struct {
	//in:body
	RegNum string
	//in:body
	Mark string
	//in:body
	Model string
	//in:body
	Year int
	//in:body
	Name string
	//in:body
	Surname string
	//in:body
	Patronymic string
}

// swagger:response CreateCarReponse
type CreateCarResponse struct {
	// in:body
	RegNums []int `json:"regNums"`
}

// swagger:route PATCH /api/v1/car/{Id} PatchCar idOfPatchCarEndpoint
// Patch car
//
// Responses:
//
//  204:
// 	400: ErrorResponse
//  404: ErrorResponse
//  500: ErrorResponse

// swagger:parameters idOfPatchCarEndpoint
type PatchCarParams struct {
	//in:path
	Id int
	//in:body
	RegNum string
	//in:body
	Mark string
	//in:body
	Model string
	//in:body
	Year int
	//in:body
	Name string
	//in:body
	Surname string
	//in:body
	Patronymic string
}

// swagger:route DELETE /api/v1/car/{Id} DeleteCar idOfDeleteCarEndpoint
// Delete car
//
// Responses:
//
//  204:
// 	400: ErrorResponse
//  404: ErrorResponse
//  500: ErrorResponse

// swagger:parameters idOfDeleteCarEndpoint
type DeleteCarParams struct {
	//in:path
	Id int
}
