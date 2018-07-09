package models

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	User UserClaims `json:"user"`
	jwt.StandardClaims
}

type UserClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
