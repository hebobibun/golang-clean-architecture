package helper

import "github.com/golang-jwt/jwt"

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userid := claims["userID"].(float64)
		return int(userid)
	}
	return -1
}