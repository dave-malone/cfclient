package cf

import "github.com/dave-malone/go-uaac"

type Client interface {
	AssociateUserWithOrg(orgGuid string, guid uaa.UserGuid) error
	CreateUser(guid uaa.UserGuid) error
	FindOrg(orgName string) (OrgsResponse, error)
	FindSpace(orgGuid string, spaceName string) (SpacesResponse, error)
	GetOrgByName(orgName string) (*OrgResource, error)
	GetSpaceByName(orgGuid string, spaceName string) (*SpaceResource, error)
	GetUsers() (UsersResponse, error)
	GetUserSummary(userResource *UserResource) (UserSummaryResource, error)
	SetOrgRole(orgGuid string, guid uaa.UserGuid, orgRole string) error
	SetOrgRoles(guid uaa.UserGuid, orgRoles []*OrgRole) error
	SetSpaceRole(username string, spaceRole *SpaceRole) error
	SetSpaceRoles(username string, spaceRoles []*SpaceRole) error
}
