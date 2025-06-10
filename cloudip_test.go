package cloudip

import (
	"net/netip"
	"testing"
)

func TestCloudip(t *testing.T) {
	checkIPs := []struct {
		IP       netip.Addr
		Provider Provider
	}{
		{
			IP:       netip.MustParseAddr("34.1.208.21"),
			Provider: GCP,
		},
		{
			IP:       netip.MustParseAddr("2600:1900:4280::1"),
			Provider: GCP,
		},
		{
			IP:       netip.MustParseAddr("103.21.244.21"),
			Provider: Cloudflare,
		},
		{
			IP:       netip.MustParseAddr("2405:8100::1"),
			Provider: Cloudflare,
		},
		{
			IP:       netip.MustParseAddr("54.74.0.27"),
			Provider: AWS,
		},
		{
			IP:       netip.MustParseAddr("2a05:d03a:8000::1"),
			Provider: AWS,
		},
		{
			IP:       netip.MustParseAddr("140.82.112.3"),
			Provider: GitHub,
		},
		{
			IP:       netip.MustParseAddr("2606:50c0:8000::153"),
			Provider: GitHub,
		},
	}

	for _, check := range checkIPs {
		got := Lookup(check.IP)
		if check.Provider == "" {
			if got != nil {
				t.Errorf("expected lookup to fail for %s but got %+v", check.IP, got)
			}
		}
		if got.Provider != check.Provider {
			t.Errorf("%s: expected provider %s but got %s", check.IP, check.Provider, got.Provider)
		}
	}
}
