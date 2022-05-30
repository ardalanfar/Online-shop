package contract

//Encrypt package Contract (interface)

type Encrypt interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
