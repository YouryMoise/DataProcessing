package main
import (
	"testing"
	"slices"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"math"
	"fmt"
)

func TestCryptoHashIdentical(t *testing.T) {
	ip := &ImageProcessor{}

	file1 := "data/example5.png"
	file2 := "data/example5Identical.png"
	
	hash1 := ip.dedupCryptoHash(file1)
	hash2 := ip.dedupCryptoHash(file2)
	
	if !slices.Equal(hash1, hash2) {
		t.Errorf("Not equal")
	}
}

func TestCryptoHashNonIdentical(t *testing.T) {
	fmt.Printf("hi\n")
	
	ip := &ImageProcessor{}

	file1 := "data/example5.png"
	file2 := "data/example7.png"
	
	hash1 := ip.dedupCryptoHash(file1)
	hash2 := ip.dedupCryptoHash(file2)
	
	if slices.Equal(hash1, hash2) {
		t.Errorf("Equal")
	}
}

// passes
func TestDCTUVBoth0(t *testing.T) {
	ip := &ImageProcessor{}

	rect := image.Rect(0, 0, 32, 32)
	img := image.NewNRGBA(rect)
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var val uint8
			val = 129
			img.Set(x, y, color.RGBA{R: val, G: val, B: val, A: 255})
		}
	}

	expectedTotal := 32.0*32.0 // if u = v = 0, then just adding cos(0)*cos(0) = 1 32x32 times
	multiplier := 2/math.Sqrt(32.0*32.0)*0.5 // C_u = C_v = 1/sqrt(2) since u = v = 0
	expectedOutput := float64(expectedTotal*multiplier)

	dctOutput, err := ip.DCT(img)
	if err != nil {
		t.Errorf("DCT returned error %v\n", err)
	}

	if math.Abs(dctOutput[0][0] - expectedOutput) > 1e-6 {
		t.Errorf("Expected %v, got %v\n", expectedOutput, dctOutput[0][0])
	}

	// C_u and C_v disappear because 0

}

// passes
func TestDCTU_0_V_Non0(t *testing.T) {
	ip := &ImageProcessor{}

	rect := image.Rect(0, 0, 32, 32)
	img := image.NewNRGBA(rect)
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var val uint8
			val = 129
			img.Set(x, y, color.RGBA{R: val, G: val, B: val, A: 255})
		}
	}

	// u := 0
	v := 2

	multiplier := 2/math.Sqrt(32.0*32.0)*1/math.Sqrt(2.0) // C_u = 1/sqrt(2) since u = 0, C_v = 1 since v > 0
	expectedTotal := 0.0
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			cos_u := 1.0
			cos_v := math.Cos((2.0*float64(y)+1.0)*float64(v)*math.Pi/(2.0*32.0))
			expectedTotal += cos_u*cos_v
		}
	}
	expectedOutput := float64(expectedTotal*multiplier)

	dctOutput, err := ip.DCT(img)
	if err != nil {
		t.Errorf("DCT returned error %v\n", err)
	}

	if math.Abs(dctOutput[2][0] - expectedOutput) > 1e-6 {
		t.Errorf("Expected %v, got %v\n", expectedOutput, dctOutput[2][0])
	}

	// C_u and C_v disappear because 0

}

