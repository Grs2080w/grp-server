package getTag

import (
	"context"
	"encoding/json"
	"errors"

	cDy "github.com/Grs2080w/grp_server/core/db/dynamo/clientDb"

	getEbook "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/getEbook"
	qf "github.com/Grs2080w/grp_server/core/db/dynamo/files/query"
	getPass "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/getPassword"
	getTask "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/getTask"

	modelE "github.com/Grs2080w/grp_server/core/db/dynamo/ebooks/.model"
	modelF "github.com/Grs2080w/grp_server/core/db/dynamo/files/.model"
	modelP "github.com/Grs2080w/grp_server/core/db/dynamo/passwords/.model"
	modelT "github.com/Grs2080w/grp_server/core/db/dynamo/tasks/.model"
)


type Response struct {
	Files []modelF.Files `json:"files"`
	Ebooks []modelE.Ebook `json:"ebooks"`
	Passwords []modelP.Passwords `json:"passwords"`
	Tasks []modelT.Tasks `json:"tasks"`
}

type Files struct {}
type Ebook struct {}
type Passwords struct {}
type Tasks struct {}

type SuccessResponse struct {
	Files []Files `json:"files"`
	Ebooks []Ebook `json:"ebooks"`
	Passwords []Passwords `json:"passwords"`
	Tasks []Tasks `json:"tasks"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Tags struct {
	Pk    string                 `dynamodbav:"pk"`
	Sk    string                 `dynamodbav:"sk"`
	Domain string                 `dynamodbav:"domain"`
	Item_id string                 `dynamodbav:"item_id"`
	Tag string                 `dynamodbav:"tag"`
	Date string                 `dynamodbav:"date"`
	Username string  `dynamodbav:"username"`
}

type GetDomain struct {
	Id string
	Username string
}

func (g *GetDomain) GetFile() (modelF.Files) {

	filesRes, _ := qf.TableBasics{DynamoDbClient :cDy.CDB.DynamoClient, TableName: cDy.CDB.TableName}.Query(context.TODO(), g.Username + "#FILES")

	var files []modelF.Files
	json.Unmarshal([]byte(filesRes), &files)

	var fileToReturn modelF.Files


	for _, file := range files {

		versions := file.Versions

		for _, version := range versions {
			if version.Id == g.Id {
				fileToReturn = file
			}
		}
	}

	return fileToReturn
}

func (g *GetDomain) GetEbook() (modelE.Ebook) {
	ebookRes, _ := (&getEbook.TableBasics{DynamoDbClient :cDy.CDB.DynamoClient, TableName: cDy.CDB.TableName}).GetEbook(context.TODO(), g.Username + "#EBOOKS", "EBOOKS#" + g.Id)

	var ebook modelE.Ebook
	json.Unmarshal([]byte(ebookRes), &ebook)

	return ebook
}

func (g *GetDomain) GetPass() (modelP.Passwords) {

	passRes, _ := (&getPass.TableBasics{DynamoDbClient :cDy.CDB.DynamoClient, TableName: cDy.CDB.TableName}).GetPassword(context.TODO(), g.Username + "#PASWORDS", "PASSWORDS#" + g.Id)

	var pass modelP.Passwords
	json.Unmarshal([]byte(passRes), &pass)

	return pass
}

func (g *GetDomain) GetTask() (modelT.Tasks) {

	taskRes, _ := (&getTask.TableBasics{DynamoDbClient :cDy.CDB.DynamoClient, TableName: cDy.CDB.TableName}).GetTask(context.TODO(), g.Username + "#TASKS", "TASKS#" + g.Id)

	var task modelT.Tasks
	json.Unmarshal([]byte(taskRes), &task)

	return task
}

func ParseTags(obj string) ([]Tags, error) {
	var tags []Tags
	err := json.Unmarshal([]byte(obj), &tags)
	if err != nil {
		return nil, errors.New("error parsing tags")
	}

	return tags, nil
}

func UnparseResponse(obj Response) []byte {
	data, _ :=json.Marshal(obj)
	return data
}

func ParseResponse(obj string) Response {
	var response Response
	json.Unmarshal([]byte(obj), &response)
	return response
}




