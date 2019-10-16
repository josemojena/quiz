package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	question := make([]string, 0)
	answer := make([]int, 0)
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		question = append(question, parts[0])
		answer = append(answer, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalQuestion := len(question)
	correctAnswer := 0
    monitor := make(chan bool)
	go func(question []string) {
		for i := 0; i < len(question); i++ {
			var value int
			fmt.Println(question[i])
			fmt.Scanln(&value)

			if value == answer[i] {
				correctAnswer++
			}
		}
		monitor <- true
	}(question)

	select {
    case <- monitor:
    	break
	case <-time.After(30 * time.Second):
		fmt.Println("Time over")
		break
	}
	//close(doing)
	fmt.Printf("Total questions: %d\n", totalQuestion)
	fmt.Printf("Correct answer: %d\n", correctAnswer)

}
