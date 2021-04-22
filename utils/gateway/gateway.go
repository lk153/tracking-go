package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

//DefaultGateMuxOpts ...
func DefaultGateMuxOpts() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
					UseEnumNumbers:  true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
		runtime.WithErrorHandler(ErrorHandler),
	}
}

//ErrorHandler ...
func ErrorHandler(ctx context.Context,
	mux *runtime.ServeMux,
	m runtime.Marshaler,
	resp http.ResponseWriter,
	req *http.Request, err error) {
	if err == nil {
		resp.WriteHeader(http.StatusNoContent)
		return
	}
	runtime.DefaultHTTPErrorHandler(ctx, mux, m, resp, req, err)
}
