package main

import "testing"
import "github.com/golang/mock/gomock"

func Test_realImpl_foo(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockmyInterface(ctrl)

	m.EXPECT().foo(gomock.Eq(1), gomock.Eq("abcd")).Return()

}
