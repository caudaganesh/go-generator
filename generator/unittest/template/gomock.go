package template

func NewGomockTemplate() string {
	return `
{{$func := .Func -}}
{{$CalledMethods := .CalledMethods -}}
// {{ $func.TestName }} is generated by gospore
func {{ $func.TestName }}(t *testing.T) {
	{{ if and $func.IsMethod $func.HasReceiverFields }}
		type fields struct {
			{{ range $field:= fields $func }} {{$field}}
			{{ end }}
		}
        {{range $i, $val := $CalledMethods}}
			type m{{$val.MethodName}} struct {
				{{range $j, $arg := $val.Args -}}
				  in{{$j}} {{$arg}}
				{{end -}}
                {{range $k, $res := $val.Results -}}
				  out{{$k}} {{$res}}
				{{end -}}
				  times int
			}
		{{end}}
	{{ end }}
	{{- if (gt $func.NumParams 0) }}
		type args struct {
			{{ range $param := params $func }}
				{{- $param}}
			{{ end }}
		}
	{{ end -}}
	tests := []struct {
		name string
		{{ if (gt $func.NumParams 0) }} args args {{ end }}
		{{ if $func.HasReceiverFields }} fields fields {{ end }}
		{{ range $result := results $func}} {{ want $result }}; {{ end -}}
		{{ if $func.ReturnsError }} wantErr bool {{ end }}
        {{range $i, $val := $CalledMethods -}}
           m{{$val.MethodName}} m{{$val.MethodName}}
        {{end}}
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{ if $func.IsMethod }}
				{{if $func.HasMockableField $func.ReceiverFields}}
					ctrl := gomock.NewController(t)
                    defer ctrl.Finish()
				{{end}}
				{{ range $receiverField:= $func.ReceiverFields }} 
					{{if and $receiverField.Mockable $receiverField.IsCalled}}
						mock{{$receiverField.StructName}}:=NewMock{{$receiverField.StructName}}(ctrl)
						tt.fields.{{$receiverField.Name}} = mock{{$receiverField.StructName}}
					    {{- range $i, $val := $CalledMethods}}
					    	{{- if eq $val.FieldName $receiverField.Name }}
							   mock{{$receiverField.StructName}}.EXPECT().
								{{$val.MethodName}}(
									{{ range $j, $arg := $val.Args -}}
									  tt.m{{$val.MethodName}}.in{{$j}},
									{{end}}).
								Return(
									{{range $j, $res := $val.Results -}}
									  tt.m{{$val.MethodName}}.out{{$j}},
									{{end}}).
								Times(tt.m{{$val.MethodName}}.times)
							{{end}}
						{{end}}
					{{end}}
				{{ end }}				

				executor := {{$func.ReceiverInstance}} {
					{{ range $receiverField:= $func.ReceiverFields }} {{$receiverField.Name}} : tt.fields.{{$receiverField.Name}},
					{{ end }}
				}
				{{ if (gt $func.NumResults 0) }}{{ join $func.ResultsNames ", " }} := {{end}}executor.{{$func.Name}}(
					{{- range $i, $pn := $func.ParamsNames }}
						{{- if not (eq $i 0)}},{{end}}tt.args.{{ $pn }}{{ end }})

			{{ else }}
				{{ if (gt $func.NumResults 0) }}{{ join $func.ResultsNames ", " }} := {{end}}{{$func.Name}}(
					{{- range $i, $pn := $func.ParamsNames }}
						{{- if not (eq $i 0)}},{{end}}tt.args.{{ $pn }}{{ end }})
			{{end}}
			{{ range $result := $func.ResultsNames }}
				{{ if (eq $result "err") }}
				if tt.wantErr {
				  assert.Error(t, err)
				}
				{{ else }}
				  assert.Equal(t, tt.{{ want $result }}, {{ $result }})
				{{end -}}
			{{end -}}
		})
	}
}
`
}