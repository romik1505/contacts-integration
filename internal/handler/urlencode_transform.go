package handler

import (
	"github.com/tidwall/sjson"
	"net/url"
	"strconv"
	"strings"
)

func FormUrlEncodeToJSON(values url.Values) (string, error) {
	mp := make(map[string]string)
	for k, v := range values {
		mp[k] = v[0]
	}

	res := ""
	for k, v := range mp {
		s := SplitDotJoin(k)
		num, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			res, err = sjson.Set(res, s, num)
			if err != nil {
				return "", err
			}
		} else {
			res, err = sjson.Set(res, s, v)
			if err != nil {
				return "", err
			}
		}
	}

	return res, nil
}

func SplitDotJoin(s string) string {
	fn := NextKey(s)
	arr := make([]string, 0)

	var v string
	v = fn()

	for v != "" {
		arr = append(arr, v)
		v = fn()
	}
	return strings.Join(arr, ".")
}

func NextKey(s string) func() string {
	str := s
	last := false
	return func() string {
		if last {
			return ""
		}
		open := strings.Index(str, "[")
		closeBr := strings.Index(str, "]")
		if open == -1 || closeBr == -1 {
			last = true
			return str
		}
		curKey := str[0:open]

		nextKey := str[open+1 : closeBr]
		str = nextKey + str[closeBr+1:]
		return curKey
	}
}
