package wbhandler

import (
	"app/app/models"
	app "app/app/validate"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
)

func WebHookHandler(wr http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(wr, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer req.Body.Close()

	var payload models.Payload
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)

	if err != nil {
		http.Error(wr, "Bad request", http.StatusBadRequest)
		return
	}
	err = app.ValidateJsonData(&payload)
	if err != nil {
		http.Error(wr, "Invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	tests := payload.Tests
	pass_rate, err := CountPassRate(tests)
	if err != nil {
		http.Error(wr, "Error in tests: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(pass_rate)

	response := models.Response{
		Received: true,
		BuildId:  payload.BuildID,
		PassRate: pass_rate,
	}
	wr.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(wr).Encode(response)

	if err2 != nil {
		http.Error(wr, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func CountPassRate(test_data []models.Test) (float64, error) {
	if len(test_data) == 0 {
		return 0, errors.New("No tests provided")
	}
	passed := 0
	total_number := len(test_data)

	for _, test_case := range test_data {
		if test_case.Name == "" {
			return 0, errors.New("Test name cannot be empty")
		}
		if test_case.Passed {
			passed += 1
		}
	}
	pass_rate := float64(passed) / float64(total_number) * 100
	pass_rate = math.Round(pass_rate*100) / 100
	return pass_rate, nil
}
