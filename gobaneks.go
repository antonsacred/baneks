package gobaneks

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"strings"
)

func RandomBAnek() (string, error) {
	resp, err := http.Get("http://baneks.ru/random")
	if err != nil {
		return "", fmt.Errorf("get banek: %w", err)
	}

	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("charset reader: %w", err)
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("parse html: %w", err)
	}

	var findArticle func(*html.Node) *html.Node
	findArticle = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.DataAtom == atom.Article {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result := findArticle(c); result != nil {
				return result
			}
		}
		return nil
	}

	section := findArticle(doc)
	if section == nil {
		return "", errors.New("banek not found")
	}

	var extractText func(*html.Node) string
	extractText = func(n *html.Node) string {
		if n.Type == html.TextNode {
			return n.Data
		}
		var result strings.Builder
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result.WriteString(extractText(c))
		}
		return result.String()
	}

	text := extractText(section)
	text = strings.Replace(text, "\\-", "-", -1)

	return text, nil
}
