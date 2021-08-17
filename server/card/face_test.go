package card

import "testing"

func TestFaces(t *testing.T) {
	faces := Faces()

	if len(faces) != 13 {
		t.Fatalf("Expected 13 faces but got %v", len(faces))
	}
}

func TestFace_String(t *testing.T) {
	testFaces := []string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}

	faces := Faces()

	for i, face := range faces {
		if face.String() != testFaces[i] {
			t.Fatalf("Expected %s but got %s", testFaces[i], face.String())
		}
	}
}
