// Package osquery Modified by harshit98 on 2025-05-10
// Changes: Added function score support
package osquery

type FunctionScoreQuery struct {
	query     Mappable
	functions []Function
	boostMode string
	scoreMode string
	maxBoost  *float32
	minScore  *float32
	boost     *float32
}

type Function interface {
	Map() map[string]interface{}
}

type RandomScoreFunction struct {
	seed  *int64
	field string
}

type ScriptScoreFunction struct {
	script *ScriptField
}

// Additional functions can be added as per usecase in future like:
// - ScriptScoreFunction
// - DecayFunction (with variants for geo, date, numeric)
// - WeightFunction
//
// For ref: https://docs.opensearch.org/docs/latest/query-dsl/compound/function-score/

func FunctionScore(query Mappable) *FunctionScoreQuery {
	return &FunctionScoreQuery{query: query}
}

func (q *FunctionScoreQuery) Function(f Function) *FunctionScoreQuery {
	q.functions = append(q.functions, f)
	return q
}

func (q *FunctionScoreQuery) BoostMode(boostMode string) *FunctionScoreQuery {
	q.boostMode = boostMode
	return q
}

func (q *FunctionScoreQuery) ScoreMode(scoreMode string) *FunctionScoreQuery {
	q.scoreMode = scoreMode
	return q
}

func (q *FunctionScoreQuery) MaxBoost(maxBoost float32) *FunctionScoreQuery {
	q.maxBoost = &maxBoost
	return q
}

func (q *FunctionScoreQuery) MinScore(minScore float32) *FunctionScoreQuery {
	q.minScore = &minScore
	return q
}

func (q *FunctionScoreQuery) Boost(boost float32) *FunctionScoreQuery {
	q.boost = &boost
	return q
}

func (q *FunctionScoreQuery) Map() map[string]interface{} {
	m := make(map[string]interface{})

	if q.query != nil {
		m["query"] = q.query.Map()
	}

	if len(q.functions) > 0 {
		funcs := make([]map[string]interface{}, len(q.functions))
		for i, f := range q.functions {
			funcs[i] = f.Map()
		}
		m["functions"] = funcs
	}

	if q.boostMode != "" {
		m["boost_mode"] = q.boostMode
	}

	if q.maxBoost != nil {
		m["max_boost"] = *q.maxBoost
	}

	if q.scoreMode != "" {
		m["score_mode"] = q.scoreMode
	}

	if q.minScore != nil {
		m["min_score"] = *q.minScore
	}

	if q.boost != nil {
		m["boost"] = *q.boost
	}

	return map[string]interface{}{
		"function_score": m,
	}
}

func RandomScore() *RandomScoreFunction {
	return &RandomScoreFunction{}
}

func (f *RandomScoreFunction) Seed(seed int64) *RandomScoreFunction {
	f.seed = &seed
	return f
}

func (f *RandomScoreFunction) Field(field string) *RandomScoreFunction {
	f.field = field
	return f
}

func (f *RandomScoreFunction) Map() map[string]interface{} {
	m := make(map[string]interface{})

	if f.seed != nil {
		m["seed"] = *f.seed
	}

	if f.field != "" {
		m["field"] = f.field
	}

	return map[string]interface{}{
		"random_score": m,
	}
}

func FunctionScriptScore(script *ScriptField) *ScriptScoreFunction {
	return &ScriptScoreFunction{script: script}
}

func (f *ScriptScoreFunction) Map() map[string]interface{} {
	if f.script == nil {
		return map[string]interface{}{
			"script_score": map[string]interface{}{},
		}
	}
	scriptMap := f.script.Map()["script"].(map[string]interface{})
	return map[string]interface{}{
		"script_score": map[string]interface{}{
			"script": scriptMap,
		},
	}
}
