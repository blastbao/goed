// Code generated by go-bindata.
// sources:
// res/.DS_Store
// res/Readme.md
// res/actions/Readme.md
// res/default/actions/find_files.sh
// res/default/actions/goed.rc
// res/default/actions/goed.sh
// res/default/actions/goed_helper.ank
// res/default/actions/goimports.sh
// res/default/actions/search.ank
// res/default/actions/search_text.sh
// res/default/actions/stats.sh
// res/default/actions/vt100_size.sh
// res/default/bindings.toml
// res/default/config.toml
// res/default/themes/acme.toml
// res/default/themes/default.toml
// res/resources_version.txt
// res/themes/Readme.md
// DO NOT EDIT!

package core

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

var _resDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x97\xcf\x6e\xd3\x40\x10\xc6\x67\xd3\xb4\x71\xd2\xcb\x22\xa8\x54\x89\x8b\x8f\x20\x45\x55\x50\x9b\x56\xdc\x42\x28\x48\x1c\x90\xa2\x06\xb5\x42\x14\x05\x3b\x36\x64\x25\xe3\x8d\xec\x75\x02\x8a\x22\xe5\x86\xc4\x13\x70\xe6\xff\x7b\x20\x38\xf5\x01\xe0\x51\x38\xc1\xac\x3d\x2d\xc6\x09\x07\x4e\x8d\x60\x3f\x6b\xf3\x9b\x24\xdf\x8c\x13\x8f\xbd\xf6\x02\x00\x6b\x27\xde\x35\x00\x8e\xa1\x05\x19\xcb\x35\x58\x28\x8b\xc6\x9c\x4a\x39\x03\x4b\x6b\x48\x08\x21\xbe\x13\xc8\xbe\x1b\x2c\xae\x65\x64\x64\xb4\x44\xd2\xd7\x6e\x05\x1c\xe8\x83\x02\x91\xbf\x7e\xa5\x0b\xd9\xb4\x70\x1b\xc7\x95\x1f\xa9\xd2\x8c\x0a\x78\xe0\xc3\x63\xcc\x49\x20\x00\x55\x70\x9f\x14\xdc\xab\xe8\xd4\x55\x55\xb1\x2e\x3b\xd0\x4e\x1c\x6c\xde\xe9\x8e\xe3\x21\x39\xbf\xbb\xc3\x40\xc4\xaa\xd1\xf8\xc6\x4a\x2b\xe5\xd5\xb5\x8a\x55\xad\xad\xeb\x8d\x5f\xa8\x3d\xec\x0e\xe4\xb8\xab\x1c\x95\xc4\x6d\x27\x7a\x90\xbe\x13\x9e\xef\x52\xdc\x71\xd4\xe0\x34\xbe\x27\x65\x70\x16\x3b\xee\xa1\xf0\xc7\x3d\x7e\xe9\xa6\x0c\x95\x23\x42\x3f\xca\xa5\x1e\x1f\x89\xd0\x93\xe3\xb6\x4c\x42\x2f\x3e\xa6\x0f\x8f\x84\xa7\x06\x3d\xbe\xd1\x89\xfc\x11\xa6\x76\x9c\xd0\x3f\x14\xb1\x70\x45\x20\xd4\x73\xab\xaa\xb7\x1e\xdf\x9c\x4c\x9a\xbb\xcd\xba\xbd\x7d\x7d\x67\x5a\xb7\x27\x7b\x7b\x8d\xba\xbd\xb3\xbd\x3b\x9d\xf2\x17\x96\x75\xf9\x6a\xf3\xc6\xdd\xfb\xc3\xe9\xcb\xd7\x6f\xde\xbe\x7b\xff\xe1\xe3\xa7\xcf\xd9\xf1\x67\x8c\x1a\x71\xb1\xd0\x98\x2f\xbf\x1f\x92\x51\x37\x0a\x03\x19\x3e\x81\x74\xaa\x85\x2a\x1c\x60\x13\x1c\x6c\xc5\x53\xe4\x16\xbe\x7a\x85\xc3\xfb\xb5\xd0\x88\x0d\x88\xd0\x19\x63\xbd\x04\xa3\x7e\x1a\xf7\x60\x84\x8c\x30\xca\x5a\xbf\x85\x7b\x7a\x36\xd7\x52\x7d\x02\xbc\xca\x55\x5a\x43\xd7\x00\xf3\xf4\x9e\x8b\x27\xcb\x49\xc1\x6b\xb4\xc4\xa2\x53\xcf\x5a\x3f\xdf\x9f\x61\x64\x64\xb4\x84\xd2\xf3\x83\x4d\x6c\x11\x67\x19\x19\x7d\x5f\x22\x96\x73\x39\x9c\x68\x13\x5b\xc4\x59\x46\x46\xbe\x12\xb1\x4c\xb4\x88\x9c\x68\x13\x5b\xc4\x59\x46\x9a\xb4\x18\x2d\x3e\x18\xed\x99\xd1\x0a\x85\x71\xa2\x4d\x6c\xfd\xe5\x9f\x36\x32\xfa\x4f\xb4\x92\x81\xeb\xfb\xff\xad\x3f\xaf\xff\x8d\x8c\x8c\xfe\x61\xb1\xf2\x7e\x77\xbf\x0d\x67\x0b\x82\x39\xe9\x7b\xad\x8d\xe3\xd1\x69\x02\x2c\x7e\x10\x20\xaf\xbe\x15\x6f\xc2\x2f\xaf\x4d\x6c\x11\x67\x19\xcd\x83\x80\x91\xd1\x79\xe9\x67\x00\x00\x00\xff\xff\xd1\x0c\x4c\xbd\x04\x18\x00\x00")

func resDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_resDs_store,
		"res/.DS_Store",
	)
}

