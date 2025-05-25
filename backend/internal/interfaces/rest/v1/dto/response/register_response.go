package response

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type RegisterResponse struct {
	AccessToken string 			`json:"accessToken"`
	User   	    common.BriefUserResult `json:"user"`
}