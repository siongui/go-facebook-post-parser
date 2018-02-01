package parsefb

import (
	"bufio"
	"errors"
	"strings"
	"unicode"
)

func TitleWidth(t string) int {
	l := 0
	for _, r := range t {
		unicode.Is(unicode.Thai, r)
		l++
	}
	return l
}

// First two line of the string may look like:
//
//   Hello World
//   #######
//
// will become
//
//   Hello World
//   ###########
//
// after processing of this func
func SetRstDocTitleWidth(s string) (r string, err error) {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return r, err
	}

	if strings.TrimLeft(lines[1], "#") != "" {
		return r, errors.New("second line does not consist of all #")
	}

	lines[1] = strings.Repeat("#", TitleWidth(lines[0]))
	r = strings.Join(lines, "\n")
	r += "\n"

	return
}
