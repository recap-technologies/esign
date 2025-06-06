// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package click implements the DocuSign SDK
// category Click.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/click-api/reference/accounts
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	)
//	...
//	clickService := click.New(esignCredential)
package click // import "github.com/jfcote87/esign/click"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
)

// Service implements DocuSign Click API operations
type Service struct {
	credential esign.Credential
}

// New initializes a click service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CreateClickwrap creates a clickwrap for an account.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/createclickwrap
//
// SDK Method Click::createClickwrap
func (s *Service) CreateClickwrap(clickwrapRequest ClickwrapRequest) *CreateClickwrapOp {
	return &CreateClickwrapOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "clickwraps",
		Payload:    clickwrapRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// CreateClickwrapOp implements DocuSign API SDK Click::createClickwrap
type CreateClickwrapOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateClickwrapOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CreateClickwrapVersion creates a new clickwrap version.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/createclickwrapversion
//
// SDK Method Click::createClickwrapVersion
func (s *Service) CreateClickwrapVersion(clickwrapID string, clickwrapRequest ClickwrapRequest) *CreateClickwrapVersionOp {
	return &CreateClickwrapVersionOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions"}, "/"),
		Payload:    clickwrapRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// CreateClickwrapVersionOp implements DocuSign API SDK Click::createClickwrapVersion
type CreateClickwrapVersionOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateClickwrapVersionOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CreateHasAgreed checks if a user has agreed to a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/createhasagreed
//
// SDK Method Click::createHasAgreed
func (s *Service) CreateHasAgreed(clickwrapID string, userAgreementRequest UserAgreementRequest) *CreateHasAgreedOp {
	return &CreateHasAgreedOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "agreements"}, "/"),
		Payload:    userAgreementRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// CreateHasAgreedOp implements DocuSign API SDK Click::createHasAgreed
type CreateHasAgreedOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateHasAgreedOp) Do(ctx context.Context) (UserAgreementResponse, error) {
	var res UserAgreementResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteClickwrap deletes a clickwrap and all of its versions.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/deleteclickwrap
//
// SDK Method Click::deleteClickwrap
func (s *Service) DeleteClickwrap(clickwrapID string) *DeleteClickwrapOp {
	return &DeleteClickwrapOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"clickwraps", clickwrapID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// DeleteClickwrapOp implements DocuSign API SDK Click::deleteClickwrap
type DeleteClickwrapOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteClickwrapOp) Do(ctx context.Context) (ClickwrapVersionsDeleteResponse, error) {
	var res ClickwrapVersionsDeleteResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Versions is a comma-separated list of versions to delete.
func (op *DeleteClickwrapOp) Versions(val string) *DeleteClickwrapOp {
	if op != nil {
		op.QueryOpts.Set("versions", val)
	}
	return op
}

// DeleteClickwrapVersion deletes a specific version of a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/deleteclickwrapversion
//
// SDK Method Click::deleteClickwrapVersion
func (s *Service) DeleteClickwrapVersion(clickwrapID string, versionID string) *DeleteClickwrapVersionOp {
	return &DeleteClickwrapVersionOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// DeleteClickwrapVersionOp implements DocuSign API SDK Click::deleteClickwrapVersion
type DeleteClickwrapVersionOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteClickwrapVersionOp) Do(ctx context.Context) (ClickwrapVersionDeleteResponse, error) {
	var res ClickwrapVersionDeleteResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteClickwrapVersionByNumber deletes a clickwrap version from a clickwrap by specifying its version number.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/deleteclickwrapversionbynumber
//
// SDK Method Click::deleteClickwrapVersionByNumber
func (s *Service) DeleteClickwrapVersionByNumber(clickwrapID string, versionNumber string) *DeleteClickwrapVersionByNumberOp {
	return &DeleteClickwrapVersionByNumberOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionNumber}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// DeleteClickwrapVersionByNumberOp implements DocuSign API SDK Click::deleteClickwrapVersionByNumber
type DeleteClickwrapVersionByNumberOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteClickwrapVersionByNumberOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteClickwrapVersions deletes the versions of a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/deleteclickwrapversions
//
// SDK Method Click::deleteClickwrapVersions
func (s *Service) DeleteClickwrapVersions(clickwrapID string) *DeleteClickwrapVersionsOp {
	return &DeleteClickwrapVersionsOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// DeleteClickwrapVersionsOp implements DocuSign API SDK Click::deleteClickwrapVersions
type DeleteClickwrapVersionsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteClickwrapVersionsOp) Do(ctx context.Context) (ClickwrapVersionsDeleteResponse, error) {
	var res ClickwrapVersionsDeleteResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ClickwrapVersionIds is a comma-separated list of clickwrap version IDs to delete.
func (op *DeleteClickwrapVersionsOp) ClickwrapVersionIds(val string) *DeleteClickwrapVersionsOp {
	if op != nil {
		op.QueryOpts.Set("clickwrapVersionIds", val)
	}
	return op
}

// DeleteClickwraps deletes clickwraps for an account.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/deleteclickwraps
//
// SDK Method Click::deleteClickwraps
func (s *Service) DeleteClickwraps() *DeleteClickwrapsOp {
	return &DeleteClickwrapsOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "clickwraps",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// DeleteClickwrapsOp implements DocuSign API SDK Click::deleteClickwraps
type DeleteClickwrapsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteClickwrapsOp) Do(ctx context.Context) (ClickwrapsDeleteResponse, error) {
	var res ClickwrapsDeleteResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ClickwrapIds is a comma-separated list of clickwrap IDs to delete.
func (op *DeleteClickwrapsOp) ClickwrapIds(val string) *DeleteClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("clickwrapIds", val)
	}
	return op
}

// GetAgreement gets a specific agreement for a specified clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getagreement
//
// SDK Method Click::getAgreement
func (s *Service) GetAgreement(agreementID string, clickwrapID string) *GetAgreementOp {
	return &GetAgreementOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "agreements", agreementID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetAgreementOp implements DocuSign API SDK Click::getAgreement
type GetAgreementOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetAgreementOp) Do(ctx context.Context) (UserAgreementResponse, error) {
	var res UserAgreementResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetAgreementPdf gets the completed user agreement PDF.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getagreementpdf
//
// SDK Method Click::getAgreementPdf
func (s *Service) GetAgreementPdf(agreementID string, clickwrapID string) *GetAgreementPdfOp {
	return &GetAgreementPdfOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "agreements", agreementID, "download"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetAgreementPdfOp implements DocuSign API SDK Click::getAgreementPdf
type GetAgreementPdfOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetAgreementPdfOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetClickwrap gets a  single clickwrap object.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrap
//
// SDK Method Click::getClickwrap
func (s *Service) GetClickwrap(clickwrapID string) *GetClickwrapOp {
	return &GetClickwrapOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapOp implements DocuSign API SDK Click::getClickwrap
type GetClickwrapOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapOp) Do(ctx context.Context) (ClickwrapVersionResponse, error) {
	var res ClickwrapVersionResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetClickwrapAgreements get user agreements
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrapagreements
//
// SDK Method Click::getClickwrapAgreements
func (s *Service) GetClickwrapAgreements(clickwrapID string) *GetClickwrapAgreementsOp {
	return &GetClickwrapAgreementsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapAgreementsOp implements DocuSign API SDK Click::getClickwrapAgreements
type GetClickwrapAgreementsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapAgreementsOp) Do(ctx context.Context) (ClickwrapAgreementsResponse, error) {
	var res ClickwrapAgreementsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ClientUserID is the client ID.
func (op *GetClickwrapAgreementsOp) ClientUserID(val string) *GetClickwrapAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("client_user_id", val)
	}
	return op
}

// FromDate optional. The earliest date to return agreements from.
func (op *GetClickwrapAgreementsOp) FromDate(val string) *GetClickwrapAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val)
	}
	return op
}

// PageNumber optional. The page number to return.
func (op *GetClickwrapAgreementsOp) PageNumber(val string) *GetClickwrapAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("page_number", val)
	}
	return op
}

