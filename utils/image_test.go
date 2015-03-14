package utils

import (
	"image"
	"os"
	"testing"
)

func TestThumbnails(t *testing.T) {
	srcImg1 := "image1.JPG"
	dstImg1 := "dst1.jpg"
	img1File, err := os.Open(srcImg1)
	if err != nil {
		t.Fatalf("Open file (%s) err: %s\n", srcImg1, err)
	}
	defer img1File.Close()

	dstImg1File, err := os.Create(dstImg1)
	if err != nil {
		t.Fatalf("Create file (%s) err: %s\n", dstImg1, err)
	}
	defer dstImg1File.Close()

	rect1 := image.Rect(100, 10, 600, 900)

	err = Thumbnails(dstImg1File, img1File, rect1, 250, 400, IMG_JPEG)
	if err != nil {
		t.Fatalf("Thumbnails image (%s) err: %s\n", srcImg1, err)
	} else {
		t.Logf("Thumbnails image (%s) done\n", srcImg1)
	}

	srcImg2 := "image2.png"
	dstImg2 := "dst2.png"
	img2File, err := os.Open(srcImg2)
	if err != nil {
		t.Fatalf("Open file (%s) err: %s\n", srcImg2, err)
	}
	defer img2File.Close()

	dstImg2File, err := os.Create(dstImg2)
	if err != nil {
		t.Fatalf("Create file (%s) err: %s\n", dstImg2, err)
	}
	defer dstImg2File.Close()

	rect2 := image.Rect(10, 100, 700, 500)
	err = Thumbnails(dstImg2File, img2File, rect2, 400, 250, IMG_PNG)
	if err != nil {
		t.Fatalf("Thumbnails image (%s) err: %s\n", srcImg2, err)
	} else {
		t.Logf("Thumbnails image (%s) done\n", srcImg2)
	}
}
