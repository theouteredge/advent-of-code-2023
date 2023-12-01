package adventofcode2023

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

func GetInput(url string) (string, error) {
	client := &http.Client{}
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		log.Errorf("Error creating request: %v", err)
		return "", err

	} else {
		req.Header.Add("Cookie", "session=53616c7465645f5ff6d7b44d1963d1262891ac8a5e80a8442d9ab0571d1fe2b008b14f75abbff6e86831a986ee74ea524528eaba85f3f09f97d7eb7a6335fd27")
		if resp, err := client.Do(req); err != nil {
			log.Errorf("Error getting response: %v", err)
			return "", err

		} else {
			defer resp.Body.Close()
			if body, err := io.ReadAll(resp.Body); err != nil {
				log.Errorf("Error reading response body: %v", err)
				return "", err
			} else {
				return string(body), nil
			}
		}
	}
}

func GetInputFromFile(filename string) (string, error) {
	// open input file
	if file, err := os.ReadFile(filename); err != nil {
		log.Errorf("Error opening file [%s]: %v", filename, err)
		return "", err
	} else {
		return string(file), nil
	}
}
