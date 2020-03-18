package tlv_solver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Board represents the state of a Sudoku.
type UserInput struct {
	Types	string
	Length	int
	Value	string
}

type UserOutput struct {
	Types	string
	Value	string
}

// GetUserInput takes an string and returns a UserInput with the input data.
// Invalid or malformed inputs will be rejected.
func GetUserInput(source []string) ([]*UserInput, error) {
	if len(source) != 1 {
		return nil, fmt.Errorf("only accept 1 input")
	}

	var stringSource string = source[0]
	input, err := ParseInput(stringSource)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func GetUserInputFromFile(source io.Reader) ([]*UserInput, error) {
	scanner := bufio.NewScanner(source)

	var input []*UserInput
	curRow := 0

	for ; scanner.Scan(); curRow++ {
		// remove all spaces of the current row
		inputRow := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		parseResult, err := ParseInput(inputRow[0])
		if err != nil {
			return nil, err
		}

		input = append(input, parseResult...)
	}
	return input, nil
}

func ParseInput(source string) ([]*UserInput, error) {
	source = strings.Trim(source, "")
	_, matches := CheckSubstrings(source, UPPRCS, REPLCE)

	if matches > 1 {
		// Handle Multiple
		results, err := ParseMultiple(source)
		if err != nil {
			return nil, err
		}

		var input []*UserInput

		for _, result := range results {
			userInput, err := splitStr(result)
			if err != nil {
				return nil, err
			}
			input = append(input, userInput)
		}
		return input, err
	} else {
		// Handle Single
		result, err := ParseSingle(source)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func splitStr(source string) (*UserInput, error) {
	splitSource := strings.Split(source, "-")

	if len(splitSource) != 3 {
		return nil, fmt.Errorf("invalid input")
	}

	strLength, err := strconv.Atoi(splitSource[1])
	if err != nil {
		return nil, err
	}
	return &UserInput{
		Types:  splitSource[0],
		Length: strLength,
		Value:  splitSource[2],
	}, nil
}

func ParseSingle(source string) ([]*UserInput, error) {
	var result []*UserInput
	userInput, err := splitStr(source)
	if err != nil {
		return nil, err
	}
	result = append(result, userInput)
	return result, nil
}

func ParseMultiple(str string) ([]string, error) {
	countHypen := strings.Count(str, "-")
	splitWords := strings.Split(str, "-")
	newWords := make([]string, (2 / 3) * countHypen )
	for _, word := range splitWords {
		if strings.Contains(word, UPPRCS) {
			splits := strings.Split(word, UPPRCS)
			split := splits[0]
			if split == "" {
				newWords = append(newWords, UPPRCS)
			} else {
				newWords = append(newWords, split)
				newWords = append(newWords, UPPRCS)
			}
		} else if strings.Contains(word, REPLCE) {
			splits := strings.Split(word, REPLCE)
			split := splits[0]
			if split == "" {
				newWords = append(newWords, REPLCE)
			} else {
				newWords = append(newWords, split)
				newWords = append(newWords, REPLCE)
			}
		} else {
			newWords = append(newWords, word)
		}
	}

	if countHypen % 2 != 0 {
		return nil, fmt.Errorf("invalid input")
	}

	var results []string
	for index := 0; index < countHypen / 2; index++ {
		value := strings.Join(newWords[:3], "-")
		results = append(results, value)
		if countHypen/2 > len(newWords){
			return nil, fmt.Errorf("invalid input")
		}
		newWords = newWords[3:]
	}
	return results, nil
}

func CheckSubstrings(str string, subs ...string) (bool, int) {
	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += strings.Count(str, sub)
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}