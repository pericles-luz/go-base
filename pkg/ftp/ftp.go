package ftp

import (
	"errors"
	"io"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/secsy/goftp"
)

type Client struct {
	Config *Config
	engine *goftp.Client
}

func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) GetConfig() map[string]interface{} {
	return c.Config.GetConfig()
}

func (c *Client) Connect() error {
	if c.engine != nil {
		return nil
	}
	config := goftp.Config{
		User:     c.Config.Username,
		Password: c.Config.Password,
	}
	engine, err := goftp.DialConfig(config, c.Config.URL)
	if utils.ManageError(err) != nil {
		panic(err)
	}
	c.engine = engine
	return nil
}

func (c *Client) Disconnect() {
	c.engine.Close()
	c.engine = nil
}

func (c *Client) ReadFiles(dirname string) ([]string, error) {
	if c.engine == nil {
		err := c.Connect()
		if err != nil {
			return []string{}, err
		}
	}
	files, err := c.engine.ReadDir(dirname)
	if utils.ManageError(err) != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, nil
}

func (c *Client) Retrieve(path string, dst io.Writer) error {
	if c.engine == nil {
		err := c.Connect()
		if err != nil {
			return err
		}
	}
	return c.engine.Retrieve(path, dst)
}

func (c *Client) Delete(path string) error {
	if c.engine == nil {
		err := c.Connect()
		if err != nil {
			return err
		}
	}
	stat, err := c.engine.Stat(path)
	if utils.ManageError(err) != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("path is a directory")
	}
	return c.engine.Delete(path)
}
