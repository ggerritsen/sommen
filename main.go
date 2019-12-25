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

outer:
	for {
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

type operator struct {
	op              string
	compute         func(int, int) int
	generateOperand func() int
}

var operators = []operator{
	{"+", func(a, b int) int { return a + b }, func() int { return rand.Int() % 51 }},
	{"-", func(a, b int) int { return a - b }, func() int { return rand.Int() % 51 }},
	{"x", func(a, b int) int { return a * b }, func() int { return rand.Int() % 11 }},
}

func generateSom() (string, int) {
	o := operators[rand.Int() % 3]
	a := o.generateOperand()
	b := o.generateOperand()

	// exclude negative answers
	if o.op == "-" {
		if a < b {
			tmp := a
			a = b
			b = tmp
		}
	}

	return fmt.Sprintf("%d %s %d", a, o.op, b), o.compute(a, b)
}
