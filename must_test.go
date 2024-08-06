package must

import (
	"net/url"
	"reflect"
	"testing"
)

func TestMustURL(t *testing.T) {
	type args[T any] struct {
		v   T
		err error
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		want      T
		wantPanic bool
	}
	tests := []testCase[*url.URL]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Must(tt.args.v, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Must() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMust2(t *testing.T) {
	type args[T1 any, T2 any] struct {
		v1  T1
		v2  T2
		err error
	}
	type testCase[T1 any, T2 any] struct {
		name  string
		args  args[T1, T2]
		want  T1
		want1 T2
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Must2(tt.args.v1, tt.args.v2, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Must2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Must2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMust3(t *testing.T) {
	type args[T1 any, T2 any, T3 any] struct {
		v1  T1
		v2  T2
		v3  T3
		err error
	}
	type testCase[T1 any, T2 any, T3 any] struct {
		name  string
		args  args[T1, T2, T3]
		want  T1
		want1 T2
		want2 T3
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Must3(tt.args.v1, tt.args.v2, tt.args.v3, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Must3() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Must3() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Must3() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
