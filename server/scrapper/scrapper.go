package scrapper

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"bookstore/books"
	"bookstore/item"

	"github.com/gocolly/colly/v2"
	"github.com/nfnt/resize"
)

func scrapBooks() {
	var wg sync.WaitGroup
	wg.Add(2)

	go scrapBotanica(&wg, "distribuidoralabotica.json")
	go scrapMeli(&wg, "meli.json")

	wg.Wait()
}

func scrapMeli(wg *sync.WaitGroup, fileName string) {
	if wg != nil {
		defer wg.Done()
	}

	allBooks := make([]*books.Book, 0)
	from := 0

	for {
		c := colly.NewCollector()
		fmt.Fprintf(os.Stderr, "\n count: %d", len(allBooks))
		initialLen := len(allBooks)
		// On every a element which has href attribute call callback
		c.OnHTML("ol", func(e *colly.HTMLElement) {
			e.ForEach("li", func(_ int, pub *colly.HTMLElement) {
				book := &books.Book{Item: &item.Item{}}
				book.Name = pub.ChildAttr("a", "title")
				strprice := pub.ChildText("span.price-tag-fraction")
				strprice = strings.ReplaceAll(strprice, ".", "")
				price, _ := strconv.Atoi(strprice)
				book.Price = int64(price) * 100
				book.PictureURL = pub.ChildAttr("img", "data-src")
				allBooks = append(allBooks, book)
			})
		})

		c.Visit(fmt.Sprintf("https://m4editorialml.mercadoshops.com.ar/libros/_Desde_%d", from))
		c.Wait()

		if len(allBooks) == initialLen {
			break
		}

		from += 50
	}

	saveBooks(allBooks, fileName)
	fmt.Fprintf(os.Stderr, "\n meli done")
}

func scrapBotanica(wg *sync.WaitGroup, filename string) {
	if wg != nil {
		defer wg.Done()
	}

	// Instantiate default collector
	c := colly.NewCollector()

	allBooks := make([]*books.Book, 0)

	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			book := &books.Book{Item: &item.Item{}}
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				switch el.Index {
				case 1:
					url := el.ChildAttr("img", "src")
					book.PictureURL = getPictureUrl(url)
				case 3:
					book.Name = strings.Title(strings.ToLower(el.Text))
				case 4:
					book.Author = []string{strings.Title(strings.ToLower(el.Text))}
				case 5:
					book.Publisher = []string{strings.Title(strings.ToLower(el.Text))}
				case 6:
					book.Category = []string{strings.Title(strings.ToLower(el.Text))}
				case 7:
					strPrice := strings.ReplaceAll(el.Text, "$", "")
					strPrice = strings.ReplaceAll(strPrice, ".", "")
					price, _ := strconv.Atoi(strPrice)
					book.Price = int64(price)
				}
			})
			allBooks = append(allBooks, book)
			fmt.Fprintf(os.Stderr, "\n count: %d", len(allBooks))
		})
	})

	c.Visit("http://www.distribuidoralabotica.com.ar/libro/CatalogoLibrosC")
	c.Wait()

	saveBooks(allBooks, filename)
	fmt.Fprintf(os.Stderr, "\n botanica done")
}

func saveBooks(books []*books.Book, filename string) {
	bytes, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the whole body at once
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func downloadIMG(URL string) ([]byte, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return []byte{}, fmt.Errorf("error downloading %s", URL)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("error downloading %s", URL)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading %s", URL)
	}
	return body, nil
}

func resizeImg(url string, imgBytes []byte) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return "", fmt.Errorf("error decoding img %s", url)
	}

	newImage := resize.Resize(200, 300, img, resize.Lanczos3)

	var b bytes.Buffer
	newImageBytes := bufio.NewWriter(&b)

	err = jpeg.Encode(newImageBytes, newImage, nil)
	if err != nil {
		return "", fmt.Errorf("error writing new resized img %s", url)
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(b.Bytes())
	return "data:image/jpg;base64," + imgBase64Str, nil
}

func getPictureUrl(URL string) string {

	originalImg, err := downloadIMG(URL)
	if err != nil {
		return ""
	}
	resizedImg, err := resizeImg(URL, originalImg)
	if err != nil {
		return ""
	}

	return resizedImg
}
