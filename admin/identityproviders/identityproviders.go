// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package identityproviders implements the DocuSign SDK
// category IdentityProviders.
//
// Methods to get a list of identity providers.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/admin-api/reference/IdentityProviders
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/admin"
//	)
//	...
//	identityprovidersService := identityproviders.New(esignCredential)
package identityproviders // import "github.com/jfcote87/esignadmin/identityproviders"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/admin"
)

// Service implements DocuSign IdentityProviders API operations
type Service struct {
	credential esign.Credential
}

// New initializes a identityproviders service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetIdentityProviders returns the list of identity providers for an organization.
//
// https://developers.docusign.com/docs/admin-api/reference/identityproviders/identityproviders/getidentityproviders
//
// SDK Method IdentityProviders::getIdentityProviders
func (s *Service) GetIdentityProviders(organizationID string) *GetIdentityProvidersOp {
	return &GetIdentityProvidersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"", "v2", "organizations", organizationID, "identity_providers"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.AdminV2,
	}
}

// GetIdentityProvidersOp implements DocuSign API SDK IdentityProviders::getIdentityProviders
type GetIdentityProvidersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetIdentityProvidersOp) Do(ctx context.Context) (*admin.IdentityProvidersResponse, error) {
	var res *admin.IdentityProvidersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
