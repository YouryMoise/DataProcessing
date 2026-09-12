package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"crypto/sha256"
	"slices"
	"errors"
	"math"
)


type ImageProcessor struct {

}
// Notes
// Grayscale images are represented as NRGBA but with r, g, and b channels being the same value

func (ip *ImageProcessor) dedup() {
	// chooses one of them by default
}

func (ip *ImageProcessor) dedupCryptoHash(filepath string) []byte {
	// identical only (including file type)
	// just returns the hash it computes
	
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Crypto Hash: Failed to read file: %s", err)
	}
	h := sha256.New()
	h.Write(data)
	fmt.Printf("%x\n", h.Sum(nil))
	return h.Sum(nil)

}


func (ip *ImageProcessor) DCT(img *image.NRGBA) ([32][32]float64, error) { 
	// center the data around 0
	// returning raw slice instead of img because grayscale would have to be ints
	// do 2DDCT on each pixel (do not split into 8x8 blocks)
	// use formula F(u,v) = 2/sqrt(MN) * C(u)*C(v)*Sum over i,j p(i,j)cos((2i+1)u*pi/2M)cos((2j+1)v*pi/2N)
	// u,v is matrix index, i,j is image index while looping
	// C: 1/sqrt(2) if input > 0, else 1
	// do brute force for now, add some caching later on for efficiency (do tests to compare speed)

	b := img.Bounds()
	startX := b.Min.X
	width := b.Max.X // this is already width, not width-1

	startY := b.Min.Y
	height := b.Max.Y // this is already height, not height-1

	var dctOutput [32][32]float64

	if width != 32 || height != 32 {
		fmt.Printf("Here width %v height %v\n", width, height)
		return dctOutput, errors.New("Images should be size 32")
	}

	CENTER_OFFSET := 128


	MN_multiplier := 2/math.Sqrt(32.0*32.0)
	
	for v := 0; v < 8; v++ {
		C_v := 1.0/math.Sqrt(2.0)
		if v > 0 {
			C_v = 1
		}

		for u := 0; u < 8; u++ {
			C_u := 1.0/math.Sqrt(2.0)
			if u > 0 {
				C_u = 1
			}

			multiplier := MN_multiplier*C_u*C_v
			// if u == 0 && v == 0 {
			// 	fmt.Printf("MN %v C_u %v C_v %v mult %v\n", MN_multiplier, C_u, C_v, multiplier) // checks out
			// }
			total := 0.0
			for y := startY; y < height; y++ {
				cos_v := math.Cos((2.0*float64(y)+1.0)*float64(v)*math.Pi/(2.0*32.0))
				for x := startX; x < width; x++ {
					pixel,_,_,_ := img.At(x,y).RGBA() // value 0-255
					pixel >>= 8 // because it's one int containing rgba concatenated
					cos_u := math.Cos((2.0*float64(x)+1.0)*float64(u)*math.Pi/(2.0*32.0))
					// if u == 0 && v == 0 {
					// 	fmt.Printf("cos_u %v cos_v %v\n", cos_u, cos_v) // checks out
					// 	fmt.Printf("pix %v dif %v\n", pixel, float64(int(pixel)-CENTER_OFFSET))
					// 	fmt.Printf("inc %v\n", float64(int(pixel)-CENTER_OFFSET)*cos_u*cos_v)
					// }
					
					total+=float64(int(pixel)-CENTER_OFFSET)*cos_u*cos_v
				}
			}
			dctOutput[v][u] = multiplier*total // YOURY: assuming u is col since tied to x
		}	
	}

	return dctOutput, nil



}

func (ip *ImageProcessor) dedupPerceptualHash(img *image.NRGBA) ([8][8]float64, error) {
	// some resizing and color
	// img, err := imaging.Open(filepath)
	// if err != nil {
	// 	fmt.Printf("failed to open image: %v", err)
	// }

	// TODO: return error if not 32x32
	// TODO: replace 32 with macro

	// downscale to 32x32
	downsampled := imaging.Resize(img, 32, 32, imaging.Lanczos)

	// grayscale
	grayscale := imaging.Grayscale(downsampled)

	// DCT, 32x32
	dct, err := ip.DCT(grayscale)
	var mask [8][8]float64

	if err != nil {
		return mask, err
	}

	// keep only top left 8x8 block
	topLeft := make([]float64, 64)
	for i := 0; i < 8; i++{
		for j := 0; j < 8; j++{
			topLeft[i*8+j] = dct[i][j]
		}
	}
	slices.Sort(topLeft)
	median := (topLeft[31] + topLeft[32])/2

	// find median of 64 pixels, set pixels above median to 1, rest 0
	// for i := 0; i < len(mask); i++ {
	// 	mask[i] = make([]float64, 8)
	// }

	for i := 0; i < len(mask); i++ {
		for j := 0; j < len(mask[0]); j++ {
			if dct[i][j] > median {
				mask[i][j] = 1
			} else {
				mask[i][j] = 0
			}
		}
	}

	return mask, nil


	// will use hamming distance between this and other bitmask to find similarity
	

	// The image is now loaded and ready for processing
	// fmt.Printf("Image opened successfully! Size: %dx%d Type %T\n", downsampled.Bounds().Dx(), downsampled.Bounds().Dy(), downsampled)
	



}

func (ip *ImageProcessor) comparePerceptualHashMasks(mask1 [8][8]float64, mask2 [8][8]float64) (bool, uint8) {
	var total uint8
	total = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if mask1[i][j] != mask2[i][j] {
				total++
			} 
		}
	}

	return total <= 14, total // 
}



func (ip *ImageProcessor) dedupORB(filepath string) {
	// rotation and heavy crops
}


func hi(){
	fmt.Printf("Hi")
}