package mapper

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression"
)

type mergeMapper struct {
	mappers []expression.Expr
}

const (
	Merge = "@merge"
)

func (mm *mergeMapper) Eval(scope data.Scope) (interface{}, error) {
	var combinedResult []interface{}
	for _, mapper := range mm.mappers {
		evalResult, err := mapper.Eval(scope)
		if err != nil {
			return nil, err
		}
		switch result := evalResult.(type) {
		case []interface{}:
			combinedResult = append(combinedResult, result...)
		default:
			combinedResult = append(combinedResult, result)
		}
	}
	return combinedResult, nil
}

func createMergeMapper(merges []interface{}, ef expression.Factory) (*mergeMapper, error) {
	mergeMapper := &mergeMapper{mappers: make([]expression.Expr, 0)}
	for _, v := range merges {
		mapper, err := NewObjectMapper(v, ef)
		if err != nil {
			return nil, err
		}
		mergeMapper.mappers = append(mergeMapper.mappers, mapper)
	}
	return mergeMapper, nil

}
