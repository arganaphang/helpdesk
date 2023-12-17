package jewete

import (
	"reflect"
	"testing"
)

type Data struct {
	Name string `json:"name"`
}

func TestJWTDecrypt(t *testing.T) {
	user := Data{Name: "John Doe"}
	token, _ := JWTEncrypt(user)
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "Decrypt",
			args:    args{token: *token},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JWTDecrypt(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
