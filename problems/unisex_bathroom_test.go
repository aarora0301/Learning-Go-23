package problems

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// A bathroom is being designed for the use of both males and females in an office but requires the following
//constraints to be maintained:
//
//There cannot be men and women in the bathroom at the same time.
//There should never be more than three employees in the bathroom simultaneously.
//The solution should avoid deadlocks. For now, though, don’t worry about starvation.

type gender string

const (
	WOMEN gender = "women"
	MEN   gender = "men"
	NONE  gender = "none"
)

type UnisexBathroom struct {
	inUseBy      gender
	maxEmps      chan struct{}
	bathRoomCond *sync.Cond
	waitGroup    sync.WaitGroup
}

func NewUnisexBathroom(n int) *UnisexBathroom {
	b := &UnisexBathroom{
		inUseBy:      NONE,
		maxEmps:      make(chan struct{}, n),
		bathRoomCond: sync.NewCond(&sync.Mutex{}),
	}
	return b
}

func (u *UnisexBathroom) useBathroom(name string) {
	fmt.Printf("\n%s using bathroom. Current employees in bathroom = %d %d\n", name, len(u.maxEmps), time.Now().UnixNano())
	time.Sleep(3 * time.Second)
	fmt.Printf("\n%s done using bathroom %d\n", name, time.Now().UnixNano())
}

func (u *UnisexBathroom) maleUseBathroom(name string) {
	u.bathRoomCond.L.Lock()

	for u.inUseBy == WOMEN && len(u.maxEmps) > 0 {
		u.bathRoomCond.Wait()
	}

	u.inUseBy = MEN
	u.bathRoomCond.L.Unlock()

	u.maxEmps <- struct{}{}
	fmt.Println("maleUseBathroom locked by ", name)
	u.useBathroom(name)

	fmt.Println("maleUseBathroom unlocked by ", name)

	<-u.maxEmps // release the resource

	u.bathRoomCond.L.Lock()
	if len(u.maxEmps) == 0 {
		u.inUseBy = NONE
		u.bathRoomCond.Broadcast()
	}
	u.bathRoomCond.L.Unlock()

	u.waitGroup.Done()
}

func (u *UnisexBathroom) femaleUseBathroom(name string) {
	u.bathRoomCond.L.Lock()
	for u.inUseBy == MEN && len(u.maxEmps) > 0 {
		u.bathRoomCond.Wait()
	}

	u.inUseBy = WOMEN
	u.bathRoomCond.L.Unlock()

	u.maxEmps <- struct{}{}
	fmt.Println("femaleUseBathroom locked by ", name)
	u.useBathroom(name)
	fmt.Println("femaleUseBathroom unlocked by ", name)
	<-u.maxEmps // release the resource

	u.bathRoomCond.L.Lock()
	if len(u.maxEmps) == 0 {
		u.inUseBy = NONE
		u.bathRoomCond.Broadcast()
	}
	u.bathRoomCond.L.Unlock()

	u.waitGroup.Done()
}

func TestUsageFlow(t *testing.T) {
	unisexBathroom := NewUnisexBathroom(3) //  // only 3 people can use the bathroom at a time
	fmt.Println("len", len(unisexBathroom.maxEmps))

	unisexBathroom.waitGroup.Add(6)

	go func() {
		unisexBathroom.femaleUseBathroom("Lisa")
	}()
	go func() {
		unisexBathroom.femaleUseBathroom("Ray")
	}()
	go func() {
		unisexBathroom.maleUseBathroom("John")
	}()
	go func() {
		unisexBathroom.maleUseBathroom("Bob")
	}()
	go func() {
		unisexBathroom.maleUseBathroom("Anil")
	}()
	go func() {
		unisexBathroom.maleUseBathroom("Wentao")
	}()

	unisexBathroom.waitGroup.Wait()
}
