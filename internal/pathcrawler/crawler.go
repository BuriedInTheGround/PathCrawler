package pathcrawler

// An HashString is a string representation of an hash.
type HashString string

// A FilePath is a string of a path to a file.
type FilePath string

// A Crawler finds all files inside a directory and its subdirectories and
// calculate their hashes.
//
// Crawl returns a mapping of the hashes and lists of files with that hashes,
// and an error that may happen when looking for the files.
type Crawler interface {
	Crawl(rootPath string) (map[HashString][]FilePath, error)
}
