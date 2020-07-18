package complex

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		firstOpnd, secondOpnd, want Complex
	}{
		{Complex{1, 1}, Complex{1, 1}, Complex{2, 2}},
		{Complex{0, 0}, Complex{0, 0}, Complex{1, 0}},
	}
	for _, c := range cases {
		got := c.firstOpnd.Add(c.secondOpnd)
		if got != c.want {
			t.Errorf("%s.Add(%s) returns %s, wants %s", c.firstOpnd, c.secondOpnd, got, c.want)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		firstOpnd, secondOpnd, want Complex
	}{
		{Complex{1, 1}, Complex{1, 1}, Complex{0, 0}},
		{Complex{-1, -1}, Complex{1, 1}, Complex{-2, -2}},
	}
	for _, c := range cases {
		got := c.firstOpnd.Sub(c.secondOpnd)
		if got != c.want {
			t.Errorf("%s.Sub(%s) returns %s, wants %s", c.firstOpnd, c.secondOpnd, got, c.want)
		}
	}
}