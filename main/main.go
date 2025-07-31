package main

import (
	"html_targeter/tui"
	"html_targeter/html"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter URL to parse source code of: ")
	reader := bufio.NewReader(os.Stdin)
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	tags, attrs := html.Html_parse(url)
	tui.Tui_html(tags, attrs)
}
