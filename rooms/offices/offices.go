// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package offices implements the DocuSign SDK
// category Offices.
//
// This section shows you how to create and manage offices.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Offices
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/rooms"
//	)
//	...
//	officesService := offices.New(esignCredential)
package offices // import "github.com/jfcote87/esignrooms//offices"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign Offices API operations
type Service struct {
	credential esign.Credential
}

// New initializes a offices service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CreateOffice creates an office.
//
// https://developers.docusign.com/docs/rooms-api/reference/offices/offices/createoffice
//
// SDK Method Offices::CreateOffice
func (s *Service) CreateOffice(body *rooms.OfficeForCreate) *CreateOfficeOp {
	return &CreateOfficeOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "offices",
		Payload:    body,
		Accept:     "application/json-patch+json, application/json, text/json, application/*+json, application/xml, text/xml, application/*+xml",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// CreateOfficeOp implements DocuSign API SDK Offices::CreateOffice
type CreateOfficeOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOfficeOp) Do(ctx context.Context) (*rooms.Office, error) {
	var res *rooms.Office
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteOffice deletes an office.
//
// https://developers.docusign.com/docs/rooms-api/reference/offices/offices/deleteoffice
//
// SDK Method Offices::DeleteOffice
func (s *Service) DeleteOffice(officeID string) *DeleteOfficeOp {
	return &DeleteOfficeOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"offices", officeID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// DeleteOfficeOp implements DocuSign API SDK Offices::DeleteOffice
type DeleteOfficeOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOfficeOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetOffice gets information about an office.
//
// https://developers.docusign.com/docs/rooms-api/reference/offices/offices/getoffice
//
// SDK Method Offices::GetOffice
func (s *Service) GetOffice(officeID string) *GetOfficeOp {
	return &GetOfficeOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"offices", officeID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetOfficeOp implements DocuSign API SDK Offices::GetOffice
type GetOfficeOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOfficeOp) Do(ctx context.Context) (*rooms.Office, error) {
	var res *rooms.Office
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetOffices gets offices.
//
// https://developers.docusign.com/docs/rooms-api/reference/offices/offices/getoffices
//
// SDK Method Offices::GetOffices
func (s *Service) GetOffices() *GetOfficesOp {
	return &GetOfficesOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "offices",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetOfficesOp implements DocuSign API SDK Offices::GetOffices
type GetOfficesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOfficesOp) Do(ctx context.Context) (*rooms.OfficeSummaryList, error) {
	var res *rooms.OfficeSummaryList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetOfficesOp) Count(val int) *GetOfficesOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// StartPosition is the starting zero-based index position of the results set from which to begin returning values. The default value is `0`.
func (op *GetOfficesOp) StartPosition(val int) *GetOfficesOp {
	if op != nil {
		op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val))
	}
	return op
}

// OnlyAccessible when **true,** the response only includes the offices that are accessible to the current user.
func (op *GetOfficesOp) OnlyAccessible() *GetOfficesOp {
	if op != nil {
		op.QueryOpts.Set("onlyAccessible", "true")
	}
	return op
}

// Search filters returned records by the specified string. The response only includes records containing this string in the office `name` field.
func (op *GetOfficesOp) Search(val string) *GetOfficesOp {
	if op != nil {
		op.QueryOpts.Set("search", val)
	}
	return op
}

// GetReferenceCounts retrieves the number and type of objects that reference an office.
//
// https://developers.docusign.com/docs/rooms-api/reference/offices/offices/getreferencecounts
//
// SDK Method Offices::GetReferenceCounts
func (s *Service) GetReferenceCounts(officeID string) *GetReferenceCountsOp {
	return &GetReferenceCountsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"offices", officeID, "reference_counts"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetReferenceCountsOp implements DocuSign API SDK Offices::GetReferenceCounts
type GetReferenceCountsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetReferenceCountsOp) Do(ctx context.Context) (*rooms.OfficeReferenceCountList, error) {
	var res *rooms.OfficeReferenceCountList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
