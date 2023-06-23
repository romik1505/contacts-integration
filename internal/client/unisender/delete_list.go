package unisender

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DeleteListResponse struct {
	Result *struct{} `json:"result,omitempty"`
	Code   string    `json:"code,omitempty"`
	Err    string    `json:"error,omitempty"`
}

func (c Client) DeleteList(ctx context.Context, apiKey string, listID uint64) error {
	reqString := fmt.Sprintf("%s/ru/api/deleteList?format=json&api_key=%s&list_id=%d", UnisenderURL, apiKey, listID)
	resp, err := c.httpClient.Get(reqString)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respSt := &CreateListResponse{}
	if err := json.Unmarshal(data, respSt); err != nil {
		return err
	}

	if respSt.Result == nil {
		return respSt
	}

	return nil
}
