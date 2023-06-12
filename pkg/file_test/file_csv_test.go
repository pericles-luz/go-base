package file_test

import (
	"io"
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/file"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestFileCSV(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	file := file.NewFileCSV(utils.GetBaseDirectory("csv") + "/file.csv")
	err := file.Open()
	require.NoError(t, err)
	defer file.Close()
	for {
		data, err := file.ReadLine()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		t.Log(data)
	}
}
