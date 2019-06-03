package generators

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/templateutil"
	"github.com/pkg/errors"
)

// EvaluateFileContent ...
func EvaluateFileContent(filePath, projectPath string, withDb bool) (string, error) {
	templateBox, err := rice.FindBox("../templates/api")
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
	err := filepath.Walk("../templates/api", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}

		if ext := filepath.Ext(path); ext == ".gotemplate" {
			newPath := strings.TrimPrefix(path, "../templates/api/")
			filename := strings.TrimSuffix(newPath, ext)

			content, err := EvaluateFileContent(newPath, projectPath, withDb)
			if err != nil {
				return errors.WithStack(err)
			}
			dirToCreate := filepath.Dir(filename)
			err = os.MkdirAll(dirToCreate, os.ModePerm)
			if err != nil {
				return errors.WithStack(err)
			}
			if err := fileutil.WriteStringToFile(filename, content); err != nil {
				return errors.Wrapf(err, "Failed to write evaluated template into file (%s)", filename)
			}
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
