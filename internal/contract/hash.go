package contract

//Hash Contract (interface)

type Hash interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
