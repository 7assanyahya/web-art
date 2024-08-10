package asciiZ

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiART(tex, ban string) string {

	file, err := os.Open("art/" + ban + ".txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	var reader *bufio.Scanner // Read the file
	tex = strings.ReplaceAll(tex, "\r", "")
	spl := strings.Split(tex, "\n")
	finalstring := ""
	for _, s := range spl {
		lines := make([]string, 8) // Store the lines of ASCII art
		stringArray := []rune(s)   // Convert the string into an array
		for i := 0; i < len(stringArray); i++ {
			charIndex := 9 * (int(stringArray[i]) - 32) // Determine the index of the character in the file
			file.Seek(0, 0)

			reader = bufio.NewScanner(file)

			for j := 0; j <= charIndex; j++ { // Find the first line for the character
				reader.Scan()
			}

			for k := 0; k < 8 && reader.Scan(); k++ { // Store the lines for the character
				lines[k] += reader.Text()
			}

			//finalstring += strings.Join(lines,"\n")
		}
		if finalstring != "" {
			finalstring += "\n"
		}
		finalstring += strings.Join(lines, "\n")
	}
	return finalstring
}
