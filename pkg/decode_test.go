package resp

import (
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
			want1: 3,
		},
		{
			name: "10",
			args: args{
				input: []byte(":+10\r\n"),
			},
			want:  10,
			want1: 4,
		},
		{
			name: "393003",
			args: args{
				input: []byte(":+393003\r\n"),
			},
			want:  393003,
			want1: 8,
		},
		{
			name: "no int",
			args: args{
				input: []byte(":+\r\n"),
			},
			want:  0,
			want1: 2,
		},
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
			got, got1, err := decodeInt(tt.args.input)
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
			want1:   5,
			wantErr: false,
		},
		{
			name: "no string",
			args: args{
				input: []byte("+\r\n"),
			},
			want:    "",
			want1:   1,
			wantErr: false,
		},
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
			got, got1, err := decodeSimpleString(tt.args.input)
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
			want1:   6,
			wantErr: false,
		},
		{name: "valid bulk string long bulk",
			args: args{
				input: []byte("$15\r\nhappy halloween\r\n"),
			},
			want:    "happy halloween",
			want1:   20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := decodeBulkString(tt.args.input)
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
