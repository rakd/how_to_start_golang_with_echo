// Code generated by go-bindata.
// sources:
// data/templates/dashboard.html
// data/templates/login.html
// data/templates/welcome.html
// DO NOT EDIT!

package main

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

var _templatesDashboardHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x49\x2c\xce\x48\xca\x4f\x2c\x4a\x51\xaa\xad\xe5\x82\x73\xb8\x6c\x32\x8a\x14\xf4\xed\xb8\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xf4\x95\xec\x32\xf2\x73\x53\x6d\xf4\x13\xed\xb8\x6c\x92\xd0\xe4\x72\xf2\xd3\xf3\x4b\x4b\x94\xec\x20\x34\x58\x0d\x57\x75\x75\x6a\x5e\x4a\x6d\x2d\x17\x20\x00\x00\xff\xff\x64\x0c\x30\xe4\x6a\x00\x00\x00")

func templatesDashboardHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDashboardHtml,
		"templates/dashboard.html",
	)
}

func templatesDashboardHtml() (*asset, error) {
	bytes, err := templatesDashboardHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dashboard.html", size: 106, mode: os.FileMode(420), modTime: time.Unix(1504058636, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xbd\x6e\x84\x30\x10\x84\x7b\x9e\x62\xb4\x0f\x10\xfa\x13\xa6\x4b\x97\x22\x6f\x10\xf9\xce\x4b\xb0\x82\x6d\x64\x9b\xfc\x68\xc5\xbb\x47\x98\x73\x08\xba\x6e\x76\xbd\xf3\x69\x3c\x22\x86\x07\xeb\x19\x34\x85\x77\xeb\x69\x5d\x1b\xa0\xc8\x06\xe8\xc6\x88\xb6\x6f\x36\x35\x84\xe8\xa0\x6f\xd9\x06\xaf\xa8\xdd\x6f\xe1\x38\x8f\xc1\x28\x9a\x43\xca\xd4\x37\x00\x20\x62\x07\x3c\xdd\x52\x1c\xde\x72\xf8\x60\xbf\xae\x9d\xf5\xf3\x92\x91\x7f\x66\x56\x34\x5a\x63\xd8\x13\xbc\x76\xac\xe8\x38\x23\x7c\xea\x69\x61\x45\x22\x27\x33\xa1\xed\x45\xd8\x9b\x92\x0b\x78\x76\xda\x4e\x17\x9c\x98\x99\xbf\x73\x25\xf2\xf6\xfe\x1f\x56\x16\x3b\xa7\x00\xba\x6b\xac\xf2\x55\xa7\xf4\x15\xa2\xb9\x9c\x68\xf3\x7d\x5b\x89\xc7\xfc\x48\x38\x19\xd3\x72\x75\xf6\x2f\x48\x9d\xee\x49\x5e\xf6\xc2\x8a\xaf\x6b\xb7\x32\xfb\xa3\xdf\xfa\xc1\xdf\x00\x00\x00\xff\xff\xce\x59\xc5\xb4\x8c\x01\x00\x00")

func templatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoginHtml,
		"templates/login.html",
	)
}

func templatesLoginHtml() (*asset, error) {
	bytes, err := templatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/login.html", size: 396, mode: os.FileMode(420), modTime: time.Unix(1504057461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesWelcomeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\x4b\x0a\x03\x21\x0c\x06\xe0\xfd\x9c\xe2\xc7\x03\xd4\x0b\x58\xaf\x52\x9c\x1a\xc7\x80\x35\xa0\x2d\x5d\x84\xb9\x7b\xc1\xb1\x0f\xe8\x2a\x0f\xbe\xfc\x51\x8d\x94\xb8\x12\xcc\x93\xca\x55\x6e\x64\xf6\x7d\x01\xe6\xb0\x00\x2e\x37\x58\xbf\x00\xaa\x9c\x70\xe2\x7e\x29\xb2\x71\x1d\x0a\x70\x01\xb9\x51\x3a\x1b\x1b\x43\xcf\xab\x84\x16\x8d\xff\xb4\xce\x06\x7f\xb0\x75\x86\xfc\x5e\x14\xd9\xe4\x71\x37\xfe\xa8\xd3\xaa\x52\xe9\xf4\x97\x3e\x7e\x0e\xca\xf5\x2b\x6b\x1c\x70\xa6\xbf\x17\xaf\x00\x00\x00\xff\xff\x10\x68\xab\x71\xd3\x00\x00\x00")

func templatesWelcomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesWelcomeHtml,
		"templates/welcome.html",
	)
}

func templatesWelcomeHtml() (*asset, error) {
	bytes, err := templatesWelcomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/welcome.html", size: 211, mode: os.FileMode(420), modTime: time.Unix(1504058739, 0)}
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
	"templates/dashboard.html": templatesDashboardHtml,
	"templates/login.html": templatesLoginHtml,
	"templates/welcome.html": templatesWelcomeHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"dashboard.html": &bintree{templatesDashboardHtml, map[string]*bintree{}},
		"login.html": &bintree{templatesLoginHtml, map[string]*bintree{}},
		"welcome.html": &bintree{templatesWelcomeHtml, map[string]*bintree{}},
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

