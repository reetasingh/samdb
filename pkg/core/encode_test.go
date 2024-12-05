package core

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeInt(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "positive",
			args: args{
				input: 10,
			},
			want: []byte(":10\r\n"),
		},
		{
			name: "negative",
			args: args{
				input: -10,
			},
			want: []byte(":-10\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeInt(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "simple",
			args: args{
				err: fmt.Errorf("new error"),
			},
			want: []byte("-new error\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	type args struct {
		input string
		bulk  bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "bulk",
			args: args{
				input: "aaaa",
				bulk:  true,
			},
			want: []byte("$4\r\naaaa\r\n"),
		},
		{
			name: "simple",
			args: args{
				input: "aaaa",
				bulk:  false,
			},
			want: []byte("+aaaa\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeString(tt.args.input, tt.args.bulk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
