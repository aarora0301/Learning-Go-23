package patterns

import (
	"context"
	"errors"
)

// Implement concurrency pool

type Pool struct {
	ch chan struct{}
}

func (p *Pool) release() {
	p.ch <- struct{}{}
}

func (p *Pool) GetLease() {
	<-p.ch
}

func initPool(size int) *Pool {
	p := &Pool{
		ch: make(chan struct{}, size),
	}

	for i := 0; i < size; i++ {
		p.release()
	}
	return p
}

func (p *Pool) GetLeaseWithCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("context done")

	default:
		// to not block on above channel

	}

	select {
	case <-ctx.Done():
		return errors.New("context don now")
	case <-p.ch:
		return nil

	}
}

// this may not work in some cases as select randomly picks the channel if multiple channels are ready
//
//	the order of case does not determine the priority order
//	in our case we first want to see , if context is done/canel , then dont serve request
//
// so to prioritize contect `GetLeaseWithCtx` method is used
// first see if context is done/cancelled
// if not again check and then release connection
func (p *Pool) GetLeaseWithCtx2(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("context don now")
	case <-p.ch:
		return nil

	}
}
