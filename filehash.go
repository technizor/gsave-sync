package gsavesync

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// Filehash hashes a file and returns it
func Filehash(fpath string) string {
	h := sha256.New()
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
