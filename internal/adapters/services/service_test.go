package services

import (
	"errors"
	"testing"

	"github.com/EvertonTomalok/ports-challenge/internal/core/domain"
	"github.com/EvertonTomalok/ports-challenge/internal/ports"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	var stubData = domain.PortData{"stub": domain.Port{}}

	t.Run("service returns success", func(t *testing.T) {
		// Initialize gomock controller
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepository := ports.NewMockRepository(ctrl)
		mockRepository.
			EXPECT().
			Upsert("stub", gomock.Any()).
			Return(nil).
			Times(1)

		fakeService := NewService(mockRepository)
		err := fakeService.Upsert(stubData)

		assert.Nil(t, err)
	})

	t.Run("service returns error", func(t *testing.T) {
		// Initialize gomock controller
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepository := ports.NewMockRepository(ctrl)
		mockRepository.
			EXPECT().
			Upsert("stub", gomock.Any()).
			Return(errors.New("some error")).
			Times(1)

		fakeService := NewService(mockRepository)
		err := fakeService.Upsert(stubData)

		assert.Error(t, err)
	})
}
