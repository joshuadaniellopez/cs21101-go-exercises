package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	quizzes := flag.String("csv", "problems.csv", "database")

	var count int = 0
	flag.IntVar(&count, "n", 5, "number of questions")
	flag.Parse()

	if count <= 0 {
		fmt.Println("Invalid count. Setting a default value of 10. Good luck with your quiz!")
		count = 10
	}

	fp, err := os.Open(*quizzes)
	if err != nil {
		log.Fatal("%v", err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	lines, _ := r.ReadAll()

	correct_sum := 0

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		unique_index := rand.Intn(len(lines) - 1)
		item := lines[unique_index]

		var ans int
		fmt.Printf("%s = ", item[0])
		_, err := fmt.Scanf("%d", &ans)
		if err != nil {
			log.Fatal("Invalid Input. Integer Input is only accepted.")
		}

		correct_answer, err := strconv.Atoi(strings.TrimSpace(item[1]))

		if ans == correct_answer {
			fmt.Println("Correct!")
			correct_sum++
		} else {
			fmt.Printf("Wrong! Correct Answer is (%d)\n", correct_answer)
		}
	}
	fmt.Printf("You have answered %d out of %d questions correctly.\n", correct_sum, count)
	if correct_sum > count/2 {
		fmt.Printf("Congratulations!\n")
	} else {
		fmt.Printf("You need to do better! Keep Studying!!\n")
	}
}
