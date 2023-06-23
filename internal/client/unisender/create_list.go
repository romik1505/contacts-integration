package unisender

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CreateListResponse struct {
	Result *struct {
		ID uint64 `json:"id,omitempty"`
	} `json:"result"`
	Code string `json:"code,omitempty"`
	Err  string `json:"error,omitempty"`
}

func (c *CreateListResponse) Error() string {
	return c.Code
}

// IF code=invalid_arg - list already created
func (c Client) CreateList(ctx context.Context, apiKey string, listTitle string) (uint64, error) {
	reqString := fmt.Sprintf("%s/ru/api/createList?format=json&api_key=%s&title=%s", UnisenderURL, apiKey, listTitle)
	resp, err := c.httpClient.Get(reqString)
	if err != nil {
		return 0, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	respSt := &CreateListResponse{}
	if err := json.Unmarshal(data, respSt); err != nil {
		return 0, err
	}

	if respSt.Result == nil {
		return 0, respSt
	}
	return respSt.Result.ID, nil
}
