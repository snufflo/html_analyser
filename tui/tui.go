package tui

import (
	"log"
	"fmt"
	"strings"
	"strconv"
	st "html_targeter/shared"

	"github.com/gdamore/tcell/v2"
)

func Tui_html(tags map[string][]st.Tag_info, attrs map[string][]st.Attr_info) {
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

		desc := fmt.Sprintf("%-5s %-15s %-15s", "LINE", "ATTR", "VAL")
		printLine(screen, 0, desc)
		printLine(screen, 1, strings.Repeat("-", width))
		// Display current input
		printLine(screen, height-2, strings.Repeat("=", width))
		printLine(screen, height-1, "Type: "+input)

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
			printLine(screen, 2, "No match")
		}

		if matches {
			for i := range height-4 {
				if scroll+i >= len(lines) {
					break
				}
				printLine(screen, i+2, lines[scroll+i])
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

