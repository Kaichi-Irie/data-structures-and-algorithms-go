package dp

import "testing"

func TestRodCutting(t *testing.T) {
	tests := []struct {
		name    string
		problem RodCuttingProblem
		want    int
	}{
		{
			name: "Test Case 1",
			problem: RodCuttingProblem{
				N:      4,
				Prices: []int{1, 5, 8, 9},
			},
			want: 10,
		},
		{
			name: "Test Case 2",
			problem: RodCuttingProblem{
				N:      5,
				Prices: []int{2, 3, 7, 8, 10},
			},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.problem.SolveBottomUp()
			if got != tt.want {
				t.Errorf("SolveBottomUp() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			got := tt.problem.SolveTopDown()
			if got != tt.want {
				t.Errorf("SolveTopDown() = %v, want %v", got, tt.want)
			}
		})

	}
}
