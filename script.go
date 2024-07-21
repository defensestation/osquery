package osquery

import (
	"github.com/fatih/structs"
)

type ScriptParams map[string]interface{}

type ScriptField struct {
	name     string
	Src      string       `structs:"source,omitempty"`
	Param    ScriptParams `structs:"params,omitempty"`
	Id       string       `structs:"id,omitempty"`
	Language string       `structs:"lang,omitempty"`
}

func Script(name string) *ScriptField {
	return &ScriptField{name: name}
}

func (f *ScriptField) Source(source string) *ScriptField {
	f.Src = source
	return f
}

func (f *ScriptField) Params(params ScriptParams) *ScriptField {
	f.Param = params
	return f
}

func (f *ScriptField) ID(id string) *ScriptField {
	f.Id = id
	return f
}

func (f *ScriptField) Lang(lang string) *ScriptField {
	f.Language = lang
	return f
}

func (f *ScriptField) Name() string {
	return f.name
}

func (f *ScriptField) Map() map[string]interface{} {
	return map[string]interface{}{
		"script": structs.Map(f),
	}
}
