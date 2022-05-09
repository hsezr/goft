package goft

type IClass interface {
	Build(g *Goft)
	Name() string
}