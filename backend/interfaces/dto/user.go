package dto

type UserAuthDto struct{
	Provider 	string 	`json:"provider"`
	Id  	 	string 	`json:"id"`
	Name 	 	string 	`json:"name"`
	Email 		string 	`json:"email"`
	AvatarUrl 	string 	`json:"avatarUrl"`
}