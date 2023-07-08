package problems

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

type state string

const (
	sleeping state = "sleeping"
	ready    state = "ready"
	checking state = "checking"
	cutting  state = "cutting"
)

type barbershop struct {
	name         string
	barber       *barber
	roomSize     int
	customerChan chan *customer
	cuttingTime  time.Duration
	wakeSig      chan int
	mu           sync.Mutex
}

type barber struct {
	name  string
	state state
	mu    sync.Mutex
}
type customer struct {
	customerId int
	wg         *sync.WaitGroup
}

func (s *barbershop) reception(cus *customer) {
	log.Printf("Customer [%d] arrived. [%s] is [%s]... WR: [%d]\n",
		cus.customerId, s.barber.name, s.barber.getState(), len(s.customerChan))

	switch s.barber.getState() {
	case sleeping:
		// send wake signal
		select {
		case s.wakeSig <- cus.customerId:
			fmt.Println("sent wake signal to barber")
		default:
			log.Println("cant wake barber")
		}
	}

	// for all states
	select {
	case s.customerChan <- cus:
		log.Printf("Customer [%d] is waiting... [%s] is [%s]... WR: [%d]\n",
			cus.customerId, s.barber.name, s.barber.getState(), len(s.customerChan))
	default:
		log.Printf("Customer [%d] is leaving :(. [%s] is [%s]... WR: [%d] full\n",
			cus.customerId, s.barber.name, s.barber.getState(), len(s.customerChan))
		cus.wg.Done()
	}
}

func (s *barbershop) work() {
	log.Printf("[%s] started working with [%s]\n", s.name, s.barber.name)

	for {
		//s.barber.setState(checking)
		log.Printf("[%s] %s for customers...\n", s.barber.name, s.barber.state)

		select {
		case c := <-s.wakeSig:
			s.barber.setState(ready)
			log.Printf("[%s] wokened by customer [%d]\n", s.barber.name, c)
		case cus := <-s.customerChan:
			log.Printf("[%s] having state [%s] taking customer [%d]\n", s.barber.name, s.barber.getState(), cus.customerId)
			// haircut
			s.barber.haircut(cus, s.cuttingTime)
		default:
			s.barber.setState(sleeping)
			log.Printf("customers in the waiting room [%d]\n", len(s.customerChan))
			log.Printf("[%s] %s 😴Zzzzzzzzzzzzzzz...\n", s.barber.name, s.barber.state)
			c := <-s.wakeSig
			s.barber.setState(ready)
			log.Printf("[%s] wokened by customer [%d]\n", s.barber.name, c)
		}
	}
}

func (b *barber) getState() state {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}

func (b *barber) setState(state state) {
	b.mu.Lock()
	b.state = state
	b.mu.Unlock()
}

func (b *barber) haircut(cus *customer, t time.Duration) {
	b.mu.Lock()
	b.state = cutting
	log.Printf("[%s] %s for customer [%d]\n", b.name, b.state, cus.customerId)
	time.Sleep(t)
	b.mu.Unlock()

	// finish customer
	log.Printf("[%s] %s finished customer [%d]\n", b.name, b.state, cus.customerId)
	cus.wg.Done()
}

func TestBarberShop(t *testing.T) {
	waitingRoomSize := 5
	totalCutomer := 10
	barber1 := &barber{
		name:  "RAY",
		state: sleeping,
		mu:    sync.Mutex{},
	}
	shop := &barbershop{
		name:         "CUT & STYLE",
		barber:       barber1,
		roomSize:     waitingRoomSize,
		customerChan: make(chan *customer, waitingRoomSize),
		wakeSig:      make(chan int, 1),
		cuttingTime:  100 * time.Millisecond,
		mu:           sync.Mutex{}, // when working with mutiple barbers
	}
	// start the shop working
	go shop.work()
	var wg sync.WaitGroup
	wg.Add(totalCutomer)
	for i := 1; i <= totalCutomer; i++ {
		go shop.reception(&customer{
			customerId: i,
			wg:         &wg,
		})
	}
	wg.Wait()
	log.Println("done with all customers")
	time.Sleep(2 * time.Second)
	//now := time.Now()
	//next := now.Add(500 * time.Millisecond)
	//result := (next.UnixMilli() - now.UnixMilli()) * 10 / 1e9
}
