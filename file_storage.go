package main

import (
	"os"
	"path"
)

type FileStorage struct {
	directory string
}

func NewFileStorage(directory string) (*FileStorage, error) {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &FileStorage{directory: directory}, nil
}

func (s *FileStorage) Store(entity, id string, labels []string, data []byte) error {
	dir := path.Join(s.directory, entity)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	p := path.Join(dir, id)

	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	if len(labels) == 0 {
		return nil
	}

	labelsFile, err := os.OpenFile(p+".labels", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	defer labelsFile.Close()

	labelsString := ""

	for _, label := range labels {
		labelsString += label + ","
	}

	labelsString = labelsString[:len(labelsString)-1]

	_, err = labelsFile.WriteString(labelsString)
	if err != nil {
		return err
	}

	return nil
}
