// Code generated by go-bindata.
// sources:
// res/Readme.md
// res/actions/Readme.md
// res/resources_version.txt
// res/standard/actions/goimports.sh
// res/standard/config.toml
// res/standard/themes/default.toml
// res/themes/Readme.md
// DO NOT EDIT!

package core

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _resReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xcd\xae\x1a\x31\x0c\x85\xf7\xf3\x14\xa7\xbb\x7b\x25\x3a\xec\xef\xae\x2a\x77\x81\x54\x41\x25\x60\x8f\x49\x9c\x21\x6a\x26\x19\x25\x99\x42\xbb\xe8\xb3\x37\x09\x19\xfe\xfa\xb3\x8c\xed\xef\xd8\x3e\xce\x67\x67\x23\xdb\x18\xe0\x14\x7e\xcd\xdb\xce\xb1\xc4\x5b\x03\x7c\x84\x70\x56\xe9\xae\x8d\xae\x37\x78\xc3\x2e\xb0\xc7\x8b\x18\x43\x7a\xeb\x9f\x2c\x5f\x6b\x7e\xf4\x14\xb5\xb3\x50\xda\x70\xc1\xe2\x91\x7b\x0e\xf3\xbf\x22\x97\xdc\x0c\xc2\x33\x45\x86\xf2\xae\x47\x10\x49\x41\x1c\xe1\x7c\x52\x1c\x74\x6a\x7f\x09\x47\xb2\x92\xbc\x9c\x57\xbd\xa2\x4d\x22\xf7\xfa\x87\x78\x4d\x36\xa5\xf2\x8a\xa7\xd2\xb5\xd7\x9d\xb6\x64\x50\x96\xcb\x83\xa6\x11\xa4\x83\x75\x11\x2c\x75\x84\xd4\x9e\x45\x34\x3f\x40\x01\x27\x6d\x0c\x0e\x0c\xcf\x83\x21\x91\xea\xc7\x21\x6d\x37\x0e\x9d\x27\xc9\xa1\x7d\x14\x7f\xb4\x68\x53\xc3\x93\x73\x58\xdc\x7a\x3c\x81\x37\x93\xa6\x50\xf5\xe6\x7f\xd0\xdd\xf6\x57\xaa\xc6\x9e\xb0\xc2\x1d\x46\xa5\xd8\xa7\x72\x71\x39\x31\x3c\x9d\xa6\x20\x24\x45\x9a\x61\xb1\xc6\x6a\xbd\xc5\xfb\x62\xb9\xc5\x87\x02\x69\x9b\xa5\x05\x57\x8c\xd2\x3b\x4f\x06\x31\x7a\x9f\x35\x8a\x85\x53\xd1\x4b\x78\xc5\xa7\xaf\x4b\x04\x27\xbe\x71\xcc\xe6\x34\x57\x0f\x8a\xcd\x20\xcf\x18\x43\x42\x76\xab\x2f\xef\x9b\x0d\x08\xb7\x8b\xe1\x7b\x1a\x24\x7f\x1d\x3e\xeb\x10\xc3\xac\x09\x0e\x2a\xfd\x02\x3e\x53\x3f\x18\x86\x56\x38\xb8\x78\xc4\x7e\xda\xbb\x73\xaa\x8f\xc5\xec\x3d\x52\x17\xec\xff\x70\xe6\xbe\xa2\xaa\x4e\xf8\x43\x6e\x3a\x72\x1e\xad\x6d\x7e\x07\x00\x00\xff\xff\x52\x41\xf4\x3d\x03\x03\x00\x00")

func resReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resReadmeMd,
		"res/Readme.md",
	)
}

func resReadmeMd() (*asset, error) {
	bytes, err := resReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/Readme.md", size: 771, mode: os.FileMode(420), modTime: time.Unix(1435791068, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resActionsReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x2e\x2d\x2e\xc9\xcf\x55\x48\x4c\x2e\xc9\xcc\xcf\x2b\x56\x48\xcf\x57\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\xc4\x35\xc4\xb2\x17\x00\x00\x00")

func resActionsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resActionsReadmeMd,
		"res/actions/Readme.md",
	)
}

func resActionsReadmeMd() (*asset, error) {
	bytes, err := resActionsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/actions/Readme.md", size: 23, mode: os.FileMode(420), modTime: time.Unix(1435782197, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resResources_versionTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\x34\x31\x36\x33\xb1\x34\x34\xb4\x30\xe3\x02\x04\x00\x00\xff\xff\x0a\xb4\x5b\x5f\x0b\x00\x00\x00")

func resResources_versionTxtBytes() ([]byte, error) {
	return bindataRead(
		_resResources_versionTxt,
		"res/resources_version.txt",
	)
}

