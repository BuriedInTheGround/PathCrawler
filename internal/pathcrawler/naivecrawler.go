package pathcrawler

import (
	"encoding/hex"
	"os"
	"path/filepath"

	"golang.org/x/crypto/blake2b"
)

// NaiveCrawler is a Crawler with low corner case checking and no concurrency.
type NaiveCrawler struct{}

func (c *NaiveCrawler) Crawl(rootPath string) (map[HashString][]FilePath, error) {
	res := make(map[HashString][]FilePath)
	err := crawl(rootPath, &res)
	return res, err
}

func crawl(path string, table *map[HashString][]FilePath) error {
	// The entries variable contains the list of files and directories that are
	// inside `path`.
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var subdirs []string

	for _, entry := range entries {
		// The entryPath variable contains the relative file path.
		entryPath := filepath.Join(path, entry.Name())

		if !entry.IsDir() {
			// Read the file data.
			data, err := os.ReadFile(entryPath)
			if err != nil {
				return err
			}

			// Generate the hash of the file data.
			hashArray := blake2b.Sum512(data)
			var hashBytes []byte
			hashBytes = hashArray[:]
			hash := HashString(hex.EncodeToString(hashBytes))

			// Append the file path to the map.
			(*table)[hash] = append((*table)[hash], FilePath(entryPath))
		} else {
			subdirs = append(subdirs, entryPath)
		}
	}

	for _, dir := range subdirs {
		crawl(dir, table)
	}

	return nil
}
