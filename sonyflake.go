package sequence

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/sony/sonyflake"
)

type Sonyflake struct {
	flake *sonyflake.Sonyflake
}

var (
	_ Sequence[uint64] = &Sonyflake{}
	_ Sequence[uint64] = &SonyflakeGroup{}
)

func NewSonyflake(machineid uint16) *Sonyflake {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      time.Date(2021, 11, 8, 0, 0, 0, 0, time.Local),
		MachineID:      func() (uint16, error) { return machineid, nil },
		CheckMachineID: func(u uint16) bool { return u == machineid },
	})
	return &Sonyflake{flake: flake}
}

func (g *Sonyflake) Next(ctx context.Context) (uint64, error) {
	return g.flake.NextID()
}

type SonyflakeGroup struct {
	index  uint64
	flakes []*sonyflake.Sonyflake
}

func NewSonyflakeGroup(start uint16, size int) *SonyflakeGroup {
	var flakes []*sonyflake.Sonyflake

	for i := 0; i < size; i++ {
		current := start + uint16(i)
		settings := sonyflake.Settings{
			StartTime:      time.Date(2021, 11, 8, 0, 0, 0, 0, time.Local),
			MachineID:      func() (uint16, error) { return current, nil },
			CheckMachineID: func(u uint16) bool { return u == current },
		}
		flake := sonyflake.NewSonyflake(settings)
		flakes = append(flakes, flake)

	}

	return &SonyflakeGroup{flakes: flakes}
}

func (g *SonyflakeGroup) Next(ctx context.Context) (uint64, error) {
	return g.flakes[atomic.AddUint64(&g.index, 1)%uint64(len(g.flakes))].NextID()
}
