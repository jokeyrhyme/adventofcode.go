package day3

import (
	"strconv"
	"strings"
)

type Walker struct {
	ns int
	ew int
}

func NewWalker() *Walker {
	return &Walker{ns: 0, ew: 0}
}

func (w *Walker) Coords() string {
	return strconv.Itoa(w.ns) + "," + strconv.Itoa(w.ew)
}

func (w *Walker) Move(ch rune) {
	if ch == '^' {
		w.ns++
	}
	if ch == 'v' {
		w.ns--
	}
	if ch == '>' {
		w.ew++
	}
	if ch == '<' {
		w.ew--
	}
}

func answerB(input string) int {
	reader := strings.NewReader(input)

	houses := make(map[string]bool)

	santa := NewWalker()
	roboSanta := NewWalker()

	houses[santa.Coords()] = true

	var (
		walker *Walker
		ch     rune
		err    error
	)
	counter := 0
	for err == nil {
		ch, _, err = reader.ReadRune()

		if counter%2 == 0 {
			walker = santa
		} else {
			walker = roboSanta
		}
		walker.Move(ch)
		houses[walker.Coords()] = true

		counter++
	}

	return len(houses)
}
