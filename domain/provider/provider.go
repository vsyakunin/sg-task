package provider

import (
	"fmt"
	"io/ioutil"

	myerrs "sg-task/domain/models/errors"

	log "github.com/sirupsen/logrus"
)

const (
	folderName = "data"

	internalErr = "internal error"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) DownloadFile(fileKey string) ([]byte, error) {
	const funcName = "provider.DownloadFile"

	fileName := fmt.Sprintf("%s/%s", folderName, fileKey)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("%s: error while reading file for file key %s error: %v", funcName, fileKey, err)
		return nil, myerrs.NewServerError(internalErr, err)
	}

	return file, nil
}
