// xpathchecker is test tool for confirming extraction by XPath.
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/xmlpath.v2"
)

func init() {
	flag.Var(&xpaths, "xpaths", "xpaths is multiple xpath for extracting content. space separated.")
}

var (
	xpaths XPaths
	rawurl = flag.String("url", "http://example.com", "url to fetch and extract")
)

type XPaths []*xmlpath.Path

func (x *XPaths) String() string {
	return fmt.Sprint(*x)
}

func (x *XPaths) Set(value string) error {
	if len(*x) > 0 {
		return errors.New("xpaths flag already set")
	}
	for _, v := range strings.Split(value, " ") {
		*x = append(*x, xmlpath.MustCompile(strings.TrimSpace(v)))
	}
	return nil
}

func main() {
	flag.Parse()
	if err := realmain(); err != nil {
		panic(err)
	}
}

func realmain() error {
	resp, err := http.Get(*rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	rr, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}
	root, err := xmlpath.ParseHTML(rr)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	for _, xpath := range xpaths {
		iter := xpath.Iter(root)
		for iter.Next() {
			buf.Write(iter.Node().Bytes())
			buf.Write([]byte("\n"))
		}
	}
	fmt.Printf("%s", buf.String())
	return nil
}
