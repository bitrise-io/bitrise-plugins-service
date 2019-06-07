package generators_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bitrise-io/bitrise-plugins-service/generators"
	"github.com/stretchr/testify/require"
)

func Test_EvaluateFileContent(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	projectRoot := strings.TrimSuffix(currentDir, "/generators")

	t.Run("ok when generating with database", func(t *testing.T) {
		generatedContent, err := generators.EvaluateFileContent("main.go.gotemplate", "github.com/my-github-account/my-project-name", true)
		require.NoError(t, err)
		require.Equal(t, getTestData(t), generatedContent)
	})

	t.Run("when no file exists with the given name", func(t *testing.T) {
		_, err := generators.EvaluateFileContent("non-existing.file", "github.com/my-github-account/my-project-name", true)
		filePath := filepath.Join(projectRoot, "templates/api/non-existing.file")
		require.EqualError(t, err, fmt.Sprintf("open %s: no such file or directory", filePath))
	})
}
