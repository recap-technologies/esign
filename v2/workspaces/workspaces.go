// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package workspaces implements the DocuSign SDK
// category Workspaces.
//
// Workspaces creation and management.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/Workspaces
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/v2/model"
//	)
//	...
//	workspacesService := workspaces.New(esignCredential)
package workspaces // import "github.com/jfcote87/esignv2/workspaces"

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign Workspaces API operations
type Service struct {
	credential esign.Credential
}

// New initializes a workspaces service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// ItemsCreateFIle creates a workspace file.
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/createfile
//
// SDK Method Workspaces::createWorkspaceFile
func (s *Service) ItemsCreateFIle(folderID string, workspaceID string, media io.Reader, mimeType string) *ItemsCreateFIleOp {
	return &ItemsCreateFIleOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID, "files"}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsCreateFIleOp implements DocuSign API SDK Workspaces::createWorkspaceFile
type ItemsCreateFIleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsCreateFIleOp) Do(ctx context.Context) (*model.WorkspaceItem, error) {
	var res *model.WorkspaceItem
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ItemsDeleteFolderItems deletes workspace one or more specific files/folders from the given folder or root.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/deletefolderitems
//
// SDK Method Workspaces::deleteWorkspaceFolderItems
func (s *Service) ItemsDeleteFolderItems(folderID string, workspaceID string, workspaceItemList *model.WorkspaceItemList) *ItemsDeleteFolderItemsOp {
	return &ItemsDeleteFolderItemsOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID}, "/"),
		Payload:    workspaceItemList,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsDeleteFolderItemsOp implements DocuSign API SDK Workspaces::deleteWorkspaceFolderItems
type ItemsDeleteFolderItemsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsDeleteFolderItemsOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// ItemsGetFile get Workspace File
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/getfile
//
// SDK Method Workspaces::getWorkspaceFile
func (s *Service) ItemsGetFile(fileID string, folderID string, workspaceID string) *ItemsGetFileOp {
	return &ItemsGetFileOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID, "files", fileID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsGetFileOp implements DocuSign API SDK Workspaces::getWorkspaceFile
type ItemsGetFileOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsGetFileOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// IsDownload when set to **true**, the Content-Disposition header is set in the response. The value of the header provides the filename of the file. Default is **false**.
func (op *ItemsGetFileOp) IsDownload() *ItemsGetFileOp {
	if op != nil {
		op.QueryOpts.Set("is_download", "true")
	}
	return op
}

// PdfVersion when set to **true** the file returned as a PDF.
func (op *ItemsGetFileOp) PdfVersion() *ItemsGetFileOp {
	if op != nil {
		op.QueryOpts.Set("pdf_version", "true")
	}
	return op
}

// ItemsListFilePages list File Pages
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/listfilepages
//
// SDK Method Workspaces::listWorkspaceFilePages
func (s *Service) ItemsListFilePages(fileID string, folderID string, workspaceID string) *ItemsListFilePagesOp {
	return &ItemsListFilePagesOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID, "files", fileID, "pages"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsListFilePagesOp implements DocuSign API SDK Workspaces::listWorkspaceFilePages
type ItemsListFilePagesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsListFilePagesOp) Do(ctx context.Context) (*model.PageImages, error) {
	var res *model.PageImages
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to be returned by this request.
func (op *ItemsListFilePagesOp) Count(val int) *ItemsListFilePagesOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// Dpi number of dots per inch for the resulting image. The default if not used is 94. The range is 1-310.
func (op *ItemsListFilePagesOp) Dpi(val int) *ItemsListFilePagesOp {
	if op != nil {
		op.QueryOpts.Set("dpi", fmt.Sprintf("%d", val))
	}
	return op
}

// MaxHeight sets the maximum height (in pixels) of the returned image.
func (op *ItemsListFilePagesOp) MaxHeight(val int) *ItemsListFilePagesOp {
	if op != nil {
		op.QueryOpts.Set("max_height", fmt.Sprintf("%d", val))
	}
	return op
}

// MaxWidth sets the maximum width (in pixels) of the returned image.
func (op *ItemsListFilePagesOp) MaxWidth(val int) *ItemsListFilePagesOp {
	if op != nil {
		op.QueryOpts.Set("max_width", fmt.Sprintf("%d", val))
	}
	return op
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ItemsListFilePagesOp) StartPosition(val int) *ItemsListFilePagesOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// ItemsListFolderItems list Workspace Folder Contents
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/listfolderitems
//
// SDK Method Workspaces::listWorkspaceFolderItems
func (s *Service) ItemsListFolderItems(folderID string, workspaceID string) *ItemsListFolderItemsOp {
	return &ItemsListFolderItemsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsListFolderItemsOp implements DocuSign API SDK Workspaces::listWorkspaceFolderItems
type ItemsListFolderItemsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsListFolderItemsOp) Do(ctx context.Context) (*model.WorkspaceFolderContents, error) {
	var res *model.WorkspaceFolderContents
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to be returned by this request.
func (op *ItemsListFolderItemsOp) Count(val int) *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// IncludeFiles when set to **true**, file information is returned in the response along with folder information. The default is **false**.
func (op *ItemsListFolderItemsOp) IncludeFiles() *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("include_files", "true")
	}
	return op
}

// IncludeSubFolders when set to **true**, information about the sub-folders of the current folder is returned. The default is **false**.
func (op *ItemsListFolderItemsOp) IncludeSubFolders() *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("include_sub_folders", "true")
	}
	return op
}