func TestDCTU_Non0_V_0(t *testing.T) {
	ip := &ImageProcessor{}

	rect := image.Rect(0, 0, 32, 32)
	img := image.NewNRGBA(rect)
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var val uint8
			val = 255
			img.Set(x, y, color.RGBA{R: val, G: val, B: val, A: 255})
		}
	}

	u := 2
	// v := 0

	multiplier := 2/math.Sqrt(32.0*32.0)*1/math.Sqrt(2.0) // C_u = 1 since u = 2, C_v = 1/sqrt(2) since v = 0
	expectedTotal := 0.0
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			cos_u := math.Cos((2.0*float64(x)+1.0)*float64(u)*math.Pi/(2.0*32.0))
			cos_v := 1.0
			expectedTotal += cos_u*cos_v
		}
	}
	expectedOutput := float64(expectedTotal*multiplier)

	dctOutput, err := ip.DCT(img)
	if err != nil {
		t.Errorf("DCT returned error %v\n", err)
	}

	

	// always passes because expectedOutput in e-15, maybe that just actually means it's 0?
	if math.Abs(dctOutput[0][2] - expectedOutput) > 1e-6 {
		t.Errorf("Expected %v, got %v\n", expectedOutput, dctOutput[0][2])
	}

	// fmt.Printf("dctOutput %v\n", dctOutput)

	// C_u and C_v disappear because 0

}

func TestDCTU_Non0_V_Non0(t *testing.T) {
	ip := &ImageProcessor{}

	rect := image.Rect(0, 0, 32, 32)
	img := image.NewNRGBA(rect)
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var val uint8
			val = 255
			img.Set(x, y, color.RGBA{R: val, G: val, B: val, A: 255})
		}
	}

	u := 2
	v := 2

	multiplier := 2/math.Sqrt(32.0*32.0) // C_u = C_v = 1
	expectedTotal := 0.0
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			cos_u := math.Cos((2.0*float64(x)+1.0)*float64(u)*math.Pi/(2.0*32.0))
			cos_v := math.Cos((2.0*float64(y)+1.0)*float64(v)*math.Pi/(2.0*32.0))

			expectedTotal += cos_u*cos_v
		}
	}
	expectedOutput := float64(expectedTotal*multiplier)

	dctOutput, err := ip.DCT(img)
	if err != nil {
		t.Errorf("DCT returned error %v\n", err)
	}

	

	// always passes because expectedOutput in e-15, maybe that just actually means it's 0?
	if math.Abs(dctOutput[2][2] - expectedOutput) > 1e-6 {
		t.Errorf("Expected %v, got %v\n", expectedOutput, dctOutput[2][2])
	}

	// fmt.Printf("dctOutput %v\n", dctOutput)

	// C_u and C_v disappear because 0

}

func TestPerceptualHashIdentical(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := imaging.Clone(img1)
		
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if !(equal && score == 0)  {
		t.Errorf("Images were not found equal, hamming distance %v > 0\n", score)
	}
}

func TestPerceptualHashDifferent(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := loadImage("data/example7.png")
	
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if equal {
		t.Errorf("Images were found equal, hamming distance %v < 14\n", score)
	}
	// t.Errorf("l")


}


func TestPerceptualHashDoubleSize(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	b := img1.Bounds()
	width := b.Max.X
	height := b.Max.Y

	img2 := imaging.Resize(img1, width*2, height*2, imaging.Lanczos)
	
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if !equal {
		t.Errorf("Images were not found equal, hamming distance %v > 14\n", score)
	}
	// t.Errorf("l")


}


func TestPerceptualHashBrightness(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/house.png")
	img2 := imaging.AdjustBrightness(img1, 74.7)

		
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if !equal {
		t.Errorf("Images were not found equal, hamming distance %v > 14\n", score)
	}

}

func TestPerceptualHashCropped(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/house.png")
	b := img1.Bounds()
	width := b.Max.X
	height := b.Max.Y

	croppedTopMargin := int(float64(width)*0.076) // best it can find
	croppedSideMargin := int(float64(height)*0.076) // best it can find
	
	cropRect := image.Rect(croppedSideMargin, croppedTopMargin, width-croppedSideMargin, height-croppedTopMargin)
	img2 := imaging.Crop(img1, cropRect)
		
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if !equal {
		t.Errorf("Images were not found equal, hamming distance %v > 14\n", score)
	}

}

