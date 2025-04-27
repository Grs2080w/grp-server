package downloadFile

import (
	"context"

	"github.com/gin-gonic/gin"

	client "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"
	getURL "github.com/Grs2080w/grp_server/core/db/s3/PresignedUrl/Get"
)

// DownloadFileHandler godoc
// @Summary Download a file
// @Description Return a PRESIGNED URL for Download of a file with the query id and ext, valid for files and ebooks
// @Tags download
// @Accept json
// @Produce json
// @Param id query string true "id file"
// @Param ext query string true "extension file"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /download [get]
func DownloadFileHandler(c *gin.Context) {

	id_file := c.Query("id")
	ext_file := c.Query("ext")

	if id_file == "" {
		c.JSON(400, gin.H{"error": "id file is required"})
		return
	}
	
	if ext_file == "" {
		c.JSON(400, gin.H{"error": "extension is required"})
		return
	}

	url, err := getURL.Presigner{PresignClient: client.CDB.PresignedClient}.GetObject(context.TODO(), client.CDB.BucketName, id_file+"."+ext_file, 60)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not get presigned url"})
		return
	}

	c.JSON(200, gin.H{"url": url.URL})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	URL string `json:"url"`
}