package scrapper

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	quranCookieName = "qnet"
)

var savedCookie *http.Cookie

// var ayatModel *models.AyatModel

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

func saveCookie(resp *http.Response) {
	cs := resp.Cookies()
	if cs != nil {
		for _, cookie := range cs {
			if cookie.Name == quranCookieName {
				savedCookie = cookie
			}
		}
	}
}

// AyatID : scrape ayat with ID
func AyatID(id string) (string, string, []string) {
	client := &http.Client{}
	values := make(map[int]string)
	read := ""
	textIndo := ""
	penjelasan := []string{}

	data := url.Values{}
	data.Set("idayat", id)
	data.Add("lang", "ina")
	data.Add("autoplay", "true")

	req, err := http.NewRequest("POST", "http://quran.bacalah.net/content/surat/GetContentAyat.php", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	// save the cookie
	saveCookie(resp)

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	n := 1
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		band, err := s.Find("font").Html()
		if err != nil {
			fmt.Println("Error font")
			fmt.Println(err)
			panic(err)
		}
		if n == 3 {
			values[n] = band
			read = band

			// fmt.Println(n)
			// fmt.Println(band)
		}
		td, err := s.Find("td").Html()
		if err != nil {
			fmt.Println("Error td")
			fmt.Println(err)
			panic(err)
		}
		values[n] = strings.TrimSpace(td)
		if n >= 9 {
			if n == 9 {
				textIndo = strings.ReplaceAll(td, "\n                            ", "")
			}
			// fmt.Println(n)
			// fmt.Println(strings.TrimSpace(td))
		}
		n++
	})
	for v := 10; v <= n; v++ {
		if len(values[v]) > 44 {
			// fmt.Println(values[v])
			penjelasan = append(penjelasan, values[v])
		}
	}

	return read, textIndo, penjelasan
}

// TafsirID : scrape tafsir ayat with ID
func TafsirID(id string) string {
	client := &http.Client{}

	data := url.Values{}
	data.Set("idayat", id)

	req, err := http.NewRequest("POST", "http://quran.bacalah.net/content/surat/GetTafsir.php", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(savedCookie)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// make selection
	sel := doc.Find(`div[id=content]`)

	// remove all found elements from selection
	sel.Find(`script`).Each(func(i int, s *goquery.Selection) {
		RemoveNode(sel.Get(0), s.Get(0))
	})

	// print html
	html, _ := sel.Html()
	if !strings.Contains(html, "error") {
		return html

		// fmt.Println(html)
	}

	return ""
}

// AsbabunNuzulID : get asbabun nuzul quran by ID
func AsbabunNuzulID(id string) string {
	client := &http.Client{}

	data := url.Values{}
	data.Set("idayat", id)

	req, err := http.NewRequest("POST", "http://quran.bacalah.net/content/surat/GetAsbnuz.php", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(savedCookie)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// make selection
	sel := doc.Find(`div[id=content]`)

	// remove all found elements from selection
	sel.Find(`script`).Each(func(i int, s *goquery.Selection) {
		RemoveNode(sel.Get(0), s.Get(0))
	})

	// print html
	html, _ := sel.Html()
	if !strings.Contains(html, "error") {
		return html

		// fmt.Println(html)
	}

	return ""
}

// RemoveNode : Searching node siblings (and child siblings and so on) and after successfull found - remove it
func RemoveNode(rootNode *html.Node, removeMe *html.Node) {
	foundNode := false
	checkNodes := make(map[int]*html.Node)
	i := 0

	// loop through siblings
	for n := rootNode.FirstChild; n != nil; n = n.NextSibling {
		if n == removeMe {
			foundNode = true
			n.Parent.RemoveChild(n)
		}

		checkNodes[i] = n
		i++
	}

	// check if removing node is found
	// if yes no need to check childs returning
	// if no continue loop through childs and so on
	if foundNode == false {
		for _, item := range checkNodes {
			RemoveNode(item, removeMe)
		}
	}
}
