package services

import (
	"autentification-service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendUserToDataService(theme models.User) error {
	jsonData, _ := json.Marshal(theme)
	resp, err := http.Post("http://localhost:8081/api/users/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to save set, status: %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("data-service вернул статус: %s", resp.Status)
	}
	return nil
}
