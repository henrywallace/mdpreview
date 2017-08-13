// Code generated by go-bindata.
// sources:
// static/index.html
// static/preview.js
// DO NOT EDIT!

package server

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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4f\x6f\xd4\x30\x10\xc5\xef\xf9\x14\x83\x4f\x20\x91\x78\x8b\x28\x85\xc5\x8e\x90\x80\x33\x1c\xb8\x70\xf4\xda\xd3\x64\x68\x62\x5b\xf6\xec\x3f\x56\xfb\xdd\x91\x93\x6c\xbb\x05\x91\x8b\x95\x79\x99\xf7\x9b\x89\x9f\x7a\xf1\xe5\xdb\xe7\x1f\x3f\xbf\x7f\x85\x9e\xc7\xa1\xad\x54\x39\x60\x30\xbe\xd3\x02\xbd\x68\xab\x4a\xf5\x68\x5c\x5b\x01\x00\x28\x26\x1e\xb0\x3d\x9d\xa0\x89\x86\x7b\x38\x9f\x95\x9c\x4b\xb3\x3c\x90\x7f\x80\x84\x83\x16\x64\x83\x17\xc0\xc7\x88\x5a\xd0\x68\x3a\x94\xd1\x77\x02\xfa\x84\xf7\x5a\xdc\x9b\x5d\xd1\x9b\x52\x6a\x2b\x25\x67\x80\x1a\x91\x0d\x78\x33\xa2\x16\x3b\xc2\x7d\x0c\x89\x05\xd8\xe0\x19\x3d\x6b\xb1\x27\xc7\xbd\x76\xb8\x23\x8b\xf5\xf4\xf2\x1a\xc8\x13\x93\x19\xea\x6c\xcd\x80\xfa\xa6\x98\x3d\x8d\x90\xf9\x38\x60\xee\x11\xf9\xc2\xed\x99\x63\x5e\x4b\x99\xcc\xbe\x23\x6e\x6c\x18\x65\x26\xef\x12\xe6\x90\xfa\x6d\x96\x1d\x71\xbf\xdd\xd4\xa3\x49\x0f\x2e\xec\x7d\x6d\x73\x96\x5d\x5f\x47\xd3\xe1\x3f\x62\x63\x73\x2e\xbc\x89\x32\xaf\xdf\x3c\x36\x6e\x82\x3b\xc2\x69\x2a\x96\x67\x13\x0e\x75\xa6\xdf\xe4\xbb\x35\x6c\x42\x72\x98\xea\x4d\x38\x7c\x7c\xd4\x47\xf2\xf3\x46\x6b\x78\xb3\x5a\xc5\x6b\xc5\x1c\x2e\xca\x87\xf7\x7f\x29\xa9\x23\xbf\x86\x15\x98\x2d\x87\xa7\x7a\x34\xce\x4d\xa0\xb7\xb7\x97\xcf\xcf\xd5\x74\x7c\x1a\xd1\x91\x81\x97\x57\x9e\x77\xef\xee\xe2\xe1\xd5\xd5\xa4\xff\x5d\xe1\x99\xf5\xcd\xed\xf5\x24\xe7\x05\xa2\xe4\xf2\x2f\x2a\x55\x9a\x97\x48\x98\xc4\x64\x07\x04\x72\x5a\xc4\x84\xe5\x62\x05\xd8\xc1\xe4\xac\xc5\x33\xd8\x92\x96\x29\x85\x4a\x2e\x6d\x8b\x49\xb6\x89\x22\x43\x4e\x56\x0b\xb9\xb8\x34\xbf\xb2\x68\x95\x9c\xa5\x92\xa2\x19\x5a\xe2\x54\x2c\xfe\x04\x00\x00\xff\xff\xdd\x90\x71\xa2\xd8\x02\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 728, mode: os.FileMode(420), modTime: time.Unix(1502652820, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticPreviewJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x41\x4b\x33\x31\x10\x86\xef\xfd\x15\x2f\xbd\xec\x2e\x85\xf4\xfe\x95\x5c\x3e\x11\x14\xf4\xa4\xe0\x39\x26\x63\xbb\x98\x9d\x29\xc9\xec\x46\x91\xfe\x77\xc9\xba\xb5\x1e\x3c\x98\x53\x48\x9e\xf7\x99\x79\xdb\x97\x91\xbd\xf6\xc2\x68\x3b\x7c\xac\x00\x60\x72\x09\x63\x8a\xb0\x68\x4a\xfe\xb7\xdd\x36\xd8\xa0\xf4\x1c\xa4\x98\x28\xde\x55\xd8\x1c\x24\xeb\x2f\xcf\x47\xa7\x07\x76\x03\x61\x53\xb3\xcd\xee\xdb\x77\x4c\x34\xf5\x54\x60\x11\xc4\x8f\x03\xb1\x9a\x3d\xe9\x75\xa4\x7a\xfd\xff\x7e\x1b\xda\xf5\x82\xac\xbb\x4b\xca\x0b\x33\x2c\x98\x0a\x9e\xe8\xf9\x41\xfc\x2b\x69\x3b\xa6\xd8\xed\x56\x33\x53\xff\x8d\xb0\x8f\x92\x09\x16\x97\x2a\x34\x11\xeb\xb9\x4f\x3d\x8b\xdc\x28\xbd\xe9\x95\xb0\x12\x6b\xed\x57\x05\xf4\x95\x99\x25\x61\x59\xf9\xf4\x53\x3f\x50\xce\x6e\xff\xc7\x01\x3d\x33\xa5\x9b\xc7\xfb\x3b\x58\xcc\x90\x09\x4e\xdd\xd9\x7a\xea\xda\xee\x33\x00\x00\xff\xff\xe4\x18\x5f\x51\x72\x01\x00\x00")

func staticPreviewJsBytes() ([]byte, error) {
	return bindataRead(
		_staticPreviewJs,
		"static/preview.js",
	)
}

func staticPreviewJs() (*asset, error) {
	bytes, err := staticPreviewJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/preview.js", size: 370, mode: os.FileMode(420), modTime: time.Unix(1502652745, 0)}
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
	"static/index.html": staticIndexHtml,
	"static/preview.js": staticPreviewJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"preview.js": &bintree{staticPreviewJs, map[string]*bintree{}},
	}},
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
