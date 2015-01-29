package main

import (
	"../models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEntryModelSpec(t *testing.T) {
	Convey("create Entry", t, func() {
		entry := model.Entry{
			Title:    "うひょひょひょー",
			Content:  "ひょひょひょのひょーーーーーー",
			Basename: "post-1",
		}

		So(entry.Title, ShouldEqual, "うひょひょひょー")
		So(entry.Content, ShouldEqual, "ひょひょひょのひょーーーーーー")
		So(entry.Basename, ShouldEqual, "post-1")
	})
}
