package requestOtp

// "github.com/Grs2080w/grp_server/core/utils/requestOtp"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	c "github.com/Grs2080w/grp_server/core/config"
)

type Response struct {
	Message string `json:"message"`
}

func RequestOtp(code, email string) Response {
	url := c.GetValueEnv("OTP_URL")
	// string to payload
	payloadTxt := fmt.Sprintf(`{"code": "%s", "email": "%s"}`, code, email)
	payload := []byte(payloadTxt)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Print(err, "error on request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err, "error on response")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Print(err, "error on unmarshal")
	}

	return response
}
