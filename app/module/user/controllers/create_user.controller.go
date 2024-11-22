package user_controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (uc *userController) RegisterUser(c *fiber.Ctx) error {

	payload := new(user_model.UserRegisterStorage)
	if err := c.BodyParser(payload); err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	validate := context.AppContext().GetValidator()
	if err := validate.Struct(payload); err != nil {
		return common_response.NewAppError(c).IsValidationError(err).SendJSON()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return common_response.NewAppError(c).IsInternalServerError(err).SendJSON()
	}
	expirationTime := time.Now().Add(90 * 24 * time.Hour).Unix()
	finalPassword := fmt.Sprintf("%s::%d", string(hashedPassword), expirationTime)
	payload.Password = finalPassword

	// Generate activationKey và expirationTime
	activationKey, err := generateActivationKey(32)
	if err != nil {
		return common_response.NewAppError(c).IsInternalServerError(err).SendJSON()
	}
	activationExpirationTime := time.Now().Add(24 * time.Hour).Unix()
	finalActivationKey := fmt.Sprintf("%s::%d", activationKey, activationExpirationTime)
	payload.ActivationKey = finalActivationKey

	user, err := uc.userServices.RegisterUser(c.Context(), payload)
	if err != nil {
		return common_response.NewAppError(c).IsDatabaseError(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).WithData(user).SendJSON()

}

func generateActivationKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
