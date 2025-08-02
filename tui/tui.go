package tui

import (
	"log"
	"fmt"
	"strings"
	"strconv"
	st "html_targeter/shared"

	"github.com/gdamore/tcell/v2"
)

func TuiHtml(tags map[string][]st.TagInfo, attrs map[string][]st.AttrInfo, url string) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer screen.Fini()

	input := ""
	scroll := 0

	for {
		var matches bool
		var lines []string
		screen.Clear()
		width, height := screen.Size()

		offset := 0
		printLineHighlight(screen, 0, "Showing results for: " + url, offset, tcell.ColorBlack, tcell.ColorPink)
		desc := fmt.Sprintf("%-5s %-15s %-15s", "LINE", "ATTR", "VAL")
		printLine(screen, 1, desc)
		printLine(screen, 2, strings.Repeat("-", width))
		// Display current input
		printLine(screen, height-2, strings.Repeat("=", width))
		tagText := "Tag: "
		printLine(screen, height-1, tagText)
		printLineHighlight(screen, height-1, input, len(tagText), tcell.ColorBlack, tcell.ColorWhite)

		// Check for match
		if val, ok := tags[input]; ok {
			matches = ok
			for _, t := range val {
				for i, a := range t.Attr {
					t_line := strconv.Itoa(int(t.Line))
					lines = append(lines, fmt.Sprintf("%-5s %-15s %-15s", t_line, a, t.Value[i]))
				}
			}
		} else {
			printLine(screen, 3, "No match")
		}

		if matches {
			for i := range height-5 {
				if scroll+i >= len(lines) {
					break
				}
				printLine(screen, i+3, lines[scroll+i])
			}
		}

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			case tcell.KeyUp:
				if scroll > 0 {
					scroll--
				}
			case tcell.KeyDown:
				if scroll < len(lines)-height {
					scroll++
				}
			case tcell.KeyRune:
				input += string(ev.Rune())
			}
		}
	}
}

func printLine(s tcell.Screen, y int, text string) {
	width, _ := s.Size()
	runes := []rune(text)

	for i := 0; i < len(runes); i += width {
		// wrap text
		end := i + width
		if end > len(runes) {
			end = len(runes)
		}
		line := string(runes[i:end])

		style := tcell.StyleDefault
		for x, r := range line {
			s.SetContent(x, y, r, nil, style)
		}
		y++
	}
}

func printLineHighlight(s tcell.Screen, y int, text string, offset int, fg, bg tcell.Color) {
	width, _ := s.Size()
	style := tcell.StyleDefault.Foreground(fg).Background(bg)

	for x := offset; x < width; x++ {
		var r rune = ' '
		if x-offset < len(text) {
			r = rune(text[x-offset])
		}
		s.SetContent(x, y, r, nil, style)
	}
}
