package main
import (
	"fmt"
	"image"
	"github.com/disintegration/imaging"

)

// TODO
/*
change masks used in hamming distance to be actual bits instead of arrays of bools

*/
func loadImage(filepath string) *image.NRGBA{
	img, err := imaging.Open(filepath)
	if err != nil {
		fmt.Printf("failed to open image: %v", err)
	}
	nrgbaImg := imaging.Clone(img)
	return nrgbaImg
}

func main(){
	fmt.Printf("main")
	// ip := &ImageProcessor{}
	// mask, _ := ip.dedupPerceptualHash("data/example5.png")
	// fmt.Printf("mask %v\n", mask)
	// sum := ip.dedupCryptoHash("data/example5.png")
	// fmt.Printf("sum %x\n", sum)
}