package pathcrawler

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func hashFromFile(name string) HashString {
	hash, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintln("failed to read hash file:", err))
	}
	return HashString(strings.Trim(string(hash), "\r\n"))
}

var hashExampleTxt = hashFromFile("testdata/example-hash.txt")

func TestNaiveCrawl(t *testing.T) {
	crawler := NaiveCrawler{}

	res, err := crawler.Crawl("testdata")
	if err != nil {
		t.Fatalf("unexpected error; err = %v", err)
	}

	hits, ok := res[hashExampleTxt]
	if !ok {
		t.Fatalf("hash of example.txt (%q) not found in the results", hashExampleTxt)
	}

	hitsMap := filepathSliceToMap(t, hits)
	expectedHits := []FilePath{"testdata/example.txt", "testdata/example-symlink.txt"}
	for _, want := range expectedHits {
		if _, ok = hitsMap[want]; !ok {
			t.Fatalf("expected hit %q not found", want)
		}
	}
}

func filepathSliceToMap(t *testing.T, slice []FilePath) map[FilePath]struct{} {
	t.Helper()
	res := make(map[FilePath]struct{}, len(slice))
	for _, v := range slice {
		res[v] = struct{}{}
	}
	return res
}

func BenchmarkNaiveCrawl(b *testing.B) {
	crawler := NaiveCrawler{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		crawler.Crawl("testdata")
	}
}
