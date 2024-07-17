package osquery

type Params map[string]interface{}

type ScriptField struct {
	field  string
	source string
	params Params
}

type ScriptFields []*ScriptField

func (s ScriptFields) Map() map[string]interface{} {
	fields := map[string]interface{}{}
	for _, field := range s {
		r := map[string]interface{}{}
		if field.source != "" {
			r["source"] = field.source
		}
		if len(field.params) > 0 {
			r["params"] = field.params
		}
		fields[field.field] = r
	}
	return fields
}

func Script(field, source string, params Params) *ScriptField {
	return &ScriptField{
		field:  field,
		source: source,
		params: params,
	}
}
