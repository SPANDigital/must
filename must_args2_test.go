package must

import (
	"context"
	"errors"
)

type args2 struct {
	arg1 any
	arg2 any
	err  error
}

type result2 struct {
	value1   any
	value2   any
	panicked bool
}

func values2IntAndNoError(ctx context.Context, arg1 int, arg2 int) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args2{arg1, arg2, error(nil)}), nil
}

func must2FunctionIsCalled(ctx context.Context) (context.Context, error) {
	passedArgs, ok := ctx.Value(argsKey{}).(args2)
	if !ok {
		return ctx, errors.New("args2 not found in context")
	}

	panicked := true
	var value1 any
	var value2 any
	func() {
		defer func() {
			if r := recover(); r == nil {
				panicked = false
			}
		}()
		value1, value2 = Must2(passedArgs.arg1, passedArgs.arg2, passedArgs.err)
	}()
	return context.WithValue(ctx, resultKey{}, result2{value1: value1, value2: value2, panicked: panicked}), nil
}

func theResult2ShouldBeInt(ctx context.Context, expected1, expected2 int) (context.Context, error) {
	r, ok := ctx.Value(resultKey{}).(result2)
	if !ok {
		return ctx, errors.New("result2 not found in context")
	}
	intResult1, ok := r.value1.(int)
	if !ok {
		return ctx, errors.New("result2 is not an int")
	}
	if intResult1 != expected1 {
		return ctx, errors.New("expected1 is not an int")
	}
	intResult2, ok := r.value2.(int)
	if !ok {
		return ctx, errors.New("result2 is not an int")
	}
	if intResult2 != expected2 {
		return ctx, errors.New("expected2 is not an int")
	}
	return ctx, nil
}

func itShouldPanic2(ctx context.Context, result result2) (context.Context, error) {
	if !result.panicked {
		return ctx, errors.New("expected to panic, but it did not")
	}
	return ctx, nil
}

func itShouldNotPanic2(ctx context.Context, result result2) (context.Context, error) {
	if result.panicked {
		return ctx, errors.New("expected not to panic, but it did")
	}
	return ctx, nil
}

func twoIntValuesAndAnError(ctx context.Context, arg1 int, arg2 int, err string) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args2{arg1, arg2, errors.New(err)}), nil
}
