package utils

import (
	"encoding/json"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


func GenerateJwtToken(data interface{}) (string,string) {

	user,_ := json.Marshal(data)
	payload := string(user)
	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(10 * time.Minute).Unix(), 0)),
		Subject:  payload,
	}).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(24 * time.Hour).Unix(), 0)),
	Subject:  payload,
	}).SignedString([]byte(os.Getenv("REFRESH_JWT_SECRET_KEY")))
	
		
	return  accessToken	,refreshToken
}


func ValidateToken(token string,user interface{}) error {
	claims := jwt.MapClaims{}
	 _  , err := jwt.ParseWithClaims(token,claims, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET_KEY")), nil
    })

	if err != nil{
		return err
	}
	
	for key, val := range claims {
		if key == "sub"{
			 json.Unmarshal([]byte(val.(string)), user)
		}
	}

	return nil

}

func ValidateRefreshToken(token string,user interface{}) error {
	claims := jwt.MapClaims{}
	 _  , err := jwt.ParseWithClaims(token,claims, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("REFRESH_JWT_SECRET_KEY")), nil
    })

	if err != nil{
		return err
	}
	
	for key, val := range claims {
		if key == "sub"{
			 json.Unmarshal([]byte(val.(string)), user)
		}
	}

	return nil

}
