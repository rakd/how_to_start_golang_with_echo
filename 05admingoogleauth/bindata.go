// Code generated by go-bindata.
// sources:
// data/templates/admin_index.html
// data/templates/dashboard.html
// data/templates/login.html
// data/templates/partial_header.html
// data/templates/partials_admin_header.html
// data/templates/partials_flash.html
// data/templates/signup.html
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

var _templatesAdmin_indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x6a\xc4\x20\x18\x84\xef\x3e\xc5\xf0\x3f\x40\x65\xef\xd6\x57\x09\xb6\x4e\x50\xf8\x37\x59\xd4\xb0\x05\xc9\xbb\x97\x98\x6e\x2f\x85\x42\x2e\xea\x8c\xdf\xcc\x61\x7a\x8f\x9c\xf3\x42\x48\x88\xf7\xbc\x4c\x79\x89\xfc\x92\x7d\x37\x40\xef\x68\xbc\x3f\x34\x34\x42\x1e\xa1\xb4\x1c\xb4\x4e\x27\x96\x18\x22\x8b\xe0\x0d\x03\x75\xe9\xe6\xc7\x87\xb3\xe9\xe6\xcd\xe1\x6c\xea\x0d\x00\x38\xcd\xe7\x03\x70\x01\xa9\x70\x7e\x17\x3b\x58\xbb\x55\x96\x2a\x7e\x5c\xce\x86\x1f\xde\xbe\x02\xff\x24\x3f\xd6\x56\xc5\x1f\xe7\xb5\x5c\xfd\x4c\x8c\x9b\x32\x4e\xed\x49\x1e\x1d\xbf\x0e\x4e\xe7\x5a\xdf\xbc\xaa\xae\x4f\x96\xa9\x6c\xca\x2a\xfe\xa5\x31\xf4\x9f\x2e\x67\x8f\x55\xc6\x5e\x05\xd6\x1b\xd3\x3b\x97\xb8\xef\xe6\x3b\x00\x00\xff\xff\xf1\x2f\x1b\x7f\x86\x01\x00\x00")

func templatesAdmin_indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesAdmin_indexHtml,
		"templates/admin_index.html",
	)
}

func templatesAdmin_indexHtml() (*asset, error) {
	bytes, err := templatesAdmin_indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/admin_index.html", size: 390, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDashboardHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcb\x41\xaa\xc3\x30\x0c\x04\xd0\xbd\x4f\x21\x74\x80\xef\x0b\xe4\xfb\x2a\x45\x41\x93\x2a\xe0\xc4\x41\x75\x57\xc2\x77\x2f\x24\xa1\xab\xae\xa4\x61\xde\x44\x28\x96\x75\x07\xb1\xca\xcb\xe6\x26\xae\x3c\x46\x22\x22\x8a\xe8\xd8\x8e\x2a\x1d\xc4\x87\x78\x5f\xa5\x3e\x0c\xa2\x70\xa6\xbf\xdb\x7c\x47\x67\x9a\xcc\x29\x97\xeb\x15\x32\xc7\xf2\xcf\x99\x8b\xb5\x0d\x53\x96\xbb\x98\x7f\x98\xda\x9e\xed\xdd\xb9\x5c\xf7\xb4\x29\x02\xbb\x8e\x91\x3e\x01\x00\x00\xff\xff\x82\x6d\xef\xe2\xa2\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/dashboard.html", size: 162, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x4b\x6e\xb4\x30\x10\x84\xf7\x9c\xa2\xd4\x07\x18\xf6\x23\x60\xf7\xef\xfe\x45\x6e\x30\xf2\x8c\x9b\x60\x05\x3f\x64\x37\x79\xc8\xe2\xee\x11\x66\x1c\x82\xb2\xeb\x6e\x57\x7d\x2a\x57\xce\x9a\x47\xe3\x18\x34\xfb\x57\xe3\x68\x5d\x1b\x20\x67\x61\x1b\x66\x25\x0c\x0a\x2a\x8a\x51\xf3\x6d\x62\xa5\x39\x12\x2e\x45\x51\xc4\x0d\xd0\x4d\x11\xed\xd0\x6c\xd3\xe8\xa3\x85\x7a\x88\xf1\xae\xa7\x76\xa7\xc1\xb2\x4c\x5e\xf7\x14\x7c\x12\x1a\x1a\x60\x83\x9b\x11\x97\x47\x8a\xe3\x4d\xfc\x1b\xbb\x75\xed\x8c\x0b\x8b\x40\xbe\x02\xf7\x34\x19\xad\xd9\x11\x9c\xb2\xdc\xd3\x21\x23\xbc\xab\x79\xe1\x9e\x72\x3e\x99\x09\xed\x90\x33\x3b\x5d\x72\x01\xff\xac\x32\xf3\x15\x27\xa6\xf0\xa7\x54\x22\x6f\xef\xbf\x61\xe5\xb0\x73\x0a\xa0\xbb\xc7\x3a\xbe\xa8\x94\x3e\x7c\xd4\xd7\x13\x2d\x3c\xaf\x95\x78\xec\x7f\x09\x27\x63\x5a\xee\xd6\xfc\x04\xa9\xdb\x33\xc9\xff\xbd\xb0\xe2\xeb\xda\xad\xcc\xe1\xe8\xb7\x7e\xf0\x3b\x00\x00\xff\xff\x48\xc3\x4e\xde\xae\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/login.html", size: 430, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartial_headerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x4f\x8e\x87\x20\x0c\x85\xf7\x9e\xa2\xe9\x01\x7e\x5e\xc0\xe1\x2a\xa6\x13\x8a\x34\x41\x30\x80\x2b\xe4\xee\x13\xff\xa1\x9b\xc9\x0c\x9b\xd7\x84\xef\xf5\x4b\x5a\x8a\x66\x23\x9e\x01\x17\x8a\x59\xc8\x8d\x96\x49\x73\xc4\x5a\x3b\x00\x80\x81\xc0\x46\x36\x5f\xd8\xa3\xb2\x61\xe6\xa1\x27\x75\x7c\x94\x22\x06\x3e\x92\x46\x17\x26\xf1\x17\xbd\xbf\xad\x4d\x4f\x57\x53\xb2\xdf\x81\xa2\x46\xd5\xc6\xb6\xe9\xb7\x92\x0b\x53\x58\x33\xaa\x33\x5f\x62\x76\x89\xff\x10\x26\x99\xfc\xba\xa0\x3a\xf3\x3f\x2a\xf1\x87\x49\xfc\x5b\xe4\xf5\xe5\xd9\x5e\x30\xe9\x79\x87\x8f\x68\xf0\x79\x2c\x1b\xa1\xbf\xcb\x99\xe7\xc5\x51\x7e\x2e\x9b\x46\xe3\x28\x59\x84\x4f\xad\x5d\x77\x6f\xff\x09\x00\x00\xff\xff\x02\x3e\x5c\x4a\x81\x01\x00\x00")

