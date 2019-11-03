package imgur

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type imgurResponseData struct {
	Link string
}

type imgurResponse struct {
	Data imgurResponseData
}

// Upload method to upload
func Upload(name, path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("image", path)
	if err != nil {
		return "", err
	}

	io.Copy(part, file)
	writer.WriteField("title", name)

	client := &http.Client{}
	request, _ := http.NewRequest(
		"POST",
		"https://api.imgur.com/3/image",
		body,
	)
	request.Header.Set("Authorization", "Client-ID "+"ea6c0ef2987808e")
	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := &imgurResponse{}
	json.Unmarshal(data, response)

	return response.Data.Link, nil
}