func resDs_store() (*asset, error) {
	bytes, err := resDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1469570529, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\xb1\x6e\x1b\x31\x0c\xdd\xef\x2b\x5e\xb7\x04\x70\xed\x3d\x5b\x51\x67\x08\x50\xc4\x05\x9c\xec\xa6\x25\xea\x2c\x54\x77\x3a\x48\x54\x93\x76\xe8\xb7\x57\x92\x75\x71\xce\x68\x33\x8a\x7c\xef\xf1\xf1\x51\x5f\xfd\x28\x3c\x4a\x84\x37\xf8\xb3\x59\xf7\x9e\x35\xee\x3a\xe0\x33\x94\x1f\x8d\xed\xd7\xe2\x07\x87\x3b\x3c\x47\x0e\xb8\x51\x29\xe6\xb7\xfd\xcd\xfa\xb6\xf5\x53\x20\xb1\x7e\x84\xb1\x8e\x2b\x4d\x4e\x3c\x70\xdc\xfc\x93\x72\xee\xad\xa0\x02\x93\x30\x4c\xf0\x03\xa2\xca\x0a\xea\x04\x1f\xb2\xe2\x64\xf3\xf8\x73\x59\x68\xd4\x14\xf4\xa6\xe9\x55\x6d\x52\x65\xd6\x7f\xc4\x5b\xb3\xab\x48\xcd\x86\x92\x93\x82\xdc\x05\xdb\xdb\x91\x1c\xea\x6e\xc5\x67\x76\xa0\x3d\x46\x2f\x60\x6d\x05\xda\x06\x56\xe2\x7e\x81\x22\x5e\xac\x73\x38\x32\x02\x4f\x8e\x54\xc6\xa7\x29\x2f\x97\xa6\x3e\x90\xe6\xb8\x5e\x68\x2f\x03\xda\x37\xc3\x73\x6e\xd8\x5e\x46\x2c\x79\x97\x84\xe6\x25\x5b\x30\x1f\x70\xde\x6d\xfe\x46\x6a\xb5\x2b\x56\xa5\x1d\x93\x31\x1c\x32\x5c\x9d\xcf\x8b\x40\x2f\x73\x11\x9a\x84\x56\xd8\xee\xf0\xb8\x7b\xc2\xfd\xf6\xe1\x09\x9f\x2a\xc9\x8e\x45\x5a\x71\xa3\x51\x7e\x17\x63\x50\x29\x84\xa2\x51\xf3\x9b\x41\x37\xf1\x16\x5f\xbe\x3f\x20\x7a\xf5\x83\xa5\x24\xd3\xbd\x25\x50\x33\x06\x05\x46\x8a\x99\xf2\xfc\xf8\xed\x7e\xbf\x07\xe1\x72\x2d\xfc\xcc\x46\xca\xb7\xe1\x57\x1b\x25\xae\xba\xe8\x61\xf2\x0f\xe0\x57\x1a\x26\xc7\xb0\x06\x47\x2f\x27\x1c\xe6\xbd\x7b\x6f\x06\xa9\x51\x1f\x90\xa7\xe0\x70\x1d\xcc\x7b\x40\x13\x9d\xd9\x8b\xde\x7c\xe0\xe2\x6c\xdd\xfd\x0d\x00\x00\xff\xff\x4d\x01\x5a\xb0\xfe\x02\x00\x00")

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

	info := bindataFileInfo{name: "res/Readme.md", size: 766, mode: os.FileMode(420), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
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

	info := bindataFileInfo{name: "res/actions/Readme.md", size: 23, mode: os.FileMode(420), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsFind_filesSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xbd\x8e\x83\x30\x10\x84\xeb\xdb\xa7\x98\x5b\x5c\xdc\x21\x71\x88\x2b\xf3\xf7\x12\x94\x08\x21\x9b\xd8\xd8\x12\x71\x22\xec\x48\xf0\xf6\xc1\xd0\x24\x5b\x8e\x76\xbe\xf9\xb2\xef\x52\x39\x5f\x2a\x19\x2c\x51\x86\x5a\xcb\xa9\xb7\x30\x6e\xd4\x01\x6a\x81\x97\x37\x8d\x87\x8c\x51\x4f\x1e\x3f\xc6\xf9\xeb\x2f\x91\x33\x68\xc0\x22\x63\x14\x63\x44\x85\xf6\x88\x68\xb5\x27\xac\xa7\x7b\x7b\x07\xd7\x8b\x8f\x72\x3e\xc0\xe0\x94\x50\x5d\xc2\x74\xe1\xa9\x42\x9c\x9c\x1f\x2e\x68\x56\xa4\x6d\x79\x6f\xcc\x6e\x85\x90\x71\x44\x29\x3d\xf3\x1f\x7f\x4c\x0c\xef\x13\x5f\xfb\x8b\xf8\xe7\xad\x90\x84\x20\x52\x86\x62\x53\xe5\x5c\x54\x39\xbf\x02\x00\x00\xff\xff\xea\x65\xf8\x2e\xd5\x00\x00\x00")

func resDefaultActionsFind_filesShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsFind_filesSh,
		"res/default/actions/find_files.sh",
	)
}

func resDefaultActionsFind_filesSh() (*asset, error) {
	bytes, err := resDefaultActionsFind_filesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/find_files.sh", size: 213, mode: os.FileMode(493), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsGoedRc = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x52\xb1\x6e\xdb\x30\x10\x9d\xcd\xaf\x38\x40\x9c\x82\x24\x4a\x3a\x16\xd0\x10\x34\x42\xe2\xa1\x76\x51\x1b\xf5\xc8\xd2\x22\x59\x1d\xaa\x92\x86\x78\xb6\x6a\x0b\xfa\xf7\x92\x94\x5b\x5b\x5e\x3a\x91\xb8\xf7\xee\xbd\xc7\x3b\x66\xf0\xe6\xb4\x82\xb6\x02\xb4\x48\x8c\xbd\x2d\xcb\x57\x31\x5f\xac\xd6\x2f\x8b\x4f\x65\xc1\x9f\xc7\xc2\xb7\x79\xb9\x29\xf8\x07\xc6\x32\x58\x69\xd9\x56\x35\xfc\x88\x5d\xd5\xde\x93\xfb\x95\x6f\xf7\xd8\x10\x5a\x20\xe7\x1a\x0f\x06\x5b\x4f\xc0\xbe\xbc\xac\xdf\x0b\xfe\xbe\xfc\x5c\xe6\x8f\x91\x9c\xcb\x8a\xd0\x59\x9f\x7f\xbc\x2e\x2a\x6d\xe4\xbe\xa1\x2b\x30\xf6\x31\x66\x6c\x72\x10\x6e\xa7\x2d\xf4\x6c\x96\x41\xba\xc9\xa0\xde\xe8\x10\x35\xa1\x6c\x96\x52\x3c\x3c\xc8\x1d\x8e\x38\x9f\xc4\x87\xef\xfd\xae\x53\x03\x84\x57\x0c\x17\x49\xad\x90\x46\xc9\x74\x9b\x4a\x82\xb4\x0a\x3a\x19\xea\xc6\xb5\x10\x0e\x72\xb0\xd5\xe0\xe5\x21\x60\x39\x54\x8d\xf3\x37\xbe\x49\xe4\x3f\xbe\x95\x1a\x0d\xc3\x89\x36\x28\x4a\x50\xd8\xea\x8a\x5c\x7b\xbc\x4f\x8e\xd6\x11\x9a\xe3\x98\xc0\x19\xa0\x5a\x83\xd5\x5d\x64\xb1\xd9\xdf\xe9\x86\x66\x7e\x37\xb1\x3e\xa0\xee\x84\xd7\x24\x3a\xd7\x8a\x9f\x81\x7c\x9b\x83\xff\x5b\xde\x39\x12\x1b\x80\xb1\xf2\x75\xbe\x5e\x7e\x2d\x2e\xc3\xc8\x60\x53\x87\xe1\x75\x48\xf5\x79\x0c\xf7\xc0\x47\x16\xa0\x1f\x47\x1d\x9f\xe1\xa0\xbf\x2c\x85\xdf\x0d\x30\xd9\x4a\x1e\xd3\x06\x96\x87\xde\xa7\x3f\x22\x48\xff\xa6\x47\x5f\x47\x6a\x04\x0c\xf4\x06\xad\x12\x91\xec\xcf\x75\x48\x1d\x27\xe8\x0f\xf4\xfc\xf4\x24\x3c\x9e\x74\x40\x62\xca\x49\x81\xb1\x3f\x01\x00\x00\xff\xff\x8a\xdb\xb8\xe1\xa8\x02\x00\x00")

func resDefaultActionsGoedRcBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoedRc,
		"res/default/actions/goed.rc",
	)
}

