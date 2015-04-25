package zonedb

import (
	"testing"
	"unsafe"
)

func TestSizeofZone(t *testing.T) {
	var z Zone
	t.Logf("sizeof Zone = %d", unsafe.Sizeof(z))
}

func TestTLDs(t *testing.T) {
	t.Logf("%d top-level domains (%s to %s)", len(TLDs), TLDs[0].Domain, TLDs[len(TLDs)-1].Domain)
}

func TestTags(t *testing.T) {
	if numTags != len(TagStrings) {
		t.Errorf("numTags (%d) != len(TagStrings) (%d)", numTags, len(TagStrings))
	}
	if len(TagStrings) != len(TagValues) {
		t.Errorf("len(TagStrings) (%d) != len(TagValues) (%d)", len(TagStrings), len(TagValues))
	}
}

func TestZone_IsTLD(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"um":     true,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsTLD()
		if g != v {
			t.Errorf(`Expected Zones["%s"].IsTLD() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsDelegated(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  true,
		"org.za": true,
		"db.za":  false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsDelegated()
		if g != v {
			t.Errorf(`Expected Zones["%s"].IsDelegated() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsInRootZone(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"net":    true,
		"org":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsInRootZone()
		if g != v {
			t.Errorf(`Expected Zones["%s"].IsInRootZone() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_AllowsRegistration(t *testing.T) {
	data := map[string]bool{
		"com":          true,
		"net":          true,
		"org":          true,
		"ck":           false,
		"yu":           false,
		"arpa":         false,
		"cadillac":     false,
		"amazon":       false,
		"co.uk":        true,
		"in-addr.arpa": false,
	}
	for k, v := range data {
		g := ZoneMap[k].AllowsRegistration()
		if g != v {
			t.Errorf(`Expected Zones["%s"].AllowsRegistration() == %t, got %t`, k, v, g)
		}
	}
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