func resResources_versionTxt() (*asset, error) {
	bytes, err := resResources_versionTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/resources_version.txt", size: 11, mode: os.FileMode(420), modTime: time.Unix(1436491186, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resStandardActionsGoimportsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x41\x4e\x86\x30\x10\x85\xf7\x73\x8a\x31\x74\xfb\x87\x13\xb0\x30\x4a\x0c\x1b\x5c\x68\x74\x59\x0a\x4c\x71\x92\xd2\x21\x6d\x15\xbd\xbd\x2d\xc1\xc8\x72\xde\xbc\xef\x7b\xd5\x5d\x3d\xb2\xaf\x47\x13\x3f\x20\x52\xc2\x1b\x7d\x63\x65\x0d\x3b\x24\x13\xdc\x0f\xc0\x17\xd3\xde\xa8\xa7\xe7\xf6\x51\xbf\x75\xed\x3b\xb0\x8f\xe9\xbc\xbb\xfe\xe5\xf5\xbe\x7f\x68\x01\x2a\x5c\x32\x6b\xd9\x11\x3a\x99\x4c\x62\xf1\x10\xc3\xd4\x0c\x8b\xd0\xac\xcd\xc6\x58\x34\x3a\x47\x3a\xff\x51\x15\x09\xaa\x92\x0d\x58\x60\xe1\x75\x93\x90\x22\x8a\x3f\x2c\xf0\x9f\xdc\x76\x54\x99\xcb\xad\x40\x4e\xcc\x7c\x98\x70\xfc\xb4\x96\x02\xda\x20\xeb\x1f\x70\x1d\x3a\xab\x97\x1d\x84\xdf\x00\x00\x00\xff\xff\x4e\x39\x38\xf4\xea\x00\x00\x00")

func resStandardActionsGoimportsShBytes() ([]byte, error) {
	return bindataRead(
		_resStandardActionsGoimportsSh,
		"res/standard/actions/goimports.sh",
	)
}

func resStandardActionsGoimportsSh() (*asset, error) {
	bytes, err := resStandardActionsGoimportsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/standard/actions/goimports.sh", size: 234, mode: os.FileMode(493), modTime: time.Unix(1436489398, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resStandardConfigToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\x70\x49\x4d\x4b\x2c\xcd\x29\x51\x48\xce\xcf\x4b\xcb\x4c\xe7\x0a\xae\xcc\x2b\x49\xac\xf0\xc8\x4c\xcf\xc8\x01\xe2\x92\xcc\xbc\x74\xdb\x92\xa2\xd2\x54\xae\x90\x8c\xd4\xdc\x54\x5b\xa5\x14\x88\x6a\xbd\x92\xfc\xdc\x1c\x25\x2e\xdf\xc4\x0a\xe7\xdc\x14\xa7\xd2\xb4\xb4\xd4\x22\x9f\xcc\xbc\xd4\x62\x5b\x43\x03\x20\x00\x04\x00\x00\xff\xff\x79\xef\x87\x0b\x55\x00\x00\x00")

func resStandardConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_resStandardConfigToml,
		"res/standard/config.toml",
	)
}