func resDefaultActionsGoedRc() (*asset, error) {
	bytes, err := resDefaultActionsGoedRcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed.rc", size: 680, mode: os.FileMode(420), modTime: time.Unix(1475089231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsGoedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x91\x51\x6f\x82\x30\x14\x85\x9f\xed\xaf\xb8\x29\x3c\x6c\xc9\x14\xdd\xe3\x12\x92\x99\x49\xd4\x87\xe9\x32\xcd\xf6\x58\x2b\x2d\xa3\x19\x6b\x09\xbd\x88\xba\xec\xbf\x8f\x02\x26\xe8\x63\xcf\xbd\xdf\xb9\x87\x83\x07\x73\x23\x05\xec\xb9\x4d\x41\x69\x85\x84\xc8\x63\x6e\x0a\x84\xf9\x3a\x9a\xb1\xe5\x6a\xb3\x9d\xae\x5e\xa2\xd0\x9f\x5c\xe9\x1f\xcb\xe8\x33\xf4\x1f\x09\xf1\x60\x23\x79\x11\xa7\xf0\xe5\x5c\xe2\xd2\xa2\xf9\x09\xf6\xa5\xca\x50\x69\x40\x63\x32\x0b\x89\x2a\x2c\xc2\x05\x7f\x9b\x6e\x17\xa1\xbf\x58\xbf\x46\xc1\xc8\x31\x01\x8f\x51\x19\x6d\x83\xa7\xbe\x28\x64\xc2\xcb\x0c\x7b\x43\xc7\x11\x92\x94\xba\x51\x9a\x73\x2c\x16\x77\xf7\xf0\x4b\x06\x1e\xc4\xa2\x0e\x8f\x06\x38\x08\x55\xc8\x18\x4d\x71\x7a\x00\xae\x05\x68\x83\x2a\x39\xb5\xe9\x4c\x02\x98\x4a\xd0\xb2\x72\x5b\x64\x70\x89\x59\xc3\xfe\x33\x19\x34\x3b\xc3\x21\xcf\x15\x1c\x94\xac\x98\x95\xc8\x2a\x53\x7c\xb3\x7a\x19\xfc\xab\x3a\xba\xa7\x6b\x01\xe8\x2e\xaf\xc4\x8e\x92\x3f\xe8\xc5\x33\x6d\xb0\x9e\xa5\xc9\xa5\xbe\x75\xe9\x50\x77\xdd\xd1\x5d\x43\xd1\x6c\xb9\x5d\xbf\x87\xb4\x07\x4b\xa1\xf0\x16\xee\x58\xaf\x1d\xf2\xba\xe6\x4c\x12\xc2\x33\xc5\x6d\xfd\x45\x2d\x5e\x37\x44\x2f\x9a\x0d\xa9\x6d\x7e\x15\x43\x79\xc4\x91\x4d\x69\x37\x48\x42\x9a\x28\x2d\x98\x33\xb0\x3d\xdd\x9e\x43\x7a\xc0\xc9\x78\xcc\xac\x3a\xcb\x66\x40\xae\xde\xe4\x3f\x00\x00\xff\xff\xdf\x60\x48\xea\x3c\x02\x00\x00")

func resDefaultActionsGoedShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoedSh,
		"res/default/actions/goed.sh",
	)
}

