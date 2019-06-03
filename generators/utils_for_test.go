package generators_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestData(t *testing.T) string {
	testData, err := ioutil.ReadFile(filepath.Join("../testdata", filepath.FromSlash(t.Name()+".golden")))
	require.NoError(t, err)
	return string(testData)
}
