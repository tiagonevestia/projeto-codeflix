package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"winehouse/domain"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "123"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	require.Nil(t, err)
	require.NotNil(t, job)
}