// Status optional. The status of the clickwraps to return.
func (op *GetClickwrapAgreementsOp) Status(val string) *GetClickwrapAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// ToDate optional. The latest date to return agreements from.
func (op *GetClickwrapAgreementsOp) ToDate(val string) *GetClickwrapAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val)
	}
	return op
}

// GetClickwrapVersion gets a specific version from a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrapversion
//
// SDK Method Click::getClickwrapVersion
func (s *Service) GetClickwrapVersion(clickwrapID string, versionID string) *GetClickwrapVersionOp {
	return &GetClickwrapVersionOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapVersionOp implements DocuSign API SDK Click::getClickwrapVersion
type GetClickwrapVersionOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapVersionOp) Do(ctx context.Context) (ClickwrapVersionResponse, error) {
	var res ClickwrapVersionResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetClickwrapVersionAgreements gets the agreement responses for a clickwrap version.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrapversionagreements
//
// SDK Method Click::getClickwrapVersionAgreements
func (s *Service) GetClickwrapVersionAgreements(clickwrapID string, versionID string) *GetClickwrapVersionAgreementsOp {
	return &GetClickwrapVersionAgreementsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionID, "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapVersionAgreementsOp implements DocuSign API SDK Click::getClickwrapVersionAgreements
type GetClickwrapVersionAgreementsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapVersionAgreementsOp) Do(ctx context.Context) (ClickwrapAgreementsResponse, error) {
	var res ClickwrapAgreementsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ClientUserID set the call query parameter client_user_id
func (op *GetClickwrapVersionAgreementsOp) ClientUserID(val string) *GetClickwrapVersionAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("client_user_id", val)
	}
	return op
}

