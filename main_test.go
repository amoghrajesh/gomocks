package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_myinterface "gomocks/myinterface/mocks"
	"testing"
)

func TestMyInterface(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_myinterface.NewMockMyInterface(ctrl)

	m.EXPECT().IsUnder16(20).Return(false)
	assert.Equal(t, m.IsUnder16(20), false)

}
