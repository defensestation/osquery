package osquery

type Collapse struct {
	field string
	Mappable
}

func NewCollapse(field string) *Collapse {
	return &Collapse{
		field: field,
	}
}

func (c *Collapse) Map() map[string]interface{} {
	outerMap := map[string]interface{}{
		"field": c.field,
	}

	return outerMap
}
