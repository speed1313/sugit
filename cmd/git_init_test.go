package cmd

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("expect 3, but %d", got)
	}
}

func TestGitInit(t *testing.T) {
	type args struct {
		cmd_args []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success with empty args",
			args: args{
				[]string{"git", "init"},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Git_init(tt.args.cmd_args)
			if got != tt.want {
				t.Errorf("init() = %v, want %v", got, tt.want)
			}
		})
	}

}
