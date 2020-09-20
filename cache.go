package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func readCache(key string) (io.Reader, error) {
	touchFile, err := os.Open(touchFilePath(key))
	if err != nil {
		return nil, err
	}

	timestamp, err := ioutil.ReadAll(touchFile)
	if t, err := strconv.ParseInt(string(timestamp), 10, 64); err == nil && time.Now().Before(time.Unix(t, 0).Add(24*time.Second)) {
		return os.Open(cacheFilePath(key))
	}
	return nil, errors.New("cache has expired")
}

func writeCache(key string, body io.Reader) error {
	_ = os.Remove(cacheFilePath(key))
	_ = os.Remove(touchFilePath(key))

	f, err := os.Create(cacheFilePath(key))
	if err != nil {
		return err
	}

	_, err = io.Copy(f, body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(touchFilePath(key), []byte(strconv.FormatInt(time.Now().Unix(), 10)), 0644)
}

func cacheFilePath(key string) string {
	hash := md5.Sum([]byte(key))
	return cachePath + "/" + hex.EncodeToString(hash[:])
}

func touchFilePath(key string) string {
	return cacheFilePath(key) + "_timestamp"
}
