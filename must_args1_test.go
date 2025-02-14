package must

import (
	"context"
	"errors"
	"fmt"
)

type args1 struct {
	arg1 any
	err  error
}

type result1 struct {
	value    any
	panicked bool
}

func aIntValueAndNoError(ctx context.Context, arg1 int) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args1{arg1, error(nil)}), nil
}

func theMustFunctionIsCalled(ctx context.Context) (context.Context, error) {
	passedArgs, ok := ctx.Value(argsKey{}).(args1)
	if !ok {
		return ctx, errors.New("args1 not found in context")
	}

	panicked := true
	var value any
	func() {
		defer func() {
			if r := recover(); r == nil {
				panicked = false
			}
		}()
		value = Must(passedArgs.arg1, passedArgs.err)
	}()
	return context.WithValue(ctx, resultKey{}, result1{value: value, panicked: panicked}), nil
}

func theResult1ShouldBeInt(ctx context.Context, expected int) (context.Context, error) {
	r, ok := ctx.Value(resultKey{}).(result1)
	if !ok {
		return ctx, errors.New("result1 not found in context")
	}
	intResult, ok := r.value.(int)
	if !ok {
		return ctx, errors.New("result1 is not an int")
	}
	if intResult != expected {
		return ctx, fmt.Errorf("expected %v, got %v", expected, intResult)
	}
	return ctx, nil
}

func anIntValueAndAnError(ctx context.Context, arg1 int, err string) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args1{arg1, errors.New(err)}), nil
}

func itShouldPanic1(ctx context.Context, result result1) (context.Context, error) {
	if !result.panicked {
		return ctx, errors.New("expected to panic, but it did not")
	}
	return ctx, nil
}

func itShouldNotPanic1(ctx context.Context, result result1) (context.Context, error) {
	if result.panicked {
		return ctx, errors.New("expected not to panic, but it did")
	}
	return ctx, nil
}

func twoIntValuesAndAnError(ctx context.Context, arg1 int, arg2 int, err string) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args2{arg1, arg2, errors.New(err)}), nil
}