func templatesPartial_headerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartial_headerHtml,
		"templates/partial_header.html",
	)
}

func templatesPartial_headerHtml() (*asset, error) {
	bytes, err := templatesPartial_headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partial_header.html", size: 385, mode: os.FileMode(420), modTime: time.Unix(1504103954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartials_admin_headerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4d\xaa\xc3\x30\x0c\x84\xf7\x3e\x85\xd0\x01\x9e\x2f\x90\xe7\x5d\xa1\x9b\xd2\x23\x04\x81\xe5\x1f\x70\xec\xe0\x38\x2b\x55\x77\x2f\x49\xbb\x08\x74\x25\x34\x33\x1f\x33\x22\x9e\x43\xae\x0c\xb8\x52\x1f\x99\xca\x36\x93\x5f\x72\x9d\x13\x93\xe7\x8e\xaa\xc6\x18\x80\x89\x20\x75\x0e\xff\x68\xd1\xdd\x9f\x8f\xdb\x64\xc9\xc1\xeb\x22\x9f\x10\xba\xf3\xfc\x9a\xfb\x48\x36\xb6\x16\x0b\xdb\xd2\x62\xdb\xc7\x37\x09\x9f\xef\x00\x8e\x16\x91\xc1\xcb\x5a\x68\x5c\xe7\x84\x42\x5b\x42\xf8\x53\x35\x22\x5c\xbd\xaa\x79\x07\x00\x00\xff\xff\x03\x41\xa6\x9e\xb5\x00\x00\x00")

func templatesPartials_admin_headerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartials_admin_headerHtml,
		"templates/partials_admin_header.html",
	)
}

func templatesPartials_admin_headerHtml() (*asset, error) {
	bytes, err := templatesPartials_admin_headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials_admin_header.html", size: 181, mode: os.FileMode(420), modTime: time.Unix(1504103837, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartials_flashHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x41\xca\xc2\x30\x10\x05\xe0\x7d\x4f\x31\x64\xff\xff\x5e\x40\x7b\x95\x32\x34\x93\x3a\x10\xa6\x32\x53\xeb\x22\xe4\xee\x62\x8d\xa0\xb6\x92\x6e\x12\x48\xde\x7b\x5f\x4a\x9e\x02\x0b\x81\xbb\xa0\x4e\x8c\xd1\xba\x10\xd1\xce\x2e\xe7\x06\x00\x20\x25\x0e\xf0\xbf\x3c\x75\x76\xed\x7b\x32\x2b\x3f\x47\xcf\x33\xf4\x11\xcd\x4e\x0e\x23\xe9\x04\xcb\xf9\x57\x52\xae\x5d\x52\xcf\x8d\x1f\x0b\x07\xcf\x73\x5b\x18\x12\x9f\x73\xb3\x32\x59\xc2\x58\x01\x1f\x91\x4d\xed\xbd\xbb\x83\xba\xa1\x0a\xcb\x50\xd1\x4a\x6a\x13\xfc\x5a\xd8\x61\x92\xea\xa8\x15\xd1\xa3\x0c\xa4\x9b\xe0\x47\x7d\xcd\xbd\xee\x7b\x00\x00\x00\xff\xff\x8b\x68\x0b\x54\xe5\x01\x00\x00")

