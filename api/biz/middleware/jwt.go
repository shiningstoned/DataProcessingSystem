package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "please login first",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next(ctx)
	}
}

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

type JWT struct {
	Secret []byte
}

func NewJWT() *JWT {
	return &JWT{
		Secret: []byte("hdfs"),
	}
}

func (j *JWT) CreateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Issuer:    "hdfs",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.Secret)
	return signedString, err
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("can't handle this token")
}
