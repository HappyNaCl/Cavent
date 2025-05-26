package response

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type CheckMeResponse struct {
	User common.BriefUserResult `json:"user"`
}