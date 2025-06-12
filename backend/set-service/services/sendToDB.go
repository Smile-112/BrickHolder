package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"set-service/models"
)

func SendSeriesToDataService(series models.Series) error {
	jsonData, _ := json.Marshal(series)
	resp, err := http.Post("http://localhost:8081/api/lego/series", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("data-service returned status: %s", resp.Status)
	}
	return nil
}

func SendSetToDataService(set models.Set) error {
	jsonData, _ := json.Marshal(set)
	resp, err := http.Post("http://localhost:8081/api/lego/sets", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("data-service returned status: %s", resp.Status)
	}
	return nil
}

func SendMinifigToDataService(fig models.Minifig) error {
	jsonData, _ := json.Marshal(fig)
	resp, err := http.Post("http://localhost:8081/api/lego/minifigs", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("data-service returned status: %s", resp.Status)
	}
	return nil
}
