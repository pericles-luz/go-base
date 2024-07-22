package zipper

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

const (
	ZIPPER_INPUT_DESTINATION = iota
	ZIPPER_INPUT_SOURCE
)

var (
	ErrOutputNotSet = errors.New("arquivo de destino não definido")
	ErrNoFilesToZip = errors.New("nenhum arquivo para zipar")
)

type Zipper struct {
	files  []map[int]string
	output string
}

func NewZipper() *Zipper {
	return &Zipper{}
}

func (z *Zipper) AddFile(fileSource, fileDestination string) {
	for i, f := range z.files {
		if f[ZIPPER_INPUT_DESTINATION] == fileDestination {
			z.files[i][ZIPPER_INPUT_SOURCE] = fileSource
			return
		}
	}
	z.files = append(z.files, map[int]string{ZIPPER_INPUT_DESTINATION: fileDestination, ZIPPER_INPUT_SOURCE: fileSource})
}

func (z *Zipper) SetOutput(output string) {
	z.output = output
}

func (z *Zipper) Zip() error {
	if z.output == "" {
		return ErrOutputNotSet
	}
	if len(z.files) == 0 {
		return ErrNoFilesToZip
	}
	zipWriter, err := z.getZipWriter()
	if err != nil {
		return err
	}
	defer zipWriter.Close()
	for _, f := range z.files {
		fileSource, err := os.Open(f[ZIPPER_INPUT_SOURCE])
		if err != nil {
			return err
		}
		defer fileSource.Close()
		fileInfo, err := fileSource.Stat()
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		header.Name = f[ZIPPER_INPUT_DESTINATION]
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, fileSource)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *Zipper) getZipWriter() (*zip.Writer, error) {
	file, err := os.Create(z.output)
	if err != nil {
		return nil, err
	}
	zipWriter := zip.NewWriter(file)
	return zipWriter, nil
}

func (z *Zipper) Unzip(zipFile string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()
		fileWriter, err := os.Create(file.Name)
		if err != nil {
			return err
		}
		defer fileWriter.Close()
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *Zipper) FileList(zipFile string) ([]string, error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var files []string
	for _, file := range reader.File {
		files = append(files, file.Name)
	}
	return files, nil
}
