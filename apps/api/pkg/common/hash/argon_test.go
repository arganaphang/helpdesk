package hash

import "testing"

func TestHashAndCompare(t *testing.T) {
	password := "password123"
	hash, _ := Hash(password)
	type args struct {
		hash     string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Hash and Compare",
			args: args{
				hash:     *hash,
				password: password,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.hash, tt.args.password); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