func resStandardConfigToml() (*asset, error) {
	bytes, err := resStandardConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/standard/config.toml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1435780539, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resStandardThemesDefaultToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xbb\x8e\xda\x40\x18\x85\x7b\x9e\x02\xb9\x9e\x62\xc6\x36\x06\x17\x2e\xf0\x5c\x9a\x28\x4a\x61\x92\xde\xc0\x88\x58\xb2\x19\x34\x18\x11\xa4\x14\xe9\x90\xd2\xe5\xde\x90\x22\x8a\xb4\xda\xe7\xe2\x49\xf6\xf7\x5c\xcc\x2e\x8b\x70\x05\xe7\xff\xce\x99\x33\xbf\x26\x5f\x65\x01\x9f\x62\xf3\x05\x03\x01\xff\x98\xc0\x11\x16\x98\x04\x83\x7c\x55\xc8\x5a\x2e\xda\x5e\x33\x84\xd7\x9c\xcb\x70\x74\xa7\xb7\x4a\x67\x41\x48\x70\xe2\x39\xaf\x5d\x38\xaa\x9a\x46\xae\xc1\x2a\xc6\x78\x6c\xb1\xa2\xd5\xd5\x1a\x0e\x8d\x53\x6f\x7c\x23\x0f\x7b\xa5\x97\xc4\x84\xb9\x22\x4e\x0b\xb3\x20\x11\x57\x5a\x04\x5a\xea\xb5\x42\x6e\x4a\x5d\xb6\x4a\x83\x3b\xc6\x78\xe4\xce\xf0\x2a\xf8\xd3\xd1\x2b\x15\x12\x78\xd8\xab\x87\x66\xae\x6a\xb0\x53\x8a\x63\x17\x6a\x24\xf0\x4e\xd8\x4b\x09\x8c\x8c\xf4\x52\x5b\xb6\xbb\xed\xbc\xd4\xc3\x6c\x18\x9c\xff\x7e\x45\x3c\x87\x3b\xc2\xd7\xfd\x70\xdb\xed\x99\x99\xfc\xd4\x76\x5c\xce\x92\xe4\xc6\x88\x6b\x93\x82\x09\x1d\xb9\xc5\x35\xcb\xfb\xc9\x16\xb8\xc4\xc2\x42\xae\xf4\x77\x6b\x33\x11\xe6\x40\x88\xfc\x50\xc9\xbd\xcf\xfc\xf2\x03\x71\xe6\x33\x99\xcf\x74\x84\x0f\xe5\x63\x1c\xba\x47\x52\xd5\x92\xd6\xb2\x34\x89\xe7\xd3\x4f\x44\xa8\x1d\x21\x9e\xf6\x0f\x09\x18\x56\xe9\xf6\x60\x99\x3f\x08\x1e\x01\xe9\x4e\x7e\xc6\x14\x0b\xad\xea\xda\x97\xf8\xfd\x0d\x41\x37\xb3\x0e\x64\xb7\x72\x61\x66\xe5\xdc\x32\xdf\x51\x12\x7b\x26\xf6\xcc\x5b\xa5\x65\xd7\xb2\xa8\x96\xd2\xde\xe7\x01\x11\x81\x63\x5b\x69\x7a\x8d\xbd\xdf\x18\xe8\xf8\xef\x1e\xc4\xd4\xde\x5e\xef\xf8\xff\x16\x06\x85\xe8\x47\x57\xfc\xf8\x88\xb8\xb0\x83\x8e\x98\xb8\xa0\x52\xaf\x2a\x13\xf1\xf9\xd6\x98\xd6\x6a\x6b\xcb\x9e\x7e\xa1\x28\x72\x77\xe2\xd4\xcd\x9f\x02\x00\x00\xff\xff\x63\x53\xd3\x5b\x96\x03\x00\x00")

func resStandardThemesDefaultTomlBytes() ([]byte, error) {
	return bindataRead(
		_resStandardThemesDefaultToml,
		"res/standard/themes/default.toml",
	)
}

func resStandardThemesDefaultToml() (*asset, error) {
	bytes, err := resStandardThemesDefaultTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/standard/themes/default.toml", size: 918, mode: os.FileMode(420), modTime: time.Unix(1435780812, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resThemesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x2e\x2d\x2e\xc9\xcf\x55\x28\xc9\x48\xcd\x4d\x2d\x56\x48\xcf\x57\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\x26\x56\x89\xea\x16\x00\x00\x00")

func resThemesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resThemesReadmeMd,
		"res/themes/Readme.md",
	)
}

func resThemesReadmeMd() (*asset, error) {
	bytes, err := resThemesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/themes/Readme.md", size: 22, mode: os.FileMode(420), modTime: time.Unix(1435782185, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"res/Readme.md": resReadmeMd,
	"res/actions/Readme.md": resActionsReadmeMd,
	"res/resources_version.txt": resResources_versionTxt,
	"res/standard/actions/goimports.sh": resStandardActionsGoimportsSh,
	"res/standard/config.toml": resStandardConfigToml,
	"res/standard/themes/default.toml": resStandardThemesDefaultToml,
	"res/themes/Readme.md": resThemesReadmeMd,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"res": &bintree{nil, map[string]*bintree{
		"Readme.md": &bintree{resReadmeMd, map[string]*bintree{
		}},
		"actions": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resActionsReadmeMd, map[string]*bintree{
			}},
		}},
		"resources_version.txt": &bintree{resResources_versionTxt, map[string]*bintree{
		}},
		"standard": &bintree{nil, map[string]*bintree{
			"actions": &bintree{nil, map[string]*bintree{
				"goimports.sh": &bintree{resStandardActionsGoimportsSh, map[string]*bintree{
				}},
			}},
			"config.toml": &bintree{resStandardConfigToml, map[string]*bintree{
			}},
			"themes": &bintree{nil, map[string]*bintree{
				"default.toml": &bintree{resStandardThemesDefaultToml, map[string]*bintree{
				}},
			}},
		}},
		"themes": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resThemesReadmeMd, map[string]*bintree{
			}},
		}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

