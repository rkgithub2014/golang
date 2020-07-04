package event

import "testing"

func TestInitConfig(t *testing.T) {
	gogPalos := InitConfig()
	if !gogPalos.Enabled {
		t.Errorf("JSON UnMarshalling failed for boolean value")
	}
	if gogPalos.PortStart != 50001 {
		t.Errorf("JSON UnMarshalling failed for int value")
	}
}
