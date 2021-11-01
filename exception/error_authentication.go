package exception

import "encoding/json"

type ErrorAuthentication struct {
	Message interface{} `json:"message"`
}

func (e ErrorAuthentication) Error() string {
	err, _ := json.Marshal(e.Message)
	return string(err)
}
