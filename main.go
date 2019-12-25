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
	logger := log.New(os.Stdout, "", 0)
	logger.Print("Start (typ 'stop' om te stoppen)")
	rand.Seed(time.Now().Unix())

	sc := bufio.NewScanner(os.Stdin)

	outer: for {
		som, expected := generateSom()

		for {
			logger.Printf("%s = ", som)

			sc.Scan()
			if err := sc.Err(); err != nil {
				logger.Fatal(err)
			}
			s := sc.Text()

			if s == "stop" {
				break outer
			}

			answer, err := strconv.Atoi(s)
			if err != nil {
				logger.Print("Daar ging iets mis, probeer het nog eens:")
				continue
			}

			if answer == expected {
				logger.Print("Goed gedaan!")
				break
			} else {
				logger.Print("Helaas, probeer het nog eens:")
			}
		}
	}

	logger.Print("Gestopt.")
}

func generateSom() (string, int) {
	a := rand.Int() % 50
	b := rand.Int() % 50

	return fmt.Sprintf("%d + %d", a, b), a+b
}
