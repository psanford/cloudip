package cloudip

import (
	"net/netip"

	"github.com/psanford/awsip"
	"github.com/psanford/cloudflareip"
	"github.com/psanford/gcpip"
)

func Lookup(ip netip.Addr) *IPRange {
	if aip := awsip.Range(ip); aip != nil {
		ipr := IPRange{
			Provider:         AWS,
			Prefix:           aip.Prefix,
			Region:           aip.Region,
			Services:         aip.Services,
			ProviderSpecific: aip,
		}
		return &ipr
	}

	if gip := gcpip.Range(ip); gip != nil {
		ipr := IPRange{
			Provider:         GCP,
			Prefix:           gip.Prefix,
			Region:           gip.Scope,
			Services:         []string{gip.Service},
			ProviderSpecific: gip,
		}
		return &ipr
	}

	if cip := cloudflareip.Range(ip); cip != nil {
		ipr := IPRange{
			Provider:         Cloudflare,
			Prefix:           cip.Prefix,
			ProviderSpecific: cip,
		}
		return &ipr
	}

	return nil
}

type Provider string

const (
	AWS        Provider = "AWS"
	GCP        Provider = "GCP"
	Cloudflare Provider = "CLOUDFLARE"
)

type IPRange struct {
	Provider         Provider     `json:"provider"`
	Prefix           netip.Prefix `json:"prefix"`
	Region           string       `json:"region"`
	Services         []string     `json:"services"`
	ProviderSpecific any          `json:"provider_specific"`
}
