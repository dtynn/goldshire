package goldshire

type Store interface {
	Get(project string) (*Meta, error)
}

type Meta struct {
	Base string
	VCS  string
	Url  string
}
