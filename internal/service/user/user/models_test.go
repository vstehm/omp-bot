package user

import (
	"github.com/ozonmp/omp-bot/internal/model/user"
	"testing"
)

func TestUser_String(t *testing.T) {
	type fields struct {
		Id        uint64
		Lastname  string
		Firstname string
		Phone     string
		Email     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple",
			fields: fields{
				Id:        uint64(1),
				Lastname:  "user_name",
				Firstname: "user_first_name",
				Phone:     "+70000000000",
				Email:     "demo@demo.com",
			},
			want: "id: 1\nlastname: user_name\nfirstname: user_first_name\nphone: +70000000000\nemail: demo@demo.com\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := user.User{
				Id:        test.fields.Id,
				Lastname:  test.fields.Lastname,
				Firstname: test.fields.Firstname,
				Phone:     test.fields.Phone,
				Email:     test.fields.Email,
			}
			if got := u.String(); got != test.want {
				t.Errorf("String() = %v, want %v", got, test.want)
			}
		})
	}
}
