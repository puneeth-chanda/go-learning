package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	url := "https://tools.ietf.org/html/rfc2616#section-3.3.1"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)

	var re = regexp.MustCompile(`(?misU)href="/rfc/|href="/pdf/`)
	// var str = `<span class="pre noprint docinfo top">[<a href="../html/" title="Document search and retrieval page">Docs</a>] [<a href="/rfc/rfc2616.txt" title="Plaintext version of this document">txt</a>|<a href="/pdf/rfc2616" title="PDF version of this document">pdf</a>] [<a href="./draft-ietf-http-v11-spec-rev" title="draft-ietf-http-v11-spec-rev">draft-ietf-http...</a>] [<a href='https://datatracker.ietf.org/doc/rfc2616' title='IESG Datatracker information for this document'>Tracker</a>] [<a href="/rfcdiff?difftype=--hwdiff&amp;url2=rfc2616" title="Inline diff (wdiff)">Diff1</a>] [<a href="/rfcdiff?url2=rfc2616" title="Side-by-side diff">Diff2</a>] [<a href="https://www.rfc-editor.org/errata_search.php?rfc=2616">Errata</a>]</span><br />`

	for i, match := range re.FindAllString(string(html), -1) {
		fmt.Println(match, "found at index", i)
	}
}
