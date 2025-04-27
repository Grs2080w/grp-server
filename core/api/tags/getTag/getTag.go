package getTag

import (
	"context"
	"log"
	"sync"

	"github.com/Grs2080w/grp_server/core/domains/tags/getTag"
	"github.com/gin-gonic/gin"

	cDy "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	qt "github.com/Grs2080w/grp_server/core/db/dynamo/tags/queryTag"
)

// GetTagHandler godoc
// @Summary Get all domains by tags from a user
// @Description Get all domain with relationship with the user tags from a user with the header token
// @Tags tags
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param tag path string true "Tag name"
// @Success 200 {object} getTag.SuccessResponse
// @Failure 400 {object} getTag.ErrorResponse
// @Router /tags/:tag [get]
func GetTagHandler(c *gin.Context) {

	username := c.GetString("username")

	tag_name := c.Param("tag")

	if tag_name == "" {
		c.JSON(400, gin.H{"error": "tag don't be empty"})
		return
	}

	tags, err := (&qt.TableBasics{DynamoDbClient: cDy.CDB.DynamoClient, TableName: cDy.CDB.TableName}).QueryTag(context.TODO(), username + "#" + tag_name)

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	tagsParsed, err := getTag.ParseTags(tags)

	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "error internal server"})
		return
	}

	var (
		response getTag.Response
		mu       sync.Mutex
		wg       sync.WaitGroup
	)
	
	for _, tag := range tagsParsed {
		wg.Add(1)
	
		go func(tagType getTag.Tags) {
			defer wg.Done()
	
			obj := getTag.GetDomain{Id: tagType.Item_id, Username: tagType.Username}
	
			switch tagType.Domain {
			case "files":
				file := obj.GetFile()
				mu.Lock()
				response.Files = append(response.Files, file)
				mu.Unlock()
			case "ebooks":
				ebook := obj.GetEbook()
				mu.Lock()
				response.Ebooks = append(response.Ebooks, ebook)
				mu.Unlock()
			case "passwords":
				pass := obj.GetPass()
				mu.Lock()
				response.Passwords = append(response.Passwords, pass)
				mu.Unlock()
			case "tasks":
				task := obj.GetTask()
				mu.Lock()
				response.Tasks = append(response.Tasks, task)
				mu.Unlock()
			}
		}(tag)
	}
	
	wg.Wait()

	c.JSON(200, response)

}