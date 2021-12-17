package main

import (
	"github.com/golang/mock/gomock"
	"gomocks/myinterface"
	"testing"
)

func Test_realImplementation_getHttpClient(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := myinterface.NewMockMyInterface(ctrl)
	m.EXPECT().SomeMethod(1, "abcd")
}
