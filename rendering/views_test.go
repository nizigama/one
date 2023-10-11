package rendering_test

import (
	_ "embed"
	"github.com/nizigama/one/rendering"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http/httptest"
)

const templateData = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>One & only</title>
</head>
<body>
    <h1 style="text-align: center; color: #303030">Welcome to the One & only framework</h1>
    <h4 style="text-align: center; color: #303030">Served on <a style="color:green" href="http://{{ .Host }}">http://{{ .Host }}</a></h4>
</body>
</html>`

var _ = Describe("Views:", func() {

	var vTpl *rendering.Template
	var data map[string]string
	var fileReaderMock *rendering.TemplateFileReaderMock
	var basePath string
	var templateFileData []byte

	BeforeEach(func() {
		data = map[string]string{
			"name": "One & only",
		}

		fileReaderMock = &rendering.TemplateFileReaderMock{}

		vTpl = &rendering.Template{
			Name:       "testing",
			Data:       data,
			FileReader: fileReaderMock,
		}

		basePath = ""
		templateFileData = []byte("Hello world, from {{ .name }}")

		vTpl.SetBasePath(basePath)
	})

	It("Can parse a template a add data into it", func() {
		fileReaderMock.On("Read", vTpl.Name, basePath, "").Return(templateFileData, nil)

		vTpl.Parse()
		Expect(string(vTpl.GetContent())).Should(Equal("Hello world, from One &amp; only"))
	})

	It("Can set the destination of a template to the http response writer and write the content of the template to the writer", func() {
		w := httptest.NewRecorder()

		vTpl.SetWriter(w)

		fileReaderMock.On("Read", vTpl.Name, basePath, "").Return([]byte("test data"), nil)

		vTpl.Parse()

		vTpl.Write()

		data, err := io.ReadAll(w.Result().Body)

		Expect(err).To(BeNil())
		Expect(data).To(Equal([]byte("test data")))
	})

	It("Can set the status of a template write it to the writer", func() {
		w := httptest.NewRecorder()

		statusCode := 201

		vTpl.SetWriter(w)

		fileReaderMock.On("Read", vTpl.Name, basePath, "").Return([]byte("test data"), nil)

		vTpl.SetStatus(statusCode)
		vTpl.WriteStatus()

		Expect(w.Result().StatusCode).To(Equal(statusCode))
	})

	It("Can read the content of a template", func() {

		fileReaderMock.On("Read", vTpl.Name, basePath, "").Return([]byte("test data"), nil)

		vTpl.Parse()

		Expect(vTpl.GetContent()).To(Equal([]byte("test data")))
	})

	It("Can set the base path of a template", func() {
		vTpl.SetBasePath("path-here")
		Expect(vTpl.GetBasePath()).Should(Equal("path-here"))
	})

	It("Can set the headers in a template", func() {
		w := httptest.NewRecorder()

		vTpl.SetWriter(w)

		fileReaderMock.On("Read", vTpl.Name, basePath, "").Return([]byte("test data"), nil)

		vTpl.Parse()

		headers := map[string]string{
			"name":    "one",
			"surname": "only",
		}

		vTpl.SetHeaders(headers)

		Expect(w.Result().Header["Name"][0]).To(Equal("one"))
		Expect(w.Result().Header["Surname"][0]).To(Equal("only"))
	})

	It("Can read a template's file data", func() {
		reader := &rendering.Reader{}

		data, err := reader.Read("welcomeTmpl", "../init/embed", "")

		Expect(err).To(BeNil())
		Expect(data).NotTo(Equal(templateData))
	})
})
