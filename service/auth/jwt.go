package auth

import (
	"Product-Hub/config"
	"Product-Hub/types"
	"Product-Hub/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type conextKey string

const UserKey conextKey = "userID"

func CreateJwt(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func WithJwtAuth(handler http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenstr := getTokenFromRequest(r)
		fmt.Println("TOKEN:", tokenstr)

		token, err := validateToken(tokenstr)
		if err != nil {
			permissionDenied(w)
			log.Print("Error validating token: ", err)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			log.Print("invalid token")
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userIDStr := claims["userId"].(string)

		userId, err := strconv.ParseInt(userIDStr, 10, 64)

		u, err := store.GetUserById(r.Context(), int(userId))
		if err != nil {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("user not found"))
			permissionDenied(w)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)
		fmt.Println("into next")
		handler.ServeHTTP(w, r)
	}
}
func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth == "" {
		return ""
	}
	return tokenAuth
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ([]byte)(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdfromRequest(ctx context.Context) (int64, error) {
	userId, ok := ctx.Value(UserKey).(int64)
	if !ok {
		return 0, fmt.Errorf("user not found in context")
	}
	return userId, nil
}
