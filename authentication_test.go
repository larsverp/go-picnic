package picnic

import (
	"testing"
)

// TestgetMD5Hash calls authentication.getMD5Hash and checks the returned value
func TestGetMD5Hash(t *testing.T) {
	want := "e5daaa90c369adfd156862d6df632ded"
	data := getMD5Hash("qwertyuiopasdfghjklzxcvbnm")
	if data != want {
		t.Errorf("getMD5Hash() = %v, want %v", data, want)
	}
}

// TestNewUser calls authentication.NewUser and checks the returned value
func TestNewUser(t *testing.T) {
	user := NewUser("username", "password")
	want := User{Key: "username", Secret: "5f4dcc3b5aa765d61d8327deb882cf99", Client_id: 1}
	if user != want {
		t.Errorf("NewUser() = %v, want %v", user, want)
	}
}
