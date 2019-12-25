package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	log.Println("Start (typ 'stop' om te stoppen)")
	rand.Seed(time.Now().Unix())

	sc := bufio.NewScanner(os.Stdin)

	outer: for {
		som, expected := generateSom()
		log.Printf("Wat is %s?\n", som)

		for sc.Scan() {
			if err := sc.Err(); err != nil {
				log.Fatal(err)
			}
			s := sc.Text()

			if s == "stop" {
				break outer
			}

			answer, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}

			if answer == expected {
				log.Printf("Goed gedaan!\n")
				break
			} else {
				log.Printf("Helaas, probeer het nog eens?\n")
			}
		}
	}

	log.Println("Done")
}

func generateSom() (string, int) {
	a := rand.Int() % 50
	b := rand.Int() % 50

	return fmt.Sprintf("%d + %d", a, b), a+b
}
