package gocv

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestFisheyeCalibrate(t *testing.T) {

}

func TestFisheyeUndistorImage(t *testing.T) {
	img := IMRead("images/fisheyelens.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat in BilateralFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	FisheyeUndistortImage(img, dest, k, d)

	finalImg, err := dest.ToImage()
	if err != nil {
		t.Error(err)
	}

	f, err := os.Create("images/fisheye_undistorted.jpg")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()

	if err := jpeg.Encode(f, finalImg, nil); err != nil {
		t.Error(err)
	}

}

func TestFisheyeUndistorImageWithKNewMat(t *testing.T) {
	img := IMRead("images/fisheyelens.png", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat in BilateralFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	knew := NewMat()
	defer knew.Close()

	k.CopyTo(knew)

	knew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	FisheyeUndistortImageWithKNewMat(img, dest, k, d, knew)

	finalImg, err := dest.ToImage()
	if err != nil {
		t.Error(err)
	}

	f, err := os.Create("images/fisheye_undistortedknewMat.jpg")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()

	if err := jpeg.Encode(f, finalImg, nil); err != nil {
		t.Error(err)
	}

}
