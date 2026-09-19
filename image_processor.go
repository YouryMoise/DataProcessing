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

var OFFSETS [256 * 4]int = [256 * 4]int{
	8,-3,9,5,      4,2,-3,4,      -11,5,-4,5,    -2,-15,5,-12,
	-14,11,-5,10,  -13,3,-12,-9,  7,-13,12,-2,   5,-14,7,7,
	-1,-6,-4,-9,   1,4,-1,-13,    -15,-3,-14,7,  -1,10,-12,14,
	-13,-5,-4,-3,  -13,-11,-12,6, -4,-1,-2,14,   -13,12,-3,4,
	1,15,4,13,     -14,12,-9,14,  3,13,2,6,      4,1,-2,11,
	6,-11,10,3,    0,-12,-5,-5,   -13,2,-15,3,   -3,0,-15,1,
	-1,-11,1,10,   -3,-13,-11,-14,-16,-1,-14,10, -11,8,-3,14,
	-1,6,-7,12,    -3,10,2,14,    -14,-4,-3,-7,  -9,14,-1,14,
	-12,13,-10,8,  -14,8,-10,13,  -1,13,-7,14,   -14,0,-11,-5,
	-13,2,-12,11,  -13,12,-11,7,  8,-15,-2,-11,  6,-14,6,-9,
	4,-13,7,-11,   7,-12,11,-9,   1,15,2,9,      13,-15,12,-10,
	14,-13,13,-6,  12,-12,14,-7,  11,-13,12,-8,  14,-11,14,-6,
	0,-15,1,-7,    -1,-14,-1,-9,  2,-14,2,-8,    1,-13,1,-6,
	-13,10,-11,6,  -12,12,-11,7,  -13,1,-10,4,   -11,2,-10,-3,
	-12,-2,-7,-5,  -12,-2,-10,-6, -11,2,-9,-4,   -11,0,-8,-5,
	-12,-6,-7,-8,  -11,-7,-6,-10, -10,-8,-5,-9,  -9,-9,-4,-10,
	-12,-1,-12,4,  -12,1,-10,-2,  -10,-2,-10,3,  -10,3,-7,-1,
	-11,1,-6,-1,   -10,0,-4,-3,   -9,1,-5,-1,    -8,-2,-3,-4,
	-10,-4,-6,-6,  -9,-5,-4,-7,   -8,-6,-3,-8,   -7,-7,-2,-9,
	-12,-10,-9,-11,-11,-10,-7,-12,-10,-11,-6,-13,-9,-12,-4,-14,
	-13,15,-4,12,  -13,14,-4,11,  -12,14,-5,10,  -12,13,-5,9,
	-13,8,-5,7,    -12,8,-6,6,    -13,6,-7,5,    -12,5,-8,4,
	-13,1,-10,2,   -11,1,-9,3,    -10,1,-8,2,    -9,1,-7,3,
	-13,3,-9,1,    -12,2,-8,0,    -11,2,-7,-1,   -10,1,-6,-2,
	-14,-3,-9,-1,  -13,-2,-8,-1,  -12,-2,-7,-1,  -11,-2,-6,-1,
	-13,-7,-8,-5,  -12,-6,-7,-4,  -11,-5,-6,-3,  -10,-5,-5,-2,
	-14,-11,-9,-9, -13,-10,-8,-8, -12,-9,-7,-7,  -11,-8,-6,-6,
	-13,-13,-8,-10,-12,-12,-7,-9, -11,-11,-6,-8, -10,-10,-5,-7,
	-11,-15,-6,-12,-10,-14,-5,-11,-9,-13,-4,-10, -8,-12,-3,-9,
	-6,-15,-1,-11, -5,-14,0,-10,  -4,-13,1,-9,   -3,-12,2,-8,
	-1,-15,3,-11,  0,-14,4,-10,   1,-13,5,-9,    2,-12,6,-8,
	3,-15,7,-11,   4,-14,8,-10,   5,-13,9,-9,    6,-12,10,-8,
	13,-15,9,-11,  12,-14,8,-10,  11,-13,7,-9,   10,-12,6,-8,
	14,-11,10,-9,  13,-10,9,-8,   12,-9,8,-7,    11,-8,7,-6,
	14,-7,11,-6,   13,-6,10,-5,   12,-5,9,-4,    11,-4,8,-3,
	14,-3,12,-4,   13,-2,11,-3,   12,-1,10,-2,   11,0,9,-1,
	14,1,13,6,     13,2,12,5,     12,3,11,4,     11,4,10,3,
	14,5,11,6,     13,6,10,5,     12,7,9,4,      11,8,8,3,
	14,9,10,9,     13,10,9,8,     12,11,8,7,     11,12,7,6,
	14,13,13,8,    12,12,12,7,    11,11,11,6,    10,10,10,5,
	13,15,10,11,   12,14,9,10,    11,13,8,9,     10,12,7,8,
	9,15,6,12,     8,14,5,11,     7,13,4,10,     6,12,3,9,
	4,15,1,11,     3,14,0,10,     2,13,-1,9,     1,12,-2,8,
	0,15,-3,11,    -1,14,-4,10,   -2,13,-5,9,    -3,12,-6,-8,
	-4,15,-7,11,   -5,14,-8,10,   -6,13,-9,9,    -7,12,-10,8,
	-12,15,-9,11,  -11,14,-8,10,  -10,13,-7,9,   -9,12,-6,8,
	-14,13,-10,9,  -13,12,-9,8,   -12,11,-8,7,   -11,10,-7,6,
	-14,9,-11,6,   -13,8,-10,5,   -12,7,-9,4,    -11,6,-8,3,
	-14,5,-12,4,   -13,4,-11,3,   -12,3,-10,2,   -11,2,-9,1,
	-14,1,-13,-4,  -13,2,-12,-3,  -12,3,-11,-2,  -11,4,-10,-1,
	-14,-3,-11,-4, -13,-2,-10,-3, -12,-1,-9,-2,   -11,0,-8,-1,
	-14,-7,-10,-9, -13,-6,-9,-8,  -12,-5,-8,-7,  -11,-4,-7,-6,
	-14,-11,-13,-6,-12,-10,-12,-5,-11,-9,-11,-4, -10,-8,-10,-3,
	-13,-15,-10,-11,-12,-14,-9,-10,-11,-13,-8,-9, -10,-12,-7,-8,
	-9,-15,-6,-12, -8,-14,-5,-11,  -7,-13,-4,-10, -6,-12,-3,-9,
	-4,-15,-1,-11, -3,-14,0,-10,   -2,-13,1,-9,   -1,-12,2,-8,
	0,-15,3,-11,   1,-14,4,-10,   2,-13,5,-9,    3,-12,6,-8,
	4,-15,7,-11,   5,-14,8,-10,   6,-13,9,-9,    7,-12,10,-8,
	12,-15,9,-11,  11,-14,8,-10,  10,-13,7,-9,   9,-12,6,-8,
	14,-13,10,-9,  13,-12,9,-8,   12,-11,8,-7,   11,-10,7,-6,
	14,-9,11,-6,   13,-8,10,-5,   12,-7,9,-4,    11,-6,8,-3,
	14,-5,12,-4,   13,-4,11,-3,   12,-3,10,-2,   11,-2,9,-1,
	14,-1,13,4,    13,-2,12,3,    12,-3,11,2,    11,-4,10,1,
	14,3,11,4,     13,2,10,3,     12,1,9,2,      11,0,8,1}

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
	var total uint8 = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if mask1[i][j] != mask2[i][j] {
				total++
			} 
		}
	}

	return total <= 14, total // 
}


