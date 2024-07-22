package zipper_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-base/pkg/zipper"
	"github.com/stretchr/testify/require"
)

func TestZipperMustZipFiles(t *testing.T) {
	z := zipper.NewZipper()
	directory := utils.GetBaseDirectory("csv")
	z.AddFile(directory+"/file.csv", "test/file.csv")
	z.AddFile(directory+"/file.xml", "test/file.xml")
	z.SetOutput(directory + "/output.zip")
	require.NoError(t, z.Zip())
	require.FileExists(t, directory+"/output.zip")
	require.NoError(t, os.Remove(directory+"/output.zip"))
}

func TestZipperMustReturnErrorWhenOutputIsNotSet(t *testing.T) {
	z := zipper.NewZipper()
	directory := utils.GetBaseDirectory("csv")
	z.AddFile(directory+"/file.csv", "test/file.csv")
	z.AddFile(directory+"/file.xml", "test/file.xml")
	require.EqualError(t, z.Zip(), zipper.ErrOutputNotSet.Error())
}

func TestZipperMustReturnErrorWhenNoFilesAreAdded(t *testing.T) {
	z := zipper.NewZipper()
	directory := utils.GetBaseDirectory("csv")
	z.SetOutput(directory + "/output.zip")
	require.EqualError(t, z.Zip(), zipper.ErrNoFilesToZip.Error())
}

func TestZipperMustReturnErrorWhenFileDoesNotExist(t *testing.T) {
	z := zipper.NewZipper()
	directory := utils.GetBaseDirectory("csv")
	z.AddFile(directory+"/fileerror.csv", "test/file.csv")
	z.SetOutput(directory + "/output.zip")
	err := z.Zip()
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file or directory")
}
