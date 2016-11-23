package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

type alphabetType uint

const (
	lettersAlphabet alphabetType = iota
	lettersDigitsAlphabet
	lettersDigitsPuctuationAlphabet
	digitsAlphabet

	numberOfPasswords = 10
)

var (
	nchars uint
	alpha  uint

	letters     = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	digits      = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	punctuation = []rune{'!', '"', '#', '$', '%', '&', '`', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
)

func buildAlphabet(t alphabetType) []rune {
	var alphabet []rune
	switch t {
	case lettersAlphabet:
		alphabet = letters
	case lettersDigitsAlphabet:
		alphabet = append(alphabet, letters...)
		alphabet = append(alphabet, digits...)
	case lettersDigitsPuctuationAlphabet:
		alphabet = append(alphabet, letters...)
		alphabet = append(alphabet, digits...)
		alphabet = append(alphabet, punctuation...)
	case digitsAlphabet:
		alphabet = digits
	}
	return alphabet
}

func genPassword(a []rune, nchars uint) <-chan string {
	out := make(chan string)
	go func() {
		alphabetLeght := big.NewInt(int64(len(a)))
		for j := 0; j < numberOfPasswords; j++ {
			var password []rune
			for i := uint(0); i < nchars; i++ {
				idx, err := rand.Int(rand.Reader, alphabetLeght)
				if err != nil {
					panic(err)
				}
				password = append(password, a[idx.Int64()])
			}
			out <- string(password)
		}
		close(out)
	}()
	return out
}

func init() {
	flag.UintVar(&alpha, "a", 1, "alphabet used to generate the password, options ASCII letters only: 0, ASCII letters and numbers: 1, ASCII letters, numbers and punctuation: 2, numbers only: 3.")
	flag.UintVar(&nchars, "n", 32, "length of the password")
	flag.Parse()
}

func main() {
	t := alphabetType(alpha)
	if t > digitsAlphabet {
		fmt.Println("invalid alphabet")
		flag.PrintDefaults()
		return
	}
	if nchars > 512 {
		fmt.Println("invalid password length, max length 512 chars")
		flag.PrintDefaults()
		return
	}
	c := genPassword(buildAlphabet(t), nchars)
	for p := range c {
		fmt.Printf("%s\n", p)
	}
}
