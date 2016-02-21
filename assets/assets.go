// Code generated by go-bindata.
// sources:
// velox.js
// json-patch.js
// event-source.js
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _veloxJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x51\x6f\xdb\x36\x10\x7e\xcf\xaf\xb8\x08\x45\x2a\xc7\x8e\xe4\x76\xd8\x4b\xdc\x34\x1b\xb2\x3c\xb4\xe8\xda\x02\x0e\xda\x87\x2c\x2b\x68\x89\xb6\xd5\xc8\xa2\x40\x52\xb6\x85\x34\xff\x7d\x77\x24\x25\x51\x76\xdc\xb5\x03\x56\xa0\x51\x74\xbc\x3b\xde\xf1\xbe\xfb\x8e\x4a\x7c\x0a\x8b\x5c\xcc\x58\x0e\x9b\xac\x48\xc5\x66\xf4\x99\xcf\xa6\x22\xb9\xe7\x1a\x4e\xe3\xa3\x38\x86\x35\xcf\xc5\x16\xce\x60\x3d\x8e\x5e\x44\x63\xfc\x65\xa9\x75\xa9\xce\xe3\x78\x91\xe9\x65\x35\x8b\x12\xb1\x8a\xbf\x96\x59\x9e\x0b\xc9\x62\xa3\x4c\x56\x6f\x59\xb6\xe2\xf0\xd1\x8a\xe1\x55\xca\xd7\xbf\x35\x4a\x64\xf1\x1a\xfd\xfc\xf9\xe6\x06\xae\x44\x59\xcb\x6c\xb1\xd4\xf0\x72\xfc\xe2\xd7\xa3\x70\x5e\x15\x89\xce\x44\x11\xba\x60\x20\x15\x49\xb5\xe2\x85\x1e\xc0\xc3\x11\x40\x36\x0f\x8f\xed\x4a\xd4\x86\x39\x40\x39\x80\xe4\xba\x92\x05\xb0\x9c\x4b\x1d\x06\x37\xcb\x4c\xc1\x4c\x8a\x8d\xe2\x12\x5d\x70\x05\x85\xd0\xa0\xaa\xb2\x14\x52\x43\x6b\xaa\x82\xc1\x04\xad\x63\x1b\x36\x94\x52\x68\x91\x88\x1c\x53\x96\x0a\x83\xc0\xa5\x35\x93\x56\x0c\x17\x10\xac\x5f\x06\x13\x27\x5b\x2b\x14\xdc\xde\x59\xeb\xb2\x9a\xe5\x59\x02\x2b\xae\x97\x22\x6d\x14\x8c\xc7\x0b\x68\x13\xd2\x75\xc9\x47\x50\xc9\x7c\x04\x62\xf6\xd5\xa6\xe3\x54\x51\xad\xe0\x1b\xf8\x44\x26\xbb\x7a\x13\xab\xa6\xa2\xb2\x52\xcb\x70\xed\xde\x5d\xb6\x6b\x7a\x7b\x34\x31\x91\x6d\xf4\x79\x8a\xae\x1e\x3c\xc1\x46\xf9\x21\xec\x6e\xde\x78\x31\xfb\x36\x1e\xbc\xbd\xfb\xce\xa7\xd3\xeb\x1d\xef\x4a\xf1\x9f\x74\x8f\x3e\x0e\xfa\x6f\x8e\xd9\x3c\x3b\xb1\x28\xf2\xac\xa0\x7d\xb4\xac\xb8\x3d\x6f\x55\xf2\x24\x43\xc0\xae\xb8\x5c\x70\x04\x52\xb6\x28\x84\xe4\xf0\x0c\x2d\x4b\x2c\x7f\xc6\x95\x51\x83\x2d\xbc\x3a\x83\xda\x15\xc4\x2a\x7b\xe1\x6e\x47\x50\x37\xa1\x66\x73\x08\x8f\xb7\xf0\xed\x1b\xd0\xf1\x8b\x39\x9a\x1e\x5f\x60\xc5\x31\x4a\x9e\xe8\x80\x16\x8e\x6b\x6f\xb9\xee\x2d\x5b\x04\xb6\x09\xd7\x93\xb6\xb4\xf7\x93\xd6\xfd\x16\xb2\x42\x69\x56\x24\x64\xff\xbb\x94\xac\x86\x93\x13\x74\xb4\x2b\x6d\x42\xa2\x0c\x24\x5f\x89\x35\x07\xbe\xd5\xd8\x43\x3c\xe7\xd4\x08\xca\xad\x6e\x96\x59\xce\xd1\x6d\x94\xf3\x62\xa1\x97\xf0\x1a\x6a\xf7\x6b\x13\x0e\x00\x1e\xaa\x28\x43\x87\x99\x47\xf4\x80\x05\x3b\xe0\xbd\x77\x76\xf4\x6f\x2e\x24\x84\xf7\x18\x1e\x6c\x3b\x87\x94\xc8\xfd\xed\xf8\xce\xa6\xff\x2c\xa0\x14\x8e\xad\x56\x3d\xe8\xd4\x00\x52\x8c\x56\x73\xd8\xde\xde\xdf\xb9\xdd\x8f\xfa\x4e\xeb\x46\x9b\x54\xb0\x2c\xa6\x3c\x21\xbd\x60\x59\xf0\x67\x1f\xe8\xdb\x16\xe8\x48\x3b\x96\xad\xf0\xd4\x74\xa5\x20\x59\xb2\x02\xeb\x8a\x3f\x53\x6c\x7d\x54\x68\xea\x0b\xa2\xb0\x2a\x21\x5f\xb7\xfc\xb1\x07\xaa\x82\xad\xb3\x05\xd3\x42\xa2\xe8\x1d\x8a\x26\x4d\x98\x21\x95\x2f\x43\x8d\xf1\x04\x1f\xaf\x40\x6a\xe5\x8e\x17\xdf\x87\xc3\xae\x4a\x48\x49\x3d\x9f\x78\x22\xa8\x7b\x9b\xdd\x45\xac\x42\xbf\x5c\xcb\xba\x3b\x18\xb7\x62\xa4\x6d\x5d\x8e\xec\x7f\xc7\x6b\x2c\x4d\xaf\x29\xe0\x77\x99\xd2\xbc\xe0\x32\x7c\x6e\x1d\x3f\x1f\x41\x9b\x92\xb1\x3c\xac\x3f\x9f\x3b\x03\x4f\xdf\x1c\xdd\x92\xe7\x58\x64\xe5\x3a\xc2\x9c\x8b\xe1\xb1\x60\xc5\x95\x62\x0b\x1e\x8c\x02\x2e\xa5\x90\xf8\x44\x34\x14\xf8\x48\x72\xa1\x78\x70\xd7\xd0\x5e\x2e\x12\xd4\x77\x3b\xe3\x0b\xa3\x93\xf6\x29\x34\xc9\x99\x52\xd8\x92\x92\x97\x92\x2b\xe3\x9f\x81\xca\x8a\x05\x62\x75\xc3\x67\xca\x0e\x96\xf0\x4a\x14\x54\x20\xd0\x4b\x0e\x48\xd1\xc8\xb7\x67\x2a\x4b\xf9\xc0\xaf\xdf\x93\x74\xe8\x8e\x5d\x6d\x32\x9d\x2c\xcd\x62\x23\x4a\x18\x82\xbb\x21\xb1\x73\x77\xe2\x1a\xe7\x80\x25\x41\xc3\x1e\x38\x12\x38\x73\x3d\xe9\xe9\x23\x2b\xf5\x0c\x2c\xaf\xed\x59\xa4\x7c\xce\xaa\x5c\x77\xaa\x38\x5f\x20\xb8\xc1\x18\x60\x55\x29\x0d\xb3\x6e\x7f\x10\xb2\xf3\x1d\xf8\x1d\x40\x13\x0c\xb3\xe9\xf0\x83\x2f\x34\x5b\xec\xf9\xed\xa9\x86\xf1\xdf\xe1\x46\x7d\xa3\x91\x3b\x50\x97\xe7\x71\xa4\xb9\xd2\xc4\xb4\x83\xc1\xae\x0b\x2c\x47\xd4\x4e\xb0\x21\x7a\x8c\x03\x7c\x90\x74\x29\x30\xba\x21\xa9\xed\xb8\x77\xc7\xb3\xeb\x09\x7f\x22\x44\xcb\x9c\x25\x1c\xf7\xa7\xbd\xe3\x11\x04\x1b\x37\x2d\x1b\x7b\x63\xdc\x1a\x4c\xda\x90\xb1\x4c\x1e\x53\xd2\xdb\x53\x5c\xe9\x4e\xef\x4d\xb1\x66\x79\x96\x82\x5b\x9d\x74\x9e\xc9\xf0\x82\xe4\x9e\xcc\xcd\x65\xd3\x95\x9e\x66\x51\x95\x29\xd3\x3d\x72\xc7\x94\xe2\xd3\x42\x88\xf2\x34\x7e\xec\xa9\x1a\x7c\xff\x80\x66\x82\x08\xc5\x88\x78\x4a\xba\x0c\x89\x73\x7f\xcd\x36\x30\x1d\xc6\xa7\x6e\x80\x51\xda\x34\x26\x2d\xc6\xac\xe2\x79\x6f\x37\x1f\x69\x2d\x45\x78\xf3\xad\x5d\xec\xb3\xc4\xa8\xe1\x43\x59\x3f\xe9\x2f\xc9\x39\x93\x37\x78\xe9\x12\x95\x0e\x3b\xfb\x48\x0f\x26\x1d\x53\xb5\xe1\x77\x94\x64\x45\x68\x8c\xa7\x18\xfa\xba\xc7\x66\x05\x89\x9c\xd5\x3b\xda\x46\x86\x11\xbf\x18\x8f\x1b\x7d\x22\x07\x8b\x85\x16\x16\x43\x08\xe3\xbf\x2e\x1d\x62\x1b\xe9\x00\x2e\x21\x38\x09\xe0\x1c\x82\xcb\x60\x40\x30\x2d\x2f\x82\xa1\x39\xb9\x61\x70\xb2\xc6\xdf\xfd\x42\xef\x04\xee\x23\xd5\x2b\x85\xbb\x3e\xb5\xb7\x3a\xd3\x1e\x4d\x60\x3b\x53\x6f\xdf\xcc\x90\xe7\x54\x54\x12\xc1\x6e\x68\xe6\x01\xf9\x4d\x2f\xaf\x24\x4f\x71\x01\xaf\x19\xea\xdc\x94\x06\x1e\x3b\x9f\x5e\xd2\x5f\xc8\x9f\x4b\xbb\x59\xb7\xc4\x1a\xe1\x14\xb9\x66\xc8\x52\x6d\xb1\xb8\x1f\xfe\x97\x36\x90\xdb\x40\x14\xc1\x90\xd3\x14\x34\xc2\xdb\x80\xa4\x24\x89\x66\xc8\xb4\xa1\x11\x76\x9b\x0f\x7a\x20\x29\x91\x5b\x23\x8d\xa6\x8a\xeb\x37\x85\x46\x2e\x65\x79\xd8\xad\x18\x07\xc6\x7e\x04\xbf\x8c\xe1\x94\x6a\x36\xee\x43\x2a\xcd\xd4\x4f\xe1\xd4\xeb\x86\xa7\xd1\xe3\xfc\x3a\xa9\xef\xf4\x61\x17\x5d\x7d\x28\xda\x61\x3f\xf9\xdf\x4e\xb7\xa8\xf2\x7c\xff\x18\xfd\xb6\xc0\xe9\xed\x5d\x62\x3c\xb9\x77\x4b\xf3\x00\x43\xc3\xbe\xd5\xc1\x7e\x63\x69\x3d\xd5\xc4\x44\xc4\x78\x9e\x5e\x74\xf5\xee\xc3\xf4\xfa\x8f\x01\xf2\xe2\xbf\xba\xef\x3e\xbf\xbe\xe7\xbc\xd5\x6a\x5c\x3f\xd9\x19\x91\x19\xdf\xe1\x1e\x70\x7b\x3d\xe0\x1d\xca\x3e\x87\x58\x78\xf5\x0b\x8b\x63\x3d\xf5\xaa\x8a\xd4\xcb\x7a\xd7\x21\xff\x34\xe1\x3f\xe6\x78\xd1\xcb\xf1\xc3\xc7\xeb\xf7\x7e\x82\xee\x56\xd8\x59\x52\x48\x36\x90\x7e\xaa\x2e\x62\xca\xe2\x30\xb8\x8d\x71\x40\x3a\xc1\x0e\x82\xd1\xb5\xbb\x19\x79\xd6\xbd\xeb\xa4\x85\xa9\xd2\x34\x52\xcc\x42\x44\x51\x74\xd0\x82\xd0\xac\xd1\x00\xb4\x1b\x3c\x01\x72\x3b\xbd\xda\x8e\xc2\x26\xeb\x32\x6d\x27\xdb\xdb\xe9\x87\xf7\x51\xc9\x24\x96\x13\x3d\x7a\xf4\x86\x77\x30\xec\x02\x1c\x6b\x7b\x08\x70\xc3\xce\xac\x4d\x0e\x34\xd9\xa3\xd7\x91\x76\xaf\x68\x26\x52\xf3\xa9\x73\xdc\x8c\xe1\x83\x8e\x03\xc2\x8e\x1b\xdd\xed\xe5\xe0\xf0\x1e\xf8\xa1\xcc\x25\xb6\xed\xca\x65\xd5\x6d\xed\x76\xc6\xb9\x82\x15\x6c\xbd\x7c\x55\xa2\x28\x29\xbb\x88\x95\x65\x5e\x87\x4d\x3c\x23\xf0\x22\x6d\x77\x25\x9a\x6f\x4d\xed\xd7\xc4\xf7\x0d\xe2\x98\x58\xed\x0c\x3f\x1f\xaa\x9c\x49\x0f\xbe\xf6\xe2\xd2\x18\x47\xcf\xcc\xee\xb6\x86\x0d\x08\x82\x9d\x89\xd8\xe9\xf5\x46\x68\xcf\x55\x7b\x4b\xf9\xae\x23\xa7\x15\xf6\x59\xbe\xbb\xf8\xb8\x4c\x9c\xa0\xcb\x45\x55\x49\x82\x50\x9d\x57\xf8\x59\xac\x16\x40\x97\x6f\xbc\x7b\x5b\xce\x4e\x44\x45\x93\xc1\x77\xb8\x37\xc2\x3d\xc8\xd3\xd5\xff\x70\xb7\xf8\x17\xa3\xee\xca\xe2\x99\x1b\xce\xf9\x31\xfb\xfd\x51\x62\xe3\x3a\xbd\x80\x97\xbb\xf4\xdc\x8d\x20\xe4\x0c\xff\x8b\x6b\x0f\x9e\xee\xde\x63\x67\x62\x8f\xcd\xdc\xde\xbd\xa9\xe8\xdd\x72\x9e\xe4\x0e\x63\x43\x70\xf7\x29\xc0\xef\xb6\x38\x26\x1d\x25\x72\x1e\x6d\x98\x2c\xbc\x76\x33\x5f\x76\xfe\x9f\x88\x32\x03\x51\xf7\xf9\xd4\xfc\x85\xc8\x3c\x27\x47\x8f\x7b\x7f\xf2\x42\xd8\x16\xf8\xcd\x81\x29\xa6\x03\xf4\xf8\x4f\x00\x00\x00\xff\xff\x81\xcc\x8e\x48\xa9\x13\x00\x00")

