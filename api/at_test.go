package api

import "testing"

func Test_Valid_AT_No_Param(t *testing.T) {
	at := ATCommand{
		FrameID:   0x01,
		Type:      [2]byte{'N', 'I'},
		Parameter: nil,
	}

	actual, err := at.Bytes()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if len(actual) != 4 {
		t.Errorf("Expected AT frame to be 5 bytes in length, got: %d", len(actual))
	}

	expected := []byte{byte(apiTypeAT), 1, 'N', 'I'}
	for i, b := range expected {
		if b != actual[i] {
			t.Errorf("Expected 0x%02x, but got 0x%02x", b, actual[i])
		}
	}
}

func Test_Valid_AT_With_Param(t *testing.T) {
	at := ATCommand{
		FrameID:   0x01,
		Type:      [2]byte{'N', 'I'},
		Parameter: []byte{1},
	}

	actual, err := at.Bytes()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if len(actual) != 5 {
		t.Errorf("Expected AT frame to be 5 bytes in length, got: %d", len(actual))
	}

	expected := []byte{byte(apiTypeAT), 1, 'N', 'I', 1}
	for i, b := range expected {
		if b != actual[i] {
			t.Errorf("Expected 0x%02x, but got 0x%02x", b, actual[i])
		}
	}
}
