package model

import "mime/multipart"

type Question struct {
	Text    string
	Channel Channel
	Image   *multipart.FileHeader
}
