package application_test

import (
	"testing"

	"github.com/lucasmatsui/go-hexagonal/application"
	mock_application "github.com/lucasmatsui/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"

	gomock "github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInteface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")

	require.Nil(t, err)
	require.Equal(t, product, result)

}
