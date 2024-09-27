package da

import (
	"bytes"
	"context"
	"encoding/base64"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	altda "github.com/ethereum-optimism/optimism/op-alt-da"
)

type s3store struct {
	bucket     string
	path       string
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
}

var _ altda.KVStore = &s3store{}

func NewS3store(bucket, path string) altda.KVStore {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
	uploader := s3manager.NewUploader(sess)
	return &s3store{
		bucket:     bucket,
		path:       path,
		downloader: downloader,
		uploader:   uploader,
	}
}

func (s *s3store) Get(ctx context.Context, key []byte) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := s.downloader.DownloadWithContext(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filepath.Join(s.path, base64.RawURLEncoding.EncodeToString(key))),
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *s3store) Put(ctx context.Context, key []byte, value []byte) error {
	b := bytes.NewBuffer(value)
	_, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filepath.Join(s.path, base64.RawURLEncoding.EncodeToString(key))),
		Body:   b,
	})
	return err
}
