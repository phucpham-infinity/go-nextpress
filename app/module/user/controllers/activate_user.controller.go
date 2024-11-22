package user_controllers

import (
	"github.com/gofiber/fiber/v2"
	common_response "github.com/phucpham-infinity/go-nextpress/app/common/response"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
)

func (uc *userController) ActivateUser(c *fiber.Ctx) error {

	payload := new(user_model.ActivateUserPrams)
	if err := c.BodyParser(payload); err != nil {
		return common_response.NewAppError(c).IsBadRequest(err).SendJSON()
	}
	validate := context.AppContext().GetValidator()
	if err := validate.Struct(payload); err != nil {
		return common_response.NewAppError(c).IsValidationError(err).SendJSON()
	}
	return common_response.NewSuccessResponse(c).SendJSON()
}

//func (us *userService) ActivateUser(ctx context.Context, activationKeyInput string) error {
//	// Tìm người dùng theo activationKey
//	user, err := us.userRepository.FindByActivationKey(ctx, activationKeyInput)
//	if err != nil {
//		return err
//	}
//
//	// Nếu không tìm thấy người dùng, trả lỗi
//	if user == nil {
//		return errors.New("invalid activation key")
//	}
//
//	// Tách activationKey thành [activationKey] và [expirationTime]
//	parts := strings.Split(user.ActivationKey, "::")
//	if len(parts) != 2 {
//		return errors.New("invalid activation key format")
//	}
//
//	storedActivationKey := parts[0]
//	expirationTimeStr := parts[1]
//
//	// Chuyển expirationTime từ string sang int64
//	expirationTime, err := strconv.ParseInt(expirationTimeStr, 10, 64)
//	if err != nil {
//		return errors.New("invalid expiration time")
//	}
//
//	// Kiểm tra nếu activationKey đã hết hạn
//	if time.Now().Unix() > expirationTime {
//		return errors.New("activation key has expired")
//	}
//
//	// Kiểm tra nếu activationKey nhập vào khớp với activationKey lưu trữ
//	if storedActivationKey != activationKeyInput {
//		return errors.New("invalid activation key")
//	}
//
//	// Kích hoạt tài khoản
//	user.IsActive = true
//	user.ActivationKey = "" // Xóa activationKey sau khi kích hoạt
//	return us.userRepository.Update(ctx, user)
//}
