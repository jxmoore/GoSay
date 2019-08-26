package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

// quote is a struct for json unmarshiling
type quote struct {
	Author    string
	id        int
	Quote     string
	Permalink string
}

// GoSay is a function that prints a speech bubble and a gopher
func GoSay(saying *string, randomSaying *bool, randomColor *bool, localASCII *string, messageBorder *string) (err error) {

	var speechBubble, ASCIIart string
	rand.NewSource(time.Now().UnixNano())

	if *randomSaying == true || *saying == "" {
		pullMessage, err := getMessage()
		if err != nil {
			return errors.New("Error retreiving a random quote via API call")
		}
		speechBubble, err = newSpeechBubble(pullMessage.Quote, *messageBorder)
	} else {
		speechBubble, err = newSpeechBubble(*saying, *messageBorder)
	}
	if err != nil {
		return errors.New("Unkown error generating speech bubble")
	}

	if *localASCII == "" {
		ASCIIart = getGopher()
	} else {
		ASCIIart, err = readLocalASCII(*localASCII)
		if err != nil {
			return errors.New("There was an error reading the local file")
		}
		ASCIIart = "\n" + ASCIIart
	}
	if *randomColor == true {
		printRandomColors(speechBubble)
		printRandomColors(ASCIIart)
	} else {
		fmt.Println(speechBubble)
		fmt.Println(ASCIIart)
	}
	return nil
}

// getMessage is a function that returns a pointer to a quote retreived from an API call
func getMessage() (*quote, error) {

	message, err := http.Get("http://quotes.stormconsultancy.co.uk/random.json")
	if err != nil {
		return &quote{}, errors.New("Error making GET request")
	}

	defer message.Body.Close()

	body, err := ioutil.ReadAll(message.Body)
	if err != nil {
		return &quote{}, errors.New("Error reading message body")
	}

	marshalQuote := quote{}
	err = json.Unmarshal([]byte(body), &marshalQuote)
	if err != nil {
		return &quote{}, errors.New("Error reading json")
	}

	return &marshalQuote, nil
}

// textSplit takes a string and returns a slice of strings that do not exceed 62 runes in length
func textSplit(phrase string) []string {

	var buffer string
	var returnPhrase []string
	splitPhrase := strings.Split(phrase, " ")

	for x, word := range splitPhrase {
		if (len([]rune(buffer)) + len([]rune(word))) > 62 {
			returnPhrase = append(returnPhrase, buffer)
			buffer = word
		} else {
			if buffer == "" {
				buffer = word
			} else {
				buffer = buffer + " " + word
			}
		}

		if len(splitPhrase) <= x+1 {
			returnPhrase = append(returnPhrase, buffer)
		}
	}

	return returnPhrase

}

// newSpeechBubble is a renders a string into an ASCII speech bubble.
func newSpeechBubble(phrase string, border string) (string, error) {
	if len(phrase) <= 0 {
		return "", errors.New("no phrase entered")
	}
	if len(border) > 1 {
		border = "!"
	}
	txtBubble := ".-----------------------------------------------------------------.\n"

	phraseSlice := textSplit(phrase)

	for _, sentence := range phraseSlice {
		for len([]rune(sentence)) <= 63 {
			sentence = sentence + " "
		}

		txtBubble = txtBubble + border + " " + sentence + border + "\n"
	}

	txtBubble = txtBubble + "`------------------------------------------.  ,-------------------' \n                                            \\|\n                                             \\"

	return txtBubble, nil
}

// printRandomColors takes a string and prints said string char by char in a random color
func printRandomColors(textBlock string) {
	colorArray := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgBlue, color.FgMagenta,
		color.FgCyan, color.FgWhite, color.FgHiRed, color.FgHiGreen, color.FgHiYellow,
		color.FgHiBlue, color.FgHiMagenta, color.FgHiCyan, color.FgHiWhite}

	for _, c := range textBlock {
		newColor := color.New(colorArray[rand.Intn(len(colorArray))])
		schar := string(c)
		if schar == "\n" {
			fmt.Println("")
		} else {
			newColor.Print(schar)
		}
	}
}

// readLocalASCII reads a file and returns it as a string
func readLocalASCII(file string) (fileASCII string, err error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading file")
	}
	return string(contents), nil
}
