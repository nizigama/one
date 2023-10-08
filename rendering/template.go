package rendering

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nizigama/one/types"
	"html/template"
	stdHttp "net/http"
	"os"
	"strings"
)

type Template struct {
	Path      string
	content   *bytes.Buffer
	CsrfToken string
	Data      interface{}
}

const basePath = "resources/views"

func getViewFilePath(template string) string {
	paths := strings.Split(template, ",")
	path, _ := os.Getwd()

	path = fmt.Sprintf("%s/%s", path, basePath)

	for idx, step := range paths {
		if idx == len(paths)-1 {
			path = fmt.Sprintf("%s/%s.tmpl", path, step)
			continue
		}
		path = fmt.Sprintf("%s/%s/", path, step)
	}

	return path
}

func View(viewTemplate string, data interface{}) *types.Response {

	view := Template{
		Path:      viewTemplate,
		content:   &bytes.Buffer{},
		CsrfToken: "",
		Data:      data,
	}

	response := &types.Response{
		Content: "",
		Status:  stdHttp.StatusOK,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}

	filePath := getViewFilePath(view.Path)

	tpl, err := template.ParseFiles(filePath)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			response.Content = err.Error()
			response.Status = stdHttp.StatusNotFound
		} else {

			response.Content = err.Error()
			response.Status = stdHttp.StatusInternalServerError
		}

		return response
	}

	err = tpl.Execute(view.content, data)
	if err != nil {
		response.Content = err.Error()
		response.Status = stdHttp.StatusInternalServerError
		return response
	}

	response.Status = stdHttp.StatusOK
	response.Content = view.content.String()

	return response
}