// FromDate optional. The earliest date to return agreements from.
func (op *GetClickwrapVersionAgreementsOp) FromDate(val string) *GetClickwrapVersionAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val)
	}
	return op
}

// PageNumber optional. The page number to return.
func (op *GetClickwrapVersionAgreementsOp) PageNumber(val string) *GetClickwrapVersionAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("page_number", val)
	}
	return op
}

// Status clickwrap status. Possible values:
//
// - `active`
// - `inactive`
// - `deleted`
func (op *GetClickwrapVersionAgreementsOp) Status(val string) *GetClickwrapVersionAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// ToDate optional. The latest date to return agreements from.
func (op *GetClickwrapVersionAgreementsOp) ToDate(val string) *GetClickwrapVersionAgreementsOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val)
	}
	return op
}

// GetClickwrapVersionAgreementsByNumber gets the agreement responses for a clickwrap version
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrapversionagreementsbynumber
//
// SDK Method Click::getClickwrapVersionAgreementsByNumber
func (s *Service) GetClickwrapVersionAgreementsByNumber(clickwrapID string, versionNumber string) *GetClickwrapVersionAgreementsByNumberOp {
	return &GetClickwrapVersionAgreementsByNumberOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionNumber, "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapVersionAgreementsByNumberOp implements DocuSign API SDK Click::getClickwrapVersionAgreementsByNumber
type GetClickwrapVersionAgreementsByNumberOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapVersionAgreementsByNumberOp) Do(ctx context.Context) (ClickwrapAgreementsResponse, error) {
	var res ClickwrapAgreementsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ClientUserID is the client user ID.
func (op *GetClickwrapVersionAgreementsByNumberOp) ClientUserID(val string) *GetClickwrapVersionAgreementsByNumberOp {
	if op != nil {
		op.QueryOpts.Set("client_user_id", val)
	}
	return op
}

// FromDate optional. The earliest date to return agreements from.
func (op *GetClickwrapVersionAgreementsByNumberOp) FromDate(val string) *GetClickwrapVersionAgreementsByNumberOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val)
	}
	return op
}

// PageNumber optional. The page number to return.
func (op *GetClickwrapVersionAgreementsByNumberOp) PageNumber(val string) *GetClickwrapVersionAgreementsByNumberOp {
	if op != nil {
		op.QueryOpts.Set("page_number", val)
	}
	return op
}

// Status clickwrap status. Possible values:
//
// - `active`
// - `inactive`
// - `deleted`
func (op *GetClickwrapVersionAgreementsByNumberOp) Status(val string) *GetClickwrapVersionAgreementsByNumberOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// ToDate optional. The latest date to return agreements from.
func (op *GetClickwrapVersionAgreementsByNumberOp) ToDate(val string) *GetClickwrapVersionAgreementsByNumberOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val)
	}
	return op
}

// GetClickwrapVersionByNumber gets a clickwrap version by specifying its version number.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwrapversionbynumber
//
// SDK Method Click::getClickwrapVersionByNumber
func (s *Service) GetClickwrapVersionByNumber(clickwrapID string, versionNumber string) *GetClickwrapVersionByNumberOp {
	return &GetClickwrapVersionByNumberOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionNumber}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapVersionByNumberOp implements DocuSign API SDK Click::getClickwrapVersionByNumber
type GetClickwrapVersionByNumberOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapVersionByNumberOp) Do(ctx context.Context) (ClickwrapVersionResponse, error) {
	var res ClickwrapVersionResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GetClickwraps gets all the clickwraps for an account.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getclickwraps
//
// SDK Method Click::getClickwraps
func (s *Service) GetClickwraps() *GetClickwrapsOp {
	return &GetClickwrapsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "clickwraps",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetClickwrapsOp implements DocuSign API SDK Click::getClickwraps
type GetClickwrapsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetClickwrapsOp) Do(ctx context.Context) (ClickwrapVersionsResponse, error) {
	var res ClickwrapVersionsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate optional. The earliest date to return agreements from.
func (op *GetClickwrapsOp) FromDate(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val)
	}
	return op
}

// OwnerUserID optional. The user ID of the owner.
func (op *GetClickwrapsOp) OwnerUserID(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("ownerUserId", val)
	}
	return op
}

// PageNumber optional. The page number to return.
func (op *GetClickwrapsOp) PageNumber(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("page_number", val)
	}
	return op
}

// Shared set the call query parameter shared
func (op *GetClickwrapsOp) Shared(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("shared", val)
	}
	return op
}

// Status optional. The status of the clickwraps to filter by. One of:
//
// - `active`
// - `inactive`
func (op *GetClickwrapsOp) Status(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("status", val)
	}
	return op
}

// ToDate optional. The latest date to return agreements from.
func (op *GetClickwrapsOp) ToDate(val string) *GetClickwrapsOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val)
	}
	return op
}

// GetServiceInformation gets the current version and other information about the Click API.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/getserviceinformation
//
// SDK Method Click::getServiceInformation
func (s *Service) GetServiceInformation() *GetServiceInformationOp {
	return &GetServiceInformationOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/service_information",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// GetServiceInformationOp implements DocuSign API SDK Click::getServiceInformation
type GetServiceInformationOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetServiceInformationOp) Do(ctx context.Context) (ServiceInformation, error) {
	var res ServiceInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateClickwrap updates the user ID of a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/updateclickwrap
//
// SDK Method Click::updateClickwrap
func (s *Service) UpdateClickwrap(clickwrapID string, clickwrapTransferRequest ClickwrapTransferRequest) *UpdateClickwrapOp {
	return &UpdateClickwrapOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"clickwraps", clickwrapID}, "/"),
		Payload:    clickwrapTransferRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// UpdateClickwrapOp implements DocuSign API SDK Click::updateClickwrap
type UpdateClickwrapOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateClickwrapOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateClickwrapVersion updates a specific version of a clickwrap.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/updateclickwrapversion
//
// SDK Method Click::updateClickwrapVersion
func (s *Service) UpdateClickwrapVersion(clickwrapID string, versionID string, clickwrapRequest ClickwrapRequest) *UpdateClickwrapVersionOp {
	return &UpdateClickwrapVersionOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionID}, "/"),
		Payload:    clickwrapRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// UpdateClickwrapVersionOp implements DocuSign API SDK Click::updateClickwrapVersion
type UpdateClickwrapVersionOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateClickwrapVersionOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// UpdateClickwrapVersionByNumber updates a clickwrap version by specifying its version number.
//
// https://developers.docusign.com/docs/click-api/reference/accounts/clickwraps/updateclickwrapversionbynumber
//
// SDK Method Click::updateClickwrapVersionByNumber
func (s *Service) UpdateClickwrapVersionByNumber(clickwrapID string, versionNumber string, clickwrapRequest ClickwrapRequest) *UpdateClickwrapVersionByNumberOp {
	return &UpdateClickwrapVersionByNumberOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"clickwraps", clickwrapID, "versions", versionNumber}, "/"),
		Payload:    clickwrapRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.ClickV1,
	}
}

// UpdateClickwrapVersionByNumberOp implements DocuSign API SDK Click::updateClickwrapVersionByNumber
type UpdateClickwrapVersionByNumberOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateClickwrapVersionByNumberOp) Do(ctx context.Context) (ClickwrapVersionSummaryResponse, error) {
	var res ClickwrapVersionSummaryResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
