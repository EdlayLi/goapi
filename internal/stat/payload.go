package stat

// type LinkCreateRequest struct {
// 	Url string `json:"url" validate:"required,url"`
// }

// type LinkUpdateRequest struct {
// 	Url  string `json:"url" validate:"required,url"`
// 	Hash string `json:"hash"`
// }

type GetStatResponse struct {
	Period string `json:"period" `
	Sum    int    `json:"sum " `
}
