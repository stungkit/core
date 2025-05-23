package trigger

import (
	"context"
	"testing"

	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/data/expression"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/resolve"
	"github.com/stretchr/testify/assert"
)

var defResolver = resolve.NewCompositeResolver(map[string]resolve.Resolver{
	".": &resolve.ScopeResolver{},
})

type MockAction struct {
}

func (t *MockAction) Metadata() *action.Metadata {
	return nil
}

func (t *MockAction) IOMetadata() *metadata.IOMetadata {
	return nil
}

func (t *MockAction) Run(context context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {
	return inputs, nil
}

func TestNewHandler(t *testing.T) {
	actionCfg := &ActionConfig{Input: map[string]interface{}{"anInput": "input"}, Output: map[string]interface{}{"anOutput": "output"}}
	var actionCfgArr []*ActionConfig
	actionCfgArr = append(actionCfgArr, actionCfg)

	hCfg := &HandlerConfig{Name: "sampleConfig", Settings: map[string]interface{}{"aSetting": "aSetting"}, Actions: actionCfgArr}

	mf := mapper.NewFactory(defResolver)
	expf := expression.NewFactory(defResolver)

	//Action not specified
	handler, err := NewHandler(hCfg, nil, mf, expf, nil, log.RootLogger())
	assert.NotNil(t, err, "Actions not specified.")

	//HandlerSettings configured
	handler, err = NewHandler(hCfg, []action.Action{&MockAction{}}, mf, expf, nil, log.RootLogger())
	assert.NotNil(t, handler.Settings())
}

func TestHandlerContext(t *testing.T) {

	actionCfg := &ActionConfig{Input: map[string]interface{}{"anInput": "input"}, Output: map[string]interface{}{"anOutput": "output"}}
	var actionCfgArr []*ActionConfig
	actionCfgArr = append(actionCfgArr, actionCfg)

	hCfg := &HandlerConfig{Name: "sampleConfig", Settings: map[string]interface{}{"aSetting": "aSetting"}, Actions: actionCfgArr}
	hCfg.parent = &Config{Id: "sampleTrig"}

	assert.NotNil(t, NewHandlerContext(context.Background(), hCfg))

}
