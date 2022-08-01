package html

import (
	"testing"
	_ "fmt"
)

func TestHtml(t *testing.T) {

	htmlResult := Html().Child(
		Div().
			Styles("width:100px;heigth:100px;backgrund:red").
			Text("My Html bulder"), 
			Select().Options(
				Option("1", "teste 1"),
				Option("2", "teste 2")),
				Input().Type("number").Val("my ass")).
		AsHtml()

	t.Fatalf(htmlResult)

}