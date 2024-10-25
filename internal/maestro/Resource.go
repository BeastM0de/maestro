package maestro

type Resource interface {
	Name() string
	Path() string
	Url() string
	Parent() Resource
	Children() []Resource
}
