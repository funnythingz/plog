package viewmodels

type Paginate struct {
	IsEndpoint   bool
	IsFirstpoint bool
	CurrentPage  int
	PrevPage     int
	NextPage     int
}
