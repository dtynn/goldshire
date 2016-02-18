package goldshire

import (
	"text/template"
)

const tpl = `
<html>
<head>
<meta name="go-import" content="{{.Domain}}/{{.Project}} {{.VCS}} {{.Url}}">
</head>
<body>
go get {{.Domain}}/{{.Path}}
</body>
</html>
`

var t = template.Must(template.New("tpl").Parse(tpl))
