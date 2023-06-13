package states

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	structs "warn/warn/structs"
)

/*
func newWarn(row string) *warn {
	w := warn()
	return &w
}
*/

/**
 * Code lifted from: https://zetcode.com/golang/net-html/
 */

func getHtmlPage(webPage string) (string, error) {

	resp, err := http.Get(webPage)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return "", err
	}

	return string(body), nil
}

/**
 * Trim whitespace on the front of the string and collapse extra spaces in the middle of the
 * string to one space character.
 */
func formatString(text string) string {

	trimmed := strings.TrimSpace(text);
	m1 := regexp.MustCompile(`\s+`);

	return m1.ReplaceAllString(trimmed, " ");
}

/**
 * Walk through the HTML nodes to find any new ones.
 */
func parseAndShow(text string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var isTd bool
	var isHTML bool
	var n int

	var strArray [][]string
	var tempArray []string

	for isHTML != true {
		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return

		case tt == html.EndTagToken:

			t := tkn.Token()
			isHTML = t.Data == "html"
			if isHTML {
				log.Println("HTML End Tag")
				break;
			}

		case tt == html.StartTagToken:

			t := tkn.Token()
			isTd = t.Data == "td"

		case tt == html.TextToken:

			t := tkn.Token()
			if isTd {
				tempArray = append(tempArray, t.Data)
				n++
			}

			if isTd && n%8 == 0 {
				strArray = append(strArray, tempArray)
				tempArray = nil
			}
			isTd = false
		}
	}

	for _, v := range strArray {
		totalEmployees, _ := strconv.Atoi(v[5])
		test := structs.WarnNotice{
			NoticeDate: formatString(v[0]),
			NaicsCode: formatString(v[1]),
			Company: formatString(v[2]),
			Location: formatString(v[3]),
			LocalArea: formatString(v[4]),
			TotalEmployees: totalEmployees,
			EffectiveDate: formatString(v[6]),
			LayoffType: formatString(v[7]),
		}
		fmt.Printf("%v\n", test.Company)
	}
}

func MarylandHandler() (string, error) {
	webPage := "https://www.dllr.state.md.us/employment/warn.shtml"
	data, err := getHtmlPage(webPage)
	if err != nil {
		log.Fatal(err)
	}
	parseAndShow(data)
	return "", nil
}
