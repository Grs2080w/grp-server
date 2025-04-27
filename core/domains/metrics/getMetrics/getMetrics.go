package getMetrics

// "github.com/Grs2080w/grp_server/core/domains/metrics/getMetrics"

import (
	"encoding/json"
	"log"
	"strings"
)

type Metrics struct {
	Pk string `json:"Pk" dynamodbav:"pk"`
	Sk  string  `json:"Sk" dynamodbav:"sk"`
	Value float32  `json:"Value" dynamodbav:"value"`
	Username string  `json:"Username" dynamodbav:"username"`
	Domain string  `json:"Domain" dynamodbav:"domain"`
	Type string  `json:"Type" dynamodbav:"type"`
}

type Files_per_extension struct {}
type Records_per_domain struct {}
type Storage_per_type struct {}
type Storage_per_domain struct {}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Files_per_extension  map[string]int  `json:"files_per_extension"`
	Records_per_domain map[string]int    `json:"records_per_domain"`
	Storage_per_type  map[string]int  `json:"storage_per_type"`
	Storage_per_domain  map[string]int  `json:"storage_per_domain"`
}

func ParseMetrics(obj string) ([]Metrics, error) {
	var metrics []Metrics
	err := json.Unmarshal([]byte(obj), &metrics)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return metrics, nil
}


func ExtractKey(sk string) string {
	separated := strings.Split(sk, "#")
	return separated[2]
}
