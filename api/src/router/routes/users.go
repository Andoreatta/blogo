package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPost,
		Function:    controllers.SearchUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{userID}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userID}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
}
