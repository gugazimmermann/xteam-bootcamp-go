# Coursera - Programming with Google Go Specialization
https://www.coursera.org/specializations/google-golang

## Part 3 - Concurrency in Go
https://www.coursera.org/learn/golang-concurrency


## Module 3 Activity: Threads 

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.