package body

import (
	"encoding/json"

	"github.com/datshiro/crud/internal/infras/errors"
)

// func ReadRequestBody(req *http.Request) string {
// 	buf := new(strings.Builder)
// 	_, err := io.Copy(buf, req.Body)
// 	if err != nil {
// 		return ""
// 	}
// 	return buf.String()
// }

func BindRequest(body []byte, model interface{}) error {
	if len(body) == 0 {
		return errors.InvalidParamError
	}
	if err := json.Unmarshal(body, model); err == nil {
		return err
	}
	return errors.InvalidParamError
}