func templatesPartials_flashHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartials_flashHtml,
		"templates/partials_flash.html",
	)
}

func templatesPartials_flashHtml() (*asset, error) {
	bytes, err := templatesPartials_flashHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials_flash.html", size: 485, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSignupHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4b\x6e\xf3\x30\x0c\x84\xf7\x39\x05\xc1\x03\xc4\x40\x96\x41\xe4\xdd\xbf\xff\x6f\x10\x30\x11\x5d\x13\xb5\x1e\x90\xe8\x3e\x20\xf8\xee\x45\x62\xa9\x89\x81\x16\xe8\x6e\x48\x73\x3e\x6b\xc8\x52\x2c\x0f\xe2\x19\x30\xcb\x8b\x9f\x23\x2e\xcb\x0e\x00\xa0\x14\x65\x17\x27\x52\x06\x8c\x94\x54\x68\x3a\x8f\x4c\x96\x13\xc2\xbe\xce\xac\x8e\xbb\x3c\x8d\x09\xba\x7e\xb7\xea\x21\x24\x07\x74\x55\x09\xde\x60\x57\xb9\xe0\x58\xc7\x60\x0d\xc6\x90\x15\xfb\xfa\x13\x19\x60\x7f\xcd\x69\x38\x6b\x78\x65\xbf\x2c\x27\xf1\x71\x56\xd0\xcf\xc8\x06\x47\xb1\x96\x3d\x82\x27\xc7\x06\x1f\x63\x08\x6f\x34\xcd\x6c\xb0\x94\x8d\x19\xa1\xeb\x4b\x61\x6f\xeb\xfb\xfe\x39\x92\xe9\x08\x1b\xa6\xf2\x87\x36\x22\xdf\xbe\x3f\xc3\xee\x8d\x95\xb3\x26\xb9\xa4\x26\xff\x53\xce\xef\x21\xd9\xe3\x86\x16\x6b\xb7\x11\x1f\xf5\xef\x84\xc3\xdf\x10\x87\x9f\x18\x1b\x67\x9e\x2f\x4e\xbe\xc3\xb4\xaa\xa6\x69\x5b\x6f\xc6\xee\x76\x94\xfe\xf9\x56\x6d\x53\x5f\x01\x00\x00\xff\xff\x97\xb6\x8d\xdb\x02\x02\x00\x00")

func templatesSignupHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSignupHtml,
		"templates/signup.html",
	)
}

func templatesSignupHtml() (*asset, error) {
	bytes, err := templatesSignupHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/signup.html", size: 514, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesWelcomeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xca\xc1\x09\x80\x30\x0c\x05\xd0\x7b\xa6\x08\x19\xc0\x95\x24\xd8\x2f\x16\xda\x5a\x4a\xc0\xc3\x27\xbb\x0b\xe2\xfd\x91\x05\x67\x1d\x50\x7b\xd0\x8e\xbb\xc3\x32\x45\x95\x0c\xf4\xd9\x3c\xa0\x36\x7d\x45\xf5\xb6\x5f\xf0\x82\x65\xba\x7d\xe2\xe7\x22\x24\x46\xc9\x94\x37\x00\x00\xff\xff\x83\x13\x1e\x80\x4a\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/welcome.html", size: 74, mode: os.FileMode(420), modTime: time.Unix(1504103209, 0)}
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
	"templates/admin_index.html": templatesAdmin_indexHtml,
	"templates/dashboard.html": templatesDashboardHtml,
	"templates/login.html": templatesLoginHtml,
	"templates/partial_header.html": templatesPartial_headerHtml,
	"templates/partials_admin_header.html": templatesPartials_admin_headerHtml,
	"templates/partials_flash.html": templatesPartials_flashHtml,
	"templates/signup.html": templatesSignupHtml,
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
		"admin_index.html": &bintree{templatesAdmin_indexHtml, map[string]*bintree{}},
		"dashboard.html": &bintree{templatesDashboardHtml, map[string]*bintree{}},
		"login.html": &bintree{templatesLoginHtml, map[string]*bintree{}},
		"partial_header.html": &bintree{templatesPartial_headerHtml, map[string]*bintree{}},
		"partials_admin_header.html": &bintree{templatesPartials_admin_headerHtml, map[string]*bintree{}},
		"partials_flash.html": &bintree{templatesPartials_flashHtml, map[string]*bintree{}},
		"signup.html": &bintree{templatesSignupHtml, map[string]*bintree{}},
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

