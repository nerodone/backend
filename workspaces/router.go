package workspaces

import (
	"backend/server"

	"github.com/go-chi/chi/v5"
)

func workspaceRouter(s *server.Server) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", listWorkspaces(s))
	r.Post("/", createWorkspace(s))
	r.Get("/{workspace_id}", getWorkspaceByID(s))
	r.Put("/{workspace_id}", updateWorkspace(s))
	r.Delete("/{workspace_id}", deleteWorkspace(s))

	return r
}

func WorkspacesRoutes(s *server.Server) server.Route {
	return server.Route{
		Endpoint:     "/workspaces",
		Handler:      workspaceRouter(s),
		IsAuthorized: true,
	}
}
