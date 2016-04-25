package main

import (
//	h "net/http" // mock
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	//do nothing
//	ctrl := gomock.NewController(t)
	h.EXCEPT().getSite(url).Returns(200)
}

func TestGetSite(t *testing.T) {
	//	url := "https://www.baidu.com"
	//	getSite(url)
	h.MOCK().SetController(ctrl)
	h.EXCEPT().getSite(url).Returns(200)
	//	So()
	//	should
	//	Convey()
}
