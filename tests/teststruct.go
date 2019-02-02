package tests

//go:generate gcgt S1 templates/template1.go tests_S1_template1.go
type S1 struct {
	I int
	S string
	O OtherStruct
}

type OtherStruct struct {
	I int
}
