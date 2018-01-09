package mac_wireless

import (
	"testing"
)

func TestLoadWirelessTable(t *testing.T) {
	// line of input
	input := `agrCtlRSSI: -42
     agrExtRSSI: 0
    agrCtlNoise: -92
    agrExtNoise: 0
          state: running
        op mode: station
     lastTxRate: 300
        maxRate: 450
lastAssocStatus: 0
    802.11 auth: open
      link auth: wpa2-psk
          BSSID: 5c:99:99:99:9:99
           SSID: Foo Bar
            MCS: 15
        channel: 157,1`
	// the headers we expect from that line of input

	// the map of data we expect.
	parsed := map[string]interface{}{
		"agrCtlRSSI":      int(-42),
		"agrExtRSSI":      int(0),
		"agrCtlNoise":     int(-92),
		"agrExtNoise":     int(0),
		"lastTxRate":      int(300),
		"maxRate":         int(450),
		"lastAssocStatus": int(0),
		"MCS":             int(15),
	}
	// load the table from the input.
	got, _, err := loadWirelessTable([]byte(input), false)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatalf("want %+v, got %+v", parsed, got)
	}
	for key := range parsed {
		if parsed[key].(int) != got[key].(int) {
			t.Fatalf("want %+v, got %+v", parsed[key], got[key])
		}
	}
}
