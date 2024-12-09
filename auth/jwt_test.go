package auth

import "testing"

func TestJWT(t *testing.T) {
	var tests = []struct{
		name string
		input []byte
		userID int
	}{
		{"test1", []byte("secret"), 3 },
		{"test2", []byte("notsecret"), 37 },
		{"test3", []byte("password"), 56 },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CreateJWT(tt.input, tt.userID)
			if err != nil {
				t.Fatal(err)
			}

			if token == "" {
				t.Errorf("expected a token")
			}
		})
	}
}