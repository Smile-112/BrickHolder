package services

import (
        "encoding/json"
        "fmt"
        "strconv"

        "net/http"
        "time"

        "set-service/models"
)

type RebrickableClient struct {
	BaseURL    string
	APIKey     string
	HttpClient *http.Client
}

func NewRebrickableClient(apiKey string) *RebrickableClient {
	return &RebrickableClient{
		APIKey: apiKey,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Получить список наборов Lego (пример с ограничением 10)
func (c *RebrickableClient) GetLegoSets(pageSize int) (*models.SetsResponse, error) {
	url := fmt.Sprintf("https://rebrickable.com/api/v3/lego/sets/?page_size=%d", pageSize)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "key "+c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Rebrickable API returned status %d", resp.StatusCode)
	}

	var setsResponse models.SetsResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&setsResponse); err != nil {
		return nil, err
	}

	return &setsResponse, nil
}

// Получения списка всех серий
func FetchAllSeries(apiKey string) ([]models.Series, error) {
	var allSeries []models.Series
	url := "https://rebrickable.com/api/v3/lego/themes/"

	client := &http.Client{}

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "key "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch series: %s", resp.Status)
		}

		var sr models.SeriesResponse
		if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
			return nil, err
		}

		allSeries = append(allSeries, sr.Results...)
		url = sr.Next // ссылка на следующую страницу или пустая строка, если последняя
		resp.Body.Close()
	}
	return allSeries, nil
}

// Получения списка всех наборов
func FetchAllSets(apiKey string) ([]models.Set, error) {
	var allSets []models.Set
	url := "https://rebrickable.com/api/v3/lego/sets/"

	client := &http.Client{}

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "key "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch set: %s", resp.Status)
		}

		var sr models.SetsResponse
		if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
			return nil, err
		}

		allSets = append(allSets, sr.Results...)
		url = sr.Next // ссылка на следующую страницу или пустая строка, если последняя
		resp.Body.Close()
	}
	return allSets, nil
}

// Получения списка всех минифигурок
func FetchAllMinifigs(apiKey string) ([]models.Minifig, error) {
        return FetchMinifigsChunk(apiKey, 1, 1000, 0)
}

// FetchMinifigsChunk загружает указанное количество страниц минифигурок.
// Если pages == 0, то загружаются все доступные страницы начиная со startPage.
func FetchMinifigsChunk(apiKey string, startPage, pageSize, pages int) ([]models.Minifig, error) {
        var allMinifigs []models.Minifig
        url := fmt.Sprintf("https://rebrickable.com/api/v3/lego/minifigs/?page=%d&page_size=%d", startPage, pageSize)
        client := &http.Client{}

        fetched := 0
        wait := 200 * time.Millisecond

        for url != "" {
                req, err := http.NewRequest("GET", url, nil)
                if err != nil {
                        return nil, err
                }
                req.Header.Set("Authorization", "key "+apiKey)

                resp, err := client.Do(req)
                if err != nil {
                        return nil, err
                }

                if resp.StatusCode == http.StatusTooManyRequests {
                        waitDur := parseRetryAfter(resp.Header.Get("Retry-After"))
                        resp.Body.Close()
                        if waitDur == 0 {
                                waitDur = time.Second
                        }
                        time.Sleep(waitDur)
                        if wait < time.Minute {
                                wait *= 2
                        }
                        continue
                }

                if remaining := resp.Header.Get("X-RateLimit-Remaining"); remaining == "0" {
                        waitDur := parseRetryAfter(resp.Header.Get("Retry-After"))
                        if waitDur == 0 {
                                waitDur = time.Minute
                        }
                        resp.Body.Close()
                        time.Sleep(waitDur)
                        continue
                }

                wait = 200 * time.Millisecond

                if resp.StatusCode != http.StatusOK {
                        return nil, fmt.Errorf("failed to fetch minifigs: %s", resp.Status)
                }

                var mr models.MinifigsResponse
                if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
                        resp.Body.Close()
                        return nil, err
                }
                resp.Body.Close()

                allMinifigs = append(allMinifigs, mr.Results...)
                url = mr.Next

                fetched++
                if pages > 0 && fetched >= pages {
                        break
                }

                time.Sleep(wait)
        }
        return allMinifigs, nil
}

// parseRetryAfter преобразует значение заголовка Retry-After в продолжительность ожидания.
func parseRetryAfter(v string) time.Duration {
        if v == "" {
                return 0
        }
        if sec, err := strconv.ParseFloat(v, 64); err == nil {
                if sec <= 0 {
                        return 0
                }
                return time.Duration(sec * float64(time.Second))
        }
        return 0
}