func (ip *ImageProcessor) fastReject(img *image.NRGBA, t, currentPixel uint32, x, y int, minCount uint32) (bool, error) {
	if x < 3 || y < 3 {
		return false, errors.New("X or Y too small\n")
	}
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	
	if x + 3 >= width || y + 3 >= height {
		return false, errors.New("X or Y too large\n")
	}

	topLeft,_,_,_ := img.At(x-3, y-3).RGBA()
	topRight,_,_,_ := img.At(x+3, y-3).RGBA()
	bottomLeft,_,_,_ := img.At(x-3, y+3).RGBA()
	bottomRight,_,_,_ := img.At(x+3, y+3).RGBA()
	
	cardinalPixels := []uint32{topLeft, topRight, bottomLeft, bottomRight}
	
	var countBrighter, countDarker uint32 = 0,0
	// countBrighter := 0
	// countDarker := 0

	for _, pixel := range(cardinalPixels) {
		if pixel >= currentPixel+t {
			countBrighter++
		} else if int(pixel) <= int(currentPixel)-int(t) { // having these be uint causes wraparound bug
			countDarker++
		}
	}

	// if x == 3 && y == 3 {
	// 	fmt.Printf("fast: current %v cardinal %v countBrighter %v countDarker %v\n", currentPixel, cardinalPixels, countBrighter, countDarker)
	// }

	return countBrighter >= minCount || countDarker >= minCount, nil
}

