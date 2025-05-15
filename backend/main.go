package main

import v1 "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1"

// @title Cavent API
// @version 1.0
// @description This is the Cavent backend API documentation.
// @BasePath /api/v1
func main(){
	r := v1.Init()
	r.Run(":8080")
}