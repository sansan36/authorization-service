package core

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	commonv1 "github.com/sansan36/authorization-service/gen/common/v1"
	"google.golang.org/grpc/status"
)

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	errorCode := runtime.HTTPStatusFromCode(status.Code(err))
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(errorCode)

	errMessage := status.Convert(err).Message()
	if errorCode >= 500 {
		errMessage = "internal server error"
	}

	body := &commonv1.ErrorResponse{
		Message: errMessage,
		HttpStatus: &commonv1.StandardResponse{
			Status:       "error",
			Code:         uint64(errorCode),
			ErrorMessage: status.Code(err).String(),
		},
	}

	jErr := json.NewEncoder(w).Encode(body)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
