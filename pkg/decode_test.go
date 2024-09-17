package resp

import "testing"

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