func (ip *ImageProcessor) checkConsecutive(values [16]bool, n uint32) bool {
	index := 0
	
	var countConsecutive uint32 = 0
	passedZero := false
	for {
		if countConsecutive >= n {
			return true
		}
		if !values[index] {
			countConsecutive = 0
			if index == 15 || passedZero {
				return false
			}
			index=(index+1)%16
			continue
		}
		countConsecutive++
		if index == 15 {
			passedZero = true
		}
		index=(index+1)%16
	}
}

func (ip *ImageProcessor) fullCircle(img *image.NRGBA, t uint32, x,y int, n uint32) bool {
	FAST_CIRCLE_OFFSETS := [16][2]int{
		{ 0, -3}, { 1, -3}, { 2, -2}, { 3, -1},
		{ 3,  0}, { 3,  1}, { 2,  2}, { 1,  3},
		{ 0,  3}, {-1,  3}, {-2,  2}, {-3,  1},
		{-3,  0}, {-3, -1}, {-2, -2}, {-1, -3}}

	currentPixel,_,_,_ := img.At(x,y).RGBA()

	var brighterMask, darkerMask [16]bool

	// fast reject already handles the pixels that are within offset pixels of edge
	for index, offset := range(FAST_CIRCLE_OFFSETS){
		xOffset := offset[0]
		yOffset := offset[1]

		offsetPixel,_,_,_ := img.At(x+xOffset, y+yOffset).RGBA()

		brighterMask[index] = offsetPixel >= currentPixel+t
		darkerMask[index] = int(offsetPixel) <= int(currentPixel)-int(t)
	}

	return ip.checkConsecutive(brighterMask, n) || ip.checkConsecutive(darkerMask, n)
}

func (ip *ImageProcessor) generateDescriptor(img *image.NRGBA, centerX,centerY int) [256]bool {
	// look at the 31x31 patch around x,y
	// get moments
	// 	m10 = sum over all pixels, x*pixelValue
	//  m01 = sum over all pixels, y*pixelValue
	var m10, m01 uint32 = 0,0

	b := img.Bounds()
	startX := b.Min.X
	startY := b.Min.Y
	width := b.Max.X
	height := b.Max.Y
	
	for y := startY; y < height; y++ {
		for x := startX; x < width; x++ {
			pixelValue,_,_,_ := img.At(x,y).RGBA()
			m10 += uint32(x)*pixelValue
			m01 += uint32(y)*pixelValue
		}
	}
	

	angle := math.Atan2(float64(m01), float64(m10))

	// use the hardcoded list of coordinate pairs
	// 	rotate everything by that angle (how does this work since they're offsets)
	// 	the coordinates are defined as offsets from the center
	// 	x' = xcos(theta)-ysin(theta)
	// 	y' = xsin(theta)+ycos(theta) 

	var rotatedOffsets [256*4]int
	for index := 0; index < len(OFFSETS); index+=2 {
		x := OFFSETS[index]
		y := OFFSETS[index+1]
		rotatedOffsets[index] = int(math.Round(float64(x)*math.Cos(angle) - float64(y)*math.Sin(angle)))
		rotatedOffsets[index+1] = int(math.Round(float64(x)*math.Sin(angle) + float64(y)*math.Cos(angle)))
	}


	// for each coord pair
	// get average of 5x5 window around each coord
	// set bit to 1 if first average is greater, else 0

	var mask [256]bool 

	for index := 0; index < len(rotatedOffsets); index+=4 {
		x1 := rotatedOffsets[index]+centerX
		y1 := rotatedOffsets[index+1]+centerY
		x2 := rotatedOffsets[index+2]+centerX
		y2 := rotatedOffsets[index+3]+centerY

		var total1, total2 uint32 = 0,0

		for i := -2; i <= 2; i++ {
			for j := -2; j <=2; j++ {
				patch1Pixel,_,_,_ := img.At(x1+i, y1+j).RGBA()
				patch2Pixel,_,_,_ := img.At(x2+i, y2+j).RGBA()

				total1 += patch1Pixel
				total2 += patch2Pixel
			}
		}

		avg1 := total1/25
		avg2 := total2/25

		if avg1 > avg2 {
			mask[index/4] = true
		} else {
			mask[index/4] = false
		}

	}

	return mask

}

