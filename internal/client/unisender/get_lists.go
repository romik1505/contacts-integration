package unisender

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GetListsResponse struct {
	Result []struct {
		ID    uint64 `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	} `json:"result"`
	Code string `json:"code,omitempty"`
	Err  string `json:"error,omitempty"`
}

func (g *GetListsResponse) Error() string {
	return g.Code
}

func (c Client) GetLists(ctx context.Context, key string) (GetListsResponse, error) {
	url := fmt.Sprintf("%s/ru/api/getLists?format=json&api_key=%s", UnisenderURL, key)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return GetListsResponse{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetListsResponse{}, err
	}

	sResp := GetListsResponse{}

	if err := json.Unmarshal(data, &sResp); err != nil {
		return GetListsResponse{}, err
	}

	if sResp.Result == nil {
		return GetListsResponse{}, &sResp
	}

	return sResp, nil
}
