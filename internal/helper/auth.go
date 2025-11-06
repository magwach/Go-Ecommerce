package helper

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/schema"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Auth struct {
	Secret string
}

type JWTRequirements struct {
	UserID uuid.UUID
	Email  string
	Role   string
}

func InitializeAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (r Auth) GenerateJWT(requirements JWTRequirements) (string, error) {
	var validate = validator.New()
	if err := validate.Struct(requirements); err != nil {
		return "", errors.New("required inputs missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": requirements.UserID.String(),
		"email":   requirements.Email,
		"role":    requirements.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	signedToken, err := token.SignedString([]byte(r.Secret))

	if err != nil {
		return "", errors.New("failed to sign the token")
	}

	return signedToken, nil
}

func (r Auth) VerifyJWT(token string) (schema.User, error) {
	tokenArr := strings.Split(token, " ")

	if len(tokenArr) != 2 {
		return schema.User{}, errors.New("invalid authorization header")
	}

	if tokenArr[0] != "Bearer" {
		return schema.User{}, errors.New("invalid authorization header")
	}

	parsedToken, err := jwt.Parse(tokenArr[1], func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method: %v", t.Header)
		}
		return []byte(r.Secret), nil
	})

	if err != nil {
		return schema.User{}, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return schema.User{}, errors.New("expired token")
		}

		user := schema.User{}

		id, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return schema.User{}, errors.New("invalid user ID in token")
		}
		user.ID = id
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)
		return user, nil
	}

	return schema.User{}, errors.New("token verification failed")
}

func (r Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	user, err := r.VerifyJWT(authHeader)

	if err == nil && len(user.ID) > 1 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "authorization failed",
			"error":   err.Error(),
		})
	}
}

func (r Auth) GetCurrentUser(ctx *fiber.Ctx) schema.User {
	currentUser := ctx.Locals("user").(schema.User)
	return currentUser
}
