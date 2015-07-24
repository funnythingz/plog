package viewmodels

import (
	"github.com/funnythingz/plog/models"
)

type NewViewModel struct {
	Entry  model.Entry
	Error  []string
	Theme  string
	MetaOg MetaOg
	Colors []string
}
