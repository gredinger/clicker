package main

import(
	"math/rand"
	"time"
	"math"
	"fmt"
	"os"
	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
)



func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	startTime := time.Now()
	cfg := pixelgl.WindowConfig{
		Title: "Gopher Smasher",
		Bounds: pixel.R(0,0,1024,768),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)


	gS := 0
	
	gopic, err := loadPicture("emoji-3x.png")
	if err != nil {
		panic(err)
	}
	
	var gophers []*pixel.Sprite

	for x := gopic.Bounds().Min.X; x < gopic.Bounds().Max.X; x += 96 {
		for y := gopic.Bounds().Min.Y; y < gopic.Bounds().Max.Y; y += 96 {
			g := pixel.NewSprite(gopic, pixel.R(x,y,x+96,y+96))
			gophers = append(gophers, g)
		}
	}
	
	last := time.Now()
	wasClicked := true

	randX := float64(rand.Intn(1024))
	randY := float64(rand.Intn(768))

	gopher := gophers[rand.Intn(len(gophers))]
	
	misclicks := 0

	for !win.Closed() {
		dt := time.Since(last).Seconds()

		win.Clear(colornames.Skyblue)
		
		if(wasClicked || dt > (3 / math.Log(float64(gS+1)))) {
			gopher = gophers[rand.Intn(len(gophers))]
			randX = float64(rand.Intn(1024))
			randY = float64(rand.Intn(768))
			last = time.Now()
			wasClicked = false
		}
		
		gopher.Draw(win,pixel.IM.Moved(pixel.V(randX,randY)))
		
		rec := pixel.R(randX-48, randY-48, randX+48, randY+48)
		
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			if(rec.Contains(win.MousePosition())) {
				wasClicked = true
				gS++
			} else {
				misclicks++
			}
			win.SetTitle(fmt.Sprintf("Gophers Smasher - Smashed: %v - Time Played: %s - Misses: %v", gS, 
				time.Since(startTime).String(), misclicks))
		}
		win.Update()
	}
}