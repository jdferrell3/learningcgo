// Outputs file hashes
//
// > hashes.exe hashes.exe
// size:   1855488
// SHA256: c376478d0db9b8f11b55b1a5a6445b981543591e9dc565f943bf8ee2b699370a
// SHA1:   7c947e6705d16ffee423b31391fdb785ea1a6b6b
// MD5:    e03593e324ea9163b31481e42dd761a9

package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("\nUsage:\n\t%s <file_to_hash>", os.Args[0])
	}

	// Open file for hashing
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if nil != err {
		log.Fatalf("os.Open() err= %s", err)
	}
	defer file.Close()

	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()

	// new Reader, use OS's memory page size
	reader := bufio.NewReaderSize(file, os.Getpagesize())

	// multi writer to write all three algorithms at once
	writer := io.MultiWriter(md5, sha1, sha256)

	// buffered read, copy bytes to the writer (hash writers)
	size, err := io.Copy(writer, reader)
	if nil != err {
		log.Fatalf("io.Copy() err= %s", err)
	}

	fmt.Printf("size:   %d\n", size)
	fmt.Printf("SHA256: %s\n", hex.EncodeToString(sha256.Sum(nil)))
	fmt.Printf("SHA1:   %s\n", hex.EncodeToString(sha1.Sum(nil)))
	fmt.Printf("MD5:    %s\n", hex.EncodeToString(md5.Sum(nil)))
}