func resDefaultActionsGoedSh() (*asset, error) {
	bytes, err := resDefaultActionsGoedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed.sh", size: 572, mode: os.FileMode(420), modTime: time.Unix(1475089176, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsGoed_helperAnk = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\xc1\x4e\xfb\x30\x0c\xc6\xcf\xcd\x53\xf8\xdf\xff\x25\x95\x50\x81\x07\xd8\x61\xda\xaa\xa9\x97\x21\x31\x04\x07\x84\x50\xd4\xba\x10\xd1\x3a\x95\x93\x0e\x24\xb4\x77\xc7\x69\xc3\xe0\xc8\x2d\xfe\x3e\xfb\x67\x7f\xf9\x0f\x3b\x87\x2d\xbc\x62\x3f\x22\xc3\xe0\xda\xa9\x47\xe8\x1c\x83\xa1\x37\x07\xbe\x61\x3b\x06\x5f\xaa\xa3\x61\xc0\x0f\x6c\x60\x05\x76\x18\x1d\x07\x9d\x3b\x7f\x19\x95\xbc\x50\x2a\x8d\xcd\xa4\x4f\x95\x25\xe6\xd1\xe2\x3b\xd4\x5b\x50\x59\x37\x51\x33\x97\xba\x88\x7e\xc6\x18\x26\x26\x08\xae\xa6\xa0\x9d\x2f\x77\x18\x90\x8e\x3a\xdf\xdd\x54\xdb\xe7\xfb\xba\x7a\xc8\x8b\x42\x65\x27\x95\x9d\x59\x96\x7c\x30\xd4\xa0\xf0\x12\xee\x5b\xf9\x0b\xb2\xde\x1f\xee\xd6\xfb\x4d\xf5\x1b\x7b\x3b\x91\x07\x03\x8d\x1b\x06\x43\xed\x85\xa4\xb3\x41\x5f\x17\x60\x3b\xe8\x8c\xed\x7d\x5a\xc3\x13\x69\xc3\x2f\x7e\xd9\xd2\x0c\xad\xfc\x40\x8c\x5d\x6e\x96\xc1\xd9\x7c\xbc\x7a\x2a\x16\xb7\x5c\x4b\x29\x2d\x51\x15\x65\x64\x4b\xa1\x4f\x84\xd4\x71\x08\xad\x9b\x82\xf4\xc8\x95\xcb\xfb\xc7\x40\xe6\xb3\x21\x6f\x31\x16\x25\xda\x72\xb0\x8e\x0c\x39\x30\x8a\xff\x56\x40\xb6\x9f\xaf\x3a\xaf\x11\x3d\x76\x64\x02\xa8\x96\x38\x52\x9d\x62\xe4\x93\xfa\x0a\x00\x00\xff\xff\x98\x29\x53\x99\xe9\x01\x00\x00")

func resDefaultActionsGoed_helperAnkBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoed_helperAnk,
		"res/default/actions/goed_helper.ank",
	)
}

func resDefaultActionsGoed_helperAnk() (*asset, error) {
	bytes, err := resDefaultActionsGoed_helperAnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed_helper.ank", size: 489, mode: os.FileMode(420), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsGoimportsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\x41\x4b\x33\x31\x10\x86\xef\xf9\x15\xef\x47\xf7\xf8\xad\x0b\xde\x2b\x88\x16\xe9\xa5\x82\x8a\x1e\xdb\x6c\x76\xd2\x0d\x26\x99\x92\x64\x5b\x84\xfe\x78\x27\xb6\xc8\xe2\x71\x26\x4f\xde\xf7\x99\xc5\xbf\xae\x77\xb1\xeb\x75\x1e\x55\xa6\x82\x96\xb0\xb0\xda\x79\x90\x4e\xfe\x4b\xa9\x05\x5e\xa6\x88\x3d\xbb\x70\xe0\x54\x32\x38\x42\xcb\x88\xcc\x53\x32\x04\xeb\x3c\xfd\x17\xc8\x6a\xef\x33\x7a\x6d\x3e\x51\x58\xde\x6d\x28\x70\x76\xf6\x2f\x72\x81\x3e\x4a\xb0\xee\x3d\xdd\x28\x75\x74\x74\x5a\x36\x4f\xcf\xab\xc7\xed\xfb\x7a\xf5\xa1\x5c\xcc\xe5\x3a\xaf\x37\xaf\x6f\xf7\x9b\x87\x55\x2d\xdf\x8b\x52\xed\x80\x67\xa3\x8b\xe3\xa8\x72\x32\xcb\xdd\x9e\x69\x40\xdb\xea\x83\x43\x0d\xda\xca\x72\x2b\x04\x9a\x1a\x83\xa6\xee\x76\xca\x84\x61\xf9\x2b\xa0\x4e\xa3\x33\xe3\x4c\xe8\xf6\x0e\xdd\x40\xc7\x2e\x4e\xde\xe3\x7c\xc6\x85\x16\x6f\x69\x4d\x72\x72\x19\x09\x86\x43\xd0\x71\xa8\x47\xd7\xb1\x8a\xa8\x46\x40\xb4\x27\x34\xd2\x59\x51\xf2\xac\x87\x1f\x0b\xf4\x93\xb5\x94\x60\x13\x87\x0b\xfb\x57\xf3\x0a\xcf\x2c\xa1\xbe\x03\x00\x00\xff\xff\x23\x8f\x45\x09\x81\x01\x00\x00")

func resDefaultActionsGoimportsShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoimportsSh,
		"res/default/actions/goimports.sh",
	)
}

func resDefaultActionsGoimportsSh() (*asset, error) {
	bytes, err := resDefaultActionsGoimportsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goimports.sh", size: 385, mode: os.FileMode(493), modTime: time.Unix(1472885273, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsSearchAnk = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xcf\x4a\x03\x31\x10\x87\xcf\xce\x53\x8c\xe9\x25\xc1\xed\x9f\x17\x58\x41\x45\xbc\x8a\x3d\x96\x22\x61\x3b\xdd\x06\xd3\x64\x99\x24\x52\x11\xdf\xdd\x49\x44\xad\x5e\x16\x66\xe6\xfb\x7d\x33\x9b\xd9\xa5\x0d\x2f\x11\x60\x86\xf7\x27\x7b\x9c\x3c\x61\xad\x31\x0d\xec\xa6\xdc\xe1\x50\x98\x29\x64\xff\x86\x25\x94\x44\x3b\x00\x1f\xed\x4e\xab\x31\xd2\xee\xf9\x40\x7e\x22\x5e\x48\x40\x19\x80\x57\xcb\x18\x53\x87\x7b\x6f\xc7\x0e\xe9\x44\x03\xf6\xe8\x8e\x53\xe4\xac\x55\x4c\xca\x74\x3f\x55\x45\xce\xeb\x98\x96\x95\xaf\x96\x19\x3e\x5a\x4e\x72\x05\x8f\x09\x97\x4d\x96\xc0\x8d\x21\x32\xdd\x59\xe9\xf7\xad\xb5\xb8\x8d\xd1\x6b\xe5\x94\xac\xb3\x3e\x51\x87\xea\x97\x11\x4d\x63\x9a\x48\x1b\xb9\x78\xf8\x8e\xdd\x88\x55\x9b\xcd\x6a\x0b\xcd\xdf\xe3\x46\x8d\x4c\x93\x68\xd4\x3c\xb4\xaf\x57\x5b\x70\x7b\x3c\xdf\xd8\x63\xe6\x42\xf8\x0e\x17\x2d\x74\x55\x53\x73\x27\xdc\x47\x25\x3d\x05\x2d\x1b\x0c\x5e\xe3\xea\x1f\xc3\x62\x94\x91\x80\x48\x72\xe4\x9f\xe9\x57\xbf\xfe\xef\x53\x09\x98\xc8\xf2\x70\x80\x07\x79\xd5\x05\x97\xa0\x73\x5c\x67\x76\x61\x5c\x7b\x37\x90\xae\x21\x63\xe0\x33\x00\x00\xff\xff\x7e\x24\xb8\x7b\xab\x01\x00\x00")

func resDefaultActionsSearchAnkBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsSearchAnk,
		"res/default/actions/search.ank",
	)
}

func resDefaultActionsSearchAnk() (*asset, error) {
	bytes, err := resDefaultActionsSearchAnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/search.ank", size: 427, mode: os.FileMode(493), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsSearch_textSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\x3d\x8e\x84\x30\x0c\x46\xeb\xf5\x29\xbe\x35\x14\xbb\x45\x16\xb1\xe5\xfc\x5d\x82\x12\x51\x04\x26\x10\x4b\x28\xa0\xc4\x05\x73\xfb\x09\x43\x33\xe3\xca\xb2\xde\x7b\x72\xf1\x5d\xf5\x12\xaa\xde\x26\x4f\x54\xa0\x71\x36\x0e\x1e\xea\x36\xc5\x6a\x55\x5d\x0c\x90\x80\x51\x66\x97\x50\xe1\x2e\x31\xe1\x67\x8a\x6e\xfd\x25\x92\x11\x2d\xb8\x2c\x18\x66\x56\xd4\xe8\xce\x50\xef\x02\x21\x8f\x1b\xfc\x02\x6e\x1e\x41\xed\x76\x42\xc2\x25\xc7\x72\xeb\x86\x36\x2f\xbe\xe3\x03\xda\x24\x7b\x34\x0a\xd1\x7e\xbd\xf2\x1f\x7f\x54\xa7\xf7\xea\xd7\x81\x94\xff\xfc\x12\xf6\x1f\x60\x62\x10\x18\x33\x2c\xf3\x12\xb3\x53\x33\xca\x9d\xc2\x33\x00\x00\xff\xff\xb7\x62\x20\x19\xd5\x00\x00\x00")

func resDefaultActionsSearch_textShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsSearch_textSh,
		"res/default/actions/search_text.sh",
	)
}

func resDefaultActionsSearch_textSh() (*asset, error) {
	bytes, err := resDefaultActionsSearch_textShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/search_text.sh", size: 213, mode: os.FileMode(493), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsStatsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\x4a\x4d\xce\xc8\x57\xd0\xd5\x55\x70\x0b\x72\x75\x05\xd2\x5c\x69\x45\xa9\xa9\x70\x41\x17\x37\x90\x50\x4a\x9a\x82\x6e\x85\x42\x49\x6e\x41\x5a\x31\x5c\x26\xc4\x3f\x40\x21\x20\xc8\xdf\x39\x18\xa4\xa0\xa0\x58\x21\xb1\xb4\x42\xa1\x46\xa1\x38\xbf\xa8\x44\x41\xb7\x28\x5b\xc1\x58\xc7\x18\xc8\xcd\x48\x4d\x4c\x51\xd0\xcd\x53\x30\x03\x04\x00\x00\xff\xff\x85\x91\x8c\xbb\x71\x00\x00\x00")

func resDefaultActionsStatsShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsStatsSh,
		"res/default/actions/stats.sh",
	)
}

func resDefaultActionsStatsSh() (*asset, error) {
	bytes, err := resDefaultActionsStatsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/stats.sh", size: 113, mode: os.FileMode(493), modTime: time.Unix(1464737891, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultActionsVt100_sizeSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xb1\x0e\x82\x30\x10\x86\xf7\x7b\x8a\x13\x3a\x68\x22\xe1\x09\x18\x8c\x12\xc3\x82\x03\x04\x47\x28\x58\xa5\x09\x52\x43\x1b\x08\x3e\xbd\x5c\x25\x24\xca\xd2\xb4\x5f\xbf\xfb\x73\x77\xee\xc6\x2f\x65\xeb\x97\x5c\xd7\xa0\x85\x41\x4f\xa0\x7b\xe7\xb2\x41\xc1\xbb\x66\x04\x70\x31\x99\xa8\x31\x23\x6a\xf9\x16\x68\x14\x3e\xb9\xa9\x6a\x7c\x28\x71\xc3\x5e\x8a\xe1\xcb\xb7\x9a\x94\x4e\x0d\x7a\x8f\xf6\x5a\xa9\x46\xef\x00\xc8\x08\xd8\xf9\x12\x9e\xf2\x2c\x0a\xaf\x20\x5b\x6d\xe6\x77\x14\x27\xe9\x21\x3e\x86\x00\x54\x16\x14\x36\xd1\xf3\xf8\x4b\xda\xdc\x9c\x28\x32\x2a\x40\x46\xa0\x00\xca\x5c\x7b\x44\x7f\x3d\xf8\x57\xa6\xc1\xf2\xde\xac\x4c\x64\x44\x60\x69\x1d\x19\x9d\xb0\xf4\x3f\xff\x83\xa8\x6a\x85\x0e\xed\x21\x4b\x97\x35\x58\x77\x9e\xd8\x7a\xb6\xc4\xf9\x04\x00\x00\xff\xff\xcf\x77\x15\x3b\x50\x01\x00\x00")

func resDefaultActionsVt100_sizeShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsVt100_sizeSh,
		"res/default/actions/vt100_size.sh",
	)
}

