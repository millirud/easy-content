package transform

import (
	"io"
	"io/ioutil"
	"os"
	"video-api/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type VideoTransformer struct {
	log     *zap.Logger
	storage fileStorage
}

func New(
	log *zap.Logger,
	storage fileStorage,
) *VideoTransformer {
	return &VideoTransformer{
		log:     log,
		storage: storage,
	}
}

type (
	fileStorage interface {
		Download(bucket string, filename string) (io.Reader, func() error, error)
	}
)

func (v *VideoTransformer) TransformVideo(fileInfo entity.UploadedFile) error {
	v.log.Sugar().Infof("VideoTransformer.TransformVideo: file: %+v", fileInfo)

	contentReader, downloadCloser, err := v.storage.Download(fileInfo.GetBucket(), fileInfo.GetFilename())

	if err != nil {
		v.log.Sugar().Error(err.Error())
		return err
	}

	defer func() {
		if err := downloadCloser(); err != nil {

			v.log.Error(err.Error())
		}
	}()

	filepath, localCloser, err := v.tempWrite(contentReader)

	if err != nil {
		v.log.Sugar().Error(err.Error())
		return err
	}

	v.log.Sugar().Infof("file written to %s", filepath)

	defer func() {
		if err := localCloser(); err != nil {
			v.log.Error(err.Error())
		}
	}()

	return nil
}

func (v *VideoTransformer) tempWrite(reader io.Reader) (string, func() error, error) {

	file, err := ioutil.TempFile("", v.randomFilename())

	if err != nil {
		v.log.Error(err.Error())
		return "", nil, err
	}

	return file.Name(), func() error { return os.Remove(file.Name()) }, nil
}

func (v *VideoTransformer) randomFilename() string {
	val, err := uuid.NewUUID()

	if err != nil {
		v.log.Sugar().Warn(err.Error())
		return v.randomFilename()
	}

	return val.String()
}
