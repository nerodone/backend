package workspaces

import (
	"backend/database"
	"backend/types"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestMatchProjectsToWorkspaces(t *testing.T) {
	w1ID := uuid.New()
	w1 := database.Workspace{
		ID:   w1ID,
		Name: "w1",
	}
	w2ID := uuid.New()
	w2 := database.Workspace{
		ID:   w2ID,
		Name: "w2",
	}
	p1ID := uuid.New()
	p1 := database.Project{
		ID:        p1ID,
		Name:      "p1",
		Workspace: w1ID,
	}
	p2ID := uuid.New()
	p2 := database.Project{
		ID:        p2ID,
		Name:      "p2",
		Workspace: w1ID,
	}
	p3ID := uuid.New()
	p3 := database.Project{
		ID:        p3ID,
		Name:      "p3",
		Workspace: w2ID,
	}
	tests := []struct {
		name string
		args []database.GetAllWorkspacesWithProjectsRow
		want []types.Workspace
	}{
		{
			name: "empty row",
			args: []database.GetAllWorkspacesWithProjectsRow{},
			want: []types.Workspace{},
		},
		{
			name: "2 worksapce 3 projects",
			args: []database.GetAllWorkspacesWithProjectsRow{
				{Workspace: w1, Project: p1},
				{Workspace: w1, Project: p2},
				{Workspace: w2, Project: p3},
			},
			want: []types.Workspace{
				{
					ID:   w1ID,
					Name: "w1",
					Projects: []types.Project{
						{
							ID:          p1ID,
							WorkspaceID: w1ID,
							Name:        "p1",
						},
						{
							ID:          p2ID,
							WorkspaceID: w1ID,
							Name:        "p2",
						},
					},
				},
				{
					Name: "w2",
					ID:   w2ID,
					Projects: []types.Project{
						types.ProjectFromDB(p3),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchProjectsToWorkspaces(tt.args)
			for _, w := range got {
				fmt.Println(w)
			}
			fmt.Printf("len of got : %v\n", len(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchProjectsToWorkspaces() =  \ngot %+v, \nwant %+v", got, tt.want)
			}
		})
	}
}
