package config

import (
	"fmt"
	"mime/multipart"
)

func uploadToSupabase(file multipart.File, fileName string, folder string, contentType string){
	fileName = fmt.Sprintf("%s/%s", folder, fileName)

	// uploadUrl = 
}