package parse

import (
	"net/url"
	"strings"
)

func ParseForm(body string) (map[string]string, error) {
	form := make(map[string]string)
	pairs := strings.Split(body, "&")
	for _, p := range pairs {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key, err1 := url.QueryUnescape(kv[0])
		val, err2 := url.QueryUnescape(kv[1])
		if err1 != nil || err2 != nil {
			continue
		}
		form[key] = val
	}
	return form, nil
}
