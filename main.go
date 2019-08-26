// GoSay is a simple go implementation of CowSay.
package main

import (
	"flag"
	"fmt"

	"github.com/jxmoore/GoSay/models"
)

func main() {

	var saying = flag.String("s", "", "The string the gopher will say")
	var randomSaying = flag.Bool("r", false, "Determines if the message the gopher says is random of if the contents of -s should be used.")
	var flagNoColor = flag.Bool("st", false, "Enables or disable 'tv static' (rainbow) like text output")
	var localASCII = flag.String("la", "", "A local file containing an ASCII image that can be used in place of a gopher.")
	var messageBorder = flag.String("b", "!", "A charecter to use in the border of the message.")

	flag.Parse()
	err := models.GoSay(saying, randomSaying, flagNoColor, localASCII, messageBorder)
	if err != nil {
		fmt.Println("There was an error generating the output.")
	}
}
