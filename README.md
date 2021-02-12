# Coursera - Programming with Google Go Specialization
https://www.coursera.org/specializations/google-golang

## Part 3 - Concurrency in Go
https://www.coursera.org/learn/golang-concurrency


## Module 4 Activity: Dining Philosophers 

![Dining Philosopher](dining-philosophers.png)

Implement the dining philosopher’s problem with the following constraints/modifications.

1) There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2) Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3) The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4) In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5) The host allows no more than 2 philosophers to eat concurrently.
6) Each philosopher is numbered, 1 through 5.
7)  When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8) When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
