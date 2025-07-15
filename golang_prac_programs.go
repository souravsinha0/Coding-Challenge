package main

import (
	"fmt"
	"sync"
)

// -------------------------------------
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrime(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
	close(ch)
}

func printEven(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range n {
		if i%2 == 0 {
			ch <- i
		}
	}
	close(ch)
}

func printOdd(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range n {
		if i%2 != 0 {
			ch <- i
		}
	}
	close(ch)
}

func runChannelRoutineExample() {
	var wg sync.WaitGroup
	oddCh := make(chan int)
	evenCh := make(chan int)
	primeCh := make(chan int)

	limit := 10

	wg.Add(3)
	go printEven(limit, evenCh, &wg)
	go printOdd(limit, oddCh, &wg)
	go printPrime(limit, primeCh, &wg)

	for c := range oddCh {
		fmt.Println("odd : ", c)
	}
	for c := range evenCh {
		fmt.Println("even : ", c)
	}
	for c := range primeCh {
		fmt.Println("prime : ", c)
	}

	wg.Wait()
}

//---------------------

type Animal interface {
	Speak() string
}

// type dog will implement the interface
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says, Whoop !"
}
func (c Cat) Speak() string {
	return c.Name + " says, Meow !"
}

func MakeAnimalSound(a Animal) {
	fmt.Println(a.Speak())
}

func runinterfaceExample() {
	d := Dog{Name: "Daniel"}
	c := Cat{Name: "Alisa"}
	MakeAnimalSound(d)
	MakeAnimalSound(c)
}

// -------------------
func revString(s string) string {
	runeStr := []rune(s)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

func runRevStringExample() {
	str := "Hello"
	fmt.Println("rev: ", revString(str))
}

// ----------------------------
func revArray(n []int) []int {

	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func runRevArrayExample() {
	inp := []int{1, 2, 3, 4}
	fmt.Println("rev: ", revArray(inp))
}

// -----------------------------
func findConsSeq() {}

// -----------------------------
func isPalindromeStr(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func findPalSubstrings(s string) []string {
	ans := []string{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			if isPalindromeStr(substring) {
				ans = append(ans, substring)
			}
		}
	}
	return ans
}

func findLongestPalindrome(s string) string {
	//start, maxL := 0, 1
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]
			if isPalindromeStr(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}
	return longest
}

func runPalindromeSubstringExample() {
	in := "aabac"
	fmt.Println(findPalSubstrings(in))
}

func runLongestPalindromeExample() {
	in := "aabac"
	fmt.Println(findLongestPalindrome(in))
}

func main() {

	//1. go routine and channel example:
	//runChannelRoutineExample()

	//2. interface example
	//runinterfaceExample()

	//3. rev string example
	//runRevStringExample()

	//4. rev array example
	//runRevArrayExample()

	//5. consecutive sequential char
	//TBD

	//6. palindrome
	//runPalindromeSubstringExample()
	//runLongestPalindromeExample()

	//7. fibonacci series

	//8. armstrong number

}
