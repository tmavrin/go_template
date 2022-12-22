package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/teltech/logger"
)

// ErrorFromHTTPStatus will use http status code and http.StatusText method to set message.
// It will log error and message
func JSONErrorFromHTTPStatus(log *logger.Log, w http.ResponseWriter, code int) {
	er := errorResponse{
		Code:    code,
		Message: http.StatusText(code),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Code)

	if err := json.NewEncoder(w).Encode(er); err != nil {
		log.Errorf("failed to encode error response: %s", err)
	}
}

// DecodeJSONBody is used as generic decoding method to get request body in JSON format back as an struct
func DecodeJSONBody(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(body)
}

// WriteJSONResponse is used as generic encoding method to write response into response writer
func WriteJSONResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
