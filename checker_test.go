package paxi

import "testing"

// examples from https://pdos.csail.mit.edu/6.824/papers/fb-consistency.pdf
func TestLinerizabilityChecker(t *testing.T) {
	c := newChecker(nil)

	// single operation is liearizable
	ops := []*operation{
		&operation{42, nil, 0, 24},
	}
	if c.linearizable(ops) != true {
		t.Fatal("expected operations to be linearizable")
	}

	// concurrent operation is linearizable
	ops = []*operation{
		&operation{42, nil, 0, 5},
		&operation{nil, 42, 3, 10},
	}
	if c.linearizable(ops) != true {
		t.Fatal("expected operations to be linearizable")
	}

	// no dependency in graph is linearizable
	ops = []*operation{
		&operation{1, nil, 0, 5},
		&operation{nil, 2, 6, 10},
		&operation{3, nil, 11, 15},
		&operation{nil, 4, 16, 20},
	}
	if c.linearizable(ops) != true {
		t.Fatal("expected operations to be linearizable")
	}

	// concurrent reads
	ops = []*operation{
		&operation{100, nil, 0, 100},
		&operation{nil, 100, 5, 35},
		&operation{nil, 0, 30, 60},
	}
	if c.linearizable(ops) != true {
		t.Fatal("expected operations to be linearizable")
	}

	// concurrent reads
	ops = []*operation{
		&operation{0, nil, 0, 0},
		&operation{100, nil, 0, 100},
		&operation{nil, 100, 5, 25}, // reads 100, write time is cut to 25
		&operation{nil, 0, 30, 60},  // happends after previous read, but read 0
	}
	if c.linearizable(ops) == true {
		t.Error("expected operations to NOT be linearizable")
	}

	// read misses a previous write is not linearizable
	ops = []*operation{
		&operation{1, nil, 0, 5},
		&operation{2, nil, 6, 10},
		&operation{nil, 1, 11, 15},
	}
	if c.linearizable(ops) == true {
		t.Fatal("expected operations to NOT be linearizable")
	}

	// cross reads is not linearizable
	ops = []*operation{
		&operation{1, nil, 0, 5},
		&operation{2, nil, 0, 5},
		&operation{nil, 1, 6, 10},
		&operation{nil, 2, 6, 10},
	}
	if c.linearizable(ops) == true {
		t.Fatal("expected operations to NOT be linearizable")
	}
}
