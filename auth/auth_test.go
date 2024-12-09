package auth

import "testing"

func TestAuth(t *testing.T) {

	var tests = []struct {
		name  string
		input string
	}{
		{"test1", "mukuna"},
		{"test2", "john"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			hashedPassword, err := HashPassword(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if hashedPassword == ""{
				t.Errorf("password should not be empty")
			}
		})
	}
}


func TestComparePassword(t *testing.T)  {

	hashedPassword, err := HashPassword("mypassword")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("test1", func(t *testing.T) {

		if !ComparePassword(hashedPassword, []byte("mypassword")){
			t.Errorf("expected passwords to match")
		}

	})


	
}