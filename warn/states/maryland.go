package states

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func parseAndShow(text string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var isTd bool
	var isHTML bool
	var n int

	var strAray [][]string
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
				strAray = append(strAray, tempArray)
				tempArray = nil
			}
			isTd = false
		}
	}

	// fmt.Print("Output:")
	// fmt.Printf("%v", strAray)

	for _, v := range strAray {

		totalEmployees, _ := strconv.Atoi(v[5])
		test := structs.WarnNotice{
			NoticeDate: strings.TrimSpace(v[0]),
			NaicsCode: strings.TrimSpace(v[1]),
			Company: strings.TrimSpace(v[2]),
			Location: strings.TrimSpace(v[3]),
			LocalArea: strings.TrimSpace(v[4]),
			TotalEmployees: totalEmployees,
			EffectiveDate: strings.TrimSpace(v[6]),
			LayoffType: strings.TrimSpace(v[7]),
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


