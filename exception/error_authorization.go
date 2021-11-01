package exception

import "encoding/json"

type ErrorAuthorization struct {
	Message interface{} `json:"message"`
}

func (e ErrorAuthorization) Error() string {
	err, _ := json.Marshal(e.Message)
	return string(err)
}
