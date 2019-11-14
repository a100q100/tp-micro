// Code generated by go-bindata.
// sources:
// .gitignore
// __tp-micro__gen__.lock
// api/handler.go
// api/router.go
// args/const.go
// args/type.go
// args/var.go
// doc/APIDoc.md
// doc/README.md
// doc/databases.md
// errs/errs.go
// sdk/rpc.go
// sdk/rpc_test.go
// DO NOT EDIT!

package tpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
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

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x4f\x4b\xc4\x30\x10\xc5\xef\xef\xa3\x0c\x38\xa0\xa0\x78\xf5\xe0\x45\xf0\x0f\x5e\x45\x42\x9b\x4e\x63\xd7\x6d\x67\x4c\xd2\xb5\xee\x52\x3f\xbb\xc4\x54\xf0\xf2\x92\xf7\x63\xf8\xcd\x10\x2b\x88\x1b\x10\x27\x85\xd3\x76\x07\x97\x25\x65\x10\xbf\x5c\x5e\x5d\x1f\x3e\x5e\xb1\xbd\xac\x73\xa1\x3e\xe8\x39\x07\xad\xbf\x0b\xf6\x70\x3e\xa8\xeb\xa4\x9f\xa7\xbf\x12\x34\x7f\x99\xa4\x32\xf5\xdb\x65\x31\x8d\x99\xa9\x9a\xc7\x66\x98\xaa\x40\x16\xa9\xf9\x0d\xe2\x6d\xa9\x45\xed\x41\x1c\x9b\x08\xe2\xe3\x60\x20\x0e\xc7\xc2\x53\x07\xe2\x76\x2c\xe9\xfb\x50\x88\xe5\x05\xc4\x7b\x0d\xa0\x49\xdf\x66\xdb\x0e\x3c\x24\xaf\x5d\x11\xa7\xb9\xdd\x0f\xa3\x9c\x59\xd4\x9d\xf8\xfc\x8f\x7c\x6a\x7c\x4f\xd6\x78\xc1\xe9\xf4\xf4\xfc\x78\xe7\x1e\x6e\xee\x6f\xd7\x15\x3f\x01\x00\x00\xff\xff\x78\x2c\x45\x8a\x0c\x01\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 268, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var ___tpMicro__gen__Lock = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xc1\x0d\x82\x21\x0c\x06\xd0\xbb\x53\x7c\x32\x00\x03\x78\xf7\xaa\x17\x06\xc0\x60\x0b\x4d\xa0\x4d\x28\x2a\x6e\x6f\xfc\xaf\x2f\xef\x76\x4f\xd7\xcb\x09\xa9\x89\x23\x0f\x29\xd3\x50\x49\x33\x8a\x8d\xf1\xd0\x27\x4c\xfb\x17\xc5\xde\x34\x1d\x2c\x9d\x1c\x1f\x59\x0d\xab\x11\x42\xac\xa4\xb1\x5a\x80\xbf\x98\x65\x43\xf8\xf0\x7f\x03\x6d\xf1\xe5\xe7\x5f\x00\x00\x00\xff\xff\x83\x15\xc3\x44\x5f\x00\x00\x00")

func __tpMicro__gen__LockBytes() ([]byte, error) {
	return bindataRead(
		___tpMicro__gen__Lock,
		"__tp-micro__gen__.lock",
	)
}

