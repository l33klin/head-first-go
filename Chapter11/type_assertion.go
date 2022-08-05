package main

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

type TapePlayer string

// TapePlayer type definition here
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder string

// TapeRecorder type definition here
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func TryOut(player Player) {
	player.Play("Song~")
	player.Stop()
	// type assertion
	if tr, ok := player.(TapeRecorder); ok {
		tr.Record()
	}
}

func main() {
	tr := TapeRecorder("t")
	TryOut(tr)
	tp := TapePlayer("t")
	TryOut(tp)
}