// IncludeThumbnails when set to **true**, thumbnails are returned as part of the response.  The default is **false**.
func (op *ItemsListFolderItemsOp) IncludeThumbnails() *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("include_thumbnails", "true")
	}
	return op
}

// IncludeUserDetail set to **true** to return extended details about the user. The default is **false**.
func (op *ItemsListFolderItemsOp) IncludeUserDetail() *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("include_user_detail", "true")
	}
	return op
}

// StartPosition is the position within the total result set from which to start returning values.
func (op *ItemsListFolderItemsOp) StartPosition(val int) *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// WorkspaceUserID if set, then the results are filtered to those associated with the specified userId.
func (op *ItemsListFolderItemsOp) WorkspaceUserID(val string) *ItemsListFolderItemsOp {
	if op != nil {
		op.QueryOpts.Set("workspace_user_id", val)
	}
	return op
}

// ItemsUpdateFile update Workspace File Metadata
// If media is an io.ReadCloser, Do() will close media.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaceitems/updatefile
//
// SDK Method Workspaces::updateWorkspaceFile
func (s *Service) ItemsUpdateFile(fileID string, folderID string, workspaceID string, media io.Reader, mimeType string) *ItemsUpdateFileOp {
	return &ItemsUpdateFileOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"workspaces", workspaceID, "folders", folderID, "files", fileID}, "/"),
		Payload:    &esign.UploadFile{Reader: media, ContentType: mimeType},
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ItemsUpdateFileOp implements DocuSign API SDK Workspaces::updateWorkspaceFile
type ItemsUpdateFileOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ItemsUpdateFileOp) Do(ctx context.Context) (*model.WorkspaceItem, error) {
	var res *model.WorkspaceItem
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Create create a Workspace
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaces/create
//
// SDK Method Workspaces::createWorkspace
func (s *Service) Create(workspace *model.Workspace) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "workspaces",
		Payload:    workspace,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CreateOp implements DocuSign API SDK Workspaces::createWorkspace
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.Workspace, error) {
	var res *model.Workspace
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete delete Workspace
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaces/delete
//
// SDK Method Workspaces::deleteWorkspace
func (s *Service) Delete(workspaceID string) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"workspaces", workspaceID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DeleteOp implements DocuSign API SDK Workspaces::deleteWorkspace
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) (*model.Workspace, error) {
	var res *model.Workspace
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Get get Workspace
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaces/get
//
// SDK Method Workspaces::getWorkspace
func (s *Service) Get(workspaceID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"workspaces", workspaceID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GetOp implements DocuSign API SDK Workspaces::getWorkspace
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.Workspace, error) {
	var res *model.Workspace
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// List list Workspaces
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaces/list
//
// SDK Method Workspaces::listWorkspaces
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "workspaces",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ListOp implements DocuSign API SDK Workspaces::listWorkspaces
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.WorkspaceList, error) {
	var res *model.WorkspaceList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Update update Workspace
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/workspaces/workspaces/update
//
// SDK Method Workspaces::updateWorkspace
func (s *Service) Update(workspaceID string, workspace *model.Workspace) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"workspaces", workspaceID}, "/"),
		Payload:    workspace,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateOp implements DocuSign API SDK Workspaces::updateWorkspace
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.Workspace, error) {
	var res *model.Workspace
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
