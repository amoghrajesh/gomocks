package main

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_myinterface "gomocks/myinterface/mocks"
	"math/rand"
	"testing"
)

func TestMyInterface(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_myinterface.NewMockMyInterface(ctrl)

	m.EXPECT().IsUnder16(gomock.Any()).MaxTimes(11)

	for i := 0; i < 10; i++ {
		r := rand.Int() % 20
		fmt.Println("The random number is", r)
		assert.Equal(t, m.IsUnder16(r), false)
	}

	assert.Equal(t, m.IsUnder16(20), false)

}
