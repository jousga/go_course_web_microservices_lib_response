package response

import (
	"encoding/json"
	"github.com/jousga/go_course_web_microservices_meta/meta"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

func OK(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInfo(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusResetContent)
}

func success(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
		Meta:    meta,
	}
}

func (s SuccessResponse) StatusCode() int {
	return s.Status
}

func (s SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