func veloxJsBytes() ([]byte, error) {
	return bindataRead(
		_veloxJs,
		"velox.js",
	)
}

func veloxJs() (*asset, error) {
	bytes, err := veloxJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "velox.js", size: 5033, mode: os.FileMode(420), modTime: time.Unix(1456065893, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsonPatchJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func jsonPatchJsBytes() ([]byte, error) {
	return bindataRead(
		_jsonPatchJs,
		"json-patch.js",
	)
}

func jsonPatchJs() (*asset, error) {
	bytes, err := jsonPatchJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "json-patch.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(1456039442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _eventSourceJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x58\xff\x4f\x1b\x3b\x12\xff\x3d\x7f\x85\x41\xa7\xee\xe6\x48\x96\xb4\xd7\xd3\x49\xc9\x71\x08\xd1\xa8\x45\xa2\xe1\x54\x52\x5d\xa5\xd2\xab\x4c\xd6\x24\x7e\x6c\xec\xbc\xb5\x97\x80\x20\xff\xfb\x9b\xb1\xbd\xbb\xf6\x26\x4b\xa1\x11\x22\x59\x7b\x3c\xdf\x67\xfc\x99\x1d\xc5\x37\x85\x98\x69\x2e\x05\x89\xe7\x99\xbc\xa6\x59\x97\x3c\x76\x3a\xfc\x86\xc4\xfb\xe3\x3b\x26\xf4\xa5\x2c\xf2\x19\xdb\x27\x5c\x90\x72\x3f\x67\xba\xc8\xc5\xa8\xd3\xb9\xa3\x39\x3c\x4c\x73\xbe\x24\x47\xe4\xf0\xff\xf1\x95\x7a\xba\x2a\x06\x83\x93\x41\xf7\xe0\xc9\x7f\xf8\xdb\xe1\xdc\x51\x7b\x2c\xe1\x48\x2d\xba\xc8\x8d\x5c\x42\x90\x88\x21\x91\x2a\x89\xf4\x82\xab\x1e\xec\xe0\x87\x0b\xcd\xf2\x3b\x9a\xc1\xf2\x3f\x07\x83\x1e\x39\x3c\x24\x2b\x99\x65\x5c\xcc\xab\x2d\x47\x99\x51\xa5\x8d\xb0\xb3\x14\x88\x45\x91\x65\x25\x8f\x19\x9d\x2d\x90\x6f\x14\x81\x4e\xc0\x12\x2c\xdd\x03\xf1\xe4\xe9\x89\xe8\x87\x15\x93\x37\x04\x9f\xf6\x80\x40\xe9\x1c\x38\x47\x56\x31\x02\x8a\xe4\x72\x4d\x04\x5b\x93\xcb\x07\xa1\xe9\xfd\x38\xcf\x65\x1e\x47\x13\xa9\x09\x13\xb2\x98\x2f\x08\xcd\xe7\xc5\x12\x75\x8f\xba\x23\x38\xb2\x41\xfe\xa8\x7e\xf2\xf5\xcb\x39\x48\x04\xbe\xa3\x72\x25\x67\x34\x7d\xb8\xd4\x54\x97\x26\x26\xa7\x17\x93\xc9\xf8\x74\x7a\x36\xf9\x58\x11\xfd\x44\xe3\xa6\x7c\xc9\x72\x67\x43\xbd\x73\xbf\xa8\xd7\x60\xb1\xf2\x24\x9e\x38\x99\x53\x2e\xe2\xd2\x21\xa5\xfa\x9e\x57\x43\xc6\x8a\x69\xfc\x29\x0b\xed\xe5\x42\x79\x8a\x18\x8e\xc9\x8c\x66\x59\xec\x71\x30\xf6\x81\x85\xbd\xca\xef\xb5\xc5\x81\x2e\x15\x23\x9d\x3f\x90\x47\x0c\xd8\x8d\xc4\xb8\x2e\x78\x8a\x51\x03\x6f\x6b\x88\x06\x43\x57\x92\x25\x53\x8a\xce\x59\x92\x24\xc0\x55\x51\xc1\x8e\xcb\xb0\x43\x8c\x7c\xf5\x7d\xe7\x1d\x05\x86\x9d\x9e\x5f\x5c\x8e\x3f\x78\x29\x6a\x19\x80\xd8\xc9\xc5\x74\x3c\x24\x67\xe3\x7f\x11\x2a\x52\x52\xac\xd6\x34\x4f\x15\x51\xc5\x6a\x25\x73\xed\xc8\x30\xf9\x9c\x63\x21\xca\xdf\x3e\x9f\x7f\xd2\x7a\xf5\x85\xfd\x59\x30\xa5\x63\x67\x32\x41\x8a\x44\xae\x98\x88\xa3\x8f\xe3\x69\xd4\x0b\xe4\x43\xa0\x7b\x60\x6a\xc1\x02\x6a\xf0\xb0\xe3\xf2\x09\x34\x67\x90\x34\x27\xb3\x19\x5b\x69\x38\x1c\x69\x76\xaf\x0f\x0d\x8b\x3e\xe4\x1b\xa3\xcb\xe8\x17\x47\x4f\x31\x7f\xfb\xa7\x52\xe8\x5c\x66\xc8\x41\xc8\xbe\xc9\xe9\xfa\x20\x98\xbb\x66\x64\x59\x28\x4d\x96\xf4\x96\x91\x42\x31\xeb\x69\xae\x08\xc4\x05\x3d\xae\x20\x68\x10\x7d\xc5\x53\x86\xee\x5d\xb3\x28\x67\x64\x2d\xf3\x5b\x8c\xca\x9a\xeb\x05\x39\x11\x69\x2e\x79\x4a\xfa\xe4\x9a\xcd\x28\xb2\x80\x73\x0f\x24\x95\x22\xd2\x60\x22\x9f\xcf\x59\x5e\xcb\x33\x21\x99\x2d\xa8\x98\x83\x38\xa1\x79\xe6\x4b\x99\x49\x21\x98\xcd\x09\xd0\x60\x96\x49\xc5\xd2\x67\x6d\xfc\xd6\x77\x0b\x2c\xed\xff\x0f\x94\x41\x33\xc3\x70\xa0\xb1\x5e\x76\xf8\xe5\xbe\x67\xeb\xa2\xdb\xc2\xfb\x1c\x48\xfb\x86\xb6\x7f\xf6\x01\x18\x7b\x47\x2b\x07\x36\x7a\x44\xa9\xa8\xb6\x65\x62\xbb\xcf\x60\x10\x24\x84\x30\x1e\x50\x98\x94\xce\x0d\x7e\x7f\xab\xcb\xc9\xea\xbb\xd5\x03\x8e\xc8\x3f\xb0\x03\xed\xda\x78\x4f\xde\xbc\xb1\x55\x8f\xdc\x0b\x85\x6b\xef\x06\x83\xae\xcf\xd3\xc4\x00\x04\xa9\x02\x12\x4b\x29\x6f\xfd\x15\xb5\x53\x35\xa0\x90\x73\xd8\x3a\x82\xc6\xe5\x6f\x5c\xfc\x77\x3c\x19\xb5\x1e\x4b\xb9\x5a\x51\x3d\x5b\x18\x4f\xc7\x11\xd6\x0f\xf8\xfe\xd1\xb4\xdc\x21\xb1\xcf\x64\xd3\xf5\x19\x6c\x3a\xde\x83\xbd\x6a\xd4\x4a\x0a\xc5\xa6\x50\x32\x2e\x38\x35\x81\xe9\x2f\x81\xf8\x06\xb9\xf3\xac\xb7\x06\xfe\x0e\x79\x6c\x20\xf0\xa0\x23\x38\x0c\x3c\x10\x88\xc7\x8b\x26\x97\xe8\xda\x6d\x3e\x0d\x2d\x57\x34\xd7\x10\xa3\x40\x7c\xa2\x8a\x6b\xa8\xee\xd8\xe4\x55\x92\x31\x31\xd7\x8b\x6e\xa2\x56\x19\xd7\xf1\xfe\x95\xd8\xef\xf6\x02\xcd\x9d\xeb\xa6\xe0\x1b\xb4\xd3\xf5\xc5\xa8\x49\x94\x52\x4d\x61\xff\xfb\x8f\xe6\x06\x87\xd5\x41\x73\x11\x6e\xc9\x46\x4a\xe3\xa7\xcc\x74\x5f\xdb\x51\xc3\xf2\xe9\xc5\x87\x0b\x02\x39\x9d\x66\x8c\x44\x46\xb3\x88\xc4\xd0\xc6\xc9\x75\x71\x73\x03\xe5\x2d\xe8\x92\x75\x7b\xd8\x73\xf3\x07\xef\x24\x52\xc4\x23\x50\xe6\xdf\xd6\x27\xce\x6e\x58\x39\x38\x68\x66\x98\x53\xce\xd0\x7d\xe7\x3f\xc0\xbf\xab\x8c\xce\x58\x6c\xd1\x05\x54\x7f\xd4\x0d\x73\xcb\x14\x3d\x1c\x4a\xb8\x48\xd9\xfd\xc5\x4d\xec\x14\xeb\x62\x56\x0f\x9a\xec\x43\x87\x9a\x73\xa5\x04\xdb\x7b\x87\xc7\x57\xea\xef\x87\x3b\xe4\x6c\x08\xcb\x14\xdb\x21\xce\x58\xdb\x2a\xce\xec\x5a\x83\x14\x3b\x83\x8c\x0f\x65\x9a\x6d\x4f\x66\x43\x28\x9a\x17\xef\x71\x35\xa1\x93\xd8\x90\x62\xad\xfb\xe0\xc7\x2c\x8e\xa0\x42\x5e\xa4\x2a\xe6\x49\xab\xa6\xb8\x99\xac\x0a\xb5\x68\xa8\x88\xeb\xed\x1a\xb6\xca\xe2\xe9\xb0\x55\x54\x88\xc9\x42\x71\x70\xee\xd5\x21\xe0\x69\x2d\x0a\xf3\xd4\xdc\x6f\x90\xc8\x4c\x2b\x73\xfb\xf0\xf4\x59\x05\x4a\x40\xd5\x2a\x09\x79\x47\xd1\xb6\x1d\xb8\x6f\xdc\xe6\x0a\x79\x8b\xc0\x83\xb1\x0e\x4b\x7c\xb6\x25\x6c\xbb\x9f\x39\xfb\x87\x04\x8c\x16\x5d\x89\xa8\x1b\x22\x08\x00\x89\xbb\xef\xa4\xfa\xd3\xde\x57\xab\x2c\x77\x3c\x77\x1c\xae\x9a\x46\x0b\xdf\x66\xd3\x69\x92\x35\x72\xae\xad\x5f\xb7\x5c\x72\xef\xbb\xbb\x10\xea\x28\xec\x37\x16\x62\x08\xc6\x52\xa2\xa5\xa1\x27\x14\x0f\xf4\x2a\x24\x62\xe1\x4a\x26\xc5\xbc\x9f\x49\x8a\x48\xb2\xb3\x23\x86\x2d\xd7\xd6\x5e\x0b\x66\x7c\x7c\x91\xfa\x26\xd3\x10\x42\x2a\xb9\x04\x54\x05\x79\x06\x28\xca\xdc\xcb\x81\x63\xd0\x0c\x17\x1a\x8b\x6e\x5f\x7f\xa1\x86\x03\xc1\x4b\xc2\x1f\x19\x49\xfe\xbd\x6a\x17\x1a\x17\x2b\xf9\x55\x0c\x3c\x1f\xee\x70\x42\x59\x6e\x19\xbf\x65\xd9\x03\xa1\xd7\x00\xa0\x59\xfa\x1b\xfc\x1f\x3b\xbb\x12\xa9\xfc\xb5\x09\xa0\x97\x62\x22\x8d\x6b\xcc\xf7\x8b\x99\xc5\x05\x10\x70\x38\x5e\xf2\x78\xbe\x81\xb5\x2c\x38\x34\xaa\xd7\xc0\x1e\xa6\x19\x0f\xe5\xd5\xc2\x82\xc9\xc9\xce\x07\xf0\xdf\x6d\x07\x98\x01\xdd\x02\x83\x32\x85\x3f\xad\xd9\x72\xa5\x31\x81\x15\x87\x3e\x31\x63\xf5\xa0\xa3\x76\xf0\x7d\x59\x20\x7b\xa6\x7c\x87\x84\x25\xae\x3c\x31\xb2\x28\xf4\xf8\xd8\xce\x4a\xe8\x3c\xeb\x38\x3b\x7c\x8d\xac\x46\x1c\xca\x49\xae\x3b\xb8\xe3\x4d\xe0\x09\x40\x1a\x2d\xb5\x2d\x78\x74\x9d\x41\xe6\xc3\x1d\xb0\x15\x98\x98\x3d\xdb\x57\x3d\x38\xdf\xc7\x2c\xa7\xd7\x66\x02\xc7\x2d\x37\x8d\xbb\x79\x79\xf7\x9c\x6b\xca\xcd\xfa\x7c\x96\x31\x9a\x9f\xb9\x14\x89\x1b\x33\xaf\x8b\x4b\x35\xef\xfa\xe1\xda\x20\xba\xa9\x0b\x64\x68\xe1\x0e\x42\xd0\x21\x79\x6b\xf6\x8c\x94\x21\x79\x87\x0f\x81\x73\x7d\xfb\xb4\xd7\x2a\x9d\xa9\xd8\xb9\x2d\xd6\xc9\x95\x53\xf9\x7b\xf4\x33\x22\x07\x26\x12\xf0\x15\x7d\x72\xbb\x91\x6b\xa1\x98\x6a\xe5\x89\x3a\x05\x0d\xfa\x41\x66\x06\x8c\x59\x18\x54\x52\xb5\x21\xa1\x72\x1f\x11\x90\x99\xb8\xcd\xbb\x8f\x46\x2b\xdf\xb8\x40\x57\xb2\xad\x8a\xd0\x33\x9d\x8e\x3f\x6a\x96\x5b\x5b\x6d\x6c\x37\xa5\x4f\x69\x9a\x1a\x2f\x9d\x73\x18\xbe\x04\xcb\xb7\xbd\xe5\x94\x2c\x85\x98\xf7\x27\xcf\x7a\xa9\xa9\x4e\x0b\x99\x77\x29\x39\xe3\x9e\x25\xb7\x98\xa5\x54\xa6\xca\x89\x9c\x2d\xe5\x1d\x7b\x9d\x09\xbf\x17\xf2\xbd\xed\x98\x97\x6f\x1d\xea\x30\x05\x59\xd0\x08\x3f\x14\xcf\x5b\xcc\x8b\xff\x98\xfc\xe8\xf7\x79\xcd\xc8\x4f\x29\x48\x06\x68\x59\x47\x4d\xa5\xfd\x7c\x31\x93\x04\xa0\x28\xde\x23\x6f\xbd\x46\x7b\x0d\xc5\x77\xdb\x4c\x1b\xe7\x27\x29\x4c\x47\x19\x56\xaf\xc6\xa4\x70\x2d\xc5\x5f\xc2\xb1\xac\x7e\xae\x6b\xd9\xd5\xdb\xd7\x2f\xe7\xd0\x9a\x22\xd3\x57\xd0\x48\x1f\xe5\x04\xe3\x2f\x76\xad\x1e\x91\x39\x9f\xe3\x2d\xee\xa3\x1b\x63\x8d\x29\x71\x07\x4c\xf0\xab\x7a\xcd\x65\x4f\xc0\xaa\xfd\x51\xad\x37\x90\xa4\xf7\xe4\x66\x3a\xd4\xc8\xd7\x66\xab\xd5\xd9\x3e\x5a\x5a\xe6\x9a\xac\x3f\x68\x79\x4c\xd1\x46\xe3\x0f\xa3\x44\x65\x31\x06\x29\x5a\xca\xb4\xc8\x58\xe4\xbf\x1e\xb5\x4b\x09\xbb\xc7\x37\x4b\x98\x50\x5e\xcb\x1d\x75\x2c\x55\x12\xbe\x08\x0d\x28\x3a\x9b\xae\xa9\x50\x88\xe4\x5f\x01\x00\x00\xff\xff\x52\x76\x4d\xb8\x9e\x15\x00\x00")

func eventSourceJsBytes() ([]byte, error) {
	return bindataRead(
		_eventSourceJs,
		"event-source.js",
	)
}

func eventSourceJs() (*asset, error) {
	bytes, err := eventSourceJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "event-source.js", size: 5534, mode: os.FileMode(420), modTime: time.Unix(1456039471, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"velox.js": veloxJs,
	"json-patch.js": jsonPatchJs,
	"event-source.js": eventSourceJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"event-source.js": &bintree{eventSourceJs, map[string]*bintree{}},
	"json-patch.js": &bintree{jsonPatchJs, map[string]*bintree{}},
	"velox.js": &bintree{veloxJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

