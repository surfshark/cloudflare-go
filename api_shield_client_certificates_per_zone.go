package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// PerZoneAPIShieldClientCertificateAuthorityDetails represents the metadata for a per zone API Shield mTLS client certificates authority.
type PerZoneAPIShieldClientCertificateAuthorityDetails struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PerZoneAPIShieldClientCertificateDetails represents the metadata for a per zone API Shield mTLS client certificate.
type PerZoneAPIShieldClientCertificateDetails struct {
	Certificate          string                                            `json:"certificate"`
	CertificateAuthority PerZoneAPIShieldClientCertificateAuthorityDetails `json:"certificate_authority"`
	CommonName           string                                            `json:"common_name"`
	Country              string                                            `json:"country"`
	CSR                  string                                            `json:"csr"`
	ExpiresOn            time.Time                                         `json:"expires_on"`
	FingerprintSha256    string                                            `json:"fingerprint_sha256"`
	ID                   string                                            `json:"id"`
	IssuedOn             time.Time                                         `json:"issued_on"`
	Location             string                                            `json:"location"`
	Organization         string                                            `json:"organization"`
	OrganizationalUnit   string                                            `json:"organizational_unit"`
	SerialNumber         string                                            `json:"serial_number"`
	Signature            string                                            `json:"signature"`
	SKI                  string                                            `json:"ski"`
	State                string                                            `json:"state"`
	Status               string                                            `json:"status"`
	ValidityDays         int                                               `json:"validity_days"`
}

// PerZoneAPIShieldClientCertificatesResponse represents the response from the per zone API Shield mTLS client certificate list endpoint.
type PerZoneAPIShieldClientCertificatesResponse struct {
	Response
	Result []PerZoneAPIShieldClientCertificateDetails `json:"result"`
}

// ListPerZoneAPIShieldClientCertificates returns a list of API Shield mTLS client certificates per zone.
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-client-certificates
func (api *API) ListPerZoneAPIShieldClientCertificates(ctx context.Context, zoneID string) ([]PerZoneAPIShieldClientCertificateDetails, error) {
	uri := fmt.Sprintf("/zones/%s/client_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PerZoneAPIShieldClientCertificateDetails{}, err
	}
	var r PerZoneAPIShieldClientCertificatesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []PerZoneAPIShieldClientCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
