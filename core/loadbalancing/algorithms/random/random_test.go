package random

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/netm4ul/netm4ul/core/communication"
)

func TestRandom_NextExecutionNodes(t *testing.T) {

	rand.Seed(123) // fixed seed so we can expect the same random every time.
	type fields struct {
		Nodes []communication.Node
	}

	type args struct {
		cmd communication.Command
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []communication.Node
	}{
		{
			name: "Testing next node execution",
			args: args{cmd: communication.Command{Name: "test_cmd"}},
			fields: fields{
				Nodes: []communication.Node{
					{
						ID:          "1",
						Project:     "TestProject",
						Modules:     []string{"A", "B", "C"},
						IsAvailable: true,
					},
					{
						ID:          "2",
						Project:     "TestProject",
						Modules:     []string{"A", "C", "E"},
						IsAvailable: true,
					},
				},
			},
			want: []communication.Node{
				{
					ID:          "2",
					Project:     "TestProject",
					Modules:     []string{"A", "C", "E"},
					IsAvailable: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Random{
				Nodes: tt.fields.Nodes,
			}
			if got := r.NextExecutionNodes(tt.args.cmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Random.NextExecutionNodes() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
