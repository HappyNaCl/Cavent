package repo

type UserRepository interface {
	Login(email, password string) (string, error)
}