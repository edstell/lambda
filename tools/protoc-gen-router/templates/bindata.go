// Code generated by go-bindata.
// sources:
// router.gotmpl
// DO NOT EDIT!

package templates

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

var _routerGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x72\xd3\x30\x10\x86\xcf\xd2\x53\x2c\x3e\xb4\x76\xc6\x63\xdf\xc3\xf4\x02\x65\x86\x1e\x9a\x30\x25\x3c\x80\x62\x6f\x12\x83\x2d\x29\x2b\x99\x36\xa3\xd1\xbb\x33\x92\x6c\x97\x4e\x38\xc0\x70\x8a\xb2\xde\xfd\xfe\xff\x5f\xa9\xae\xe1\x7e\x0b\x9b\xed\x0e\x3e\xdd\x3f\xec\xd6\xf0\xa4\x46\x8b\x04\xcf\xc2\x80\x18\xad\x3a\xa2\x44\x12\x16\x5b\x38\x90\x1a\xe0\xd6\xb9\x6a\x23\x06\xf4\xfe\x96\x6b\xd1\xfc\x10\x47\x04\xe7\xaa\x2f\xe9\xe8\x3d\xe7\xdd\xa0\x15\x59\xc8\x39\xcb\x1a\x25\x2d\xbe\xd8\x8c\xb3\x0c\x65\xa3\xda\x4e\x1e\xeb\xef\x46\xc9\x8c\x73\x96\x1d\x3b\x7b\x1a\xf7\x55\xa3\x86\x1a\x5b\x63\xb1\xef\xeb\x5e\x0c\xfb\x56\xd4\x7d\xb7\x27\x41\x1d\x9a\x9a\x74\x93\xf1\x82\x3b\x47\x42\x1e\x11\xaa\xaf\x48\x3f\xbb\x26\xc8\xd8\x8b\x46\xf8\x2c\x64\xdb\x23\x41\x27\x2d\xd2\x41\x34\xc1\xca\xdc\xfa\x88\xf6\xa4\x5a\xef\x39\x5b\x1c\xe7\x93\x9f\xea\x63\xfa\x2d\x83\xf3\x07\xa9\x47\xbb\xbb\x68\xf4\xbe\x80\x7c\xe5\x5c\xb5\x1d\xed\x52\x2a\x01\x89\x14\x15\xce\xa1\x0c\x30\xcf\x93\xf2\xb4\x24\x63\x69\x6c\x2c\x38\xce\x56\xa4\x9b\x2a\x55\x43\xd3\x61\x94\x0d\x6c\xf0\x39\x55\xf2\xd3\x64\x74\x32\x5c\xc0\x6a\x02\x38\xce\x28\x9d\xd6\x77\x10\x10\xaf\x33\xc5\x1f\xa2\xa4\xde\xa4\x93\x67\x4b\xb0\x2c\x26\x99\x42\x4e\x5a\xd5\x52\x29\x16\xf7\x8c\xd0\x8e\x24\xe1\x26\x49\x38\xce\x58\x3a\xad\x21\x91\x4b\xce\x42\xc6\x6b\xe5\x98\xe7\x4a\x03\x42\xf9\x3f\xb6\x5a\xc4\xc8\xf3\x2d\xba\xc5\x5f\xc2\xda\x17\xb8\x42\x13\x9e\xe3\xcc\x13\x9e\x47\x34\x36\xb0\xd3\x5f\xa3\x95\x34\x38\x83\x03\x8b\xed\x55\x7b\x09\x6b\xbd\x79\xeb\xc7\x79\xce\x58\x77\x08\x9d\xe1\x6b\x78\x8f\xd5\x37\x39\x08\x32\x27\xd1\xe7\x84\xe7\xea\x83\x6a\x2f\x25\x84\xf1\xe2\x7d\x6c\x7b\x77\x07\xb2\xeb\x23\x74\xb6\x28\xbb\x3e\x8a\x71\xc6\x3c\xe7\x8c\x91\xd1\xe5\x8c\x9c\xb6\x13\x12\x94\xb0\x8a\x9c\x57\xc9\xbf\x60\xed\x2f\x16\x4d\xf9\xc6\xe0\xe3\x6c\xcf\xe8\x7f\x63\xcd\x37\xfe\xfb\x96\x62\x7f\x08\xb9\x86\x24\x15\x9a\xcb\x30\x1a\xaf\x7f\x7a\x2d\xf3\xa3\xf9\x15\x00\x00\xff\xff\x45\xc6\xb5\x55\x20\x04\x00\x00")

func routerGotmplBytes() ([]byte, error) {
	return bindataRead(
		_routerGotmpl,
		"router.gotmpl",
	)
}

func routerGotmpl() (*asset, error) {
	bytes, err := routerGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router.gotmpl", size: 1056, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"router.gotmpl": routerGotmpl,
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
	"router.gotmpl": &bintree{routerGotmpl, map[string]*bintree{}},
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