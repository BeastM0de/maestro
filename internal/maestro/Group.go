package maestro

type Group struct {
	name         string
	path         string
	url          string
	Repositories []Repository
}

func (g Group) Name() string {
	return g.name
}

func (g Group) Path() string {
	return g.path
}

func (g Group) Url() string {
	return g.url
}

func (g Group) Parent() Resource {
	return nil
}

func (g Group) Children() []Resource {
	var resources []Resource
	for _, repo := range g.Repositories {
		resources = append(resources, repo)
	}
	return resources
}
