package core

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_decodeInt(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "0",
			args: args{
				input: []byte(":+0\r\n"),
			},
			want:  0,
			want1: 5,
		},
		{
			name: "10",
			args: args{
				input: []byte(":+10\r\n"),
			},
			want:  10,
			want1: 6,
		},
		{
			name: "393003",
			args: args{
				input: []byte(":+393003\r\n"),
			},
			want:  393003,
			want1: 10,
		},
		// {
		// 	name: "no int",
		// 	args: args{
		// 		input: []byte(":+\r\n"),
		// 	},
		// 	want:  0,
		// 	want1: 4,
		// },
		{
			name: "invalid char",
			args: args{
				input: []byte(":+-\r\n"),
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DecodeInt(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeSimpleString(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name: "basic string",
			args: args{
				input: []byte("+test\r\n"),
			},
			want:    "test",
			want1:   7,
			wantErr: false,
		},
		// {
		// 	name: "no string",
		// 	args: args{
		// 		input: []byte("+\r\n"),
		// 	},
		// 	want:    "",
		// 	want1:   3,
		// 	wantErr: false,
		// },
		{
			name: "invalid string",
			args: args{
				input: []byte("-\r\n"),
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DecodeSimpleString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeSimpleString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeSimpleString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeSimpleString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeBulkString(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{name: "valid bulk string",
			args: args{
				input: []byte("$2\r\nhi\r\n"),
			},
			want:    "hi",
			want1:   8,
			wantErr: false,
		},
		{name: "valid bulk string long bulk",
			args: args{
				input: []byte("$15\r\nhappy halloween\r\n"),
			},
			want:    "happy halloween",
			want1:   22,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DecodeBulkString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeBulkString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeBulkString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeBulkString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeOne(t *testing.T) {
	type args struct {
		input []byte
	}
	longText := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Curabitur pretium tincidunt lacus. Nulla gravida orci a odio, et interdum diam laoreet quis. Suspendisse potenti. Nulla facilisi. Phasellus malesuada nulla at lectus dictum, in sodales eros hendrerit. In aliquet purus ut velit volutpat, ut lacinia velit interdum. Vivamus euismod dui id turpis faucibus, vitae scelerisque purus viverra. Sed vehicula orci vel erat bibendum, vel scelerisque odio vehicula.

Integer in nibh sed odio cursus euismod. Proin ut magna at nisi egestas venenatis. Sed vitae quam vitae urna bibendum feugiat. Fusce in lacus libero. Nulla facilisi. Integer venenatis nunc ac dolor aliquet fermentum. Maecenas finibus mauris eu erat gravida, et congue felis blandit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Duis quis tortor nunc.

Praesent id est turpis. Curabitur at metus sit amet orci tincidunt tincidunt. Morbi sed magna ac neque facilisis suscipit. Phasellus convallis lorem ut nisl suscipit, et varius justo euismod. Aliquam erat volutpat. Integer gravida odio nec nisi pellentesque, vel tempor nisi aliquam. Nulla vulputate lectus sit amet odio lacinia, eu tincidunt nunc consectetur.

Vivamus a urna nisl. Aenean non purus turpis. Sed id orci sit amet magna malesuada placerat. Cras fringilla, nunc non convallis scelerisque, nunc libero tincidunt nunc, non cursus leo lectus eu justo. Vestibulum et ante dolor. Mauris ac varius metus. Sed et cursus orci. Ut ac risus nisi. Morbi volutpat, urna ac consectetur ultrices, nunc lacus efficitur nisl, ut cursus quam arcu ut nunc.`
	tests := []struct {
		name    string
		args    args
		want    any
		want1   int
		wantErr bool
	}{
		{
			name: "decode int",
			args: args{
				input: []byte(":+0\r\n"),
			},
			want:    0,
			want1:   5,
			wantErr: false,
		},
		{name: "decode bulk string",
			args: args{
				input: []byte(fmt.Sprintf("$2004\r\n%s\r\n", longText)),
			},
			want:    longText,
			want1:   2013,
			wantErr: false,
		},
		{
			name: "decode simple string",
			args: args{
				input: []byte("+test\r\n"),
			},
			want:    "test",
			want1:   7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DecodeOne(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeOne() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeOne() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeArray(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		want1   int
		wantErr bool
	}{
		{
			name: "valid bulk string array",
			args: args{
				input: []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"),
			},
			want:    []any{"hello", "world"},
			want1:   26,
			wantErr: false,
		},
		{
			name: "valid int array",
			args: args{
				input: []byte("*3\r\n:1\r\n:2\r\n:3\r\n"),
			},
			want:    []any{1, 2, 3},
			want1:   16,
			wantErr: false,
		},
		{
			name: "valid mixed array",
			args: args{
				input: []byte("*5\r\n:1\r\n$1\r\nA\r\n:2\r\n:3\r\n$5\r\nworld\r\n"),
			},
			want:    []any{1, "A", 2, 3, "world"},
			want1:   34,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := DecodeArray(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeArray() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeArray() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
