package encrypt

import (
	"testing"
)

//Unit test
func TestHashPassword(t *testing.T) {
	out, _ := HashPassword("salam")

	if out == "" {
		t.Fatal("error hash password 1")
	}

	if out == "salam" {
		t.Log("error hash password 2")
		//t.FailNow()
	}
}

//Table-driven
// func TestCheckPasswordHash(t *testing.T) {
// 	//table test
// 	type Tcase struct {
// 		input1 string
// 		input2 string
// 	}

// 	Tcases := []Tcase{
// 		Tcase{ //True
// 			input1: "123456",
// 			input2: "$2a$14$m7kidDZOvh0M5J4kF8rCaOGvjTqrYt1jWJOpHQ0yJ/KD7Q/FhVk0.",
// 		},
// 		Tcase{ //Error
// 			input1: "12345",
// 			input2: "$2a$14$m7kidDZOvh0M5J4kF8rCaOGvjTqrYtfnhfghjfghfgh1jWJOpHQ0y",
// 		},
// 		Tcase{ //Error
// 			input1: "123456",
// 			input2: "123456",
// 		},
// 	}

// 	for _, c := range Tcases {
// 		result := CheckPasswordHash(c.input1, c.input2)

// 		if result != nil {
// 			t.Fatal("error Check Password Hash")
// 		}
// 	}
// }

//Subtests
func TestHashPasswordSub(t *testing.T) {
	t.Run("test Hash Password Sub 1", func(te *testing.T) {
		out1, _ := HashPassword("salam")
		if out1 == "" {
			te.Error("Error Hash Password 1")
		}
	})

	t.Run("test Hash Password Sub 2", func(te *testing.T) {
		out2, _ := HashPassword("salam")
		if out2 == "salam" {
			te.Error("Error Hash Password 2")
		}
	})
}
