package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"winehouse/domain"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.CreatedAt = time.Now()
	video.FilePath = "path"
	video.ResourceID = "123"
	video.ID = "abc"

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.CreatedAt = time.Now()
	video.FilePath = "path"
	video.ResourceID = "123"
	video.ID = uuid.NewV4().String()

	err := video.Validate()
	require.Nil(t, err)
}
