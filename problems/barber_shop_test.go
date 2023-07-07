package problems

import "sync"

type state string

const (
	sleeping state = "sleeping"
	ready    state = "ready"
	checking state = "checking"
	cutting  state = "cutting"
)

type barbershop struct {
	name string
}

type barber struct {
	name  string
	state state
	mu    sync.Mutex
}
type customer struct {
	customerId int
}
