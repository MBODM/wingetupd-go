package prettify

import (
	"strings"
)

// Not all possible search or list calls produce the same
// output. So therefore some conditionals are in use here.

func removeProgressBarChars(output string) string {
	if strings.Contains(output, "\b") {
		output = strings.NewReplacer(
			"\b|", "",
			"\b/", "",
			"\b-", "",
			"\b\\", "",
			"\b", "",
		).Replace(output)
		output = strings.TrimSpace(output)
	}
	return output
}

func removeLeadingReturn(output string) string {
	runes := []rune(output)
	firstChar := string(runes[0:1])
	if firstChar == "\r" {
		output = strings.Replace(output, "\r", "", 1)
		output = strings.TrimSpace(output)
	}
	return output
}

// Todo: Temporary disabled.
/*
func removeDownloadGraphs(output string) string {
	// A successful search or list output contains "Name " as first text,
	// or sometimes after weird download graphs. Solution: Remove graphs.
	// Don´t confuse that download graphs with the \b progress bar chars.
	namePos := strings.Index(output, "Name ")
	if namePos != -1 && namePos != 0 {
		runes := []rune(output)
		tail := runes[namePos:]
		output = string(tail)
	}
	return output
}
*/

// This is some example of above mentioned download graphs:
//
// ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒  1%\r ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒  2%\r ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 3%\r
// █▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒  5%\r ██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒  9%\r █████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 18%\r
// █████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 19%\r █████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 30%\r █████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 31%\r
// █████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 32%\r █████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 33%\r ██████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 34%\r
// ██████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 35%\r ██████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 48%\r ██████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 49%\r
// ███████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 50%\r ███████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 51%\r ███████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 52%\r
// ███████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 53%\r ████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 54%\r ████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 55%\r
// ████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒▒ 56%\r █████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒ 56%\r █████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒ 57%\r
// █████████████████▒▒▒▒▒▒▒▒▒▒▒▒▒ 59%\r ██████████████████▒▒▒▒▒▒▒▒▒▒▒▒ 60%\r ██████████████████▒▒▒▒▒▒▒▒▒▒▒▒ 61%\r
// ██████████████████▒▒▒▒▒▒▒▒▒▒▒▒ 62%\r ██████████████████▒▒▒▒▒▒▒▒▒▒▒▒ 63%\r ███████████████████▒▒▒▒▒▒▒▒▒▒▒ 64%\r
// ███████████████████▒▒▒▒▒▒▒▒▒▒▒ 65%\r ███████████████████▒▒▒▒▒▒▒▒▒▒▒ 66%\r ████████████████████▒▒▒▒▒▒▒▒▒▒ 67%\r
