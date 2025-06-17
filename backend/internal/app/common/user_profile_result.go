package common

type UserProfileResult struct {
	Id           string  `json:"id"`
	Name 	   	 string  `json:"name"`
	AvatarUrl    string  `json:"avatarUrl"`
	Description  *string `json:"description"`
	PhoneNumber  *string `json:"phoneNumber"`
	Address      *string `json:"address"`
}