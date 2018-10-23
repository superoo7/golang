# LEARNING GOOGLE GO
This repo is meant for notes taking for golang

_`learn-go`_ learning some go stuff

## learn-go
* 01 so test file

* 02 Templating
    - 07  concatenation
    
    - 08  [text/template](https://golang.org/pkg/text/template/)
    - 09 Passing data to templates with `{{.}}`
    - 10 Variable in templates with `{{$test := .}}`
    - 11 Passing in composite data structure (`slice`, `map`, `struct`) into templates
        ```html
        {{range .}}
            <li> {{.}} </li>
        {{end}}
        ```
    - 12 Function in Templates up
        ```go
            type FucMap map[string]interface{}
            [in html]
            func .var
        ```