func __tpMicro__gen__Lock() (*asset, error) {
	bytes, err := __tpMicro__gen__LockBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "__tp-micro__gen__.lock", size: 95, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiHandlerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\xc8\xe4\x02\x04\x00\x00\xff\xff\x0c\x0c\x0a\x62\x0c\x00\x00\x00")

func apiHandlerGoBytes() ([]byte, error) {
	return bindataRead(
		_apiHandlerGo,
		"api/handler.go",
	)
}

func apiHandlerGo() (*asset, error) {
	bytes, err := apiHandlerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/handler.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiRouterGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8c\x31\x0e\x02\x31\x0c\x04\x6b\xfc\x0a\xeb\xaa\x3b\x8a\x8b\x44\xc1\x43\xf8\x41\x30\x26\x89\xb8\xc4\x91\xe3\x20\x21\xc4\xdf\x51\x80\x82\x72\x67\x77\xa7\x7a\xba\xf9\xc0\xe8\x6b\x02\x48\xb9\x8a\x1a\xce\xb0\x9b\x42\xb2\xd8\xcf\x2b\x49\x76\x91\x8b\x3e\x36\xe6\x03\x15\xc7\x5a\xc9\xdd\x8f\x13\x2c\x00\xce\x21\xf5\x66\x92\x4f\xd2\x8d\x51\x39\xa4\x66\xac\xed\x47\x31\xfa\x72\xd9\x46\x36\x41\x1d\x13\x5d\xe1\xda\x0b\xfd\xbf\xe6\x6f\x81\xfb\x21\x5e\x3f\x48\x17\x7c\xc2\x0b\xde\x01\x00\x00\xff\xff\xfc\xfb\x49\xc0\x99\x00\x00\x00")

func apiRouterGoBytes() ([]byte, error) {
	return bindataRead(
		_apiRouterGo,
		"api/router.go",
	)
}

func apiRouterGo() (*asset, error) {
	bytes, err := apiRouterGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/router.go", size: 153, mode: os.FileMode(420), modTime: time.Unix(1573737148, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsConstGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x4a\xce\xcf\x2b\x2e\x51\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\xf5\x9a\x10\x2f\x17\x00\x00\x00")

func argsConstGoBytes() ([]byte, error) {
	return bindataRead(
		_argsConstGo,
		"args/const.go",
	)
}

func argsConstGo() (*asset, error) {
	bytes, err := argsConstGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/const.go", size: 23, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsTypeGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x2a\xa9\x2c\x48\x55\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\x61\x1b\x80\x25\x16\x00\x00\x00")

func argsTypeGoBytes() ([]byte, error) {
	return bindataRead(
		_argsTypeGo,
		"args/type.go",
	)
}

func argsTypeGo() (*asset, error) {
	bytes, err := argsTypeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/type.go", size: 22, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _argsVarGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\x4a\x2f\xe6\xe2\x2a\x4b\x2c\x52\xd0\xd0\xe4\x02\x04\x00\x00\xff\xff\xa5\xca\xdc\xfb\x15\x00\x00\x00")

func argsVarGoBytes() ([]byte, error) {
	return bindataRead(
		_argsVarGo,
		"args/var.go",
	)
}

func argsVarGo() (*asset, error) {
	bytes, err := argsVarGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "args/var.go", size: 21, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docApidocMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\x51\x6b\xd3\x50\x18\x7d\xff\x7e\xc5\x21\x7d\x0b\xde\x85\xce\x17\xa9\xf8\x24\xe2\x8b\x4f\x32\x9f\xd7\xd8\x66\x6d\x64\x4d\x42\x9b\x49\x86\x77\xd0\x8d\x0d\x1d\x65\x4d\xc1\x59\x6b\xd5\x41\x51\xe7\x40\x30\x0e\x26\x65\xb3\x2c\x7f\x66\xf7\x36\x7d\xda\x5f\x90\x9b\xa4\x5b\xdf\x76\x21\x90\xf3\xdd\xef\x3b\xdf\x39\xe7\xa2\x80\xe9\xe1\x89\x7c\x37\x96\x27\x23\xf1\xf5\x23\x11\x0a\x05\x14\x97\x8a\x98\xb5\x87\x49\xfc\x56\xfc\x7c\x2f\x7a\x07\x49\x77\x2c\xc2\x7e\xde\xd7\xfd\x21\xc2\x6f\x44\xba\x3e\xfd\xdd\x4e\x8e\xb7\x65\x18\x26\xf1\x9f\xeb\xc9\x50\xd7\x41\xc4\xb0\xd8\x9b\xfc\xdd\x95\xfd\x4f\x57\xf1\x48\x6e\x47\x6a\x22\x89\xc6\xf2\x74\xe7\xc5\xf3\x67\x79\x3b\x43\x19\x75\xdf\xf7\x4a\x86\x11\x04\xc1\x92\xfa\x2a\x6e\xc3\xd8\x68\x59\x4d\xe3\x75\xd1\xf0\x9a\xee\x9a\xbd\x6e\x19\x35\xcb\x5f\x7d\xb9\xb9\xea\x98\x0d\x8b\x80\x5b\x22\xd9\x3f\x17\x93\x30\xe5\x22\x86\xa7\x4f\x56\x14\xe3\x63\xd7\xf1\x2d\xc7\x67\x2b\x9b\x9e\x55\x82\xe9\x79\xeb\x76\xc5\xf4\x6d\xd7\x31\x5e\xb5\x5c\xe7\x61\xa5\x6e\x36\x5b\x96\xff\x68\xc3\x5f\x63\x0f\x14\x95\x08\x77\xe4\x87\x1b\xfd\x3c\x83\xa2\x77\xc0\x45\xbc\x37\x6b\xef\xf3\xe9\xe9\x3f\x71\xd4\xe1\x49\x74\x26\x07\x5d\x4e\xbc\xc4\x18\x63\x00\x90\xfe\x65\x90\x81\xb3\xbc\xcc\x89\x43\x09\x05\x87\x1c\x44\xe0\x2d\xbf\x69\x3b\x35\xf0\x3c\xe5\x2c\xcf\x74\x9c\x08\xf3\xed\xd3\xef\x17\x57\x97\x1d\x5d\x27\x2a\x97\xcb\x50\x3a\x09\x78\x43\x48\x8f\xa6\xe8\xb4\x12\xb4\x86\x5d\x69\xba\x1a\x01\x5b\xaa\x2d\x1d\x4f\xe2\x43\xf1\xf9\x28\x23\xc9\x14\x66\x2e\x80\x1b\x1f\xe0\x72\x10\x89\xde\xb1\x88\xf7\xc4\xe8\x97\x5a\x9c\xbd\x58\xae\x95\x2d\x1c\xdc\x01\x89\xc3\xae\xa6\x7c\x3c\x73\x63\x57\x55\xcd\xac\x59\x8b\x45\x71\x7e\x36\xbb\xdc\x55\x17\x5e\xdd\x75\xac\x85\x1b\xb9\xdf\x91\x5f\x2e\x44\x38\xe6\xe9\x1b\xa6\xda\xe7\xd6\xaf\x27\x43\x22\xdc\xda\xcf\xdd\x6b\x76\x55\x2b\xa1\xb8\x7c\xff\x5e\x8e\xcd\x9a\xca\x62\x79\x0e\xd3\x15\x2a\x9c\x20\x08\x02\x8d\xb0\x85\x8c\xe5\x7f\x00\x00\x00\xff\xff\x69\x55\x52\x40\xd7\x02\x00\x00")

func docApidocMdBytes() ([]byte, error) {
	return bindataRead(
		_docApidocMd,
		"doc/APIDoc.md",
	)
}

func docApidocMd() (*asset, error) {
	bytes, err := docApidocMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc/APIDoc.md", size: 727, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func docReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_docReadmeMd,
		"doc/README.md",
	)
}

func docReadmeMd() (*asset, error) {
	bytes, err := docReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc/README.md", size: 0, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docDatabasesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x5d\x6b\x1a\x4d\x14\xc7\xef\xe7\x53\x1c\xf0\x62\x5d\x48\x1e\xdc\x87\x06\x0a\xc5\x8b\x8d\x6e\x5b\xa9\xae\xc5\xac\x81\x5c\x39\x6b\x76\x6b\x04\x5d\x13\x5d\x03\x85\xb9\x68\x9b\x96\x1a\x42\x6b\xa1\x2f\x49\x89\x84\x5a\x84\x4a\x28\x4d\x0a\xd2\x50\x2d\xed\x97\xe9\x8c\xfa\x2d\xca\xec\xaa\xb3\xda\x36\xa1\x73\x33\xe3\xe1\xfc\xfe\xe7\xc5\xfd\x87\x42\x90\xba\xbf\xb6\x53\x42\x28\x14\x02\xe5\x3f\x05\x86\xaf\xba\xac\x71\xf1\xf3\x47\x9b\x3d\x3c\x1b\xb5\xbb\xe1\x7a\xcd\xae\xca\x88\x00\x6d\x3e\x62\xaf\xcf\x81\xc0\xe8\xac\xc7\x8e\x9e\x03\x81\xe1\xe7\x01\x3d\x39\x00\x02\xe3\xc1\xd1\xe8\x53\x87\x3e\xf8\xc6\x83\xfb\x0d\xd6\xfa\x08\x04\x91\x65\xef\x5c\x7a\x21\x02\x45\x8b\x4b\x3e\x3d\xa5\xef\x4f\x12\x16\x01\x9c\x2f\x16\x8a\x8e\x1b\xfe\x3f\x22\x63\x20\xb0\x0c\x04\xd6\x15\xae\x06\x8e\x59\xb6\xb9\xbe\xd7\x1e\xfd\xf0\x92\xbe\x78\x06\x04\xf0\xae\x59\xdd\xdc\x32\xab\x53\x20\x22\x00\xb3\x10\xc8\xff\xda\x1b\x7f\x7f\xcc\xf3\xb9\xb8\xf2\x5b\xee\xf6\x56\xc5\x11\xd9\x6c\xff\x80\xb5\xfa\xb4\x79\x11\x2c\xa0\xac\x78\x90\x24\x4d\x29\x84\x31\xae\xed\x94\x50\x2c\xa3\xa9\x86\x06\x86\xba\x9a\xd4\x00\xf3\x75\x61\x08\x23\x00\x5c\xb4\x30\x88\x71\x40\x4f\x1b\xa0\x67\x93\x49\x50\xb3\x46\x3a\x97\xd0\x63\x19\x2d\xa5\xe9\x06\xc4\xd2\x29\xef\x96\x26\x5b\x88\x4b\x4b\x9c\xe6\xf3\x62\x08\x8c\x27\x04\xe2\xda\x4d\x35\x9b\x34\x40\x8a\x48\x82\xf6\x57\xe2\xb3\x66\xc1\xc6\x30\x99\xf4\x0f\x58\x90\xf2\x16\xe3\x53\xde\x12\x44\x49\x65\xe5\x0a\x76\xb6\x26\x1f\xaf\x6f\x5b\xa6\x6b\x5b\x39\xd3\x9d\x8d\xad\x28\x57\x75\xcd\x8e\x7b\xec\xcd\x39\x3b\xfc\x32\x3e\xec\xf9\x32\x9b\x55\xfb\x52\x19\xd1\x7a\xe3\x98\x0e\xfa\x41\xd4\xb2\x4b\x36\x47\xdd\xda\x3f\x74\x40\x1b\xef\xc6\x6f\x3b\xbe\x4c\x38\x32\x6a\x77\x87\x9d\x3e\x6b\x9d\xfa\x61\xd9\x13\xbe\x9b\x49\xa4\xd4\xcc\x06\xdc\xd1\x36\x20\xcc\xff\x56\x79\x09\xc9\xa0\xe9\xb7\x12\xba\x16\x4d\x38\x4e\x25\xbe\x3a\xd3\x8e\xdd\x56\x33\x6b\x9a\x11\xad\xbb\xf7\xae\x97\xf3\xd7\xa6\x75\xa2\xd2\x82\xaf\xa4\x1b\xde\x07\x84\x3c\xdf\xa5\x2a\x4e\xa1\x62\xe5\x17\x3c\x48\x9f\xec\xcd\xdb\x30\x57\xb6\x5d\x53\x46\xc2\x8c\x20\xec\x08\xde\x8f\x89\x25\xc5\xf9\x9b\x39\xc1\x77\x20\x4c\x1f\x73\xcf\x85\xc0\xe4\xe1\x9b\x6a\xd7\x74\xcd\xaa\xf0\x55\xa7\x47\xf7\x9a\xf3\x3e\x9c\x9a\x6b\x52\x7f\x5d\xe1\x17\x42\xbf\x02\x00\x00\xff\xff\x00\x07\x32\xfd\x69\x04\x00\x00")

func docDatabasesMdBytes() ([]byte, error) {
	return bindataRead(
		_docDatabasesMd,
		"doc/databases.md",
	)
}

func docDatabasesMd() (*asset, error) {
	bytes, err := docDatabasesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc/databases.md", size: 1129, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errsErrsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcc\xb1\xce\x82\x30\x10\xc0\xf1\x99\x7b\x8a\x4b\x27\x48\x08\x2d\xdf\xf0\x6d\x4e\x4e\x2e\xc6\xc4\x27\x38\xeb\x05\x1a\xa1\xd4\xeb\x81\xf1\xed\x0d\x06\x5d\x9c\xff\xbf\xfc\x13\xf9\x1b\x75\x8c\x2c\x92\x01\xc2\x98\x26\x51\x2c\xa1\x30\x5d\xd0\x7e\xbe\x34\x7e\x1a\x6d\xcf\x51\x9e\x03\xf3\x9f\x8f\x96\x25\x79\xbb\xfc\x1b\xa8\x00\x16\x92\x95\x5a\x8b\x87\xb8\xd0\x10\xae\x27\x12\x1a\x59\x59\x30\x2b\xe9\x9c\xa1\xf8\x09\x3b\x5c\x0f\xcd\x91\x1f\xe7\x37\x29\x5b\xe7\x9c\x6b\x6b\x34\x1b\xc5\xaf\x35\x35\x9a\xfd\x14\x95\x42\xcc\x18\xb6\x2a\x7c\x9f\x39\x2b\xa6\x8f\xca\xa6\x82\x0a\x5e\x01\x00\x00\xff\xff\x68\xaa\xf8\xb3\xc7\x00\x00\x00")

func errsErrsGoBytes() ([]byte, error) {
	return bindataRead(
		_errsErrsGo,
		"errs/errs.go",
	)
}

func errsErrsGo() (*asset, error) {
	bytes, err := errsErrsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errs/errs.go", size: 199, mode: os.FileMode(420), modTime: time.Unix(1573737148, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpcGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpcGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpcGo,
		"sdk/rpc.go",
	)
}

func sdkRpcGo() (*asset, error) {
	bytes, err := sdkRpcGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpc_testGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpc_testGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpc_testGo,
		"sdk/rpc_test.go",
	)
}

