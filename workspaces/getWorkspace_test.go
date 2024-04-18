package workspaces

import (
	"backend/database"
	"backend/types"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNestProjectInWokrspace(t *testing.T) {
	w1ID := uuid.New()
	type args struct {
		rows []database.GetWorkspaceWithProjectsByIDRow
	}
	tests := []struct {
		name string
		args args
		want types.Workspace
	}{
		{
			name: "test1 :",
			args: args{
				rows: []database.GetWorkspaceWithProjectsByIDRow{
					{Workspace: database.Workspace{ID: w1ID}, Project: database.Project{Workspace: w1ID, Name: "p1_1"}},
					{Workspace: database.Workspace{ID: w1ID}, Project: database.Project{Workspace: w1ID, Name: "p1_2"}},
				},
			},
			want: types.Workspace{
				ID: w1ID,
				Projects: []types.Project{
					{Name: "p1_1", WorkspaceID: w1ID},
					{Name: "p1_2", WorkspaceID: w1ID},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NestProjectInWokrspace(tt.args.rows)
			if !reflect.DeepEqual(got.Projects, tt.want.Projects) {
				t.Errorf("NestProjectInWokrspace() got : %+v, want %+v", got, tt.want)
			}
		})
	}
}
