package main

import "html/template"

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Paragraph}}
    <p>{{.}}</p>
  {{end}}
  <ul>
    {{ range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>
`

func main() {

	// tpl := template.New("").Parse(defaultHandlerTmpl)
	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

}
