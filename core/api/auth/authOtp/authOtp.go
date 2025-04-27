package authOtp

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/domains/auth/authOtp"

	clientDb "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getU "github.com/Grs2080w/grp_server/core/db/dynamo/users/getUser"
	r "github.com/Grs2080w/grp_server/core/db/redis"

	auth "github.com/Grs2080w/grp_server/core/utils/authToken"
	rc "github.com/Grs2080w/grp_server/core/utils/ramdomCode"
	req "github.com/Grs2080w/grp_server/core/utils/requestOtp"
)

// AuthHandler godoc
// @Summary Send otp code
// @Description This route is called when the user wants to change a sensitive information, it sends an otp code for the user email
// @Tags auth
// @Accept json
// @Produce json
// @Param user body authOtp.UserAuthOtp true "User token in headers"
// @Success 200 {object} authOtp.SuccessResponse
// @Failure 400 {object} authOtp.ErrorResponse
// @Router /auth/otp [post]
func AuthHandler(c *gin.Context) {

	user_username, err := auth.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := (&getU.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).GetUser(context.TODO(), user_username + "#" + "PROFILE", "USERS" + "#" + user_username)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	userParsed := authOtp.ParseUnmarshal(userResponse)


	otpCode := rc.RandomCode()
	r.R_set(user_username + "#OTP", otpCode, 300)

	response := req.RequestOtp(otpCode, userParsed.Email)

	if response.Message != "OTP sent successfully" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "otp code not sent"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "otp code sent to " + userParsed.Email[0:12] + "..."})
}
