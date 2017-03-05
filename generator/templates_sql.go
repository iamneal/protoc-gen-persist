// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import "github.com/tcncloud/protoc-gen-persist/persist"

func SetupTemplates() {
	UnaryTemplateString[persist.PersistenceOptions_SQL] = `
func (s *{{.GetServiceImplName}}) {{.GetMethod}}(ctx context.Context, req *{{.GetInputType}}) (*{{.GetOutputType}}, error) {
	var (
		{{range $field := .GetSafeResponseFields}}
		{{$field.K}} {{$field.V}} {{end}}
	)
	err := s.DB.QueryRow(
		"{{.GetQuery}}",{{range $qParam := .GetQueryParams}}
		_utils.ToSafeType(req.{{$qParam}}),
		{{end}}).
		Scan({{range $field := .GetSafeResponseFields}} &{{$field.K}},
		{{end}})

	if err != nil {
		return nil, ConvertError(err, req)
	}
	result := &{{.GetOutputType}}{}
	{{range $local, $go := .GetResponseFieldsMap}}
	_utils.AssignTo(&result.{{$go}}, {{$local}}) {{end}}

	return result, nil
}
`
	ServerTemplateString[persist.PersistenceOptions_SQL] = `
func (s *{{.GetServiceImplName}}) {{.GetMethod}}(ctx context.Context, req *{{.GetInputType}}) (*{{.GetOutputType}}, error) {
	var (
		{{range $field := .GetSafeResponseFields}}
		{{$field.K}} {{$field.V}} {{end}}
	)
	err := s.DB.QueryRow(
		"{{.GetQuery}}",{{range $qParam := .GetQueryParams}}
		_utils.ToSafeType(req.{{$qParam}}),
		{{end}}).
		Scan({{range $field := .GetSafeResponseFields}} &{{$field.K}},
		{{end}})

	if err != nil {
		return nil, ConvertError(err, req)
	}
	result := &{{.GetOutputType}}{}
	{{range $local, $go := .GetResponseFieldsMap}}
	_utils.AssignTo(&result.{{$go}}, {{$local}}) {{end}}

	return result, nil
}
`
	ClientTmplateString[persist.PersistenceOptions_SQL] = `
func (s *{{.GetServiceImplName}}) {{.GetMethod}}(ctx context.Context, req *{{.GetInputType}}) (*{{.GetOutputType}}, error) {
	var (
		{{range $field := .GetSafeResponseFields}}
		{{$field.K}} {{$field.V}} {{end}}
	)
	err := s.DB.QueryRow(
		"{{.GetQuery}}",{{range $qParam := .GetQueryParams}}
		_utils.ToSafeType(req.{{$qParam}}),
		{{end}}).
		Scan({{range $field := .GetSafeResponseFields}} &{{$field.K}},
		{{end}})

	if err != nil {
		return nil, ConvertError(err, req)
	}
	result := &{{.GetOutputType}}{}
	{{range $local, $go := .GetResponseFieldsMap}}
	_utils.AssignTo(&result.{{$go}}, {{$local}}) {{end}}

	return result, nil
}`
	BidirTemplateString[persist.PersistenceOptions_SQL] = `
func (s *{{.GetServiceImplName}}) {{.GetMethod}}(ctx context.Context, req *{{.GetInputType}}) (*{{.GetOutputType}}, error) {
	var (
		{{range $field := .GetSafeResponseFields}}
		{{$field.K}} {{$field.V}} {{end}}
	)
	err := s.DB.QueryRow(
		"{{.GetQuery}}",{{range $qParam := .GetQueryParams}}
		_utils.ToSafeType(req.{{$qParam}}),
		{{end}}).
		Scan({{range $field := .GetSafeResponseFields}} &{{$field.K}},
		{{end}})

	if err != nil {
		return nil, ConvertError(err, req)
	}
	result := &{{.GetOutputType}}{}
	{{range $local, $go := .GetResponseFieldsMap}}
	_utils.AssignTo(&result.{{$go}}, {{$local}}) {{end}}

	return result, nil
}
`
}