func sdkRpc_testGo() (*asset, error) {
	bytes, err := sdkRpc_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc_test.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1571215291, 0)}
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
	".gitignore":             Gitignore,
	"__tp-micro__gen__.lock": __tpMicro__gen__Lock,
	"api/handler.go":         apiHandlerGo,
	"api/router.go":          apiRouterGo,
	"args/const.go":          argsConstGo,
	"args/type.go":           argsTypeGo,
	"args/var.go":            argsVarGo,
	"doc/APIDoc.md":          docApidocMd,
	"doc/README.md":          docReadmeMd,
	"doc/databases.md":       docDatabasesMd,
	"errs/errs.go":           errsErrsGo,
	"sdk/rpc.go":             sdkRpcGo,
	"sdk/rpc_test.go":        sdkRpc_testGo,
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
	".gitignore":             &bintree{Gitignore, map[string]*bintree{}},
	"__tp-micro__gen__.lock": &bintree{__tpMicro__gen__Lock, map[string]*bintree{}},
	"api": &bintree{nil, map[string]*bintree{
		"handler.go": &bintree{apiHandlerGo, map[string]*bintree{}},
		"router.go":  &bintree{apiRouterGo, map[string]*bintree{}},
	}},
	"args": &bintree{nil, map[string]*bintree{
		"const.go": &bintree{argsConstGo, map[string]*bintree{}},
		"type.go":  &bintree{argsTypeGo, map[string]*bintree{}},
		"var.go":   &bintree{argsVarGo, map[string]*bintree{}},
	}},
	"doc": &bintree{nil, map[string]*bintree{
		"APIDoc.md":    &bintree{docApidocMd, map[string]*bintree{}},
		"README.md":    &bintree{docReadmeMd, map[string]*bintree{}},
		"databases.md": &bintree{docDatabasesMd, map[string]*bintree{}},
	}},
	"errs": &bintree{nil, map[string]*bintree{
		"errs.go": &bintree{errsErrsGo, map[string]*bintree{}},
	}},
	"sdk": &bintree{nil, map[string]*bintree{
		"rpc.go":      &bintree{sdkRpcGo, map[string]*bintree{}},
		"rpc_test.go": &bintree{sdkRpc_testGo, map[string]*bintree{}},
	}},
}}
