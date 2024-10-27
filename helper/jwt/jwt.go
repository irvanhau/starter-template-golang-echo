package jwt

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"starter-template/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type JWTInterface interface {
	GenerateJWT(id uint, username, email, phoneNumber, role string) map[string]any
	RefreshJWT(accessToken string, refreshToken *jwt.Token) (map[string]any, error)
	ExtractToken(c echo.Context) (ExtractToken, error)
	GetCurrentToken(c echo.Context) *jwt.Token
	ValidateRole(c echo.Context) bool
}

type JWT struct {
	c *configs.ProgramConfig
}

type ExtractToken struct {
	ID          uint
	Username    string
	Email       string
	PhoneNumber string
	Role        string
}

func NewJWT(config *configs.ProgramConfig) JWTInterface {
	return &JWT{
		c: config,
	}
}

func (j *JWT) GenerateJWT(id uint, username, email, phoneNumber, role string) map[string]any {
	var result = map[string]any{}
	var accessToken = j.generateToken(id, username, email, phoneNumber, role)
	var refreshToken = j.generateRefreshToken()
	if accessToken == "" || refreshToken == "" {
		return nil
	}

	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return result
}

func (j *JWT) generateToken(id uint, username, email, phoneNumber, role string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["username"] = username
	claims["email"] = email
	claims["phone_number"] = phoneNumber
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.c.Secret))
	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) generateRefreshToken() string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(j.c.RefSecret))
	if err != nil {
		return ""
	}

	return refreshToken
}

func (j *JWT) RefreshJWT(accessToken string, refreshToken *jwt.Token) (map[string]any, error) {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		logrus.Error("Get Token Expiration Error, ", err.Error())
		return nil, errors.New("JWT : Token Expiration")
	}

	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = jwt.MapClaims{}
		newToken, err := jwt.ParseWithClaims(accessToken, newClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.c.Secret), nil
		})

		if err != nil {
			logrus.Error("Error Parse With Claims, ", err.Error())
			return nil, errors.New("JWT : Parse With Claims")
		}

		newClaim = newToken.Claims.(jwt.MapClaims)
		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 48).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignRefrToken, err := newRefreshToken.SignedString(refreshToken.Signature)

		if err != nil {
			logrus.Error("Error Signed String Refresh Token, ", err.Error())
			return nil, errors.New("JWT : Signed String Refresh Token")
		}

		result["access_token"] = newToken.Raw
		result["refresh_token"] = newSignRefrToken

		return result, nil
	}

	return nil, errors.New("JWT : Refresh Token Not Valid && Expired")
}

func (j *JWT) validateToken(token string) (*jwt.Token, error) {
	var authHeader = token[7:]
	parsedToken, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT : Unexpected Signing Method %v", t.Header["alg"])
		}
		return []byte(j.c.Secret), nil
	})

	if err != nil {
		logrus.Error("Parse Token Error, ", err.Error())
		return nil, err
	}

	return parsedToken, nil
}

func (j *JWT) ExtractToken(c echo.Context) (ExtractToken, error) {
	var result = new(ExtractToken)
	authHeader := c.Request().Header.Get("Authorization")

	token, err := j.validateToken(authHeader)

	if err != nil {
		logrus.Error("Validate Token Error, ", err.Error())
		return ExtractToken{}, err
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	idFloat, ok := mapClaims["id"].(float64)
	email := mapClaims["email"].(string)
	username := mapClaims["username"].(string)
	phoneNumber := mapClaims["phone_number"].(string)
	role := mapClaims["role"].(string)

	if !ok {
		return ExtractToken{}, errors.New("JWT : ID not found or not a valid number")
	}

	idInt := uint(idFloat)
	result.Email = email
	result.PhoneNumber = phoneNumber
	result.Username = username
	result.Role = role
	result.ID = idInt

	return *result, nil
}

func (j *JWT) GetCurrentToken(c echo.Context) *jwt.Token {
	currentToken := c.Get("user").(*jwt.Token)

	return currentToken
}

func (j *JWT) ValidateRole(c echo.Context) bool {
	ext, _ := j.ExtractToken(c)

	return ext.Role == "Admin"
}
