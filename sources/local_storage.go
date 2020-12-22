package sources

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

type LocalStorageSource struct {
	path string
}

func (src *LocalStorageSource) Read(key string) (io.Reader, error) {
	touchFile, err := os.Open(src.touchFilePath(key))
	if err != nil {
		return nil, err
	}

	timestamp, err := ioutil.ReadAll(touchFile)
	if t, err := strconv.ParseInt(string(timestamp), 10, 64); err == nil && time.Now().Before(time.Unix(t, 0).Add(24*time.Second)) {
		return os.Open(src.cacheFilePath(key))
	}
	return nil, errors.New("cache has expired")
}

func (src *LocalStorageSource) Write(key string, body io.Reader) error {
	_ = os.Remove(src.cacheFilePath(key))
	_ = os.Remove(src.touchFilePath(key))

	f, err := os.Create(src.cacheFilePath(key))
	if err != nil {
		return err
	}

	_, err = io.Copy(f, body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(src.touchFilePath(key), []byte(strconv.FormatInt(time.Now().Unix(), 10)), 0644)
}

func (src *LocalStorageSource) touchFilePath(key string) string {
	return src.cacheFilePath(key) + "_timestamp"
}

func (src *LocalStorageSource) cacheFilePath(key string) string {
	hash := md5.Sum([]byte(key))
	return src.path + "/" + hex.EncodeToString(hash[:])
}

func NewLocalCacheSource(path string) (*LocalStorageSource, error) {
	return &LocalStorageSource{path: path}, nil
}
