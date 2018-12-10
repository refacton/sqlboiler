// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (6.883kB)
// override/templates/singleton/mysql_upsert.go.tpl (1.07kB)
// override/templates_test/singleton/mysql_main_test.go.tpl (4.883kB)
// override/templates_test/singleton/mysql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.845kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x48\x0e\x7f\xb6\x3f\x05\xd7\xe8\xee\xca\x07\x55\xed\x01\x87\x7b\xc8\x21\x0f\xcd\x9f\x76\x73\x4d\xda\x24\x6e\x2e\xc0\x05\x41\x31\x91\x28\x67\x90\xf1\x8c\x3a\x1a\x25\xf5\xe9\xf4\xdd\x0f\xe4\x48\x96\xe4\xd8\x8e\xfb\x27\x87\x7d\x8a\x35\x43\x91\x1c\xfe\xc8\x1f\x39\x4a\x59\xbe\x84\x17\x42\x49\x91\xc3\xce\x2e\x44\x6f\xe8\x17\xe6\xd1\x27\x71\xa3\x10\xfc\x9f\xe8\x83\x98\x61\x55\x0d\x59\x34\x8f\x6f\x71\x26\xfc\x36\xbd\xd0\x4a\xc0\x7f\x21\x9a\xb4\xbb\xfc\x82\x4c\x21\x7a\x93\x24\xef\x94\xb9\x11\x0a\x5e\x56\xd5\xf0\xd5\x2b\xb8\xc8\x72\xb4\xee\x1d\x08\xe7\x70\x96\xb9\x1c\x84\x06\xa9\x69\x2d\x04\xa1\x13\x48\x0c\xf2\x5a\x91\x25\xc2\x21\x18\x0b\x72\xaa\x8d\x45\x30\x1a\x62\xa3\x53\x25\x63\x17\x0d\xd3\x42\xc7\x10\x18\xf8\x4b\x59\x7a\xff\xa3\x8b\x6c\x22\xf5\xb4\x50\xc2\x56\xd5\xb8\xb1\x12\xb0\x13\xda\x38\x88\x3e\x98\x7d\xa3\x1d\x7e\x75\x55\x15\xbb\xaf\xa4\x8a\x1e\xa2\x7a\x31\x84\xb2\x44\x9d\x90\x93\xb5\xe5\x7d\xa3\x8a\x99\xce\xc3\xda\xb9\xfa\x11\x6e\x8c\x54\x51\xfd\x30\x06\xb4\xd6\x58\x28\x87\x03\x8b\xae\xb0\x1a\x4c\xe4\x0d\x7b\xbb\x5d\x9b\xfc\xde\x3b\x74\x07\x7b\xc1\xb8\x2c\x51\xe5\xc8\x7e\x84\xd0\x6c\xd4\x92\xf5\xbe\x4e\xaa\x2a\xdc\xe8\xc9\x78\x58\x0d\x87\x0b\xa7\x87\x3e\xdc\x14\xc0\x4e\xc8\xe9\xe7\xa9\xd0\x32\x5e\x0a\xfe\xe9\x8f\x45\x1f\x58\x67\x4e\x6b\x1c\x80\xad\xe1\x38\x7d\x6e\x3c\xca\xe1\x40\xa6\xe4\x14\x65\xe7\xff\x13\x8c\x7f\xb0\xd1\x5f\x76\x41\x4b\x45\x5e\x0c\x32\x0a\x51\xc0\xfa\x2e\xad\xc8\x0e\xad\x0d\xd0\xda\xf1\x78\x38\xa8\x56\x01\xb7\x06\xa9\x55\x40\x41\x91\x4b\x3d\xa5\x67\xfc\x8a\x71\xe1\x8c\xfd\x96\xc2\xe9\xa8\xce\xbe\x0f\xc5\xd3\xc7\xf1\x24\x47\x7c\xec\x0e\x6b\x97\x3a\x51\x7d\x0c\x6d\x2b\x5e\x2f\x75\xde\x7a\x3a\xd6\xdb\x43\xbe\x22\xcf\xba\x79\x45\x6e\x3c\x1f\xac\xf7\xc2\xc2\x6c\x3e\x39\x3b\x5e\x19\xcc\x0b\x2d\xbf\x14\x8d\x55\xd8\x85\xab\xeb\xdc\x59\xa9\xa7\x25\xf3\xac\x15\x7a\x8a\xf0\x42\x86\xf0\x22\x36\xaa\xc3\xb4\xcd\x0b\x64\x61\x40\x92\x32\x65\x91\xc8\xeb\xa3\xd5\x51\x59\xf2\x8a\xa7\xed\x51\xe8\xe5\x1a\xb7\xea\xdf\x15\x7b\xbb\xc8\x85\xe7\xc8\xb2\x09\x62\x0f\x29\x48\x4c\x5c\xcc\x50\x3b\xe1\xa4\xd1\x90\x1a\x0b\xb7\xe6\x01\x9c\x81\xcc\x9a\x0c\xad\x9a\x43\x91\x63\x1f\x0e\xb6\xd8\x43\x64\xdb\x24\xfd\x73\xe5\xe8\xa2\x4d\xc8\x14\x0c\xec\xb6\xe9\x54\xb7\x0d\xde\xcf\xa3\x0f\xf8\x10\x8c\xca\x32\x3a\xbd\x9b\x7a\xf4\x76\x40\x1b\x28\xcb\x5e\x23\xa6\x70\xdd\xcb\x04\x13\x0e\x61\xc1\xa7\x1d\x71\xfe\x79\xa4\x09\x48\x45\xd0\x8c\x9c\x9c\x61\xee\xc4\x2c\xfb\xec\xa5\x3e\xdf\xa2\xca\xd0\x8e\x20\x82\xca\x4b\xb7\x35\xf2\x87\x31\x77\x75\x5a\x75\xab\x29\x31\x7b\x98\x1a\x8b\x3e\xa8\x2c\xb4\x75\x69\x3d\x2e\x9e\xf6\xb4\xe4\xee\xa0\xcd\xc5\xe1\x40\xff\xe7\x00\x53\x51\x28\xc7\x83\xc8\x97\x02\xad\xc4\x3c\xfa\x60\xf4\xbf\xd1\x9a\x7a\x6b\x82\x04\x6b\x0d\xfa\x81\x79\xd0\x2d\xec\x75\xa4\x2f\xa5\xbb\xad\x85\x43\x30\x63\x52\xeb\x0b\xe3\x09\xad\x5b\xd6\x29\xeb\xe4\x00\x29\xd4\xc1\x42\xf7\x98\x10\x7d\xbd\x0e\xcf\x58\x68\x0a\x96\x87\x00\x1e\xa4\xbb\x05\x01\x8e\x27\x28\x77\x2b\x1c\xd4\xfb\x4d\xed\x50\x1d\x09\x28\x58\x33\xc4\x6c\xb7\x41\xf7\xd5\x2b\xd8\x2b\xa4\x4a\x20\x16\xf1\x2d\xc2\x1d\xce\x41\xea\x97\x4a\x6a\x84\x62\xaa\xa4\x9a\xc3\x4b\x98\xcd\xf3\x2f\x0a\xee\x73\xc8\xe8\x6f\x66\xcd\x8d\xc2\x59\x3e\x1c\xdc\x14\x29\x85\x20\x77\x76\x26\xf4\x54\x21\x35\xb9\xbd\x22\x4d\xd1\x06\x63\xde\x8d\x2e\xad\x74\x38\x61\x12\x0a\x72\x67\x63\xa3\xef\xa3\x23\x67\x44\xd0\xcb\xf3\xe8\xbd\xd4\x09\xd1\x1d\x25\xdf\xe7\x10\x62\xd2\xea\xe9\xaa\x2f\xb7\x6f\x54\xce\x21\x59\xd6\x1d\xf3\x69\xda\xe5\xbd\xb9\xc3\xe0\xf7\xe8\xf7\xa7\xdc\xe8\xd3\xc0\x7a\x37\xfa\x72\xdf\xe3\xc6\x63\x9d\x9d\xec\xfc\x09\xba\x9a\x94\xdc\xa0\x8a\xb0\xdd\xd9\x05\xda\xad\x37\xc6\xc3\x41\x0b\xde\x69\xd1\x80\x77\x53\xa4\x63\x2e\xe5\x95\x65\xe1\xcb\x76\x9f\xd2\xe5\xa4\x70\xd1\xf9\xb1\x89\xef\x48\x13\x27\x50\xe8\xf3\x28\x21\x43\x4f\xbf\x7f\x75\x87\xf3\xeb\xad\x0d\x5d\x68\xe5\x4d\x0d\x07\xd4\x07\x89\x07\xb8\x26\x7c\xf5\xfc\x52\x1b\xa6\x00\x34\xc3\xa7\x45\x47\x8e\xf4\xd1\x3b\xea\x3c\x51\x9d\x0e\x07\x83\x75\x1e\x34\x25\xfa\xb4\x48\x97\x24\xb6\x93\x36\x85\xeb\xbe\xd0\x66\x03\x3d\x8e\x87\x83\x41\xdd\x0c\x77\x76\x97\x8a\xe0\xa2\xf3\xf4\xe3\xfe\x9f\x5a\x39\x13\x76\xfe\x1e\xe7\x1d\x61\x0a\x71\xc3\x48\xde\x78\x87\x8e\x9e\xee\x2f\x85\xf6\x4c\x64\x1a\x82\x5a\xea\x36\x21\xc4\xa6\x50\x09\xf3\xfd\x0d\x93\x4f\x7d\x56\x4f\x4d\xa0\x64\xce\xdd\x87\x09\x8a\xcc\x41\x97\x64\x26\x34\x49\xcf\x32\x85\xd4\xf7\x03\x8b\x2e\x6c\xd3\x9f\x5e\xe2\x3c\x88\x88\x97\xe7\xb0\xeb\xf5\xfb\x4c\x3a\xa3\xa5\x13\x62\xe5\x20\x91\x42\x61\xec\x42\x18\x2d\xb9\x36\x6a\x5a\x70\xd3\x7b\x5b\x8d\x16\xbd\x06\xd8\x85\x74\xe6\xa2\x49\x66\xa5\x76\x29\x87\x7f\x34\x39\x3c\x3e\xdc\xff\x04\xbf\xe6\xf0\xf6\xfc\xe3\x09\x9d\xf7\xf8\xac\xaa\x96\x74\x97\x65\x74\x7e\x56\x55\x70\xf9\xc7\xe1\xf9\x21\xfc\x9a\x8f\x18\x17\x3f\xa2\xe5\xd1\x3f\x8d\xd4\x41\x7b\xca\xa3\x04\xb5\x3b\x2b\x8c\xc3\x89\x92\x31\x36\x1e\x47\xc7\x67\x21\x34\xbf\xcf\xcf\x38\xc5\xc7\x21\x8c\xc2\xd1\xb8\xd1\x56\x2b\xb8\xbc\x45\x8b\xfb\x4a\x14\x39\x32\x3e\xe4\xd0\xc8\x1f\xf8\xdc\xff\x7c\xdd\x0d\xdc\x02\x76\x7f\xd8\x7b\xa1\x0a\x3c\x11\x59\x26\xf5\x34\xe4\x52\x6b\x5b\xdd\x9e\xd4\x49\xbd\xb5\xae\x75\x7e\x9a\x67\x18\xae\x23\x80\x85\xda\x36\xc2\xf5\x78\xd0\x69\xeb\xbd\xbe\x4e\xec\xd5\xe4\x23\x1d\x98\x04\xeb\x64\x5c\x60\xf3\xdc\xce\x92\x5d\x32\xb8\xc2\xd5\xbe\xaf\xec\x6c\xe5\xbb\x2b\x87\x91\x69\x1a\x53\x86\xec\x48\x27\xd2\x62\x4c\x79\xeb\x17\xfe\x45\x12\x1f\xd3\xc0\x50\xe3\xb9\x17\xaa\x37\x54\xf0\x66\xfe\xd6\x9a\x59\x73\x04\x56\x58\x93\x6c\x0f\xa4\xb1\x27\x45\xef\x49\x0e\x57\xd7\x52\x3b\xb4\xa9\x88\xb1\xac\x16\xd3\xc5\x72\xb0\x3a\x81\x6c\x5e\x6c\x8d\x9f\x3a\xbb\xde\x74\x47\x87\x3f\xa9\x4c\xfd\x78\x7a\x80\x37\xc5\xf4\xc4\x24\xc8\x5a\xa9\x50\xde\x72\xa1\x28\x1d\xb4\xfb\xdc\x9c\x6c\xa3\x8b\x4b\x75\xfc\xb4\x34\x45\x67\x31\x93\xbe\x88\x85\x3e\x16\xb9\xf3\x6c\x7e\x74\xd0\xbd\xcf\x2c\xed\xd4\xf7\x1a\xbe\xd5\xac\xda\x1a\x2c\x8d\xf5\x7e\xd5\x62\xce\x13\x5f\x3d\xb6\xd2\xf0\xc9\x43\x7e\xd0\x71\xda\xfb\x14\x45\xd1\x98\xb5\xd0\xe4\xbf\xf9\xe5\xda\x42\xc0\x93\xed\x06\x45\xf5\xc5\xaa\xa7\x73\xb5\x9b\x9f\x9b\x84\xff\x36\x07\x1f\xbf\xf6\xed\xae\x35\x83\xf6\x8a\x92\xe8\xb7\x08\xba\xd4\xd2\x8d\xd6\xb3\xcf\xa6\x46\x41\x93\xcd\x32\x23\x2f\x20\x5f\x0b\x20\x25\xbe\xa2\xd5\x03\x90\xda\xfd\xfd\x6f\x3d\xe7\x68\xd3\x4f\xbe\x27\x22\x83\xab\xeb\xa2\x16\xa1\xf5\x86\xfe\x78\xa0\xeb\x97\xcc\x86\x9a\x59\x74\xc2\xa9\x71\x06\x78\x3e\xa9\xef\x3a\x4f\x7a\xea\xbd\x6c\x62\xef\xb3\x24\xea\x88\x25\x34\x48\xad\x0d\xe7\xa1\xb5\x93\xb9\x8e\xdf\x0a\xa9\xda\x32\x30\x8a\xbf\x94\xf2\x98\x93\xe0\xd7\xa6\x08\x4e\xdf\xe3\x7c\x71\x4b\x7e\xdd\x42\xb6\x74\xf7\xe7\xcf\x52\xdc\x74\x17\x9a\x7a\xa2\x9f\xa4\x53\x7e\x9a\xab\xd9\x71\x49\x9a\x64\x4d\xe4\xfd\xf0\xb2\x55\x05\x3c\xfa\xc5\x46\x45\xc4\xac\x55\x15\xf8\x53\xfb\x93\xd5\x38\x31\xef\xfc\xf6\xdb\xfa\x08\xff\x95\x76\x97\x77\xae\x5e\x5f\xd3\xde\x66\xaa\xbe\x1a\xb5\x61\xa9\xaa\xd1\xf5\x7a\xa8\x7a\x97\xc5\x45\x8e\x3c\x5b\x07\xe9\x4e\x29\x3f\xa1\x64\x2c\x3a\x2b\xf1\x1e\x9b\x7b\x1d\xf3\x73\xbe\xa6\x84\x80\x8e\xdb\x4b\xf7\x4d\x5d\x66\x9b\x6e\x15\xb6\x55\x35\xfe\x31\xfe\x6f\x06\xab\x2d\x5a\x40\xf7\x04\x9e\x92\x16\x05\xb7\x4c\x8c\x1d\x7a\x63\xed\xe7\xe6\x21\xe8\xdb\x7b\xac\x2e\x9a\xc4\x82\x27\x0c\x6a\x85\x5e\x7f\x97\x34\x57\xa8\x5c\xc1\x9a\xdf\xaa\xbe\x21\xd4\x9f\x90\x12\x99\xc9\x0a\xfe\x4c\x93\xf8\xbb\xc4\xe6\x9c\xa0\xd8\x75\x4b\x62\xe7\xd1\x3d\x6a\xbb\x8b\x59\x73\x01\xdc\x42\x9c\x2f\x7c\xb0\xeb\x23\xb5\xb5\x81\xc5\xc5\x6f\xb0\xe1\x0b\xd3\xe2\x9f\x25\x89\x79\x93\x3a\xb4\xdf\xf5\x75\xa9\xa6\x84\x4e\x1f\x67\xa5\x9a\x08\xb7\xfb\x95\xf3\x7f\x01\x00\x00\xff\xff\x45\x71\x40\x0b\xe3\x1a\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xaa, 0x63, 0x2a, 0x19, 0xe1, 0x48, 0x1b, 0x3f, 0x19, 0xbb, 0xeb, 0xb4, 0x32, 0xfc, 0xc3, 0xbe, 0x87, 0x1c, 0xd8, 0xb3, 0xe6, 0x80, 0xb9, 0xf5, 0xff, 0xa6, 0x4, 0x15, 0x2b, 0x10, 0x27}}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdf\x8f\xd3\x30\x0c\xc7\x9f\x93\xbf\xc2\x54\x3a\x5d\x2b\x45\x3d\xee\x15\x69\x0f\x77\x6c\x9c\x0a\x63\xbf\x07\x42\x88\x87\x6c\x71\xb6\x48\x5d\x3a\x12\x67\x30\xa1\xfd\xef\x28\x5d\xb7\xf5\x8e\x21\xf1\xc0\x4b\xeb\xc4\xf6\x37\xfe\xd8\xbe\xbb\x83\x45\x30\xa5\x9a\x6f\x3d\x3a\x1a\x07\x74\xfb\x8f\xfb\xe9\xb8\x7f\xbc\xf5\x20\x21\x1e\x3c\x49\xc2\x0d\x5a\x02\x4f\xce\xd8\x15\x04\x1f\xbf\xb4\x46\x08\x75\x62\x57\x92\x84\xad\xab\x76\x46\xa1\xca\xb9\x0e\x76\x79\x5d\x37\x55\x46\x82\x72\x66\x87\xce\xe7\x5d\x23\x4b\x5c\x92\x00\x92\x8b\x12\x07\x72\x83\x8d\xbe\x80\xb0\x55\x92\x50\xc0\x8f\xb5\x21\x2c\x8d\x27\xf8\xfa\xed\xe8\xcb\x4e\x35\xfc\xe2\xec\xe2\xed\xc4\xdb\x8d\xb4\xab\x12\xf3\x42\xa1\xa5\x71\xa8\x08\xa7\xa5\x59\x62\x7c\x32\xef\x8f\x05\xc4\xff\x64\xdc\xd2\xcc\x38\x67\x8b\xa0\xe1\x4d\x3b\xfb\x09\xe9\x31\x68\x8d\x2e\xcd\x38\x53\xa8\xd1\xb5\x9c\xa3\x70\x72\x2e\x82\x8e\xe9\x3b\xe9\x60\x59\x95\x61\x63\x7d\x53\x17\x67\x46\x43\x89\x36\xbd\x3c\x03\xaf\x3a\xf0\x3a\xd6\xcb\x4e\xa1\x9d\x26\xd8\xe7\xef\x2b\xd3\x0a\x15\x90\x08\x48\x32\xce\x0e\xfc\xac\x73\x6c\x45\x06\x9d\x93\x88\xde\x50\xfe\x6e\xeb\x8c\x25\x9d\x72\xc6\x22\x82\x88\xff\xa4\x18\x4c\x7b\x93\x19\x14\x4f\x83\xe1\xa4\x07\xc5\x60\x36\x84\x1b\x0f\xe9\x8d\xcf\xe0\xd3\x43\x7f\xde\x9b\xd6\x76\x52\x07\x9f\x5b\x5e\x9f\x9a\xba\x6a\xbb\x45\x5b\xca\x25\xae\xab\x52\xa1\xf3\x75\x17\xe7\x1e\x0b\xab\xf0\x67\xdb\x21\x5e\xc0\x0a\xb8\x17\x70\x9f\x45\xa9\x8c\x33\xe6\x90\x82\xb3\xb0\x08\x3a\x9f\xd6\xc8\x69\x43\xf7\x82\xa2\x81\x38\x33\xfc\xa5\x78\x18\x0e\xa0\x3b\x1f\xf5\x8b\xb7\x0f\xb3\x1e\x7c\xe8\x7d\x81\xf9\xa8\x1b\xcd\x9a\xea\x19\x54\x8b\xe9\xbf\x21\xc5\x91\xeb\xca\x81\x11\xb0\x8b\x6b\xe3\xa4\x5d\x61\xb3\xac\xf5\x6c\x8c\x06\x73\x19\x77\xa4\xca\x3f\x3b\x43\xf8\xb8\x27\x4c\x6f\xc5\x6d\x6c\xc9\x81\x33\xf6\x3d\xae\xa7\x7a\xbe\x79\x97\xbd\xfd\x63\x65\x77\x19\x6f\x89\x35\x8d\x3c\x6a\x5c\xf3\x24\xd0\x69\x9a\x96\x26\xff\x98\x79\x2c\x30\xbb\x6d\xa6\x73\x6d\x6c\x07\xfe\x3b\x00\x00\xff\xff\xa9\x3a\x4a\xd3\x2e\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0x5a, 0x76, 0xec, 0x69, 0x8e, 0xfe, 0x64, 0xaa, 0x43, 0x53, 0x7b, 0x3f, 0x23, 0x98, 0xb3, 0x4d, 0xd, 0x37, 0xf, 0x1b, 0x16, 0x4a, 0xd6, 0xd8, 0x1, 0xeb, 0xf7, 0x29, 0x5b, 0x7, 0xd1}}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x6f\xdb\x38\x12\x7f\x96\xfe\x8a\x39\x03\xd9\x93\xb2\x0a\x53\x60\x81\x7b\xc8\xc2\x08\x1a\xc7\x59\x04\xdb\x7c\xd4\xce\x5d\x71\x68\x8a\x96\xb1\xc6\x09\x51\x89\x54\x48\x2a\x8e\x2f\xc8\xff\x7e\x18\x52\x9f\x4e\x64\x34\x77\x7d\xdc\xa7\x56\x9c\x1f\xe7\xe3\x37\xc3\x99\x71\x1e\xb8\x06\x7d\xfb\x78\xb6\x9e\x7f\xfc\xf0\x1d\xd7\x30\x06\x8d\xb7\xf8\x58\xb0\xb3\xd2\xd8\x89\xca\x0b\x91\x61\xf4\x2d\x3a\xcc\xe3\x28\x4a\xae\x65\x7c\x78\x6d\x7e\x9d\x5c\x9c\xcf\xaf\x66\xef\x4f\xcf\xaf\xd8\xee\xe1\xc9\xc5\x6c\x7a\xfa\xc7\x39\xfc\x39\xfd\x37\xdb\x3d\xbc\x96\xf1\xaf\xdf\xe2\x30\xb4\xeb\x02\x21\x5f\x9b\xfb\xec\x0a\x8d\x45\x0d\xc6\xea\x72\x61\xe1\x29\x0c\xd2\x9b\x89\x92\x12\x76\xcd\x7d\xc6\x8e\x8f\x42\x3a\x38\xe7\x39\x06\xc6\x6a\x21\x6f\xc3\xe0\x4e\x19\xdb\x7c\x94\x06\x75\xf3\x51\x70\x63\x9a\x0f\x63\xb2\x5c\xa5\xed\xb5\x42\x69\x1b\x08\x69\xc3\x30\x50\x85\x15\x4a\x9e\x88\x0c\xa1\x92\x86\x81\x45\x63\x8f\x8f\xc8\x50\x7d\xf6\x1c\x86\xcb\x52\x2e\x40\x48\x61\xa3\xd8\x7b\x76\xc6\x85\x84\x31\xfc\xd2\xf1\xfc\xe9\xb9\x41\x46\x39\xec\x76\x24\x31\x18\xb4\x65\x11\xc5\x80\x5a\x2b\x4d\x1a\x88\x4d\xd4\xda\x1f\x84\x61\xf0\x20\x0a\xd4\x6c\x8e\xf6\x18\x97\xbc\xcc\x6c\x34\x72\xf7\x59\xe5\xfc\x28\x81\x91\xd5\x25\x8e\xe2\x61\x28\xc5\x35\x4a\xe0\xb7\xdf\xde\xfd\x23\x0e\xc3\x20\x67\x9e\x2e\x18\x83\xbf\xf1\x07\xda\xb9\x0b\xa8\xbe\x90\xde\x48\x9e\x3b\x95\x39\x23\x2e\x87\x91\x24\xf5\x38\xa2\x79\x18\x47\x52\x8f\xa3\x0c\x0c\xe3\x48\x5a\xe1\x94\xee\xd9\x3d\x95\xfd\x78\x1c\xa8\x22\x61\x58\x5f\xcd\x12\xc5\x4d\xac\x8e\xe1\x81\x67\x9c\x1d\xe1\xad\x90\xff\xe2\x99\x48\x39\xe5\x39\x8a\x59\xf5\x81\x51\x18\x04\x0e\xe2\xf5\x9c\x2b\x3b\xcd\x0b\xbb\x8e\x7c\x80\x09\xf4\xe2\x49\x06\xc1\xc4\x4b\x03\xf6\x24\x35\xe0\x73\x65\x23\xf7\x9f\xe9\x7d\xc9\x33\x13\xf9\x58\x13\x78\xd7\x5c\xf0\x01\x6e\x51\xef\x13\xd8\xe0\xeb\x7c\x0d\x5f\xa8\x78\x68\x6e\x34\xbc\x24\x61\x10\xb3\xc9\x1d\x2e\xbe\x47\xc4\x91\x58\xba\xe2\xfb\xdb\x18\xa4\xc8\xa8\x1c\x03\x8d\xb6\xd4\x92\x4e\xc3\xe0\x39\x0c\x83\xfd\x7d\x98\x68\xe4\x16\x81\x83\xe6\x32\x55\xb9\xf8\x0f\xa6\x90\xde\x00\xf9\xc0\x28\x2b\x9d\x87\x32\x6e\x31\x6c\x6e\xf9\x4d\x86\x5e\xd0\xc4\xd0\x31\x3a\x86\x9c\xe5\xfc\x3b\x5e\x34\x6f\x2f\x8a\x7f\x1f\x76\x47\x69\xc3\x3e\x69\x5e\x44\xa8\x29\x2f\x0b\x55\x66\xa9\xfc\xbb\x05\x52\x01\xfe\xfd\xc2\x52\x64\xae\x8c\x9f\xfb\x56\x52\xad\x8a\x2b\xe7\xe4\x56\x0b\x74\xaf\x7b\x6d\xe1\xe2\xfe\xc1\x8b\x61\x90\x96\x79\x31\xc9\x53\x38\x18\x03\x3e\xe2\x82\x4d\x54\x9e\x73\x99\x56\xa5\x49\xd2\x51\x42\xce\xf8\xc7\x6a\x7c\xc0\x09\x8c\xf6\xf6\xa4\xda\x4b\xb9\xe5\x5e\x5c\xd1\x14\x78\xeb\xc3\x0a\x87\x94\x91\xa6\x1b\x6e\xd0\xc9\xdb\xd4\x10\xf1\x3a\x81\x15\x69\x13\x8a\x5d\x8a\x02\xa3\xb8\x71\x7a\x6e\x53\x8a\xee\x60\x0c\xbf\xdc\xac\x2d\x1a\x76\x54\x2e\x97\xae\x93\xb5\x7e\x0c\x63\x1a\x35\x6c\x6e\x53\x55\xd2\x3b\x5e\xf5\xce\x3c\xa3\x3d\x5b\x61\x47\x33\x41\x5c\x17\x95\xb8\x3a\xf9\x13\xd7\xc7\x68\xac\x56\x6b\xd4\x51\x67\xce\x24\xa0\xe3\x8d\x3b\x5e\xed\x86\x83\xdd\xd4\xb7\x2e\x70\x6d\xdf\x52\x5d\x4b\x2e\x32\x4c\xc1\x2a\x30\x74\x15\x9a\x14\xc2\xc2\x27\xc1\x57\x59\x6b\xa9\xeb\xd7\xcf\xb0\xd5\xb7\xf3\x4a\x48\x9f\xb8\x78\xcd\xca\x32\xb7\xec\x52\x0b\x69\x33\x49\xea\xe3\x8d\xa3\x5e\x0a\xaa\xb6\x11\xc5\xf1\x8f\xb9\xb7\xe2\xc2\xc2\x52\xe9\x21\x36\xc2\xe0\x2b\x65\x9e\x4d\x32\x65\x30\x8a\x61\x7f\x1f\xde\x2f\x69\x92\xd7\x2f\x43\x18\x48\x95\xc4\x04\x16\x84\x00\x7b\x87\xb0\xd2\xc2\x22\xa0\x4c\x41\x2d\xdd\x41\x21\x0a\x0c\x5f\x25\xf6\x7f\x8b\x78\xa3\x3a\xfe\xbf\x98\x37\xe3\xad\x14\x48\x91\x6d\x99\xf8\x26\x3b\x53\x29\x46\x6e\x70\xf9\x45\x22\xae\xfe\x25\xff\xcd\x4a\xd8\xc5\x1d\x38\xe9\x53\x18\x2c\xb8\xc1\x6a\xc2\x1f\xb4\x1e\x8e\x66\xd3\x8f\xff\x3c\x9d\x4d\x8f\x47\x35\x62\xc9\x33\xd3\x87\x1c\x9f\xce\xdf\x1f\x7d\x70\x90\xaa\x2b\x74\xa5\x97\xb3\xe9\xc9\x74\xe6\x35\x6c\x59\x4f\xfa\xfd\xa4\xe3\x66\xa5\x87\x98\x9d\x17\x44\xed\x32\xa2\x5e\x53\xc1\xf7\xa8\xf1\x8e\x77\x8c\xeb\x39\xed\x2e\x15\x0f\x1b\xda\x6c\xfc\xed\x42\x64\xf3\x22\x81\xaa\xd1\x08\x55\x5a\x91\xb1\x2b\xcc\x0b\x07\x1b\xd1\xfa\xe3\xf5\xd7\xad\x7e\xdb\x04\x1b\xcc\xaa\x2f\x8a\x57\xa7\x86\xb9\x9a\x5c\x92\x69\x47\x70\x18\x7c\x4d\xaa\x3a\x54\x86\x5e\xb6\xad\x66\xbd\x37\xac\x0c\x3b\x35\x34\x75\x1f\x85\xb1\xae\xf8\x9c\x03\x5e\xc7\x18\x28\x8b\x61\xf0\x0c\x98\x19\x84\x37\xf8\xe9\x46\x1b\x48\x65\xa9\x21\x58\xc8\x9b\xad\x8b\x1c\xa4\x0c\x9c\x14\x55\x71\x3b\xae\x46\x9f\x17\x99\x40\x69\xbf\x10\xa4\x15\x2f\x2b\x29\x5d\x1e\xef\x98\x6b\xe9\x92\x53\x39\xff\x12\x46\x3b\xc8\x78\x27\xad\x60\xf4\xf5\x2a\x8c\x16\xa1\x56\x1b\x7d\x79\x2a\x32\x94\x91\x5f\xf6\x62\x8a\xf1\x5d\xf3\x34\x37\xac\x70\x63\x56\x4a\xa7\xad\x0a\x77\xc5\xf5\xd2\x97\x68\x63\xb2\x3d\x7a\x18\x2d\xba\x79\x4c\xf5\x46\x13\x7b\xf3\x9e\xf2\xbe\xcd\x86\x9f\x42\x2b\xab\x16\x2a\x1b\xdb\x45\xb1\x8d\xc6\xa6\xad\xfd\xc5\xe4\x1b\x98\xec\x3e\x78\x2a\xfa\xbc\x60\x6e\xe3\x8b\xdb\xfe\x48\x67\xd5\x50\x18\xee\x08\xfd\x6d\xab\xed\x07\xd4\x76\xe9\x3d\x76\x3b\x4f\xf5\x7e\xeb\x55\x07\x76\xcc\xef\x2f\xd6\x9d\xda\x78\xce\x74\x29\x27\x79\x1a\x99\xfb\xac\xde\x87\x47\x5b\xfc\xe8\x2e\x8b\xdb\xbd\x20\x64\xeb\x03\x3d\x70\xea\x03\xe6\xa7\x7a\x63\x91\xeb\x54\xad\x64\xd7\x17\xb1\x74\x7b\xa2\xfb\x4d\xfc\xb2\x9f\xd4\xa2\x86\xf1\xee\xfe\x70\xf0\xc6\x75\xb8\xf1\x5b\x19\x36\xc3\x5c\x3d\x50\xc1\xfc\x50\x83\xaf\xc3\xa4\xcd\x2e\xa9\xe7\x66\x35\x50\x12\xe0\xfa\xd6\x00\x63\xac\x9e\x87\x4d\x6c\x4e\x30\x06\x5e\x14\x28\xd3\xe8\xf3\x17\x0f\x78\xda\xdc\x74\x9f\xbd\x0a\xc6\x18\x95\xd9\xe2\x95\x25\xb9\xb2\xd8\xc1\x11\xac\xd9\x34\xbd\x5e\xc3\xce\x71\x35\x43\x9e\xa2\xf6\x9e\x92\x36\xe3\x77\xd8\xd7\x76\x61\xb3\x65\x4d\xee\x6e\xbf\x5e\x45\x73\xe8\x27\x88\xbf\xdc\x4d\x05\x89\x67\xa5\x7c\x99\x85\xee\x02\x53\x8f\x2d\x5d\x4a\x29\xe4\xed\xc1\xa8\x61\xd3\xc7\x16\xf7\xe1\xde\x74\x77\xcd\xd9\x90\x6e\x2c\x41\x03\xf9\xde\xba\xd2\x2c\x94\xa4\x82\x8c\xaa\x3f\xc8\x24\x3e\x7d\xf1\x70\x6d\x6e\x94\x66\xe2\xd4\x3b\x73\xfd\x3f\x7f\x04\x2d\xa2\xe2\xec\x3e\x63\x17\x05\xca\xf6\x37\x4f\xaa\xc5\x03\x6a\xe6\x7e\x15\x1c\x95\x22\x4b\x3f\x96\xa8\xd7\x55\x40\xf5\xaf\x76\xdf\x0c\xfb\x6f\xb0\xee\xd9\x75\x53\xae\x9a\x60\xa7\xf5\xf5\x73\xd0\x12\x91\xbc\x60\xa7\x1f\xc8\x73\xf8\xdf\x00\x00\x00\xff\xff\xc3\xef\x17\xe0\x13\x13\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x29, 0x25, 0x60, 0x54, 0x1d, 0x80, 0xdd, 0x7d, 0xff, 0xe9, 0xfe, 0xa4, 0x6b, 0xdd, 0xf6, 0xe, 0xbf, 0x41, 0xcf, 0x43, 0x85, 0x4f, 0xbf, 0x35, 0xb2, 0x93, 0x40, 0xc6, 0xe1, 0x2, 0xc1}}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x69\xaf\x29\x7c\xc8\xdf\x21\x68\x6b\xb8\xb1\x7c\x2e\x18\x69\xe5\x10\xa6\x49\x95\x5c\xd5\x76\x05\xbe\x7b\x41\xc9\x76\x9c\xd8\x69\x73\x68\x0f\x39\xf8\x87\x8b\xd9\x9d\x99\x5d\x2e\xdb\xf6\x04\xfe\x97\x5a\x49\x0f\x67\x43\x10\xe7\xf1\x1f\x7a\x91\xcb\x3b\x8d\xd0\xff\x88\x91\x5c\x60\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xd3\x7a\xac\x1b\x27\x75\x08\xd3\xda\xa3\x23\x4e\xf0\x2e\x02\x94\x99\x89\x3c\x85\x96\x25\x24\xc6\xd2\x49\xad\x51\xf3\x94\xb1\x44\x55\xa0\xd1\xf0\x5d\x81\x2b\xbb\x34\x13\x65\x66\x8d\x96\x2e\x84\x4b\xab\x9b\x85\xf1\x29\x0c\x87\xbf\x83\x8d\x9d\x5a\x48\xb7\xfe\x84\xeb\x5d\x42\xcb\x92\x84\xc4\x64\xae\x6a\x3e\x88\xdf\xb5\x32\x33\xa0\xce\xc3\x52\xd1\x3d\x58\xa3\xd7\x50\xf7\x79\x30\xc7\x35\x14\x7d\xe6\x20\x65\x49\xd8\xc9\x5a\xac\x27\x5f\x3f\xef\x99\x7b\xa0\x9c\x1a\xf5\xbd\xc1\x7d\x7d\xef\xff\xc8\x69\x2c\x34\x5d\xda\x96\x0c\xc8\x42\x61\x4d\xa5\x55\x41\x60\x4d\xcf\xcd\x12\x8f\x58\xc6\xde\x3b\x69\x4a\xbb\x50\x3f\x51\x8c\x70\x39\x41\x2c\x79\xca\x92\x1f\xd2\x01\xba\xee\x63\x1d\x4b\x4e\x4f\xe1\x9c\x08\x17\x35\x01\xdd\x23\xdc\x8c\x26\xd7\xb7\x39\x78\x55\x22\xd8\x0a\xa4\x81\xe9\x38\x46\x58\x62\x63\xc5\xa3\x56\xda\xde\x6f\x2c\xba\xcf\x39\x21\xd7\x14\xc4\xa3\x98\x0c\xde\xda\x0c\x9e\x69\xfe\xd5\x45\xbe\xae\xd1\x67\x50\x49\xed\x31\xfd\xd8\x15\xfa\x6f\x08\x46\xe9\x4d\x47\xae\xa3\xd4\x8a\x0f\xa6\xa6\xeb\x05\xd9\x07\x96\xe3\x8a\xc0\x77\xdc\x67\xf0\xc6\x0f\xb2\x58\x6f\xd3\x98\xb6\x55\x15\x18\x4b\x20\x46\xf6\xd2\x1a\xc2\x15\x85\x50\xd0\x2a\x5a\x2b\xfa\xb3\xb8\x90\xc5\x7c\xe6\x6c\x63\x4a\x9e\xb6\x2d\x9a\x32\x04\x96\xf4\x90\x2f\x8d\xa7\x7c\xc5\xbb\x2a\xfb\x15\x0e\x02\x77\x56\x69\x71\x81\x33\x65\xba\x1a\xda\xe3\x7e\x2c\x5f\xf1\x82\x56\x59\x34\xb8\x65\x78\x11\x28\x65\x49\x89\x15\x3a\x88\x6b\xc3\x53\x68\xe1\x1b\x0c\x81\x56\xe2\xd6\x6a\x7d\x27\x8b\x39\x4f\x21\xc4\x11\xef\x86\x61\xc5\x66\x8b\x9e\x33\x1e\x87\x82\xa6\x84\x93\x10\x20\x9e\x3a\xfe\x1b\x53\xa1\xe3\xe9\xe3\xd3\xcb\xe6\xd2\x74\x74\xc7\x87\x72\x30\x8d\xc2\x36\x86\xba\xc0\x93\xab\xb5\x7d\x02\x78\x2a\x2e\x23\xe6\x85\xf2\x1f\x9c\x1f\xaa\xe4\x5b\xda\x08\xe9\x88\x23\xe8\xc3\x23\xc8\x60\x29\x4d\x5c\x23\x04\x87\x85\x75\x65\x06\x33\x4b\x67\x83\xac\xc7\x6f\x44\x3f\xd9\x97\xe9\xf8\xea\x3c\xbf\x3e\xb6\x2f\x7f\x6d\x23\x9e\x85\x1d\xbc\x5a\x42\x88\x7f\xba\x3e\xaf\xef\x5e\xbd\x92\x6b\x15\xd8\xaf\x00\x00\x00\xff\xff\x05\xd5\xa1\x1c\x35\x07\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x92, 0xf5, 0x7f, 0x80, 0xcc, 0x1f, 0xaf, 0x82, 0xb2, 0x12, 0x19, 0x6, 0xd4, 0x69, 0xc3, 0x57, 0xe4, 0x53, 0x13, 0x49, 0x4c, 0x87, 0xe7, 0xd9, 0xe0, 0xe1, 0xe6, 0xf5, 0x6c, 0x40, 0x20}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/mysql_upsert.go.tpl": templatesSingletonMysql_upsertGoTpl,

	"templates_test/singleton/mysql_main_test.go.tpl": templates_testSingletonMysql_main_testGoTpl,

	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl":   &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}