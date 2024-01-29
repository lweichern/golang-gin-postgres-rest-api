package controller

import "example/http-server/models"

// tokens slice to store token generated for verfication purpose
var tokens []string
var Users = make(map[string]models.User)