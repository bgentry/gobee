package api

import "testing"

func Test_ExplicitTx(t *testing.T) {
	cmd := ExplicitTxCommand{
		FrameID:         0xFF,
		Addr64:          0x0001020304050607,
		Addr16:          0x0001,
		SrcEp:           0x01,
		DstEp:           0x02,
		ClusterID:       0x1234,
		ProfileID:       0x5678,
		BroadcastRadius: 0xFF,
		Options:         0xEE,
		Data:            []byte{0x00, 0x01},
	}

	actual, err := cmd.Bytes()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expected := []byte{
		byte(apiTypeExplicitTx), 0xFF, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x00, 0x01, 0x01, 0x02, 0x12, 0x34,
		0x56, 0x78, 0xFF, 0xEE, 0x00, 0x01,
	}
	if len(expected) != len(actual) {
		t.Logf("expected: %#x\n\nactual: %#x\n", expected, actual)
		t.Fatalf("Expected ZB frame to be %d bytes in length, got: %d", len(expected), len(actual))
	}

	for i, b := range expected {
		if b != actual[i] {
			t.Errorf("Expected 0x%02x, but got 0x%02x", b, actual[i])
		}
	}
}
