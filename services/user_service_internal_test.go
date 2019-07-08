package services

import "testing"

func TestCompareHashAndPassword(t *testing.T) {
	cases := []struct {
		password       string
		hashedPassword string
	}{
		{
			"testpassword",
			"$2a$10$FJjWKtgaJkZPEe/UHJEpg./YsQLBOjT./mA969guhJi.yrA3J9zl.",
		},
	}

	for _, tc := range cases {
		same := new(userService).isHashAndPasswordSame(tc.hashedPassword, tc.password)
		if same == false {
			t.Error("Error: password when hashed should be equivalent to hashed password, found not the same")
		}
	}
}
