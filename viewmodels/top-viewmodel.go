package viewmodels

import (
	"../models"
)

type TopViewModel struct {
	Entries  []model.Entry
	Paginate Paginate
	MetaOg   MetaOg
	Entry    string
}
