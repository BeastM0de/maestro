package maestro

type Repository struct {
	name  string
	path  string
	url   string
	Group Group
}

func (r Repository) Name() string {
	return r.name
}

func (r Repository) Path() string {
	return r.path
}

func (r Repository) Url() string {
	return r.url
}

func (r Repository) Parent() Resource {
	return r.Group
}

func (r Repository) Children() []Resource {
	return nil
}
