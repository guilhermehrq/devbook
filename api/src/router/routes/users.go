package routes

import (
	"api/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:                   "/users",
		Method:                "POST",
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                "GET",
		Function:              controllers.GetUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                "GET",
		Function:              controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                "PUT",
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                "DELETE",
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
}
