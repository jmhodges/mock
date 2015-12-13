package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/mock/mockgen/ignore/mock_orig"
	"github.com/golang/mock/mockgen/ignore/orig"
)

type Thing struct {
	a orig.Adder
}

func (t Thing) CallInc(x int64) {
	t.a.Inc(x)
}

func TestAdder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	a := mock_orig.NewMockAdder(ctrl)
	a.EXPECT().Inc(2)
	th := Thing{a}
	th.CallInc(2)
}

func TestHasVariadic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	v := mock_orig.NewMockHasVariadic(ctrl)
	v.EXPECT().SomeIntVar()
}
