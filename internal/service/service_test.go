package service

import (
	"errors"
	mock_storage "github.com/Shemistan/Lesson_5/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	t.Run("request is nil", func(t *testing.T) {
		// Ожидаем вызов метода Add у мока IStorage с аргументом nil
		storage.EXPECT().Add(nil).Return(0, errors.New("some error"))

		// Вызываем метод Add у сервиса
		_, err := serv.Add(nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

}
