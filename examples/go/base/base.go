package base

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

type Job struct {
	Name string
	Unit int
}

func (j *Job) SplitJob(_ context.Context) ([]string, error) {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./%s [nums]\n", j.Name)
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi: %v\n", err)
		os.Exit(1)
	}
	idxes := split(n, j.Unit)

	commands := []string{}
	for _, idx := range idxes {
		commands = append(commands, idx.Command(j.Name))
	}

	return commands, nil
}

func fib(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

type Idx struct {
	Start int
	End   int
}

func (i *Idx) Command(name string) string {
	return fmt.Sprintf("%s %d %d", name, i.Start, i.End)
}

func split(n, by int) []Idx {
	idxes := []Idx{}
	for start := 1; start <= n; start += by {
		end := start + by - 1
		if end > n {
			end = n
		}
		idxes = append(idxes, Idx{Start: start, End: end})
	}
	return idxes
}

func Run(name string) {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [start] [end]\n", os.Args[0])
		os.Exit(1)
	}
	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi: %v\n", err)
		os.Exit(1)
	}
	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("job name: %s\n", name)
	for i := start; i <= end; i++ {
		fmt.Printf("n: %d, result: %d\n", i, fib(i))
	}
}
