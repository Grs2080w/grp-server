package downloadFile

// "github.com/Grs2080w/grp_server/core/domains/download/downloadFile/"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	URL string `json:"url"`
}