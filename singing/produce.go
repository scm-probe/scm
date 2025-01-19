package singing

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440
)

func GenerateSound() {
	generate()
}

const nsamps = 50

func generate() {
	nsamps := Duration * SampleRate
	tau := math.Pi * 2
	var angle float64 = tau / float64(nsamps)
	file := "out.bin"
	f, _ := os.Create(file)
	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:],
			math.Float32bits(float32(sample)))
		bw, _ := f.Write(buf[:])
		fmt.Printf("\rWrote: %v bytes to %s", bw, file)
	}
}
