package unisender

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
)

type ImportContactsRequest struct {
	Format     string
	ApiKey     string
	FieldNames []string
	Data       [][]string
}

func (u ImportContactsRequest) Encode() string {
	uv := url.Values{}
	uv.Set("format", u.Format)
	uv.Set("api_key", u.ApiKey)
	for i, v := range u.FieldNames {
		uv.Add(fmt.Sprintf("field_names[%d]", i), v)
	}
	for i, arr := range u.Data {
		for j, v := range arr {
			uv.Add(fmt.Sprintf("data[%d][%d]", i, j), v)
		}
	}
	return uv.Encode()
}

type ImportContactsResponse struct {
	Err      string  `json:"error,omitempty"`
	Code     string  `json:"code,omitempty"`
	Result   *Result `json:"result,omitempty"`
	Warnings []struct {
		Warning string `json:"warning,omitempty"`
	} `json:"warnings"`
}

type Result struct {
	Total     int   `json:"total,omitempty"`
	Inserted  int   `json:"inserted,omitempty"`
	Updated   int   `json:"updated,omitempty"`
	Deleted   int   `json:"deleted,omitempty"`
	NewEmails int   `json:"new_emails,omitempty"`
	Invalid   int   `json:"invalid,omitempty"`
	Log       []Log `json:"log"`
}
type Log struct {
	Index   string `json:"index,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func LogsToMap(res *Result) map[int]string {
	if res == nil {
		return nil
	}

	m := make(map[int]string)
	for _, l := range (*res).Log {
		num, err := strconv.Atoi(l.Index)
		if err != nil {
			continue
		}
		m[num] = l.Code
	}
	return m
}

func (i ImportContactsResponse) Error() string {
	return i.Code
}

func (c Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (ImportContactsResponse, error) {
	if len(req.Data) == 0 {
		log.Println("InportContacts: empty request data")
		return ImportContactsResponse{}, nil
	}

	url := fmt.Sprintf("%s/ru/api/importContacts?%s", UnisenderURL, req.Encode())

	resp, err := c.httpClient.Post(url, "application/json", nil)
	if err != nil {
		return ImportContactsResponse{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return ImportContactsResponse{}, fmt.Errorf("invalid code")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ImportContactsResponse{}, err
	}
	res := ImportContactsResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return ImportContactsResponse{}, err
	}

	if res.Result == nil {
		return ImportContactsResponse{}, res
	}

	return res, nil
}
