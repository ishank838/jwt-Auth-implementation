package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err := errors.New("Unauthorized to access this resource")
		return err
	}
	CheckUserType(c, userType)
	return nil
}

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	if userType != role {
		return errors.New("Unauthorized to access this resource")
	}
	return nil
}
