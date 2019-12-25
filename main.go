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
	log.Println("Start")
	sc := bufio.NewScanner(os.Stdin)

	som, expected := generateSom()
	log.Printf("Wat is %s?\n", som)
	for sc.Scan() {
		s := sc.Text()
		answer, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		if answer == expected {
			log.Printf("Goed gedaan!\n")
		} else {
			log.Printf("Helaas, volgende keer beter\n")
		}

		som, expected = generateSom()
		log.Printf("Wat is %s?\n", som)
	}

	log.Println("Done")
}

func generateSom() (string, int) {
	rand.Seed(time.Now().Unix())
	a := rand.Int() % 50
	b := rand.Int() % 50

	return fmt.Sprintf("%d + %d", a, b), a+b
}
