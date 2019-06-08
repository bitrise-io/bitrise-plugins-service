package generators

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/GeertJohan/go.rice/embedded"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/templateutil"
	"github.com/pkg/errors"
)

// EvaluateFileContent ...
func EvaluateFileContent(filePath, projectPath string, withDb bool) (string, error) {
	templateBox, err := rice.FindBox("./templates/api")
	if err != nil {
		return "", errors.WithStack(err)
	}

	tmpContent, err := templateBox.String(filePath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	evaluatedContent, err := templateutil.EvaluateTemplateStringToString(tmpContent, nil, template.FuncMap{
		"DatabaseRequired": func() bool { return withDb },
		"ProjectPath":      func() string { return projectPath },
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return evaluatedContent, nil
}

// GenerateAPI ...
func GenerateAPI(projectPath string, withDb bool) error {
	box := embedded.EmbeddedBoxes["./templates/api"]
	for _, file := range box.Files {
		if ext := filepath.Ext(file.Filename); ext == ".gotemplate" {
			filename := strings.TrimSuffix(file.Filename, ext)
			content, err := EvaluateFileContent(file.Filename, projectPath, withDb)
			if err != nil {
				return errors.WithStack(err)
			}
			dirToCreate := filepath.Dir(filename)
			err = os.MkdirAll(dirToCreate, os.ModePerm)
			if err != nil {
				return errors.WithStack(err)
			}
			if err := fileutil.WriteStringToFile(filename, content); err != nil {
				return errors.Wrapf(err, "Failed to write evaluated template into file (%s)", file.Filename)
			}
		}
	}
	return nil
}
