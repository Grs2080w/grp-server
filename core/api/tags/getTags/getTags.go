package getTags

// "github.com/Grs2080w/grp_server/core/api/tags/getTags"

import (
	"context"
	"log"

	"github.com/Grs2080w/grp_server/core/domains/tags/getTags"
	"github.com/gin-gonic/gin"

	"github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	q "github.com/Grs2080w/grp_server/core/db/dynamo/tags/queryTags"
)

// GetTagsHandler godoc
// @Summary Get all tags
// @Description Get all tags with the header token
// @Tags tags
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} []string
// @Failure 400 {object} getTags.ErrorResponse
// @Router /tags [get]
func GetTagsHandler(c *gin.Context) {

	username := c.GetString("username")

	tags, err := (&q.TableBasics{DynamoDbClient: clientDb.CDB.DynamoClient, TableName: clientDb.CDB.TableName}).QueryTags(context.TODO(), username + "#TAGS")

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	tagsParsed, err := getTags.ParseTags(tags)

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	tagSet := make(map[string]struct{})
	tagList := []string{}

	for _, tag := range tagsParsed {
		nameTag := getTags.SplitTagName(tag)

		if _, exists := tagSet[nameTag]; !exists {
			tagSet[nameTag] = struct{}{}
			tagList = append(tagList, nameTag)
		}
	}


	c.JSON(200, tagList)

}