// Code generated by go-bindata.
// sources:
// config
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

var _config = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\x5f\x6f\xe3\xc6\x11\x7f\xe7\xa7\x98\x9e\x10\xd0\x06\x64\x1e\xe5\x8b\x9d\xde\x19\x7a\x38\x5c\x50\x24\x40\xae\x2f\x97\x16\x28\xd2\xc0\x18\x92\x43\x71\xab\xe5\x2e\xb3\x3b\x2b\x59\xf7\xe9\x8b\x99\xa5\x64\xda\xb8\xa2\x79\xb2\xbc\x9c\xfd\xcd\xbf\xdf\xfc\xd9\xd5\x0a\x3e\x79\xd7\x9b\x5d\x0a\xc8\xc6\x3b\xe8\x8d\x25\xe8\x7d\x80\x4f\x1f\xff\xf9\xb1\x82\x1f\xa9\xc7\x64\x19\x0e\x68\x13\x45\xc0\x40\xd0\xfa\x71\x24\xc7\xd4\x81\x4f\x5c\x01\xfc\x23\x12\x90\xe1\x81\x02\x3c\x80\x0f\xb0\xd2\xeb\xb3\x94\x71\xbb\xaa\x28\x8a\xdf\x76\xe4\x28\xa0\xfd\xbd\x28\x56\x10\x47\xef\x79\x30\x6e\x07\xa3\xef\x68\x0d\x2d\x3a\x68\x08\x9c\x0f\x23\xda\x35\xc4\xd6\xc8\xc5\xde\xb4\x82\x76\xc4\x03\xc5\xaa\x78\x50\x59\xd8\xce\x52\x05\x14\x2b\xf8\x5b\xc0\x91\x02\x32\x9d\xad\xfc\x00\xf7\x75\x05\x1f\xdb\x96\x26\x8e\xe0\x9d\x3d\x81\xf3\xee\xc6\xd1\x0e\xd9\x1c\x68\x76\x42\xc0\xfa\xcb\xd5\x2d\x6c\x6a\x10\xab\xbe\x90\x8b\x86\xcd\xc1\xf0\x09\x8c\x03\xf8\xae\x82\x9f\x7b\x38\xf9\x04\x62\xeb\x1e\x78\x20\x68\x30\xe4\x18\xb0\x07\xeb\x8f\x6f\x03\xc5\xc9\xbb\x48\x60\xa2\x1e\x19\x66\x4b\x6b\xe0\x70\x2a\x56\x60\x5c\x1b\x08\xa3\xf8\xc9\x83\x89\x59\x3b\xdc\xd6\x35\x8c\x84\x2e\x42\xe7\x53\x63\x09\x06\x32\xbb\x81\xff\x8c\xd9\x71\x61\xe1\x16\x7e\xb8\x13\xab\x3f\x26\xf6\x72\x0e\x47\x63\x2d\x20\x33\x8d\x13\x8b\x2d\x1d\xa9\x76\x82\xe5\x2d\xd3\x43\x8b\x07\x84\x89\x70\x1f\x2b\xd8\xc0\x16\xbc\x5b\x43\x2d\x7f\xfb\xbe\x58\x81\x3f\x50\x88\x83\xf7\x0c\x68\xad\x3f\xc6\xec\x31\xfb\xc5\x87\x2b\xe3\xe0\x3b\xf0\x3d\x30\x85\xd1\x38\xb4\xb3\x07\xd7\x70\x34\x3c\xf8\xc4\x60\x9c\x61\x83\x92\x7a\xc0\xd9\xbc\xe2\xe1\xf2\x53\x22\x5e\x3c\x2c\x00\xb7\x70\x5b\x8b\x2b\xbf\x0e\x04\x2e\x8d\x0d\x05\x41\x57\xc5\x57\xf5\xcd\x6d\x5d\x5f\x57\x50\xc3\x55\x97\x73\x7c\x0d\x91\x38\x82\x51\x27\x05\x13\xae\x7a\x71\x3d\x4d\xd0\x7a\x17\xbd\xa5\xeb\xaa\x58\xc1\xd1\x74\x3c\x5c\x70\xd0\x75\x10\x27\x6c\x09\x1a\xe2\x23\x91\xcb\xc7\xc6\x2d\x14\xb6\x03\x06\x6c\x99\x42\x7c\x66\x7d\x46\x79\xb7\xb8\xbf\xa9\x0a\xbd\xba\x85\xcd\x3d\x14\x0f\x82\xf3\x98\xa5\xb6\x70\x3b\xff\x2f\x92\xe2\xbc\x38\x2a\x8e\xfd\xe2\x8f\x14\x14\x63\x30\x3b\x29\x93\x36\xb1\xef\x7b\xe8\x03\xfd\x91\xc8\xb5\x86\xa2\x96\x8c\xf5\x47\x8a\xfc\x2c\x18\x59\xcd\x2c\x56\x33\xf7\x5c\x27\x11\xd6\xd0\x0f\x04\x07\x13\x13\x5a\xf3\x95\xc2\x1a\xe6\xd8\x68\xa6\xee\x6a\xb8\x81\x4d\x5d\xd7\x35\xfc\xf4\xb5\x58\xc1\xdf\x3d\xd3\x07\xb9\x11\x94\xa6\x08\xa3\x71\x66\x4c\x23\xb0\x67\xb4\x2f\x60\xbf\x7f\xf7\x79\xf8\x0a\x4f\xaf\xb2\x50\x49\xad\x7d\x12\xda\x64\x8e\x25\xf6\x23\xb2\x69\xd1\xda\xd3\x99\xe4\xa4\x26\x3d\xbb\xa7\xfe\x99\x1e\x10\xd8\x6b\xa9\xa8\x1e\xd1\x1f\x27\x6a\x4d\x6f\xa8\x93\x32\x14\x8f\xc3\x63\x0e\xc7\xa3\x84\x03\xb6\x62\x7f\xf1\x30\x43\xbd\xfa\x94\xbd\x2a\xa4\x9f\x18\x37\x25\xd6\x6e\x32\x12\x0f\xbe\xd3\x0c\xe2\xc4\x29\x64\xda\x75\xc6\xaf\x61\xf2\x31\x1a\xa9\xb1\x2c\xa3\xa5\xfb\x01\xca\x29\xd9\x48\xe5\x1a\x4a\xb4\x11\x4b\x69\x32\x65\x6f\x7a\x5f\x0a\x6f\x7e\x5c\x44\xf2\x7f\x0a\xae\x85\x39\x3c\x20\x83\x0f\x5d\x8e\xff\x44\xae\x23\xc7\xe0\x1d\x1c\xe5\x43\x4c\xd3\xe4\x03\xe7\x72\x3b\x62\x84\x26\x19\x25\x14\x0f\xc5\x4a\xea\xd6\x5a\x50\x1f\x2e\xb6\xa5\x48\x51\xa3\x18\x71\x94\x26\x2b\x4d\x19\x0e\x18\x0c\x8a\x07\x65\xf4\x29\xb4\x54\x4a\x2e\xb4\xba\x7b\xe3\x08\x8e\x39\xab\x0c\x71\xf0\xc9\x76\xb0\x23\x56\x08\xf5\xbf\x52\x45\x42\x2c\xf5\x43\xcf\x9e\x71\x34\x97\x4d\xce\x5b\x3e\xab\x9e\xdb\x68\x29\x49\x2e\xd7\x70\x1c\x4c\x3b\x3c\x5b\x36\x7a\x67\xd8\x87\x59\xfe\x4c\xc4\x99\x7c\x10\x8d\xdb\x17\x2b\xb8\x42\x6b\x97\x1a\xe5\x38\x5e\xf9\xc4\x53\xe2\x78\x0d\x03\x1e\x08\xca\x19\xa9\x9c\xa1\xe2\x95\xc6\x22\x5e\x03\xc6\xe8\x5b\x83\x32\x5d\x94\x95\x3c\xd0\x78\x5d\xac\x60\xf6\x44\x12\xf1\x6d\x1f\x72\xf2\xc5\x9a\x83\x79\xe9\xcb\x70\xfc\xf0\x8b\xf7\x53\x83\xed\x7e\xbd\x29\x67\x20\xc9\xe3\xb7\x81\x26\x14\xb5\x5e\x25\x6e\x64\x1a\x2e\xb1\xde\xf2\x38\xbd\x1d\xa7\xae\x52\x1e\x08\x51\x67\xfa\x6d\xb3\xcb\xc5\xc3\x39\x38\x5b\x2d\x94\x62\x29\xa1\xc6\xc3\x52\xe4\x85\x69\xf0\x42\x58\x14\x2c\x45\x5f\x68\xd6\x12\xc8\x21\xd5\x1a\xf0\xe9\x99\x49\x30\xe2\x49\xc7\x69\x9b\x42\xa4\xb8\x96\x59\x92\x7f\x0a\x81\x03\x1e\xab\xc5\x89\xc9\x9d\x27\x9e\x22\xd3\x18\x33\xa7\x3b\x4f\x11\x9c\x57\x0e\x0b\x85\x67\x59\x6d\x03\x01\x8f\xb9\x89\x6c\xee\xa1\x31\x0c\x1d\x32\x42\xe4\x40\x38\x9e\xe9\xd0\x60\x98\x27\xc2\x8c\x37\x8f\xf7\x14\xa9\x93\xb8\x46\x72\xfa\xd7\xeb\xc6\x80\xd3\x64\x4d\xab\xab\x47\x9c\xf1\x97\xcd\x4c\xa6\xa5\xf4\xa0\xcb\x92\x80\xdd\x7f\x52\x14\x76\x68\x05\xce\x13\xd9\x4f\xba\xba\x60\xe3\x0f\x54\x5c\x22\x18\xf0\x28\xb1\xc9\x7d\x12\x22\x9f\xac\x44\x63\x8e\x4e\x19\x99\x02\xf9\x5c\xd4\xa3\x77\xbe\x54\xff\xbe\xe8\x29\x8c\x26\x04\x1f\x22\x34\x9e\x07\x19\x0c\xce\x91\x8d\x99\x90\xd2\xcb\x96\x8d\xdb\x38\x68\x65\x1b\x0a\xd2\x3d\x3e\x7b\xe7\x01\x0f\x14\x70\x47\xaf\x6f\x4b\xff\x9b\xab\x00\x2c\xf5\x3a\xbf\x82\x84\xe9\xdc\xf7\xd9\x5f\xda\xfe\x42\x41\xa1\x86\xc3\x56\x8a\xcf\xeb\xa2\x22\x31\xca\x40\xc0\x18\xa4\xe4\x97\x11\x8b\xdc\xf9\xc4\x6b\xc0\x4c\xf1\x33\xb3\xa5\x4b\x6b\xd8\xfa\xcb\xa5\x73\x9a\xe9\xc9\x44\x2e\x1e\x04\xf7\x71\xfe\xb6\x85\xb7\x1d\x1d\xde\x66\xac\xb3\x4e\xcd\x75\x2f\xeb\x17\x5f\xd2\xd1\x18\x87\xe1\x24\x41\xc4\xd8\x1a\x53\x88\xcc\x63\x96\x11\xc2\xeb\x59\xb1\x3a\x8b\x09\x63\x5e\x01\x94\x7f\x6d\x0c\x97\x3a\xe8\xef\xee\xae\x35\x1b\x9b\xfb\xf3\xd1\xfd\xdd\xdd\xbb\xfa\x1a\x8a\xc6\xf0\x33\xa8\x7e\x16\x50\x45\x87\x11\x9f\xf2\x82\xa4\x4d\x79\x3e\x93\x55\x31\xa0\xdb\x51\xf6\x3f\x24\x07\x7d\xf0\x23\xd4\x12\xa1\xbc\x87\x5d\xe6\x10\x48\x0f\x95\xfd\x44\xae\x3e\x8e\xf8\xf4\x98\x6f\xca\x24\xaa\xeb\x67\x45\x1d\x59\x33\x1a\xd9\x0f\x16\x9a\x66\xab\x08\xdb\x41\x89\x2f\x49\xd6\xdd\x52\x47\x1d\x4d\x18\x34\xea\xcd\x09\x70\x01\x20\x54\x49\x51\x7a\x55\x6b\x46\xb4\xb3\x45\x17\x50\xd6\x8e\x7f\x65\x2a\xaa\xe0\xee\x3d\x6c\xa1\x7c\x28\x15\x79\x23\x8b\x5a\xf9\x6f\x57\xc2\x95\x95\xf6\xdf\x13\x75\xd7\x12\x1f\x0c\x8f\x17\x74\x31\xfc\x7d\xa1\x46\xbc\x38\xdc\xd4\x85\xb6\x8d\xd6\x5b\x1f\xf2\x1e\x9e\x67\x14\x75\xa0\x67\xf3\x78\x0c\xd4\xad\x61\x17\x88\xdc\x1a\x4e\x24\x5b\xa0\x14\xcd\x8e\x1c\xe3\x1a\xda\x13\x3a\x1d\x0b\x4c\x6b\x68\x34\xea\x8d\xc5\x76\x5f\xe9\xea\xd3\xee\x77\xc1\x27\x27\xc5\xa7\xa7\xb2\x68\xfb\x40\x97\x43\xb9\x2d\x26\x5c\x1e\x00\x79\x80\x27\xcb\x66\xb2\x86\x82\xf6\x22\xa9\x6a\xe3\x98\x76\x41\xaa\xf6\xf2\x54\x68\xd1\xb6\xc9\xe6\x46\x51\xc1\xaf\xb8\xa7\x78\x7e\x94\xcc\xa9\xbd\x81\xba\x7a\xff\x5e\x62\xfb\x53\xde\x43\xe6\xcf\x79\xdb\xce\x48\x32\xac\x9b\xc4\x60\x29\x46\x98\x02\xb5\x26\x52\x95\x69\xd1\x99\x28\x71\x17\x4f\x2e\xea\xb7\x50\x57\x3f\x68\xd1\xcd\x5f\xb5\x87\x92\xcb\x3f\xf3\xe8\xbc\x91\x45\x88\x3a\x78\xf3\xd9\x3b\x69\x2a\xad\x4c\xff\xb3\xd9\x6f\x16\xc3\x63\x53\xc1\x17\xd2\x32\x7f\xad\x6f\x7c\xbe\xb9\x85\xcd\xfc\x1a\x61\xd8\x05\xd4\xad\xfd\x55\x80\xde\x74\xc1\x4f\xb2\xaf\xbf\xa9\xbe\xe9\xa9\xb6\x44\xe5\xbd\x0a\xf6\x28\xd8\xda\xdd\xfe\xef\x13\x03\xea\xea\x6e\x46\x19\xd0\xf6\x67\x0b\xd6\x70\xfb\xe2\xcd\xf2\x4d\x47\x96\x76\x15\x0f\x17\xe3\xc5\x21\xd0\x35\xf8\x67\xb7\x98\x0d\xeb\xf9\x65\x21\x03\xe2\xa8\xbb\x8b\x2e\x06\x8d\xec\xe6\xba\x15\xe6\x4f\xfa\x70\x52\x5f\xa4\x53\x35\x04\x5d\xc0\xa3\xd3\x1c\xed\x9c\x0f\x52\xa5\xba\x10\xfe\x46\x7f\xfc\x9e\xdf\x10\x46\xdc\xcb\x0f\xb2\x60\xda\xfd\xa9\x82\x7f\xf9\xa4\xdd\x46\x15\x60\x84\x31\xb5\x03\xec\xe9\x14\xe5\x1f\x79\xdf\x1d\xd1\x71\x9e\x6e\x34\x92\x6e\xbe\xec\x21\xb9\xf9\x05\x0b\xa3\xe8\xe1\x81\x9c\x02\xef\xe9\xf4\x17\xf8\x2c\x47\x0a\xb1\xcd\x9f\x33\x95\x8c\x77\x95\xee\xfb\x7e\x0f\xc8\x10\x08\xbb\x91\xaa\xb1\x93\xad\x70\x67\x78\x48\x4d\xde\x3d\x52\xd0\xd9\x47\x4f\x93\x45\x97\x29\xad\xf5\x4d\x4f\x38\x4e\x36\xbf\x50\x37\x1a\xb9\x15\x34\x18\xe5\xf5\x74\x3b\x3f\x9b\xde\xcd\xc7\xa3\xe9\xd8\x3b\xe9\x5b\xdf\xcf\x5f\xee\xe6\x2f\x1c\xa8\xb1\x54\xfc\x37\x00\x00\xff\xff\xdb\x75\xb4\xcf\xd9\x0f\x00\x00")

func configBytes() ([]byte, error) {
	return bindataRead(
		_config,
		"config",
	)
}

func config() (*asset, error) {
	bytes, err := configBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config", size: 4057, mode: os.FileMode(420), modTime: time.Unix(1504549591, 0)}
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
	"config": config,
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
	"config": &bintree{config, map[string]*bintree{}},
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
