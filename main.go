package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	maskLetters = "abcdefghijklmnopqrstuvwxyz"
	maskDigits  = "0123456789"
	maskAlph    = maskLetters + maskDigits
	maskSpecial = "-"
	maskAll     = maskSpecial + maskAlph
)

func main() {

	masksFile := flag.String("m", "", "Path to the file containing masks")
	wordlistFile := flag.String("w", "", "Path to the file containing the wordlist")
	outputFile := flag.String("o", "", "Path to the output file")
	stdout := flag.Bool("stdout", false, "Print output to stdout instead of a file")
	flag.Parse()

	var writer *bufio.Writer
	if *stdout {
		writer = bufio.NewWriter(os.Stdout)
	} else {

		output, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer output.Close()
		writer = bufio.NewWriter(output)
	}
	defer writer.Flush()

	masks, err := readLines(*masksFile)
	if err != nil {
		fmt.Println("Error reading masks file:", err)
		return
	}

	wordlist, err := readLines(*wordlistFile)
	if err != nil {
		fmt.Println("Error reading wordlist file:", err)
		return
	}

	err = GenerateWordlist(masks, wordlist, writer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !*stdout {
		fmt.Printf("Output saved to %s\n", *outputFile)
	}
}

func GenerateWordlist(masks []string, wordlist []string, writer *bufio.Writer) error {

	for _, mask := range masks {
		err := ExpandMask(mask, wordlist, writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExpandMask(word string, wordlist []string, writer *bufio.Writer) error {
	var chars string

	if strings.Count(word, "?") > 8 {
		return fmt.Errorf("Exceeded maximum mask size (8): %s", word)
	}

	parts := strings.SplitN(word, "?", 2)
	if len(parts) > 1 {
		if len(parts[1]) > 0 {
			switch parts[1][0] {
			case 'l':
				chars = maskLetters
			case 'd':
				chars = maskDigits
			case 'a':
				chars = maskAlph
			case 'A':
				chars = maskAll
			case 's':
				chars = maskSpecial
			case 'w':
				for _, w := range wordlist {
					newWord := parts[0] + w + parts[1][1:]
					err := ExpandMask(newWord, wordlist, writer)
					if err != nil {
						return err
					}
				}
				return nil
			default:
				return fmt.Errorf("Improper mask used: %s", word)
			}
			for _, ch := range chars {
				newWord := parts[0] + string(ch) + parts[1][1:]
				err := ExpandMask(newWord, wordlist, writer)
				if err != nil {
					return err
				}
			}
		}
	} else {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
