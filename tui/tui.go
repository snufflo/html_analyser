package tui

import (
	"log"
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

	for {
		screen.Clear()
		width, height := screen.Size()

		// Display current input
		printLine(screen, height-2, strings.Repeat("=", width))
		printLine(screen, height-1, "Type: "+input)

		// Check for match
		if val, ok := tags[input]; ok {
			y := 0
			for _, t := range val {
				var line string
				for i, a := range t.Attr {
					line = strconv.Itoa(int(t.Line)) + ": " + a + ":: " + t.Value[i]
					printLine(screen, y, line)
					y++
				}
			}
		} else {
			printLine(screen, 1, "No match")
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

