package maestro

type Resource interface {
	func Name() string
	func Path() string
	func Parent() Resource
	func Children() []Resource
}
