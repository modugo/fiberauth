package fiberauth

import (
	"encoding/base64"
	"strings"
)

func GetLoginFromRefreshToken(token string) string {
	t, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return ""
	}

	split := strings.Split(string(t), "/")
	if len(split) != 2 {
		return ""
	}

	return split[0]
}