func TestPerceptualHashRotated(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/house.png")
	img2 := imaging.Rotate(img1, 3.99, color.Black)
	mask1, _ := ip.dedupPerceptualHash(img1)
	mask2, _ := ip.dedupPerceptualHash(img2)
	equal, score := ip.comparePerceptualHashMasks(mask1, mask2)
	if !equal {
		t.Errorf("Images were not found equal, hamming distance %v > 14\n", score)
	}

}

func TestCheckConsecutiveFrom0(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{true,true,true,true,true,true,true,true,true,true,true,true,false,false,false,false}
	c := ip.checkConsecutive(values, n)
	if !c {
		t.Errorf("Did not find consecutive run")
	}
}

func TestCheckConsecutiveFromNon0NoWrap(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{false,true,true,true,true,true,true,true,true,true,true,true,true,false,false,false}
	c := ip.checkConsecutive(values, n)
	if !c {
		t.Errorf("Did not find consecutive run")
	}
}

func TestCheckConsecutiveFromNon0WithGap(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{false,true,true,true,true,true,true,true,true,true,true,true,false,true,false,false}
	c := ip.checkConsecutive(values, n)
	if c {
		t.Errorf("Found consecutive run when there was none")
	}
}

func TestCheckConsecutiveWrapStartAtEnd(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{true,true,true,true,true,true,true,true,true,true,true,false,false,false,false,true}
	c := ip.checkConsecutive(values, n)
	if !c {
		t.Errorf("Did not find consecutive run")
	}
}

func TestCheckConsecutiveWrapStartBeforeEnd(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{true,true,true,true,true,true,true,true,false,false,false,false,true,true,true,true}
	c := ip.checkConsecutive(values, n)
	if !c {
		t.Errorf("Did not find consecutive run")
	}
}

func TestCheckConsecutiveWrapStartBeforeEndWithGap(t *testing.T) {
	ip := &ImageProcessor{}
	var n uint32 = 12
	values := [16]bool{false,true,true,true,true,true,true,true,false,false,false,false,true,true,true,true}
	c := ip.checkConsecutive(values, n)
	if c {
		t.Errorf("Found consecutive run when there was none")
	}
}

func TestORBIdentical(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := imaging.Clone(img1)
		
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if !equal  {
		t.Errorf("Images were not found equal, match count %v <= 15\n", score)
	}
}

func TestORBNonIdentical(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := loadImage("data/example7.png")
		
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if equal  {
		t.Errorf("Images were found equal, match count %v >= 15\n", score)
	}
}

func TestORBDoubleSize(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	b := img1.Bounds()
	width := b.Max.X
	height := b.Max.Y

	img2 := imaging.Resize(img1, width*2, height*2, imaging.Lanczos)
	
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if !equal  {
		t.Errorf("Images were not found equal, match count %v <= 15\n", score)
	}
	// t.Errorf("l")


}


func TestORBBrightness(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := imaging.AdjustBrightness(img1, 96.65)

		
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if !equal  {
		t.Errorf("Images were not found equal, match count %v <= 15\n", score)
	}

}

func TestORBRotated(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	img2 := imaging.Rotate(img1, 180, color.Black) // degrees
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if !equal  {
		t.Errorf("Images were not found equal, match count < %v\n", score)
	}
}

func TestORBCropped(t *testing.T) {
	ip := &ImageProcessor{}
	img1 := loadImage("data/example5.png")
	b := img1.Bounds()
	width := b.Max.X
	height := b.Max.Y

	croppedTopMargin := int(float64(width)*0.46) // best it can find
	croppedSideMargin := int(float64(height)*0.46) // best it can find
	
	cropRect := image.Rect(croppedSideMargin, croppedTopMargin, width-croppedSideMargin, height-croppedTopMargin)
	img2 := imaging.Crop(img1, cropRect)
		
	mask1, _ := ip.dedupORB(img1)
	mask2, _ := ip.dedupORB(img2)
	equal, score := ip.compareORBDescs(mask1, mask2)
	if !equal {
		t.Errorf("Images were not found equal, match count < %v\n", score)
	}
}