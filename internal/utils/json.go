package utils

import (
	"encoding/json"
)

// ToJSON will convert interface to json string
func ToJSON(s interface{}) string {
	buf, _ := json.Marshal(s)
	return string(buf)
}
