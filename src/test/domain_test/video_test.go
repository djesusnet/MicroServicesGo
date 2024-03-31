package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVideo_Validate_IdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo("abc", "a", "path", time.Now())
	err := video.Validate()

	require.Error(t, err)
}

func TestVideo_Validate_IsValidate(t *testing.T) {
	video := domain.NewVideo(uuid.NewV4().String(), "a", "path", time.Now())
	err := video.Validate()

	require.Nil(t, err)
}
