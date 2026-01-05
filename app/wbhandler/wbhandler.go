package wbhandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"app/app/models"
)

func WebHookHandler(wr http.ResponseWriter, resp *http.Request) {
	fmt.Println("Path:", resp.URL.Path)
	body, err := io.ReadAll(resp.Body)
	fmt.Println(resp.Header)

	var payload models.Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(wr, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	tests := payload.Tests
	fmt.Println(tests)
}
