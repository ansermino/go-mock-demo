package basic

import (
	mock_basic "github.com/ansermino/go-mock-demo/basic/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreMulti(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock_basic.NewMockDb(ctrl)

	db.EXPECT().Put("a", "a").Return(nil)
	db.EXPECT().Put("b", "b").Return(nil)

	assert.Nil(t, StoreMulti([]string{"a", "b"}, []string{"a","b"}, db))
}

//func TestStoreMultiFailing(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	db := mock_basic.NewMockDb(ctrl)
//
//	db.EXPECT().Put("a", "a").Return(nil)
//	db.EXPECT().Put("b", "b").Return(nil)
//
//	assert.Nil(t, StoreMulti([]string{"a", "c"}, []string{"a","b"}, db))
//}

func TestFetchMulti(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock_basic.NewMockDb(ctrl)

	db.EXPECT().Get("a").Return("a", nil)
	db.EXPECT().Get("b").Return("b", nil)

	vals, err := FetchMulti([]string{"a","b"}, db)
	assert.Nil(t, err)
	assert.Equal(t, []string{"a","b"}, vals)
}