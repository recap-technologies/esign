// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package documents implements the DocuSign SDK
// category Documents.
//
// Information about documents.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Documents
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/rooms"
//	)
//	...
//	documentsService := documents.New(esignCredential)
package documents // import "github.com/jfcote87/esignrooms//documents"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign Documents API operations
type Service struct {
	credential esign.Credential
}

// New initializes a documents service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CreateDocumentUser grants a user access to a document.
//
// https://developers.docusign.com/docs/rooms-api/reference/documents/documents/createdocumentuser
//
// SDK Method Documents::CreateDocumentUser
func (s *Service) CreateDocumentUser(documentID string, body *rooms.DocumentUserForCreate) *CreateDocumentUserOp {
	return &CreateDocumentUserOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"documents", documentID, "users"}, "/"),
		Payload:    body,
		Accept:     "application/json-patch+json, application/json, text/json, application/*+json, application/xml, text/xml, application/*+xml",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// CreateDocumentUserOp implements DocuSign API SDK Documents::CreateDocumentUser
type CreateDocumentUserOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateDocumentUserOp) Do(ctx context.Context) (*rooms.DocumentUser, error) {
	var res *rooms.DocumentUser
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteDocument deletes a specified document.
//
// https://developers.docusign.com/docs/rooms-api/reference/documents/documents/deletedocument
//
// SDK Method Documents::DeleteDocument
func (s *Service) DeleteDocument(documentID string) *DeleteDocumentOp {
	return &DeleteDocumentOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"documents", documentID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// DeleteDocumentOp implements DocuSign API SDK Documents::DeleteDocument
type DeleteDocumentOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteDocumentOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetDocument gets information about or the contents of a document.
//
// https://developers.docusign.com/docs/rooms-api/reference/documents/documents/getdocument
//
// SDK Method Documents::GetDocument
func (s *Service) GetDocument(documentID string) *GetDocumentOp {
	return &GetDocumentOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"documents", documentID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetDocumentOp implements DocuSign API SDK Documents::GetDocument
type GetDocumentOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetDocumentOp) Do(ctx context.Context) (*rooms.Document, error) {
	var res *rooms.Document
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeContents when **true,** includes the contents of the document in the `base64Contents` property of the response. The default value is **false.**
func (op *GetDocumentOp) IncludeContents() *GetDocumentOp {
	if op != nil {
		op.QueryOpts.Set("includeContents", "true")
	}
	return op
}
