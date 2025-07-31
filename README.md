HTML-Parser
===========
This project is to be developed as a tool to aid enumerating and analyzing websites specifically to scout out potential vulnerabilities.

The current features of this project consist of:
- parsing an html source code
- constructing dictionaries that indicate which attributes and values each tag from which line has
- letting user look up tags in a tui using `tcell`

Roadmap
-------
- [x] parsing HTML source code
- [x] searching HTML tags real time
- [ ] crawl subdomains
- [ ] fingerprinting web architecture through HTTP header analysis

Usage
-----
1. After cloning this repo, run following go command:
```console
go run main/main.go
```

2. Enter URL to parse source code of

3. You can now search for tag names that were used in the html source code
    - The program will display their attributes and line number
