package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type PerZoneAPIShieldClientCertificateAuthorityDetails struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type PerZoneAPIShieldClientCertificatesResponse struct {
	Response
	Result []PerZoneAPIShieldClientCertificateDetails `json:"result"`
}

type PerZoneAPIShieldClientCertificateResponse struct {
	Response
	Result PerZoneAPIShieldClientCertificateDetails `json:"result"`
}

type ListPerZoneAPIShieldClientCertificatesParams struct {
	Status string `url:"status,omitempty"`
}

type CreatePerZoneAPIShieldClientCertificateParams struct {
	CSR          string `json:"csr"`
	ValidityDays int    `json:"validity_days"`
}

// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-client-certificates
func (api *API) ListPerZoneAPIShieldClientCertificates(ctx context.Context, zoneID string, params ListPerZoneAPIShieldClientCertificatesParams) ([]PerZoneAPIShieldClientCertificateDetails, error) {
	uri := buildURI(fmt.Sprintf("/zones/%s/client_certificates", zoneID), params)
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

// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-create-client-certificate
func (api *API) CreatePerZoneAPIShieldClientCertificate(ctx context.Context, zoneID string, params CreatePerZoneAPIShieldClientCertificateParams) (PerZoneAPIShieldClientCertificateDetails, error) {
	uri := fmt.Sprintf("/zones/%s/client_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return PerZoneAPIShieldClientCertificateDetails{}, err
	}
	var r PerZoneAPIShieldClientCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAPIShieldClientCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-client-certificate-details
func (api *API) GetPerZoneAPIShieldClientCertificateDetails(ctx context.Context, zoneID, certificateID string) (PerZoneAPIShieldClientCertificateDetails, error) {
	uri := fmt.Sprintf("/zones/%s/client_certificates/%s", zoneID, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PerZoneAPIShieldClientCertificateDetails{}, err
	}
	var r PerZoneAPIShieldClientCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAPIShieldClientCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
