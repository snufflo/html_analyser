package html

import (
	st "html_targeter/shared"
	"fmt"
	"io"
	"bytes"
	"strings"
	"net/http"
	"golang.org/x/net/html"
)


func Html_parse(url string) (map[string][]st.Tag_info, map[string][]st.Attr_info) {
	var tags = make(map[string][]st.Tag_info)
	var attrs = make(map[string][]st.Attr_info)

	resp, err := http.Get(url)
	log_err(err, "http get fail")
	defer resp.Body.Close() // close TCP connection

//	bytes, err := io.ReadAll(resp.Body)
//	log_err(err)

	parse_html(resp.Body, tags, attrs)

	fmt.Println("")
	fmt.Println("TAGS:")
	fmt.Println(strings.Repeat("=", 90))
	for key := range tags {
		fmt.Println(strings.ToUpper(key), tags[key])
		fmt.Println(strings.Repeat("-", 90))
	}

	fmt.Println("")
	fmt.Println("ATTRS:")
	fmt.Println(strings.Repeat("=", 90))
	for key := range attrs {
		fmt.Println(strings.ToUpper(key), attrs[key])
		fmt.Println(strings.Repeat("-", 90))
	}

	/*
	doc, err := html.Parse(resp.Body)
	log_err(err)
	*/

	return tags, attrs
}

func parse_html(html_src io.Reader, tags map[string][]st.Tag_info, attrs map[string][]st.Attr_info) {
	tokenizer := html.NewTokenizer(html_src)
	var line uint = 1

	html_loop:
	for {
		var tn string
		var ta []string
		var val []string
		token := tokenizer.Next()
		raw := tokenizer.Raw()
		line += uint(bytes.Count(raw, []byte{'\n'}))

		if token == html.ErrorToken {
			log_err(tokenizer.Err(), "error token")
			if tokenizer.Err() == io.EOF {
				break html_loop
			}
		}

		tn_bytes, has_attr := tokenizer.TagName()
		tn = string(tn_bytes)

		if has_attr {
			for {
				ta_bytes, val_bytes, has_more_attr := tokenizer.TagAttr()
				ta = append(ta, string(ta_bytes))
				val = append(val, string(val_bytes))

				attr := st.Attr_info{
					Tag: tn,
					Value: val[len(val)-1],
					Line: line,
				}
				var ta_indx = len(ta)-1
				attrs[ta[ta_indx]] = append(attrs[ta[ta_indx]], attr)

				if !has_more_attr {
					break
				}
			}
		}

		switch token {
		//case html.TextToken:
			// plain text detected
		case html.StartTagToken:

			tag := st.Tag_info{
				Attr: ta,
				Value: val,
				Line: line,
			}

			tags[tn] = append(tags[tn], tag)

		}
	}
	fmt.Println("number of lines:", line)
}

func log_err(err error, str string) {
	if err != nil {
		fmt.Println(str, err)
		return
	}
}
