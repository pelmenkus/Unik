package main

import (
	"github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
	"html/template"
	"net/http"
)

const INDEX_HTML = `
    <!doctype html>
    <html lang="ru">
        <head>
            <meta charset="utf-8">
            <title>Последние новости с breaking-news</title>
        </head>
        <body>
            {{if .}}
                {{range .}}
                    {{.Time}}
                    <a href="{{.Ref}}">{{.Title}}</a>
                    <br/>
                {{end}}
            {{else}}
                Не удалось загрузить новости!
            {{end}}
        </body>
    </html>
    `

var indexHtml = template.Must(template.New("index").Parse(INDEX_HTML))

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func isArticle(node *html.Node, class string) bool {
	return isElem(node, "article") && getAttr(node, "class") == class
}

func isH4(node *html.Node, class string) bool {
	return isElem(node, "h4") && getAttr(node, "class") == class
}

type Item struct {
	Ref, Time, Title string
}

func readItem(item *html.Node) *Item {
	h := item.FirstChild
	for ; !isH4(h, "storyblock_title"); h = h.NextSibling {
	}
	a := h.FirstChild
	return &Item{
		Ref:   getAttr(a, "href"),
		Title: a.FirstChild.Data,
	}
	return nil
}

func search(node *html.Node) []*Item {
	if isDiv(node, "l1-s_list") {
		var items []*Item
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isArticle(c, "storyblock") {
				for art := node.FirstChild; art != nil; art = art.NextSibling {
					if item := readItem(art); item != nil {
						items = append(items, item)
					}
				}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			return items
		}
	}
	return nil
}

func downloadNews() []*Item {
	log.Info("sending request to breaking-news")
	if response, err := http.Get("https://www.news.com.au/world/breaking-news"); err != nil {
		log.Error("request to https://www.news.com.au/world/breaking-news failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from https://www.news.com.au/world/breaking-news", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML https://www.news.com.au/world/breaking-news", "error", err)
			} else {
				log.Info("HTML from https://www.news.com.au/world/breaking-news parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}

func serveClient(response http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	log.Info("got request", "Method", request.Method, "Path", path)
	if path != "/" && path != "/index.html" {
		log.Error("invalid path", "Path", path)
		response.WriteHeader(http.StatusNotFound)
	} else if err := indexHtml.Execute(response, downloadNews()); err != nil {
		log.Error("HTML creation failed", "error", err)
	} else {
		log.Info("response sent to client successfully")
	}
}

func main() {
	http.HandleFunc("/", serveClient)
	log.Info("starting listener")
	log.Error("listener failed", "error", http.ListenAndServe("127.0.0.1:6061", nil))
}
