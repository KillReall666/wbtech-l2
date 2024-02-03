package main

import (
	"testing"
)

func TestUnpackingString(t *testing.T) {
	tests := []struct {
		name string
		arg string
		want string
		wantErr error
	}{
		{
			name: "2 nums contract",
			arg: "d2b34pk",
			want: "d2b34pk",
			wantErr: ErrIncorrectString,
	}, {
			name: "test_1",
			arg: "a4bc2d5e",
			want: "aaaabccddddde",
			wantErr: nil,

	},{
			name: "test_2",
			arg: "abcd",
			want: "abcd",
			wantErr: nil,
	},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackingString(tt.arg)
			if err != tt.wantErr {
				t.Errorf("unpackingString() error = %v, want.err %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpackingString() = %v, want %v", got, tt.want)
			}
		})
	}

}