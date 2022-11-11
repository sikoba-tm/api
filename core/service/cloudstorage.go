package service

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type CloudStorageService interface {
	UploadFile(ctx context.Context, objectPath string, objectName string, file multipart.File) error
}
type googleCloudStorage struct {
	cl *storage.Client
}

func NewGoogleCloudStorage(cl *storage.Client) *googleCloudStorage {
	return &googleCloudStorage{cl: cl}
}

func (s *googleCloudStorage) UploadFile(ctx context.Context, objectPath string, objectName string, file multipart.File) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := s.cl.Bucket(os.Getenv("GCS_BUCKET")).Object(objectPath + objectName).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
