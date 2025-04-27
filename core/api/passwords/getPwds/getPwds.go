package getPwds

import (
	"context"

	"github.com/Grs2080w/grp_server/core/domains/passwords/getPwds"
	"github.com/gin-gonic/gin"

	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	q "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/query"
)

// @Summary Get user's passwords
// @Description Retrieve all passwords associated with the user, requires a username in the header
// @Tags passwords
// @Accept  json
// @Produce  json
// @Success 200 {object} []getPwds.Passwords
// @Failure 500 {object} getPwds.ErrorResponse
// @Router /passwords [get]
func GetPwdsHandler(c *gin.Context) {

	username := c.GetString("username")

	pwdsGet, err := q.TableBasics{DynamoDbClient: client.CDB.DynamoClient,TableName: client.CDB.TableName}.Query(context.TODO(), username + "#PASSWORDS")

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query passwords"})
		return
	}

	pwds := getPwds.ParsePasswords(pwdsGet)

	c.JSON(200, pwds)
}