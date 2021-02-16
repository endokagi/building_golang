package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostUrl(res http.ResponseWriter, req *http.Request, url string) {

	body, _ := ioutil.ReadAll(req.Body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
