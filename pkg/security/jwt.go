package security

import (
	"errors"
	"grocery-management/internal/models"
	"grocery-management/pkg/logger"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var secret string // secret is the JWT secret key used for signing tokens
var exp string    // exp is the expiration time for the JWT token

func Init(log *zap.Logger) {
	secret = os.Getenv("JWT_SECRET")
	exp = os.Getenv("JWT_EXP")

	log.Info("JWT_SECRET", zap.String("secret", secret))
	log.Info("JWT_EXP", zap.String("exp", exp))
}

func GenerateJWT(user *models.User) (string, error) {
	if user == nil {
		logger.Log.Error("user cannot be nil")
		return "", errors.New("user cannot be nil")
	}

	// Create a new JWT token
	if secret == "" || exp == "" {
		logger.Log.Error("JWT_SECRET and JWT_EXP environment variables must be set")
		return "", errors.New("JWT_SECRET and JWT_EXP environment variables must be set")
	}

	// Set the claims
	claims := jwt.MapClaims{
		"sub":   user.Username,
		"exp":   exp,
		"orgId": user.GroceryID,
		"role":  user.RoleID,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logger.Log.Error("failed to sign token", zap.String("error", err.Error()))
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func GetClaims(token *jwt.Token) (string, string, string, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["sub"].(string)
		exp := claims["exp"].(string)
		orgId := claims["orgId"].(string)
		return username, exp, orgId, nil
	}
	return "", "", "", errors.New("invalid token claims")
}

func GetUserID(token *jwt.Token) (string, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string)
		return userID, nil
	}
	return "", errors.New("invalid token claims")
}
