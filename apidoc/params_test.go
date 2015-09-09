package apidoc

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type (
	User struct {
		ID    int
		Name  string
		Tags  []string
		Items []*Item
		People
	}
	People struct {
		Age      int    `json:"age"`
		Countory string `json:"contory"`
	}
	Item struct {
		Name   string `json:"name"   doc:"名称"`
		volume int    `json:"volume" doc:"数量"`
	}
	Test struct {
		*Item
		People  *People `json:"people"   doc:"人間"`
		People2 People  `json:"people2"   doc:"人間2"`
	}
)

func TestChannelImage(t *testing.T) {

	Convey("main", t, func() {

		Convey("case(1)", func() {
			item := Item{}
			data := objectAnalysis("hoge", item)
			for i, v := range data {
				fmt.Println(i, v)
			}
		})
		Convey("case(2)", func() {
			item := Test{}
			data := objectAnalysis("fuga", item)
			for i, v := range data {
				fmt.Println(i, v)
			}
		})
	})
}
