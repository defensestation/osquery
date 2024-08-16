package osquery

type ScriptParams map[string]interface{}

type ScriptField struct {
	name     string
	Src      string
	Param    ScriptParams
	Id       string
	Language string
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
	result := make(map[string]interface{})
	if f.Src != "" {
		result["source"] = f.Src
	}
	if f.Param != nil {
		result["params"] = f.Param
	}
	if f.Id != "" {
		result["id"] = f.Id
	}
	if f.Language != "" {
		result["lang"] = f.Language
	}
	return map[string]interface{}{
		"script": result,
	}
}
