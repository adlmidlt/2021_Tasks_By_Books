package ex11

import (
	"todo/todo/ex11/gadget"
)

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func PlayMain() {
	mixtype := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player gadget.Player = gadget.TapePlayer{}
	playList(player, mixtype)
	player = gadget.TapeRecorder{}
	playList(player, mixtype)
}
