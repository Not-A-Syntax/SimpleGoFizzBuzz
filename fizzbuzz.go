package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nmr = 1

func main() {

	// Leitor e Conversor Float
	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("Até quanto devo contar? ")
	objetivostr, _ := leitor.ReadString('\n')
	objetivostr = strings.TrimSpace(objetivostr)
	objetivo, _ := strconv.Atoi(objetivostr)

	//

	for nmr := nmr; nmr < objetivo+1; nmr++ {

		var texto string
		if nmr%3 == 0 {
			texto += "fizz"
		}
		if nmr%5 == 0 {
			texto += "buzz"
		}
		if nmr%7 == 0 {
			texto += "bazz"
		}
		if nmr%2 == 0 {
			texto = strings.ToUpper(texto)
		}
		if texto != "" {
			fmt.Println(texto)
		} else {
			fmt.Println(nmr)
		}

	}
}
