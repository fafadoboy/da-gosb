package ch2mc

import (
	"fmt"

	"github.com/samber/lo"
)

const (
	MAXNUM = 3
)

type MCState struct {
	wm, wc, em, ec int
	boat           bool
}

func (mc *MCState) IsLegal() bool {
	if mc.wm < mc.wc && mc.wm > 0 {
		return false
	}
	if mc.em < mc.ec && mc.em > 0 {
		return false
	}
	return true
}

func (mc *MCState) GoalTest(s MCState) bool {
	return s.IsLegal() && s.em == MAXNUM && s.ec == MAXNUM
}

func (mc *MCState) Successors(s MCState) (sucs []MCState) {
	if s.boat {
		if s.wm > 1 {
			sucs = append(sucs, NewMCState(s.wm-2, s.wc, !s.boat))
		}
		if s.wm > 0 {
			sucs = append(sucs, NewMCState(s.wm-1, s.wc, !s.boat))
		}
		if s.wc > 1 {
			sucs = append(sucs, NewMCState(s.wm, s.wc-2, !s.boat))
		}
		if s.wc > 0 {
			sucs = append(sucs, NewMCState(s.wm, s.wc-1, !s.boat))
		}
		if s.wc > 0 && s.wm > 0 {
			sucs = append(sucs, NewMCState(s.wm-1, s.wc-1, !s.boat))
		}
	} else {
		if s.em > 1 {
			sucs = append(sucs, NewMCState(s.wm+2, s.wc, !s.boat))
		}
		if s.em > 0 {
			sucs = append(sucs, NewMCState(s.wm+1, s.wc, !s.boat))
		}
		if s.ec > 1 {
			sucs = append(sucs, NewMCState(s.wm, s.wc+2, !s.boat))
		}
		if s.ec > 0 {
			sucs = append(sucs, NewMCState(s.wm, s.wc+1, !s.boat))
		}
		if s.ec > 0 && s.em > 0 {
			sucs = append(sucs, NewMCState(s.wm+1, s.wc+1, !s.boat))
		}
	}

	return lo.Filter(sucs, func(el MCState, _ int) bool { return el.IsLegal() })
}

func (s MCState) Hash() string {
	// h := fnv.New32a()
	// h.Write([]byte(fmt.Sprintf("WM%dWC%dEM%dEC%dBT%v", s.wm, s.wc, s.em, s.ec, s.boat)))
	// // If your T type is more complex or doesn't implement fmt.Stringer,
	// // you might need a more sophisticated approach to generate a hash.
	// return fmt.Sprintf("%x", h.Sum32())
	side := "E"
	if s.boat {
		side = "W"
	}
	return fmt.Sprintf("WM%d WC%d EM%d EC%d B%s", s.wm, s.wc, s.em, s.ec, side)
}

func (s *MCState) Print() {
	side := "EAST"
	if s.boat {
		side = "WEST"
	}

	fmt.Printf("On the WEST bank there are %d missionaries and %d cannibals.\n"+
		"On the EAST bank there are %d missionaries and %d cannibals\n"+
		"The boat is on the %s bank.\n\n", s.wm, s.wc, s.em, s.ec, side)
}

func NewMCState(missionaries, cannibals int, boat bool) MCState {
	return MCState{wm: missionaries, wc: cannibals, em: MAXNUM - missionaries, ec: MAXNUM - cannibals, boat: boat}
}
