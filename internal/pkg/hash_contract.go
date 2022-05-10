package pkg

//Hash Contract (interface)

type Hash interface {
	HashPassword(password string) (string, error)
	DecodePassword(password string) (string error)
}
