package viewmodels

import (
	"../models"
)

type NewViewModel struct {
	Entry  model.Entry
	Error  []string
	Theme  string
	MetaOg MetaOg
	Colors []string
}
