package controller

import (
	"net/http"

	"github.com/aristidesneto/crud-golang/src/config/logger"
	"github.com/aristidesneto/crud-golang/src/config/validation"
	"github.com/aristidesneto/crud-golang/src/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FindUserByID(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {
	c.JSON(http.StatusOK, "dfdf")
}

func StoreUser(c *gin.Context) {

	logger.Info("Store user", zap.String("journey", "storeUser"))

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "storeUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Debug("User create successfully")
	c.JSON(http.StatusOK, gin.H{
		"data": userRequest,
	})

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
