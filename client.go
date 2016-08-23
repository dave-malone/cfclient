package cf

import "github.com/pivotalservices/user-migration-plugin/uaa"

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
	SetSpaceRole(spaceGuid string, spaceRole string, guid uaa.UserGuid) error
	SetSpaceRoles(guid uaa.UserGuid, spaceRoles []*SpaceRole) error
}