func resDefaultActionsVt100_sizeSh() (*asset, error) {
	bytes, err := resDefaultActionsVt100_sizeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/vt100_size.sh", size: 336, mode: os.FileMode(493), modTime: time.Unix(1472885273, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultBindingsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x94\x3f\x6f\xdb\x3c\x10\xc6\x77\x7f\x0a\xc2\x19\x93\xd7\x42\x5e\x04\x45\x51\xa0\x53\xd2\xa1\x43\xa6\xa2\xb3\x40\x51\x67\x89\x35\xc5\x53\xc8\xa3\x1d\x77\xe8\x67\xaf\x74\x24\xf5\xaf\x49\x16\xd9\xfa\x3d\xa7\xe3\x3d\x3c\x1e\x6f\xc4\xb3\xec\xbd\x38\xc1\xb5\x42\xe9\xea\xa2\xc3\xe0\x41\xa8\x16\x5d\xed\x05\xa1\xf8\xf9\x5d\xc0\x19\x2c\xf9\xdd\x8d\xf8\x01\x20\x5a\xa2\xde\x7f\x29\x8a\x46\x53\x1b\xaa\x83\xc2\xae\x20\x85\x46\xba\xa2\x41\xa8\x8b\xca\x60\x55\x74\xd2\x13\xb8\x82\xbf\x8b\xcf\x92\xae\x3d\x1c\x1a\xdc\xdd\x0c\x69\xbe\xd5\x9a\x84\xb6\xe2\x4f\x71\x88\xdf\x68\x5b\x6b\xdb\xf8\x03\x61\x67\x06\xfd\x3f\xf1\xfc\x78\x2f\x3c\x49\x3b\x94\x70\x44\x27\x9e\xb9\xa6\x47\xa3\xd5\x49\x54\x81\x08\xad\xb8\x17\x45\x31\x3c\xb4\x17\x06\x8e\x74\x27\xfe\x17\x9d\xae\x6b\x03\x77\xe2\x41\x38\xdd\xb4\x14\xf3\x3c\xbd\x91\xe7\xc9\xc9\x66\x4a\x93\xc2\xf2\x7a\xcb\x30\x0c\x95\xd9\xae\xba\xdb\x0f\xa5\xed\xc5\x57\xb1\xf7\x40\xa5\x0a\xce\xa3\xdb\x8f\xf0\x81\x21\xf6\x60\x4b\x6d\x4b\x0b\x97\xf2\xac\xe1\xc2\xd2\xe7\x18\xaf\x1c\x1a\x53\x86\x9e\xd9\xfd\xa7\x25\xac\xf1\x62\x47\xfc\x94\x53\x1b\x50\x54\x72\x27\x18\x3f\xae\xf8\x65\x68\xcd\x80\xa5\xa1\xdb\xf1\xc3\x52\x3a\x87\x17\x0e\xb0\xf2\x9c\x73\x8d\xea\xb8\x33\x1b\x75\x44\x49\xe5\x5d\xda\xc8\xcc\x92\x1e\xfa\x8d\xc8\xa5\x57\x52\x9d\x7c\x2f\x15\x30\x9e\xdf\x76\x7b\x45\xce\xdc\x4a\xc6\x2d\x76\x13\xa9\x96\xa5\x4b\x63\x32\x57\xcc\x15\xf6\xd7\x4c\x62\x4a\xb0\x75\x06\x2d\x83\x0e\xcf\x90\xeb\x66\xfc\x6b\xc6\xb9\x5e\xe6\xa7\x99\x73\xa9\x0c\xcd\x0c\xd3\xce\x30\xb6\xef\xb5\x8b\x55\x5c\xa9\x5e\x76\xb0\x92\x5f\x58\x7e\x09\x7a\x5a\xda\x31\x71\x60\x50\x4e\xd5\xfb\x68\x5c\x9e\xa7\xad\xa0\x39\xed\x30\x1e\x5d\xc6\x81\x71\x3d\xec\x10\x41\xb9\xdc\xba\x33\x0b\xfd\x38\x4c\x19\xc5\x6e\x28\x83\x1e\xca\xcb\x30\x36\x38\x15\xf5\x1a\x95\x30\xd5\x74\x4d\x35\xd5\x98\xc9\x6f\x26\xc1\x32\x89\xeb\x2d\x96\x1e\xd9\xfa\x38\x2d\x77\x6d\x6c\xcb\xdc\x9e\x61\xa2\xc1\xa5\xf7\xf1\xdf\x40\xbc\x92\x7d\x4c\x47\xd8\x34\x06\x4a\xd5\xd5\x65\x25\x47\x8d\x3d\x2d\xce\xc5\xe6\x5c\x2e\x1b\x6c\xe1\x95\x92\xeb\x66\x5a\xbb\x77\x1a\xdd\x4c\xb9\xb9\x0e\x28\x38\xbb\x2a\x61\x7b\xa0\x57\x27\xc4\xb7\xfa\xf8\xcf\xbc\xa4\x53\x99\x96\x89\x21\xd9\x68\xd2\xa2\xdf\x28\x4d\x3e\x92\x96\xec\x44\x71\x63\x2a\x85\x24\x5b\x31\x64\x32\x97\xc4\xa5\xc7\x18\x31\x3b\x5d\x86\xb0\xe1\x18\xb0\xf5\x98\xc2\xd6\x2e\x57\x93\x9b\x22\x62\x8e\xd0\x83\xfb\xe0\xd2\x88\xfa\xfb\xd7\x46\xd4\x3f\xba\x38\x62\xc4\xdb\x57\x07\xc9\x78\x17\x8c\xbf\xbb\xfd\x2a\x66\x1a\xda\xdd\xdf\x00\x00\x00\xff\xff\x02\xe9\x71\xf0\x8f\x06\x00\x00")

func resDefaultBindingsTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultBindingsToml,
		"res/default/bindings.toml",
	)
}

func resDefaultBindingsToml() (*asset, error) {
	bytes, err := resDefaultBindingsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/bindings.toml", size: 1679, mode: os.FileMode(420), modTime: time.Unix(1474664438, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultConfigToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x4e\x4d\x6b\xc3\x30\x0c\xbd\xfb\x57\x88\xf4\x3e\x32\x06\x63\x3b\xf8\xb2\x95\x6e\x83\x15\x06\x09\xdb\xd9\xc4\x72\x2c\x88\xe5\x90\xc8\x6b\xb7\x5f\x5f\xb9\x25\x07\x21\xbd\x0f\xf4\xde\x0e\xf6\x18\x5c\x99\x04\x86\xcc\x81\x46\xd3\xfd\xb1\xb8\xf3\x3b\x8d\x71\xd2\x11\xe2\xd1\xca\x52\xd0\xf4\x11\x13\xda\xc6\xdf\xdc\x77\x92\xd3\xd4\x98\xa3\x3b\xbf\x26\xff\x52\x42\xc0\xe5\x93\x18\x57\xfb\xd0\xb6\xad\xd9\x41\x87\x02\x12\x11\x66\x27\x11\x24\x83\x83\x94\x39\xaf\xb3\x1b\x10\xfa\xfe\x00\x21\xb3\x54\xbe\xac\xa8\xda\x50\x56\xfd\x77\x25\xcd\x5b\xa1\x83\x6e\xdb\x34\xdb\xd9\xd1\x3f\xda\xfb\x76\x83\xfb\x99\xec\xf3\xa3\x66\x7c\x2d\x58\x73\xd1\x43\x22\xa6\x54\x12\xfc\x12\x9e\xe0\x44\x5e\xa2\x39\x12\x7f\x2b\xfa\xa9\xc0\x3e\xd5\x4a\x98\x27\x50\xdf\xe0\x24\x2f\xa6\x96\xbd\x6a\x1f\xec\x6f\x94\x9a\x2e\x01\x00\x00\xff\xff\x23\x64\x52\x8d\x0c\x01\x00\x00")

func resDefaultConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultConfigToml,
		"res/default/config.toml",
	)
}

func resDefaultConfigToml() (*asset, error) {
	bytes, err := resDefaultConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/config.toml", size: 268, mode: os.FileMode(420), modTime: time.Unix(1472885273, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultThemesAcmeToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xbd\x6e\xea\x30\x14\xc7\xf7\xfb\x14\x28\x73\x06\x07\x92\x40\x86\x2c\x38\xc9\x72\x75\x75\x87\x70\xef\x6e\xc0\xa2\x91\x1c\x8c\x8c\x11\x45\xea\xd0\x0d\xa9\x5b\xbf\x17\x3a\x54\x95\xaa\x3e\x17\x4f\xd2\x13\x7f\x84\x04\xa2\xc2\x72\xf2\xf7\xef\xfc\x7c\x6c\x79\xbc\x88\x9d\x34\x44\x03\x04\x3f\xe7\x57\x56\x7d\x8d\xa0\xcc\xaa\xaf\xf1\x22\xa7\x8c\xce\x64\xec\x8c\x23\x20\x32\x4d\xd8\x4c\x71\x08\x79\x15\x87\x37\x62\xcd\x45\x9b\xb3\xd9\x89\xc3\xbc\x2c\xe9\x12\x5a\xb3\x21\x1a\x6a\x2c\x97\xa2\x58\xc2\xa6\x7e\x84\x42\x9d\xfc\xa6\xbb\x2d\x17\x73\x2f\x76\xfa\x9e\x92\x79\x75\xd6\x8f\x9d\x30\x3b\xcb\x06\x90\x45\x36\xcb\xe9\x8a\x08\x22\xb9\x80\x6e\x1f\xa1\xc0\xec\x61\x53\xe8\x8f\x82\x8b\x14\x0c\x69\xbf\x4e\x77\xe5\x94\x33\x68\xc7\x18\xf9\x46\xaa\x22\xe8\x1d\x25\xed\x08\x1a\x13\xaf\x8e\x24\x91\x9b\xf5\x94\x88\x5e\xdc\x73\x8e\x6f\x77\x2e\x1e\xe8\x4b\xad\x8b\x06\x33\xa1\xd7\xb2\xe2\x4e\x57\xdd\x5a\x4a\x85\xb2\x20\x0f\xfe\x4a\x8e\xcb\x79\xc3\x1c\xa6\xc6\xdc\xb3\x95\x25\x2e\xbd\xa7\xfc\xef\xb2\xb1\x02\xce\xff\x05\xdd\x5a\xe9\xed\xe3\xd9\xb8\xaa\xd7\x10\x2d\xa9\x7e\x25\x05\xa3\x98\x51\xa2\x8c\xc7\xc3\x93\x6b\xf7\x6b\x9c\xb5\x62\x92\x42\xc8\x9d\x66\x5e\x5d\xec\xeb\x9d\x9b\xf7\x31\x13\x9c\x31\x3b\xc4\xcb\xbd\x9b\x06\xc6\xa3\x8a\x06\x33\x21\x53\xcd\x3c\xb8\xa1\x6f\x18\x55\x28\xe6\x0f\x17\xb4\x9a\x32\x2f\xe6\x54\x9f\xe7\xd3\xcd\x02\xfd\xc4\x5c\x78\xdc\xc3\x36\xf6\x6f\xa5\xa0\xfd\xfb\x4f\x50\xc2\xb7\xfa\x78\xfb\x8f\x2e\x0c\x06\xc2\x57\x66\xf0\xfd\x57\xa7\x88\x88\x45\xa1\x14\x37\x5d\xcb\x98\xf1\xb5\x1e\xf6\xf0\xdc\x75\x7f\xdf\x01\x00\x00\xff\xff\xd1\x1d\x4b\x92\x98\x03\x00\x00")

func resDefaultThemesAcmeTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultThemesAcmeToml,
		"res/default/themes/acme.toml",
	)
}

