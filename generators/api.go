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

// Config ...
type Config struct {
	DBDialect      string
	ProjectPath    string
	ProjectName    string
	ConfigFilePath string
	AWS            bool
}

// EvaluateFileContent ...
func EvaluateFileContent(filePath string, config Config) (string, error) {
	templateBox, err := rice.FindBox("./templates/api")
	if err != nil {
		return "", errors.WithStack(err)
	}

	tmpContent, err := templateBox.String(filePath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	evaluatedContent, err := templateutil.EvaluateTemplateStringToString(tmpContent, nil, template.FuncMap{
		"DatabaseRequired": func() bool { return config.DBDialect == "postgres" },
		"ProjectPath":      func() string { return config.ProjectPath },
		"ProjectName":      func() string { return config.ProjectName },
		"AWS":              func() bool { return config.AWS },
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return evaluatedContent, nil
}

// GenerateAPI ...
func GenerateAPI(config Config) error {
	box := embedded.EmbeddedBoxes["./templates/api"]
	for _, file := range box.Files {
		if ext := filepath.Ext(file.Filename); ext == ".gotemplate" {
			filename := strings.TrimSuffix(file.Filename, ext)
			content, err := EvaluateFileContent(file.Filename, config)
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
