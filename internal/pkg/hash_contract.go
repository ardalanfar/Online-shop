package pkg

//Hash Contract (interface)

type Hash interface {
	HashPassword(password string) (string, error)
}
