package line

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Liff struct {
	Uid     string
	Name    string
	Picture string
}

type liffVerifiedResponse struct {
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func VerifiedIdToken(idToken string) *Liff {
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	values := url.Values{}
	values.Add("id_token", idToken)
	values.Add("client_id", os.Getenv("LIFF_CLIENT_ID"))
	req, _ := http.NewRequest("POST", "https://api.line.me/oauth2/v2.1/verify", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var response liffVerifiedResponse
	_ = json.Unmarshal(body, &response)

	return &Liff{
		Uid:     response.Sub,
		Name:    response.Name,
		Picture: response.Picture,
	}
}
