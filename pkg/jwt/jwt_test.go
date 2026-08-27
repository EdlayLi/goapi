package jwt_test

import (
	"apigo/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "11a@a.ru"
	jwtService := jwt.NewJWT("s3cr3tK3y!@2026xYzQwErTyUiOpAsDfGhJkLzXcVbNm")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
