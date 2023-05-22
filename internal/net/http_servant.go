package net

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	once sync.Once
)

type HttpServant struct {
	Id         int
	StaticRoot string
	bufferPool sync.Pool
	mime       sync.Map
}



// Only GET Roughly Implemented
func (h *HttpServant) OnRequest(conn net.Conn, msg []byte) []byte {
	once.Do(func() {
		h.bufferPool = sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		}
		h.mime = sync.Map{}
		L("init pool/cache")
	})
	var header = h.bufferPool.Get().(*bytes.Buffer)
	header.Reset()
	defer h.bufferPool.Put(header)

	request_file,err := h.getRequestingFile(msg)
  if err!=nil{
    L("%v, %v", string(msg),err)
		header.WriteString("HTTP/1.1 404\n\n")
		return header.Bytes()
  }

	f, err := os.Open(h.StaticRoot + "/" + request_file)
	defer f.Close()

	if err != nil {
		header.WriteString("HTTP/1.1 404\n\n")
    L("Open File Error:%v, %v", request_file, err)
		return header.Bytes()
	}
	fi, _ := f.Stat()
	contentType, ok := h.mime.Load(filepath.Ext(fi.Name()))

	if !ok {
		contentType, err = h.getFileContentType(request_file)
		if err == nil {
			h.mime.Store(filepath.Ext(fi.Name()), contentType)
		}

	}
	fmt.Fprintf(header, "HTTP/1.1 200 OK\nContent-Type: %s\nContent-Length: %d\r\n\r\n", contentType, fi.Size())
//	L("Write Back")
	io.Copy(bufio.NewWriter(header), bufio.NewReader(f))
	//io.Copy(conn, bufio.NewReader(f))
	return header.Bytes()
}
var get = regexp.MustCompile(`^GET\s\/(?P<route>.*)\sHTTP`)
var post = regexp.MustCompile(`^POST\s\/(?P<route>.*)\sHTTP`)

//Only GET
func (h *HttpServant) getRequestingFile(buf []byte) (string, error) {
	m := get.FindStringSubmatch(string(buf))
	if m == nil {
		return  "", errors.New("Skip, Not a GET method")
		//panic("no match")
	}
	result := make(map[string]string)
	for i, name := range get.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = m[i]
		}
	}
	request_file := result["route"]
	if filepath.Ext(request_file) == "" {
		request_file = request_file + "/index.html"
	}
	return request_file, nil
}

func (h *HttpServant) getFileContentType(request_file string) (string, error) {
	var contentType = ""

	switch ext := filepath.Ext(request_file); ext {
	case ".html":
		contentType = "text/html"
	case ".json":
		contentType = "application/json"
	case ".js":
		contentType = "application/javascript"
	case ".css":
		contentType = "text/css"
	}
	if contentType == "" {
		f, err := os.Open(h.StaticRoot + "/" + request_file)
		defer f.Close()
		if err != nil {
			return "", err
		}
		buf := make([]byte, 512)
		f.Read(buf)
		if err != nil {
			return "", err
		}
		contentType = http.DetectContentType(buf)

		f.Seek(0, 0)
	}
	return contentType, nil
}
