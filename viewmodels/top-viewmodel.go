package viewmodels

import (
	"github.com/funnythingz/plog/models"
)

type TopViewModel struct {
	Entries  []model.Entry
	Paginate Paginate
	MetaOg   MetaOg
	Entry    string
}
