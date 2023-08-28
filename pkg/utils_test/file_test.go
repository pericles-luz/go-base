package utils_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestGetBaseDirectory(t *testing.T) {
	baseDirectory := utils.GetBaseDirectory("config")
	require.NotEmpty(t, baseDirectory)
	require.Contains(t, baseDirectory, "/config")
}

func TestFileStartFromBase64MustDetectTheExactPoint(t *testing.T) {
	start, err := utils.FileStartFromBase64(bytes.NewReader([]byte("data:image/png;base64,AAAFBfj42Pj4")))
	require.NoError(t, err)
	require.Equal(t, 22, start)
	require.Equal(t, "AAAFBfj42Pj", string([]byte("data:image/png;base64,AAAFBfj42Pj4")[start:start+11]))
}

func TestFileStartFromBase64MustReturnErrorWhenNotFound(t *testing.T) {
	_, err := utils.FileStartFromBase64(bytes.NewReader([]byte("data:image/png;base64")))
	require.Error(t, err)
}

func TestFileStartFromBase64MustReturnErrorWhenNotFound2(t *testing.T) {
	_, err := utils.FileStartFromBase64(bytes.NewReader([]byte("AAAFBfj42Pj4")))
	require.Error(t, err)
}

func TestMimeTypeFromBase64MustIdentifyCorrectly(t *testing.T) {
	mimeType, err := utils.MimeTypeFromBase64(bytes.NewReader([]byte("data:image/png;base64,AAAFBfj42Pj4")))
	require.NoError(t, err)
	require.Equal(t, "image/png", mimeType)
}

func TestMimeTypeFromBase64MustReturnErrorWhenNotFound(t *testing.T) {
	_, err := utils.MimeTypeFromBase64(bytes.NewReader([]byte("data:image/png;base64")))
	require.Error(t, err)
}

func TestMimeTypeFromBase64MustReturnErrorWhenNotFound2(t *testing.T) {
	_, err := utils.MimeTypeFromBase64(bytes.NewReader([]byte("AAAFBfj42Pj4")))
	require.Error(t, err)
}

func TestExtractFileFromBase64(t *testing.T) {
	result, err := utils.ExtractFileFromBase64(bytes.NewReader([]byte("data:image/png;base64,AAAFBfj42Pj4")))
	require.NoError(t, err)
	require.NotNil(t, result)
	t.Log(result.Name())
	t.Log(result.Stat())
	require.Contains(t, result.Name(), ".png")
}

func TestExtractFileFromBase64MustExtractTxtFile(t *testing.T) {
	result, err := utils.ExtractFileFromBase64(bytes.NewReader([]byte("data:text/plain;base64,TWFuIGlzIGRp")))
	require.NoError(t, err)
	require.NotNil(t, result)
	t.Log(result.Name())
	t.Log(result.Stat())
	require.Contains(t, result.Name(), ".plain")
	// show the file content
	content, err := utils.ReadFile(result.Name())
	require.NoError(t, err)
	require.Equal(t, "Man is di", string(content))
}

func TestFileMustBeConvertedToBsse64String(t *testing.T) {
	filePath := utils.GetBaseDirectory("csv") + "/file.xml"
	result, err := utils.FileToBase64(filePath)
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(result, "data:text/xml;base64,"))
}
