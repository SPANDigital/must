package must

import (
	"context"
	"errors"
)

type args3 struct {
	arg1 any
	arg2 any
	arg3 any
	err  error
}

type result3 struct {
	value1   any
	value2   any
	value3   any
	panicked bool
}

func values3IntAndNoError(ctx context.Context, arg1, arg2, arg3 int) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args3{arg1, arg2, arg3, error(nil)}), nil
}

func must3FunctionIsCalled(ctx context.Context) (context.Context, error) {
	passedArgs, ok := ctx.Value(argsKey{}).(args3)
	if !ok {
		return ctx, errors.New("args2 not found in context")
	}

	panicked := true
	var value1 any
	var value2 any
	var value3 any
	func() {
		defer func() {
			if r := recover(); r == nil {
				panicked = false
			}
		}()
		value1, value2, value3 = Must3(passedArgs.arg1, passedArgs.arg2, passedArgs.arg3, passedArgs.err)
	}()
	return context.WithValue(ctx, resultKey{}, result3{value1: value1, value2: value2, value3: value3, panicked: panicked}), nil
}

func itShouldPanic3(ctx context.Context, result result3) (context.Context, error) {
	if !result.panicked {
		return ctx, errors.New("expected to panic, but it did not")
	}
	return ctx, nil
}
func itShouldNotPanic3(ctx context.Context, result result3) (context.Context, error) {
	if result.panicked {
		return ctx, errors.New("expected to not panic, but it did")
	}
	return ctx, nil
}

func theResult3ShouldBeInt(ctx context.Context, expected1, expected2, expected3 int) (context.Context, error) {
	r, ok := ctx.Value(resultKey{}).(result3)
	if !ok {
		return ctx, errors.New("result3 not found in context")
	}
	intResult1, ok := r.value1.(int)
	if !ok {
		return ctx, errors.New("value1 is not an int")
	}
	if intResult1 != expected1 {
		return ctx, errors.New("expected1 is not an int")
	}
	intResult2, ok := r.value2.(int)
	if !ok {
		return ctx, errors.New("value2 is not an int")
	}
	if intResult2 != expected2 {
		return ctx, errors.New("expected2 is not an int")
	}
	intResult3, ok := r.value3.(int)
	if !ok {
		return ctx, errors.New("value3 is not an int")
	}
	if intResult3 != expected3 {
		return ctx, errors.New("expected3 is not an int")
	}
	return ctx, nil
}

func threeIntValuesAndAnError(ctx context.Context, arg1, arg2, arg3 int, err string) (context.Context, error) {
	return context.WithValue(ctx, argsKey{}, args3{arg1, arg2, arg3, errors.New(err)}), nil
}
