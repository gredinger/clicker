package main
import(
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Boring clicker",
		Bounds: pixel.R(0,0,1024,768),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)

	pressed := 0
	
	if err != nil {
		fmt.Println(err)
	}
	
	for !win.Closed() {
		win.Update()
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			pressed++
			win.SetTitle(fmt.Sprintf("Boring clicker - Clicked: %v", pressed))
		}
	}
}