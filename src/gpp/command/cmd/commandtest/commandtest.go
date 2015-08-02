package main

import (
	"github.com/pkg/profile"

	"gpp/command"
)

func main() {
	p := profile.Start(
		profile.MemProfile,
		profile.ProfilePath("."),
		profile.NoShutdownHook)
	defer p.Stop()

	u := &command.Unit{"Test-01", 0, 0}
	s := command.NewCommandStack(10)
	s.Do(command.NewMoveUnitCommand(u, 10, 10))

	for i := 0; i < 10000000; i++ {
		s.Do(command.NewMoveUnitCommand(u, i*10, i*10))
	}
}
