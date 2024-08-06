package must

import (
	"errors"
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
	tests := []testCase[int]{
		{
			name: "failure",
			args: args[int]{
				v:   1,
				err: nil,
			},
			want:      1,
			wantPanic: false,
		},
		{
			name: "success",
			args: args[int]{
				v:   0,
				err: errors.New("some error"),
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			panicked := false
			var got int
			func() {
				defer func() {
					if r := recover(); r != nil {
						panicked = true
					}
				}()
				got = Must(tt.args.v, tt.args.err)
			}()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if panicked != tt.wantPanic {
				t.Errorf("Panicked = %v, wantPanic %v", panicked, tt.wantPanic)
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
		name      string
		args      args[T1, T2]
		want1     T1
		want2     T2
		wantPanic bool
	}
	tests := []testCase[int, int]{
		{
			name: "success",
			args: args[int, int]{
				v1:  1,
				v2:  2,
				err: nil,
			},
			want1:     1,
			want2:     2,
			wantPanic: false,
		},
		{
			name: "failure",
			args: args[int, int]{
				v1:  1,
				v2:  2,
				err: errors.New("some error"),
			},
			want1:     0,
			want2:     0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			panicked := false
			var got1, got2 int
			func() {
				defer func() {
					if r := recover(); r != nil {
						panicked = true
					}
				}()
				got1, got2 = Must2(tt.args.v1, tt.args.v2, tt.args.err)
			}()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Must2() got1 = %v, want2 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Must2() got2 = %v, want2 %v", got1, tt.want2)
			}
			if panicked != tt.wantPanic {
				t.Errorf("Panicked = %v, wantPanic %v", panicked, tt.wantPanic)
			}
		})
	}
}
