package getPwd

import (
	"context"

	"github.com/Grs2080w/grp_server/core/domains/passwords/getPwd"
	"github.com/gin-gonic/gin"

	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	g "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/getPassword"
)

// @Summary Get a password
// @Description Get a password with the id in the path, this endpoint require a username in the header
// @Tags passwords
// @Accept  json
// @Produce  json
// @Param id path string true "Password id"
// @Success 200 {object} getPwd.Passwords
// @Failure 400 {object} getPwd.ErrorResponse
// @Router /passwords/{id} [get]
func GetPwdHandler(c *gin.Context) {

	username := c.GetString("username")

	id_pwd := c.Param("id")

	if id_pwd == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	pwdGet, err := (&g.TableBasics{DynamoDbClient: client.CDB.DynamoClient,TableName: client.CDB.TableName}).GetPassword(context.TODO(), username + "#PASSWORDS", "PASSWORD#" + id_pwd)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query password"})
		return
	}

	pwd := getPwd.ParsePassword(pwdGet)

	c.JSON(200, pwd)
}