func (ip *ImageProcessor) dedupORB(img *image.NRGBA) ([][256]bool, error) {
	// rotation and heavy crops
	var T uint32 = 20
	var N uint32 = 12
	var MIN_COUNT uint32 = 3 
	// fmt.Printf("fr done\n")


	grayImg := imaging.Grayscale(img)
	grayImg = imaging.Blur(grayImg, 3.5)
	
	
	b := grayImg.Bounds()
	startX := b.Min.X
	startY := b.Min.Y
	
	width := b.Max.X
	height := b.Max.Y
	
	// resizing for speed
	// ORB can compare images of different sizes
	if width > 500 {
		grayImg = imaging.Resize(grayImg, 200, 0, imaging.Lanczos)
		// fmt.Printf("In here\n")
	}
	if height > 500 {
		grayImg = imaging.Resize(grayImg, 0, 200, imaging.Lanczos)
		// fmt.Printf("In here 2\n")

	}

	b = grayImg.Bounds()
	startX = b.Min.X
	startY = b.Min.Y
	
	width = b.Max.X
	height = b.Max.Y

	// fmt.Printf("After: w %v h %v\n", width, height)
	
	var allDescs [][256]bool

	// totalPixels := width*height
	// validPixels := 0
	// fastRejected := 0
	// circleRejected := 0

	// found := false

	for y := startY; y < height; y++ {
		// fmt.Printf("y %v\n", y)
		for x := startX; x < width; x++ {
			currentPixel,_,_,_ := grayImg.At(x,y).RGBA()
			valid, _ := ip.fastReject(grayImg, T, currentPixel, x,y, MIN_COUNT)
			// fmt.Printf("fr done\n")
			if !valid {
				// fastRejected++
				continue
			}
			if ip.fullCircle(grayImg, T, x,y,N) {
				// fmt.Printf("fc done\n")
				// if !found {
				// 	fmt.Printf("First found x %v y %v\n", x,y)
				// 	found = true
				// }
				// validPixels++
				desc := ip.generateDescriptor(grayImg, x,y)
				// fmt.Printf("gd done\n")

				allDescs = append(allDescs, desc)
			} 
			// else {
			// 	circleRejected++
			// }

		}
	}
	// fmt.Printf("orb done, total %v valid %v fr %v cr %v\n", totalPixels, validPixels, fastRejected, circleRejected)

	return allDescs, nil



}

func (ip *ImageProcessor) compareORBDescs(desc1, desc2 [][256]bool) (bool, int) {
	// doing the simple version where we just count the number below 64 for now
	matchCount := 0
	for _, d1 := range(desc1) {
		distance := 0
		for _, d2 := range(desc2) {
			distance = 0
			for index := 0; index < 256; index++ {
				if d1[index] != d2[index] {
					distance++
					if distance > 50 {
						break
					}
				}
			}
		}
		if distance <= 50 {
			matchCount++
		}
	}

	// maybe 30
	return matchCount >= 15, matchCount



}


func hi(){
	fmt.Printf("Hi")
}