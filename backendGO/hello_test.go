package main

// import "testing"

// func TestHello(t *testing.T) {
// 	got := hello("Shahroz")
// 	want := "Hello, Shahroz"

// 	if got != want {
// 		t.Errorf("Failed")
// 	}
// }

// func TestIsPasswordValid(t *testing.T) {
// 	assertCorrectMessage := func(t testing.TB, got, want bool) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %t want %t", got, want)
// 		}
// 	}

// 	t.Run("Running for character less than 8", func(t *testing.T) {
// 		got := isPasswordValid("abcde7")
// 		want := false
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for character greater than 8", func(t *testing.T) {
// 		got := isPasswordValid("abcde87")
// 		want := true
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for wihout special character", func(t *testing.T) {
// 		got := isPasswordValid("abcde7")
// 		want := false
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for without digit", func(t *testing.T) {
// 		got := isPasswordValid("abcde@@@")
// 		want := false
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for character less than 8", func(t *testing.T) {
// 		got := isPasswordValid("abcde7")
// 		want := false
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for valid password", func(t *testing.T) {
// 		got := isPasswordValid("abcdhe@123#66")
// 		want := true
// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("Running for password with space", func(t *testing.T) {
// 		got := isPasswordValid("abc de7")
// 		want := false
// 		assertCorrectMessage(t, got, want)
// 	})
// }
