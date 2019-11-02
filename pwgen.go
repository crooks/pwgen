package main

import (
	"bufio"
	urand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type config struct {
	maxSepLength     int
	maxWordLength    int
	minSepLength     int
	minWordLength    int
	numPasswords     int
	suffixSepLength  int
	symbolsList      string
	wordsFile        string
	wordsPerPassword int
}

var cfg config
var flag_args []string
var flag_filename string
var flag_pwds int
var flag_short bool

// init initializes the command line flags.
func init() {
	flag.BoolVar(&flag_short, "short", false, "Generate shorter passwords")
	flag.IntVar(&flag_pwds, "passwords", 5, "Number of passwords to generate")
	flag.StringVar(&flag_filename, "file", "words.txt", "Words file to populate the generator with")
	flag.Parse()
	flag_args = flag.Args()
}

// setConfig defines some rational defaults and overrides them when required.
func setConfig() {
	// Set some reasonable defaults
	cfg.minSepLength = 3
	cfg.maxSepLength = 3
	cfg.minWordLength = 6
	cfg.maxWordLength = 9
	cfg.numPasswords = flag_pwds
	cfg.suffixSepLength = 0
	cfg.symbolsList = "123456789!$%*@"
	cfg.wordsFile = flag_filename
	cfg.wordsPerPassword = 4
	if flag_short {
		cfg.minSepLength = 2
		cfg.maxSepLength = 2
		cfg.minWordLength = 5
		cfg.maxWordLength = 5
		cfg.suffixSepLength = 2
		cfg.symbolsList = "#_123456789"
	}
	// Time for some sanity checking
	if cfg.minWordLength > cfg.maxWordLength {
		log.Fatalln("Minimum word length exceeds maximum")
	}
	if cfg.minSepLength > cfg.maxSepLength {
		log.Fatalln("Minimum seperator length exceeds maximum")
	}
	if cfg.wordsPerPassword < 1 {
		log.Fatalln("Cannot specify less than one word per password")
	}
}

// The following section defines crypto/rand as a source for functions in
// math/rand.  This means we can use many of the math/rand functions with
// a cryptographically random source.
type CryptoRandSource struct{}

func NewCryptoRandSource() CryptoRandSource {
	return CryptoRandSource{}
}

func (_ CryptoRandSource) Int63() int64 {
	var b [8]byte
	urand.Read(b[:])
	return int64(binary.LittleEndian.Uint64(b[:]) & (1<<63 - 1))
}

func (_ CryptoRandSource) Seed(_ int64) {}

// And so ends the random magic section

// randomInt returns an integer between 0 and max
func randomInt(max int) int {
	r := rand.New(NewCryptoRandSource())
	return r.Intn(max)
}

// readLines reads a text file and stores each line as a slice item.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= cfg.minWordLength && len(line) <= cfg.maxWordLength {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

// writeLines writes a slice as a text file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// randWord returns a random element of a given slice
func randWord(wordList []string) (word string) {
	listLen := len(wordList)
	word = strings.Title(wordList[randomInt(listLen)])
	return
}

// Generate a random seperator string
func seperator(minSepLen int, maxSepLen int, symbols []string) (sep string) {
	sepMax := len(symbols)
	var sepLen int
	// If min and max lengths are the same, there's no point generating a
	// random number between n and n.
	if minSepLen == maxSepLen {
		sepLen = minSepLen
	} else {
		sepLen = randomInt(maxSepLen-minSepLen+1) + minSepLen
	}
	for i := 0; i < sepLen; i++ {
		sep += symbols[randomInt(sepMax)]
	}
	return
}

func main() {
	setConfig()
	symbols := strings.Split(cfg.symbolsList, "")
	words, err := readLines(cfg.wordsFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for p := 0; p < cfg.numPasswords; p++ {
		var password string
		for i := 0; i < cfg.wordsPerPassword; i++ {
			password += randWord(words)
			password += seperator(cfg.minSepLength, cfg.maxSepLength, symbols)
		}
		if cfg.suffixSepLength > 0 {
			password += seperator(cfg.suffixSepLength, cfg.suffixSepLength, symbols)
		}
		fmt.Println(password)
	}
}
