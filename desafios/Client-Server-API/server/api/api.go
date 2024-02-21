package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type NewApi struct {}

type USDBRL struct {
	Usdbrl UsdbrlDto `json:"USDBRL,omitempty"`
}

type UsdbrlDto struct {
	Code       string `json:"code,omitempty"`
	Codein     string `json:"codein,omitempty"`
	Name       string `json:"name,omitempty"`
	High       string `json:"high,omitempty"`
	Low        string `json:"low,omitempty"`
	VarBid     string `json:"varBid,omitempty"`
	PctChange  string `json:"pctChange,omitempty"`
	Bid        string `json:"bid,omitempty"`
	Ask        string `json:"ask,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func USDBRLInit() *NewApi {
	return &NewApi{}
}

func (c *NewApi) GetDataFromApi() (*USDBRL, error) {
	payloadNormalized := USDBRL{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	payload, _ := io.ReadAll(res.Body)

	json.Unmarshal(payload, &payloadNormalized)
	if payloadNormalized.Usdbrl.Ask == "" {
		return nil, errors.New(string(payload))
	}

	return &payloadNormalized, nil
}