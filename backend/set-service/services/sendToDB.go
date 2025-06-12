package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"set-service/models"
)

func SendSeriesToDataService(theme models.Series) error {
	jsonData, _ := json.Marshal(theme)
	resp, err := http.Post("http://localhost:8081/api/lego/series", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to save theme, status: %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("data-service вернул статус: %s", resp.Status)
	}
	return nil
}

func SendSetToDataService(theme models.Set) error {
	jsonData, _ := json.Marshal(theme)
	resp, err := http.Post("http://localhost:8081/api/lego/sets", "application/json", bytes.NewBuffer(jsonData))
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
