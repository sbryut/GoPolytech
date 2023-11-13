package generate

//go:generate mockgen -destination=mock_generate.go -package=generate . Multiplication

type Multiplication interface {
	Multiplication(a, b int) int
}
