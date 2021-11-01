package exception

import "encoding/json"

type ErrorValidation struct {
	Message interface{} `json:"message"`
}

func (e ErrorValidation) Error() string {
	err, _ := json.Marshal(e.Message)
	return string(err)
}
