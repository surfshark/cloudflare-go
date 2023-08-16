package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListPerZoneAPIShieldClientCertificates(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
			  {
				"certificate": "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
				"certificate_authority": {
					"id": "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
					"name": "Cloudflare Managed CA for account"
				},
				"common_name": "Cloudflare",
				"country": "US",
				"csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
				"expires_on": "2033-02-20T23:18:00Z",
				"fingerprint_sha256": "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
				"id": "023e105f4ecef8ad9ca31a8372d0c353",
				"issued_on": "2023-02-23T23:18:00Z",
				"location": "Somewhere",
				"organization": "Organization",
				"organizational_unit": "Organizational Unit",
				"serial_number": "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
				"signature": "SHA256WithRSA",
				"ski": "8e375af1389a069a0f921f8cc8e1eb12d784b949",
				"state": "CA",
				"status": "active",
				"validity_days": 3650
			  }
			]
		  }
		`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/client_certificates", handler)
	expiresOn, _ := time.Parse(time.RFC3339, "2033-02-20T23:18:00Z")
	issuedOn, _ := time.Parse(time.RFC3339, "2023-02-23T23:18:00Z")
	want := []PerZoneAPIShieldClientCertificateDetails{
		{
			Certificate: "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoC...dhDDE\n-----END CERTIFICATE-----",
			CertificateAuthority: PerZoneAPIShieldClientCertificateAuthorityDetails{
				ID:   "568b6b74-7b0c-4755-8840-4e3b8c24adeb",
				Name: "Cloudflare Managed CA for account",
			},
			CommonName:         "Cloudflare",
			Country:            "US",
			CSR:                "-----BEGIN CERTIFICATE REQUEST-----\nMIICY....\n-----END CERTIFICATE REQUEST-----\n",
			ExpiresOn:          expiresOn,
			FingerprintSha256:  "256c24690243359fb8cf139a125bd05ebf1d968b71e4caf330718e9f5c8a89ea",
			ID:                 "023e105f4ecef8ad9ca31a8372d0c353",
			IssuedOn:           issuedOn,
			Location:           "Somewhere",
			Organization:       "Organization",
			OrganizationalUnit: "Organizational Unit",
			SerialNumber:       "3bb94ff144ac567b9f75ad664b6c55f8d5e48182",
			Signature:          "SHA256WithRSA",
			SKI:                "8e375af1389a069a0f921f8cc8e1eb12d784b949",
			State:              "CA",
			Status:             "active",
			ValidityDays:       3650,
		},
	}
	actual, err := client.ListPerZoneAPIShieldClientCertificates(context.Background(), "023e105f4ecef8ad9ca31a8372d0c353")
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
