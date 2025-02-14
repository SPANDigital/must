package must

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"testing"
)

type argsKey struct{}
type resultKey struct{}

func itShouldPanic(ctx context.Context) (context.Context, error) {
	result := ctx.Value(resultKey{})
	switch r := result.(type) {
	case result1:
		return itShouldPanic1(ctx, r)
	case result2:
		return itShouldPanic2(ctx, r)
	case result3:
		return itShouldPanic3(ctx, r)
	default:
		return ctx, errors.New("result key not found in context")
	}
}

func itShouldNotPanic(ctx context.Context) (context.Context, error) {
	result := ctx.Value(resultKey{})
	switch r := result.(type) {
	case result1:
		return itShouldNotPanic1(ctx, r)
	case result2:
		return itShouldNotPanic2(ctx, r)
	case result3:
		return itShouldNotPanic3(ctx, r)
	default:
		return ctx, errors.New("result key not found in context")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a value (\d+) and no error$`, aIntValueAndNoError)
	ctx.Step(`^the Must function is called$`, theMustFunctionIsCalled)
	ctx.Step(`^the result should be (\d+)$`, theResult1ShouldBeInt)
	ctx.Step(`^it should panic$`, itShouldPanic)
	ctx.Step(`^it should not panic$`, itShouldNotPanic)
	ctx.Step(`^a value (\d+) and an error "([^"]*)"$`, anIntValueAndAnError)
	ctx.Step(`^values (\d+) and (\d+) and no error$`, values2IntAndNoError)
	ctx.Step(`^the Must2 function is called$`, must2FunctionIsCalled)
	ctx.Step(`^the results should be (\d+) and (\d+)$`, theResult2ShouldBeInt)
	ctx.Step(`^values (\d+) and (\d+) and an error "([^"]*)"$`, twoIntValuesAndAnError)
	ctx.Step(`^values (\d+), (\d+), (\d+), and no error$`, values3IntAndNoError)
	ctx.Step(`^the Must3 function is called$`, must3FunctionIsCalled)
	ctx.Step(`^the results should be (\d+), (\d+), and (\d+)$`, theResult3ShouldBeInt)
	ctx.Step(`^values (\d+), (\d+), (\d+), and an error "([^"]*)"$`, threeIntValuesAndAnError)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
