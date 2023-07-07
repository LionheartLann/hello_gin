package cache

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetFromCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockCache(ctrl)
	m.EXPECT().Get(gomock.Eq("Math")).Return(-10, true)
	if v, _ := GetFromCache(m, "Math"); v != -10 {
		t.Fatal("expected -10,but got", v)
	}

	m.EXPECT().Get(gomock.Any()).Return(90, true)
	if v, _ := GetFromCache(m, "English"); v != 90 {
		t.Fatal("expected 90,but got", v)
	}

	m.EXPECT().Get(gomock.Any()).Return(nil, false)
	if v, ok := GetFromCache(m, "English"); ok {
		t.Fatal("expected false,but got", v)
	}
	fmt.Println("test pass")
}
