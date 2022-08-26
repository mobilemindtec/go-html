# go-html
Go Lang HTML Fluent


```

h "github.com/mobilemindtec/go-html/html"

 h.Html().Child(
		h.Div().
			Styles("width:100px;heigth:100px;backgrund:red").
			h.Text("My Html bulder"),
			h.Select().Options(
				h.Option("1", "teste 1"),
				h.Option("2", "teste 2")),
				h.Input().Type("number").Val("my ass")).
		String()

	h.P().
    Styles("font-family:Tahoma,'Helvetica Neue',Helvetica,sans-serif;font-size: 15px;color:rgba(45, 110, 155, 1);").
    Text("Texto p").
    Child(h.A().
  				  Href("https://www.mobilemind.com.br/file/%v", "12").
  				  Text("Acesse para conferir")).
  				  String()


```