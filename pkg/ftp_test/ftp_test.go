package ftp_test

import (
	"bytes"
	"testing"

	"github.com/pericles-luz/go-base/pkg/ftp"
	"github.com/stretchr/testify/require"
)

func TestClientConnection(t *testing.T) {
	t.Skip("use only if necessary")
	config := ftp.NewConfig()
	err := config.Load("ftp.discadora")
	require.NoError(t, err)
	client := ftp.NewClient(config)
	err = client.Connect()
	require.NoError(t, err)
	client.Disconnect()
}

func TestClientReadFiles(t *testing.T) {
	t.Skip("use only if necessary")
	config := ftp.NewConfig()
	err := config.Load("ftp.discadora")
	require.NoError(t, err)
	client := ftp.NewClient(config)
	err = client.Connect()
	require.NoError(t, err)
	files, err := client.ReadFiles()
	require.NoError(t, err)
	require.NotEmpty(t, files)
	client.Disconnect()
}

func TestClientRetrieveToBuffer(t *testing.T) {
	// t.Skip("use only if necessary")
	config := ftp.NewConfig()
	err := config.Load("ftp.discadora")
	require.NoError(t, err)
	client := ftp.NewClient(config)
	err = client.Connect()
	require.NoError(t, err)
	dst := bytes.NewBuffer(nil)
	err = client.Retrieve("php.ini", dst)
	require.NoError(t, err)
	require.NotEmpty(t, dst)
	client.Disconnect()
}

func TestClientRetrieveToBufferWithoutConnection(t *testing.T) {
	// t.Skip("use only if necessary")
	config := ftp.NewConfig()
	err := config.Load("ftp.discadora")
	require.NoError(t, err)
	client := ftp.NewClient(config)
	dst := bytes.NewBuffer(nil)
	err = client.Retrieve("php.ini", dst)
	require.NoError(t, err)
	require.NotEmpty(t, dst)
	client.Disconnect()
}
