package a_star

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

func TestAStar(t *testing.T) {
	type args struct {
		start  Pos
		end    Pos
		matrix [][]int
	}
	var tests []struct {
		name             string
		args             args
		wantIsAccessible bool
	}
	for i := 3; i <= 3; i++ {
		data, err := os.ReadFile("./maps/map" + strconv.Itoa(i) + ".json")
		if err != nil {
			panic(err)
		}

		var matrix [][]int
		err = json.Unmarshal(data, &matrix)
		if err != nil {
			panic(err)
		}

		tests = append(tests, struct {
			name             string
			args             args
			wantIsAccessible bool
		}{
			name: strconv.Itoa(i),
			args: args{
				start:  Pos{1, 1},
				end:    Pos{14, 10},
				matrix: matrix,
			},
			wantIsAccessible: true,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AStar(tt.args.start, tt.args.end, tt.args.matrix)
		})
	}
}
