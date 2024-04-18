package types

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestMapProjectsInWorkpsace(t *testing.T) {
	wk1ID, _ := uuid.NewRandom()
	worksapce1 := Workspace{
		ID:   wk1ID,
		Name: "wk1",
	}
	p1_1ID, _ := uuid.NewRandom()
	project1_1 := Project{
		WorkspaceID: wk1ID,
		ID:          p1_1ID,
		Name:        "p1_1",
	}

	p1_2ID, _ := uuid.NewRandom()
	project1_2 := Project{
		WorkspaceID: wk1ID,
		ID:          p1_2ID,
		Name:        "p1_2",
	}

	wk2ID, _ := uuid.NewRandom()
	worksapce2 := Workspace{
		ID:   wk2ID,
		Name: "wk2",
	}
	p2_1ID, _ := uuid.NewRandom()
	project2_1 := Project{
		WorkspaceID: wk2ID,
		ID:          p2_1ID,
		Name:        "p2_1",
	}

	p2_2ID, _ := uuid.NewRandom()
	project2_2 := Project{
		WorkspaceID: wk2ID,
		ID:          p2_2ID,
		Name:        "p2_2",
	}

	// this project doesnt match workspace 2
	p2_3ID, _ := uuid.NewRandom()
	project2_3 := Project{
		WorkspaceID: wk1ID,
		ID:          p2_3ID,
		Name:        "p2_3",
	}
	type args struct {
		w  Workspace
		ps []Project
	}
	tests := []struct {
		name  string
		args  args
		want  Workspace
		want1 []Project
	}{
		{
			name: "test1 : 1 workspace with  2 projects both of them match",
			args: args{
				w:  worksapce1,
				ps: []Project{project1_1, project1_2},
			},
			want: Workspace{
				ID:       wk1ID,
				Projects: []Project{project1_1, project1_2},
			},
			want1: []Project{},
		},
		{
			name: "test2 1 workspace with 3 projects 2 match and 1 doesnt",
			args: args{
				w:  worksapce2,
				ps: []Project{project2_1, project2_2, project2_3},
			},
			want: Workspace{
				ID:       wk2ID,
				Projects: []Project{project2_1, project2_2},
			},
			want1: []Project{project2_3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MapProjectsInWorkpsace(tt.args.w, tt.args.ps...)

			WantProjNames := " "
			GotProjNames := " "
			for _, p := range tt.want.Projects {
				WantProjNames += p.Name + " "
			}
			for _, p := range got.Projects {
				GotProjNames += p.Name + " "
			}
			if WantProjNames != GotProjNames {
				t.Errorf("expected project names : %v , got : %v", WantProjNames, GotProjNames)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MapProjectsInWorkpsace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
