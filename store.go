package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Original: hashStr,
	}
}

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	Original string
}

func (p PathKey) filename() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	data := r
	buf := new(bytes.Buffer)
	io.Copy(buf, data)

	filenameBytes := md5.Sum(buf.Bytes())
	filename := hex.EncodeToString(filenameBytes[:])
	pathAndFilename := pathKey.PathName + "/" + filename

	f, err := os.Create(pathAndFilename) 
	if err != nil {
		return err
	}

	n, err := io.Copy(f, data)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk: %s", n, pathAndFilename)

	return nil
}