package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var links []string
	var chanLinks []chan string

	links = append(links, "https://www.google.com")
	links = append(links, "https://go.dev")
	links = append(links, "https://www.youtube.com")
	links = append(links, "https://www.stackoverflow.com")
	links = append(links, "https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855")
	links = append(links, "https://medium.com/")
	links = append(links, "https://esthercrawford.medium.com/an-epilogue-to-my-time-working-at-twitter-24a126098246")
	links = append(links, "https://barackobama.medium.com/thank-you-to-americas-librarians-for-protecting-our-freedom-to-read-80ce373608b3")
	links = append(links, "https://medium.com/wise-well/special-report-extreme-heat-and-human-health-da97f08f6aa6")
	links = append(links, "https://medium.com/@mattbarros_42186/the-never-ending-search-for-the-longest-word-5b1a66f5164e")
	links = append(links, "https://ux.shopify.com/uplifting-shopify-polaris-7c54fc6564d9")
	links = append(links, "https://medium.com/microsoft-design/a-change-of-typeface-microsofts-new-default-font-has-arrived-f200eb16718d")
	links = append(links, "https://avi-loeb.medium.com/unprecedented-hearing-on-extraterrestrials-in-the-us-house-of-representative-f9a217c78c37")
	links = append(links, "https://nickfthilton.medium.com/ashes-to-ashes-dust-to-dust-twitter-to-x-9667d689533e")
	links = append(links, "https://www.uol.com.br/")
	links = append(links, "https://www.uol.com.br/esporte/futebol/colunas/juca-kfouri/2023/08/01/a-venda-de-roger-guedes.htm")
	links = append(links, "https://noticias.uol.com.br/colunas/josmar-jozino/2023/08/01/policia-identifica-outro-suspeito-de-envolvimento-na-morte-do-soldado-reis.htm")
	links = append(links, "https://www1.folha.uol.com.br/poder/2023/08/tarcisio-confirma-convite-a-mae-de-neta-de-bolsonaro-mas-nega-favor.shtml")

	for _, link := range links {
		chanLink := make(chan string)
		chanLinks = append(chanLinks, chanLink)
		go getStatus(link, chanLink)
	}

	for _, resp := range chanLinks {
		fmt.Println(<-resp)
		close(resp)
	}
}

func getStatus(link string, chanLinks chan string) {
	time_start := time.Now()
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
	}

	text := makeTextResponse(link, resp.Status, time_start)

	chanLinks <- text
}

func makeTextResponse(link string, status string, timeResponse time.Time) string {
	return link + " - " + time.Since(timeResponse).String()
}
