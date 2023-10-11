package rendering_test

import (
	"github.com/nizigama/one/rendering"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Views", func() {

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
		fileReaderMock.On("Read", vTpl.Name, basePath).Return(templateFileData, nil)

		vTpl.Parse()
		Expect(string(vTpl.GetContent())).Should(Equal("Hello world, from One &amp; only"))
	})
})
