package entity_test

import (
	"testing"

	"github.com/Euler-B/bootcampFullCycle/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateaNewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}