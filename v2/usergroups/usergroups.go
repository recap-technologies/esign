// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package usergroups implements the DocuSign SDK
// category UserGroups.
//
// Use the User Groups category to manage your permissions groups.
//
// You can:
//
// * Create and delete groups.
// * Add users to, and delete users from, your groups.
// * Manage the brand information associated with a group.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/UserGroups
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/v2/model"
//	)
//	...
//	usergroupsService := usergroups.New(esignCredential)
package usergroups // import "github.com/jfcote87/esignv2/usergroups"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign UserGroups API operations
type Service struct {
	credential esign.Credential
}

// New initializes a usergroups service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GroupBrandsDelete deletes brand information from the requested group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupbrands/delete
//
// SDK Method UserGroups::deleteBrands
func (s *Service) GroupBrandsDelete(groupID string, brandsRequest *model.BrandsRequest) *GroupBrandsDeleteOp {
	return &GroupBrandsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"groups", groupID, "brands"}, "/"),
		Payload:    brandsRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupBrandsDeleteOp implements DocuSign API SDK UserGroups::deleteBrands
type GroupBrandsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupBrandsDeleteOp) Do(ctx context.Context) (*model.BrandsResponse, error) {
	var res *model.BrandsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupBrandsGet gets group brand ID Information.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupbrands/get
//
// SDK Method UserGroups::getBrands
func (s *Service) GroupBrandsGet(groupID string) *GroupBrandsGetOp {
	return &GroupBrandsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"groups", groupID, "brands"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupBrandsGetOp implements DocuSign API SDK UserGroups::getBrands
type GroupBrandsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupBrandsGetOp) Do(ctx context.Context) (*model.BrandsResponse, error) {
	var res *model.BrandsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupBrandsUpdate adds group brand ID information to a group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupbrands/update
//
// SDK Method UserGroups::updateBrands
func (s *Service) GroupBrandsUpdate(groupID string, brandsRequest *model.BrandsRequest) *GroupBrandsUpdateOp {
	return &GroupBrandsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"groups", groupID, "brands"}, "/"),
		Payload:    brandsRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupBrandsUpdateOp implements DocuSign API SDK UserGroups::updateBrands
type GroupBrandsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupBrandsUpdateOp) Do(ctx context.Context) (*model.BrandsResponse, error) {
	var res *model.BrandsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupUsersDelete deletes one or more users from a gro
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupusers/delete
//
// SDK Method UserGroups::deleteGroupUsers
func (s *Service) GroupUsersDelete(groupID string, userInfoList *model.UserInfoList) *GroupUsersDeleteOp {
	return &GroupUsersDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"groups", groupID, "users"}, "/"),
		Payload:    userInfoList,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupUsersDeleteOp implements DocuSign API SDK UserGroups::deleteGroupUsers
type GroupUsersDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupUsersDeleteOp) Do(ctx context.Context) (*model.UsersResponse, error) {
	var res *model.UsersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupUsersList gets a list of users in a group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupusers/list
//
// SDK Method UserGroups::listGroupUsers
func (s *Service) GroupUsersList(groupID string) *GroupUsersListOp {
	return &GroupUsersListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"groups", groupID, "users"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupUsersListOp implements DocuSign API SDK UserGroups::listGroupUsers
type GroupUsersListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupUsersListOp) Do(ctx context.Context) (*model.UsersResponse, error) {
	var res *model.UsersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count number of records to return. The number must be greater than 1 and less than or equal to 100.
func (op *GroupUsersListOp) Count(val int) *GroupUsersListOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// StartPosition starting value for the list.
func (op *GroupUsersListOp) StartPosition(val int) *GroupUsersListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// GroupUsersUpdate adds one or more users to an existing group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groupusers/update
//
// SDK Method UserGroups::updateGroupUsers
func (s *Service) GroupUsersUpdate(groupID string, userInfoList *model.UserInfoList) *GroupUsersUpdateOp {
	return &GroupUsersUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"groups", groupID, "users"}, "/"),
		Payload:    userInfoList,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupUsersUpdateOp implements DocuSign API SDK UserGroups::updateGroupUsers
type GroupUsersUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupUsersUpdateOp) Do(ctx context.Context) (*model.UsersResponse, error) {
	var res *model.UsersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupsCreate creates one or more groups for the account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groups/create
//
// SDK Method UserGroups::createGroups
func (s *Service) GroupsCreate(groupInformation *model.GroupInformation) *GroupsCreateOp {
	return &GroupsCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "groups",
		Payload:    groupInformation,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupsCreateOp implements DocuSign API SDK UserGroups::createGroups
type GroupsCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupsCreateOp) Do(ctx context.Context) (*model.GroupInformation, error) {
	var res *model.GroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupsDelete deletes an existing user group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groups/delete
//
// SDK Method UserGroups::deleteGroups
func (s *Service) GroupsDelete(groupInformation *model.GroupInformation) *GroupsDeleteOp {
	return &GroupsDeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "groups",
		Payload:    groupInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupsDeleteOp implements DocuSign API SDK UserGroups::deleteGroups
type GroupsDeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupsDeleteOp) Do(ctx context.Context) (*model.GroupInformation, error) {
	var res *model.GroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// GroupsList gets information about groups associated with the account.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groups/list
//
// SDK Method UserGroups::listGroups
func (s *Service) GroupsList() *GroupsListOp {
	return &GroupsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "groups",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupsListOp implements DocuSign API SDK UserGroups::listGroups
type GroupsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupsListOp) Do(ctx context.Context) (*model.GroupInformation, error) {
	var res *model.GroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count number of records to return. The number must be greater than 1 and less than or equal to 100.
func (op *GroupsListOp) Count(val int) *GroupsListOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// GroupType set the call query parameter group_type
func (op *GroupsListOp) GroupType(val string) *GroupsListOp {
	if op != nil {
		op.QueryOpts.Set("group_type", val)
	}
	return op
}

// SearchText set the call query parameter search_text
func (op *GroupsListOp) SearchText(val string) *GroupsListOp {
	if op != nil {
		op.QueryOpts.Set("search_text", val)
	}
	return op
}

// StartPosition starting value for the list.
func (op *GroupsListOp) StartPosition(val int) *GroupsListOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// GroupsUpdate updates the group information for a group.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/usergroups/groups/update
//
// SDK Method UserGroups::updateGroups
func (s *Service) GroupsUpdate(groupInformation *model.GroupInformation) *GroupsUpdateOp {
	return &GroupsUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "groups",
		Payload:    groupInformation,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GroupsUpdateOp implements DocuSign API SDK UserGroups::updateGroups
type GroupsUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GroupsUpdateOp) Do(ctx context.Context) (*model.GroupInformation, error) {
	var res *model.GroupInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
