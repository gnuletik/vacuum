package openapi

import (
	"github.com/daveshanley/vacuum/model"
	"github.com/daveshanley/vacuum/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathParameters_GetSchema(t *testing.T) {
	def := PathParameters{}
	assert.Equal(t, "path_parameters", def.GetSchema().Name)
}

func TestPathParameters_RunRule(t *testing.T) {
	def := PathParameters{}
	res := def.RunRule(nil, model.RuleFunctionContext{})
	assert.Len(t, res, 0)
}

func TestPathParameters_RunRule_DuplicatePathCheck(t *testing.T) {

	yml := `paths:
  /pizza/{cake}/{limes}:
    parameters:
      - in: path
        name: cake
    get:
      parameters:
        - in: path
          name: limes
  /pizza/{minty}/{tape}:
    parameters:
      - in: path
        name: minty
    get:
      parameters:
        - in: path
          name: tape`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)

}

func TestPathParameters_RunRule_DuplicatePathParamCheck_MissingParam(t *testing.T) {

	yml := `openapi: 3.0.1
info:
  title: pizza-cake
paths:
  /pizza/{cake}/{cake}:
    parameters:
      - in: path
        name: cake
    get:
      parameters:
        - in: path
          name: limes
  /pizza/{minty}:
    parameters:
      - in: path
        name: minty
    get:
      parameters:`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 2)

}

func TestPathParameters_RunRule_TopParameterCheck_MissingRequired(t *testing.T) {

	yml := `paths:
  /musical/{melody}/{pizza}:
    parameters:
        - in: path
          name: melody
          required: fresh
    get:
      parameters:
        - in: path
          name: pizza
          required: true`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "/musical/{melody}/{pizza}  must have 'required' parameter that is set to 'true'", res[0].Message)
}

func TestPathParameters_RunRule_TopParameterCheck_RequiredShouldBeTrue(t *testing.T) {

	yml := `paths:
  /musical/{melody}/{pizza}:
    parameters:
        - in: path
          name: melody
          required: false
    get:
      parameters:
        - in: path
          name: pizza
          required: true`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "/musical/{melody}/{pizza}  must have 'required' parameter that is set to 'true'", res[0].Message)
}

func TestPathParameters_RunRule_TopParameterCheck_MultipleDefinitionsOfParam(t *testing.T) {

	yml := `paths:
  /musical/{melody}/{pizza}:
    parameters:
        - in: path
          name: melody
        - in: path
          name: melody
    get:
      parameters:
        - in: path
          name: pizza`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "/musical/{melody}/{pizza}  contains has a parameter 'melody' defined multiple times'", res[0].Message)
}

func TestPathParameters_RunRule_TopParameterCheck_ParamNoName(t *testing.T) {

	yml := `paths:
  /musical/{melody}/{pizza}:
    parameters:
        - in: path
          name: melody
        - in: path
    get:
      parameters:
        - in: path
          name: pizza`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestPathParameters_RunRule_TopParameterCheck_MissingParamDefInOp(t *testing.T) {

	yml := `paths:
  /musical/{melody}/{pizza}/{cake}:
    parameters:
        - in: path
          name: melody
          required: true
    get:
      parameters:
        - in: path
          name: pizza
          required: true`

	path := "$"

	nodes, _ := utils.FindNodes([]byte(yml), path)

	rule := buildOpenApiTestRuleAction(path, "path_parameters", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	def := PathParameters{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "Operation must define parameter 'cake' as expected by path '/musical/{melody}/{pizza}/{cake}'",
		res[0].Message)
}