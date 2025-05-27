package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
)

func ReadMultipartFile(file multipart.File, header *multipart.FileHeader) ([]byte, string, error) {
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, "", err
	}

	if buf.Len() >= 5*1024*1024 { // 5MB limit
		return nil, "", errors.ErrBannerMaxSize
	}

	ext := filepath.Ext(header.Filename)
	if (ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp") {
		return nil, "", errors.ErrInvalidBannerFormat
	}

	return buf.Bytes(), ext, nil
}