func resDefaultThemesAcmeToml() (*asset, error) {
	bytes, err := resDefaultThemesAcmeTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/themes/acme.toml", size: 920, mode: os.FileMode(420), modTime: time.Unix(1475266229, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resDefaultThemesDefaultToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xbb\xae\xda\x40\x10\x86\xfb\xf3\x14\x47\xae\xb7\xd8\xb5\xc1\xd8\x85\x8b\x78\x2f\x4d\x74\x94\xc2\x24\xbd\x81\x15\xb1\x64\xb3\x68\x31\x22\x48\x29\xd2\x21\xa5\xcb\xbd\x21\x45\x14\x29\xca\x73\xf1\x24\x99\xbd\x19\x27\x41\x50\x20\xfb\x9f\xef\xff\x67\x76\x3d\xe5\xba\x88\xf8\x33\x6c\x7f\xd1\x83\x80\x37\x26\x70\x82\x85\x79\x2b\xd7\x95\x6c\xe5\xb2\x1f\x6b\x62\xd0\xbc\x8b\x18\x8e\xee\xf5\x4e\xe9\x22\x8a\x09\x4e\x03\x17\xb4\x2b\x47\x55\xd7\xc9\x0d\x58\xc5\x0c\xcf\x1c\x56\xf5\xba\xd9\x40\xd3\x49\x1e\x8c\xcf\xe5\xf1\xa0\xf4\x8a\xd8\x30\xd3\x94\x0c\x5a\x5c\x44\xa9\xf8\x47\x4b\x40\xcb\x83\x56\xc9\x6d\xad\xeb\x5e\x69\x70\x4f\x30\x9e\xfa\x1e\x41\x05\x7f\x3e\xfd\x4f\x85\x04\x1e\x0f\xea\xb1\x5b\xa8\x16\xec\x94\xe2\x89\x0f\xb5\x12\x78\x33\xf6\xb7\x04\x46\x46\x06\xa9\xaf\xfb\xfd\x6e\x51\xeb\xc7\xe2\x31\xba\x7c\x7f\x8f\x78\x09\x67\x84\x9f\x79\xf0\xb7\x3b\x30\x73\xf9\xa6\x37\x5c\xc9\xd2\x34\xdc\xc3\xa8\xc4\xb5\x4d\xc1\x84\x4e\x5d\x38\xed\x56\xf7\x93\x1d\x70\x8d\x0d\x5f\xeb\xaa\xbf\xd8\xd8\x8a\xb0\x0d\x21\xf2\x55\x23\x0f\x21\xf3\xdd\x27\xc4\x59\xc8\x64\xf0\x6f\xbd\x9e\x08\xa1\x7c\x86\x63\xbf\x24\x4d\x2b\x69\x2b\x6b\x9b\x78\x39\x7f\x46\x84\x42\x49\x18\x73\x3e\x2c\x12\x30\xac\xd1\xfd\xd1\x31\xdf\x10\x2c\x01\x31\x9d\x47\x4c\xb5\xd4\xaa\x6d\xc3\x10\x5f\x3f\x98\xde\x99\xcb\x31\x0f\x23\x66\x5e\x2f\x1c\xf3\x11\x65\x6e\x73\x10\xce\xdc\xc0\xd1\xc3\x93\xd2\xd2\x0c\x59\x35\x2b\xe9\x8e\xf3\x0b\x11\x61\xbf\x0b\x24\x0d\xab\x1d\xb0\x97\x5b\x0b\x9d\x7e\xdc\x83\x98\x3a\xb8\xd3\x9d\x7e\xde\xc2\x60\x1e\xfa\xda\xcf\x7d\xfa\x8d\xb8\x70\x77\x66\x08\x3f\xf7\x53\xad\xd7\x8d\x8d\x78\x7b\xab\x4c\x5b\xb5\x73\xc3\x9e\xbf\xa0\x24\x71\xbb\x8f\x38\x0d\xf5\x3f\x01\x00\x00\xff\xff\xec\x79\x6d\x1e\x96\x03\x00\x00")

func resDefaultThemesDefaultTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultThemesDefaultToml,
		"res/default/themes/default.toml",
	)
}

func resDefaultThemesDefaultToml() (*asset, error) {
	bytes, err := resDefaultThemesDefaultTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/themes/default.toml", size: 918, mode: os.FileMode(420), modTime: time.Unix(1475253000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resResources_versionTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\x34\x31\x37\x37\x30\x37\x37\x37\xb1\xe0\x02\x04\x00\x00\xff\xff\x6d\xde\xea\x3d\x0b\x00\x00\x00")

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

	info := bindataFileInfo{name: "res/resources_version.txt", size: 11, mode: os.FileMode(420), modTime: time.Unix(1477077748, 0)}
	a := &asset{bytes: bytes, info: info}
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

	info := bindataFileInfo{name: "res/themes/Readme.md", size: 22, mode: os.FileMode(420), modTime: time.Unix(1464737891, 0)}
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
	"res/.DS_Store": resDs_store,
	"res/Readme.md": resReadmeMd,
	"res/actions/Readme.md": resActionsReadmeMd,
	"res/default/actions/find_files.sh": resDefaultActionsFind_filesSh,
	"res/default/actions/goed.rc": resDefaultActionsGoedRc,
	"res/default/actions/goed.sh": resDefaultActionsGoedSh,
	"res/default/actions/goed_helper.ank": resDefaultActionsGoed_helperAnk,
	"res/default/actions/goimports.sh": resDefaultActionsGoimportsSh,
	"res/default/actions/search.ank": resDefaultActionsSearchAnk,
	"res/default/actions/search_text.sh": resDefaultActionsSearch_textSh,
	"res/default/actions/stats.sh": resDefaultActionsStatsSh,
	"res/default/actions/vt100_size.sh": resDefaultActionsVt100_sizeSh,
	"res/default/bindings.toml": resDefaultBindingsToml,
	"res/default/config.toml": resDefaultConfigToml,
	"res/default/themes/acme.toml": resDefaultThemesAcmeToml,
	"res/default/themes/default.toml": resDefaultThemesDefaultToml,
	"res/resources_version.txt": resResources_versionTxt,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"res": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{resDs_store, map[string]*bintree{}},
		"Readme.md": &bintree{resReadmeMd, map[string]*bintree{}},
		"actions": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resActionsReadmeMd, map[string]*bintree{}},
		}},
		"default": &bintree{nil, map[string]*bintree{
			"actions": &bintree{nil, map[string]*bintree{
				"find_files.sh": &bintree{resDefaultActionsFind_filesSh, map[string]*bintree{}},
				"goed.rc": &bintree{resDefaultActionsGoedRc, map[string]*bintree{}},
				"goed.sh": &bintree{resDefaultActionsGoedSh, map[string]*bintree{}},
				"goed_helper.ank": &bintree{resDefaultActionsGoed_helperAnk, map[string]*bintree{}},
				"goimports.sh": &bintree{resDefaultActionsGoimportsSh, map[string]*bintree{}},
				"search.ank": &bintree{resDefaultActionsSearchAnk, map[string]*bintree{}},
				"search_text.sh": &bintree{resDefaultActionsSearch_textSh, map[string]*bintree{}},
				"stats.sh": &bintree{resDefaultActionsStatsSh, map[string]*bintree{}},
				"vt100_size.sh": &bintree{resDefaultActionsVt100_sizeSh, map[string]*bintree{}},
			}},
			"bindings.toml": &bintree{resDefaultBindingsToml, map[string]*bintree{}},
			"config.toml": &bintree{resDefaultConfigToml, map[string]*bintree{}},
			"themes": &bintree{nil, map[string]*bintree{
				"acme.toml": &bintree{resDefaultThemesAcmeToml, map[string]*bintree{}},
				"default.toml": &bintree{resDefaultThemesDefaultToml, map[string]*bintree{}},
			}},
		}},
		"resources_version.txt": &bintree{resResources_versionTxt, map[string]*bintree{}},
		"themes": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resThemesReadmeMd, map[string]*bintree{}},
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

