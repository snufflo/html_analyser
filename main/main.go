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
	fmt.Println("requesting html source code...")

	tags, attrs := html.HtmlParse(url)
	tui.TuiHtml(tags, attrs, url)
}
