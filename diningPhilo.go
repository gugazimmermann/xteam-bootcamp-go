package main

import (
	"fmt"
	"sync"
)

// ChopS - the chopstick representation
type ChopS struct {
	mutex sync.Mutex
}

// Philo - the philosopher representation
type Philo struct {
	id             int
	leftChopstick  *ChopS
	rightChopstick *ChopS
	eating         bool
}

// Host - the host representation
type Host struct {
	concurrent chan int
	philo      []Philo
}

func (host *Host) grantPermission(i int) bool {
	nDiners := len(host.philo)
	next := host.philo[(i+1)%nDiners].eating
	prev := host.philo[modFloor(i-1, nDiners)].eating
	return !(host.philo[i].eating || next || prev)

}

func (p *Philo) eat() {
	p.eating = true
	p.leftChopstick.mutex.Lock()
	p.rightChopstick.mutex.Lock()
	fmt.Printf("starting to eat %d\n", p.id)
	fmt.Printf("finishing eating %d\n", p.id)
	p.eating = false
	p.leftChopstick.mutex.Unlock()
	p.rightChopstick.mutex.Unlock()
}

func main() {

	nPhilosophers := 5
	nTimesToEat := 3
	maxConcurrent := 2

	chops := make([]ChopS, nPhilosophers)
	for i := 0; i < nPhilosophers; i++ {
		chops[i] = ChopS{
			mutex: sync.Mutex{},
		}
	}

	philo := make([]Philo, nPhilosophers)
	for i := 0; i < nPhilosophers; i++ {
		philo[i] = Philo{
			id:             i + 1,
			leftChopstick:  &chops[i],
			rightChopstick: &chops[(i+1)%nPhilosophers],
		}
	}

	host := Host{
		concurrent: make(chan int, 2),
		philo:      philo,
	}

	done := make(chan int)

	for conc := 0; conc < maxConcurrent; conc++ {
		go func() {
			for i := range host.concurrent {
				if host.grantPermission(i) {
					philo[i].eat()
					done <- i
				} else {
					host.concurrent <- i
				}
			}
		}()
	}

	go func() {
		for m := 0; m < nTimesToEat; m++ {
			for d := 0; d < nPhilosophers; d++ {
				host.concurrent <- d
			}
		}
	}()

	for d := 0; d < (nTimesToEat * nPhilosophers); d++ {
		<-done
	}
	close(host.concurrent)

}

func modFloor(a int, n int) int {
	return ((a % n) + n) % n
}
