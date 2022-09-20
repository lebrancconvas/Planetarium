package computer

import (
	"fmt"
)

type Hardware struct {
	cpu string
	ram uint32
	ssd uint32
}

type Software struct {
	SerialID string
	name string
	size uint32
	releaseyear string
	developer string
	publisher string
}

func (h *Hardware) StartUp(s *Software) bool {
	if(h.ssd < s.size) {
		fmt.Println("Not Enough Space to start the program");
	}

	fmt.Println("Starting up...");
	return true;
}

func (h *Hardware) Run(s *Software) {
	
}

func ShutDown(s *Software) {

}