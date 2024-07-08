package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/avran02/effective-mobile/internal/models"
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	ErrRequestFailed = fmt.Errorf("request failed")
)

func enrichUserData(enrichUserDataServiceURL, passportSerie, passportNumber string) (*models.User, error) {
	params := url.Values{}
	params.Add("passportSerie", passportSerie)
	params.Add("passportNumber", passportNumber)
	requestURL := fmt.Sprintf("%s/info?%s", enrichUserDataServiceURL, params.Encode())

	c := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, requestURL, nil)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус-код ответа
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%w (status code: %d)", ErrRequestFailed, resp.StatusCode)
		slog.Error(err.Error())
		return nil, err
	}

	// Декодируем JSON-ответ
	var enrichedUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&enrichedUser); err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return &enrichedUser, nil
}
