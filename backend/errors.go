package backend

type genericErrorResponse struct {
	Errors []string `json:"errors"`
}

func newGenericErrorResponse(errs ...string) *genericErrorResponse {
	return &genericErrorResponse{errs}
}
