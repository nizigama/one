package rendering

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type TemplateFileReader interface {
	Read(string, string) ([]byte, error)
}

type Template struct {
	Name        string
	Data        interface{}
	FileReader  TemplateFileReader
	basePath    string
	content     *bytes.Buffer
	status      int
	destination http.ResponseWriter
}

type Reader struct {
}

type TemplateFileReaderMock struct {
	mock.Mock
}

const basePath = "resources/views"

func (o *TemplateFileReaderMock) Read(fileName, basePath string) ([]byte, error) {
	args := o.Called(fileName, basePath)
	return args.Get(0).([]byte), args.Error(1)
}

func (r *Reader) Read(fileName, basePath string) ([]byte, error) {
	paths := strings.Split(fileName, ",")
	path, _ := os.Getwd()

	path = fmt.Sprintf("%s/%s", path, basePath)

	for idx, step := range paths {
		if idx == len(paths)-1 {
			path = fmt.Sprintf("%s/%s.tmpl", path, step)
			continue
		}
		path = fmt.Sprintf("%s/%s/", path, step)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(f)
}

func (v *Template) Parse() {

	if v.content == nil {
		v.content = &bytes.Buffer{}
	}

	templateData, err := v.FileReader.Read(v.Name, v.basePath)
	if err != nil {
		v.content.Reset()
		v.content.WriteString(err.Error())
		v.status = http.StatusInternalServerError
	}

	_, err = v.content.Write(templateData)
	if err != nil {
		v.content.Reset()
		v.content.WriteString(err.Error())
		v.status = http.StatusInternalServerError
	}

	tpl, err := template.New(v.Name).Parse(v.content.String())
	if err != nil {
		v.content.Reset()
		if errors.Is(err, os.ErrNotExist) {
			v.content.WriteString(err.Error())
			v.status = http.StatusNotFound
		} else {
			v.content.WriteString(err.Error())
			v.status = http.StatusInternalServerError
		}

		return
	}

	v.content.Reset()
	err = tpl.Execute(v.content, v.Data)
	if err != nil {
		v.content.Reset()
		v.content.WriteString(err.Error())
		v.status = http.StatusInternalServerError
		return
	}

	v.status = http.StatusOK
}

func (v *Template) SetWriter(w http.ResponseWriter) *Template {

	v.destination = w

	return v
}

func (v *Template) Write() (int, error) {

	return v.destination.Write(v.content.Bytes())
}

func (v *Template) SetStatus(code int) {
	v.status = code
}

func (v *Template) GetContent() []byte {
	return v.content.Bytes()
}

func (v *Template) SetBasePath(basePath string) {
	v.basePath = basePath
}

func (v *Template) WriteStatus() *Template {
	v.destination.WriteHeader(v.status)
	return v
}

func (v *Template) SetHeaders(headers map[string]string) *Template {
	for key, value := range headers {
		v.destination.Header().Set(key, value)
	}

	return v
}

func View(viewTemplate string, data interface{}) *Template {

	view := &Template{
		Name:       viewTemplate,
		Data:       data,
		FileReader: &Reader{},
		basePath:   basePath,
		content:    &bytes.Buffer{},
	}

	view.Parse()

	return view
}
