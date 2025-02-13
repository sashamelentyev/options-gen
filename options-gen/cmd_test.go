package optionsgen_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	optionsgen "github.com/kazhuravlev/options-gen/options-gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	t.Run("simple_cases", func(t *testing.T) {
		t.Parallel()

		var testDirs []string
		{
			const testdataDir = "./testdata"
			fileInfos, err := ioutil.ReadDir(testdataDir)
			require.NoError(t, err)

			for _, file := range fileInfos {
				if !file.IsDir() {
					continue
				}

				testDirs = append(testDirs, filepath.Join(testdataDir, file.Name()))
			}
		}

		for _, dir := range testDirs {
			dir := dir
			t.Run(dir, func(t *testing.T) {
				outFilename := filepath.Join(dir, "options_generated.go")
				expFilename := filepath.Join(dir, "options_generated.go.expected")
				err := optionsgen.Run(
					filepath.Join(dir, "options.go"),
					outFilename,
					"Options",
					"testcase",
				)
				assert.NoError(t, err)

				helpEqualFiles(t, expFilename, outFilename)
			})
		}
	})

	t.Run("soource_not_found", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()
		err := optionsgen.Run(
			filepath.Join(dir, "options.go"),
			filepath.Join(dir, "options_generated.go"),
			"Options",
			"testcase",
		)
		assert.Error(t, err)
	})
}

func helpEqualFiles(t *testing.T, filename1, filename2 string) {
	t.Helper()

	f1Bytes, err := ioutil.ReadFile(filename1)
	require.NoError(t, err)

	f2Bytes, err := ioutil.ReadFile(filename2)
	require.NoError(t, err)

	assert.Equal(t, string(f1Bytes), string(f2Bytes))
}
