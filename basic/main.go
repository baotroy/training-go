// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type user struct {
	name                 string
	number               string
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	fmt.Printf("existing user is %v\n", existingUser)
	fmt.Printf("Ok is %v\n", ok)
	if !ok {
		return false, errors.New("user does not exist")
	}

	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

func countDistinctWords(messages []string) int {
	var counter = make(map[string]bool)
	for _, message := range messages {
		words := strings.Split(message, " ")
		for _, word := range words {
			if _, existed := counter[strings.ToLower(word)]; !existed {
				counter[strings.ToLower(word)] = true
			}
		}
	}
	return len(counter)
}

func getNameCounts(names []string) map[rune]map[string]int {
	count := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		if _, exists := count[firstChar]; !exists {
			count[firstChar] = make(map[string]int)
		}
		count[firstChar][name]++
	}
	return count
}

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		// fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Println("Working on", n)
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	const c = 20
	jobs := make(chan int, c)
	results := make(chan int, c)
	start := time.Now()
	go worker(jobs, results)

	for i := 0; i < c; i++ {
		fmt.Printf("Sending job %d\n", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("Jobs sent")
	for j := 0; j < c; j++ {
		fmt.Println(<-results)
	}
	elapsed := time.Since(start)
	fmt.Printf("time %d", elapsed)
	// go count("sheep", c)

	// for msg := range c {
	// 	// msg, open := <-c
	// 	// if !open {
	// 	// 	break
	// 	// }
	// 	fmt.Println(msg)
	// }
	// go count("sheep")
	// go count("cow")

	// start := time.Now()
	// for i := 0; i < 1; i++ {

	// 	const privateKey2 = ""
	// 	address := cryp.PrivateKeyToAddress(privateKey2)
	// 	fmt.Println(strings.ToLower(address.String()))
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("Time taken %s\n", elapsed)

	// fmt.Printf(`%s %d`, fromAddress, len(fromAddress))
	// fmt.Print(helper.Greeting("abc"))
	// users := map[string]user{
	// 	"John": {
	// 		name:                 "John",
	// 		number:               "1234567890",
	// 		scheduledForDeletion: true,
	// 	},
	// 	"Jane": {
	// 		name:                 "Jane",
	// 		number:               "0987654321",
	// 		scheduledForDeletion: false,
	// 	},
	// }
	// user, ok := users["John1"]
	// data := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// }
	// user, ok := data["c"]

	// // res, err := deleteIfNecessary(users, "John1")
	// fmt.Printf("Result is %v\n", user)
	// fmt.Printf("Error is %v\n", ok)
	// messages := []string{"Hello world", "hello there", "General Kenobi"}
	// count := countDistinctWords(messages)
	// fmt.Print(count)
	// m := map[string]map[string]int{
	// 	"b": {
	// 		"billy": 2,
	// 		"bob":   1,
	// 	},
	// 	"j": {
	// 		"joe": 1,
	// 	},
	// }
	// fmt.Print(m)
	// names := []string{
	// 	"billy", "billy", "bob", "joe",
	// }
	// c := getNameCounts(names)
	// fmt.Print(c)
}
