package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"os"
	"time"
)

var (
	jwtKey                     string
	accessTokenDurationString  string
	refreshTokenDurationString string
)

type jwtClaims struct {
	UserID     uint `json:"user_id"`
	UserRoleID uint `json:"user_role_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userID, roleID uint) (string, error) {
	// converting string to time.Duration
	accessTokenDurationString = os.Getenv("ACCESS_TOKEN_DURATION")
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationString)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(accessTokenDuration)

	claims := &jwtClaims{
		UserID:     userID,
		UserRoleID: roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey = os.Getenv("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken() (string, time.Time, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", time.Time{}, err
	}
	refreshTokenDurationString = os.Getenv("REFRESH_TOKEN_DURATION")
	refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationString)
	if err != nil {
		return "", time.Time{}, err
	}
	return fmt.Sprintf("%x", b), time.Now().Add(refreshTokenDuration), nil
}

func ParseToken(accessToken string) (uint, uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtKey = os.Getenv("JWT_KEY")
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return 0, 0, fmt.Errorf("error get user claims from token")
	}

	return claims.UserID, claims.UserRoleID, nil
}
