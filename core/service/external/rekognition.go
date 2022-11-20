package external

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type RekognitionService interface {
	CompareImages(ctx context.Context, reference []byte, comparison []byte, THRESHOLD float64) (*rekognition.CompareFacesOutput, error)
}
type rekognitionService struct {
	cl *rekognition.Rekognition
}

func NewRekognitionClient(cl *rekognition.Rekognition) *rekognitionService {
	return &rekognitionService{cl: cl}
}

func (s *rekognitionService) CompareImages(ctx context.Context, reference []byte, comparison []byte, THRESHOLD float64) (*rekognition.CompareFacesOutput, error) {
	imageInput := &rekognition.CompareFacesInput{
		SimilarityThreshold: aws.Float64(THRESHOLD),
		SourceImage: &rekognition.Image{
			Bytes: comparison,
		},
		TargetImage: &rekognition.Image{
			Bytes: reference,
		},
	}

	result, err := s.cl.CompareFacesWithContext(ctx, imageInput)
	if err != nil {
		return nil, err
	}

	return result, nil
}
