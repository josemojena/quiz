package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {




	question := make([]string, 0)
	answer := make([]int, 0)
	file, err := os.Open("problems.csv")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	//Using golang examples
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




	for i:= 0; i< len(question) ; i++  {

		var value  int

		fmt.Println(question[i])
		fmt.Scanln(&value)

		if value == answer[i] {
			correctAnswer ++
		}
	}

	fmt.Printf("Total questions: %d\n", totalQuestion)
	fmt.Printf("Correct answer: %d\n" , correctAnswer)


}
