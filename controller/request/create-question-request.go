package request

import (
	"mime/multipart"
)

type CreateQuestionRequest struct {
	Text    string
	Channel string
	Image   MultipartFile
}

type MultipartFile struct {
	Header *multipart.FileHeader
}

func ParseCreateQuestionRequest(form *multipart.Form) CreateQuestionRequest {
	header := form.File["file"][0]

	return CreateQuestionRequest{
		Text:    form.Value["text"][0],
		Channel: "testing",
		Image:   MultipartFile{header},
	}
}
