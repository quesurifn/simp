package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	model "quesurifn/simple-auth/graph/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, token string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: Logout - logout"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string, orgID string) (*model.RefreshTokenResponse, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: ForgotPassword - forgotPassword"))
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, token string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, userID string, currentPassword string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: ChangePassword - changePassword"))
}

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: CreateOrganization - createOrganization"))
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, input model.UpdateOrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: UpdateOrganization - updateOrganization"))
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteOrganization - deleteOrganization"))
}

// CreateRoute is the resolver for the createRoute field.
func (r *mutationResolver) CreateRoute(ctx context.Context, input model.CreateRouteInput) (*model.Route, error) {
	panic(fmt.Errorf("not implemented: CreateRoute - createRoute"))
}

// UpdateRoute is the resolver for the updateRoute field.
func (r *mutationResolver) UpdateRoute(ctx context.Context, input model.UpdateRouteInput) (*model.Route, error) {
	panic(fmt.Errorf("not implemented: UpdateRoute - updateRoute"))
}

// DeleteRoute is the resolver for the deleteRoute field.
func (r *mutationResolver) DeleteRoute(ctx context.Context, id string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteRoute - deleteRoute"))
}

// BatchCreateRoutes is the resolver for the batchCreateRoutes field.
func (r *mutationResolver) BatchCreateRoutes(ctx context.Context, orgID string, routes []*model.CreateRouteInput) (*model.BatchCreateRoutesResponse, error) {
	panic(fmt.Errorf("not implemented: BatchCreateRoutes - batchCreateRoutes"))
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input model.CreateRoleInput) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: CreateRole - createRole"))
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, input model.UpdateRoleInput) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: UpdateRole - updateRole"))
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteRole - deleteRole"))
}

// AssignRole is the resolver for the assignRole field.
func (r *mutationResolver) AssignRole(ctx context.Context, userID string, roleID string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: AssignRole - assignRole"))
}

// UnassignRole is the resolver for the unassignRole field.
func (r *mutationResolver) UnassignRole(ctx context.Context, userID string, roleID string, orgID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnassignRole - unassignRole"))
}

// ValidateToken is the resolver for the validateToken field.
func (r *queryResolver) ValidateToken(ctx context.Context, token string) (*model.ValidateTokenResponse, error) {
	panic(fmt.Errorf("not implemented: ValidateToken - validateToken"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string, orgID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// ListUsers is the resolver for the listUsers field.
func (r *queryResolver) ListUsers(ctx context.Context, filter model.ListUsersFilter) (*model.ListUsersResponse, error) {
	panic(fmt.Errorf("not implemented: ListUsers - listUsers"))
}

// GetOrganization is the resolver for the getOrganization field.
func (r *queryResolver) GetOrganization(ctx context.Context, id string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: GetOrganization - getOrganization"))
}

// ListRoutes is the resolver for the listRoutes field.
func (r *queryResolver) ListRoutes(ctx context.Context, filter model.ListRoutesFilter) (*model.ListRoutesResponse, error) {
	panic(fmt.Errorf("not implemented: ListRoutes - listRoutes"))
}

// ListRoles is the resolver for the listRoles field.
func (r *queryResolver) ListRoles(ctx context.Context, filter model.ListRolesFilter) (*model.ListRolesResponse, error) {
	panic(fmt.Errorf("not implemented: ListRoles - listRoles"))
}

// CheckAccess is the resolver for the checkAccess field.
func (r *queryResolver) CheckAccess(ctx context.Context, input model.CheckAccessInput) (*model.CheckAccessResponse, error) {
	panic(fmt.Errorf("not implemented: CheckAccess - checkAccess"))
}

// BatchCheckAccess is the resolver for the batchCheckAccess field.
func (r *queryResolver) BatchCheckAccess(ctx context.Context, input model.BatchCheckAccessInput) (*model.BatchCheckAccessResponse, error) {
	panic(fmt.Errorf("not implemented: BatchCheckAccess - batchCheckAccess"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
*/
