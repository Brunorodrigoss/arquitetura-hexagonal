package application_test

import (
	"testing"

	"github.com/brunorodrigoss/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())

}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}

func TestProduct_GetId(t *testing.T) {
	expected_id := uuid.NewV4().String()
	product := application.Product{}
	product.ID = expected_id
	product.Name = "Get ID"
	product.Price = 10
	product.Status = application.ENABLED

	got := product.GetId()

	if got != expected_id {
		t.Errorf("got %q want %q", expected_id, got)
	}
}

func TestProduct_GetName(t *testing.T) {
	expected_name := "Bruno"
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Bruno"
	product.Price = 10
	product.Status = application.ENABLED

	if product.GetName() != expected_name {
		t.Errorf("got %q want %q", expected_name, product.GetName())
	}
}

func TestProduct_GetStatus(t *testing.T) {
	expected_status := application.ENABLED
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Bruno"
	product.Price = 10
	product.Status = application.ENABLED

	got := product.GetStatus()
	if got != expected_status {
		t.Errorf("got %q want %q", expected_status, got)
	}

	product.Status = application.DISABLED
	expected_status = application.DISABLED
	got = product.GetStatus()
	if got != expected_status {
		t.Errorf("got %q want %q", expected_status, got)
	}

}
