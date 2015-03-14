package utils

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

const (
	IMG_PNG  SupportedImgType = ".png"
	IMG_JPEG SupportedImgType = ".jpg"
)

var (
	ErrInvalidImgType = errors.New("invalid image type")
)

// 系统当前支持的图片类型
type SupportedImgType string

func Convert2ImgType(fileExt string) (SupportedImgType, error) {
	switch fileExt {
	case ".png":
		return IMG_PNG, nil
	case ".jpg":
		return IMG_JPEG, nil
	}
	return "", ErrInvalidImgType
}

// 对图片的指定区域进行裁剪。
// w 图片裁剪后的输出目的地
// r 源图片
// rect 裁剪区域
// imgType 支持的图片类型
func Crop(w io.Writer, r io.Reader, rect image.Rectangle, imgType SupportedImgType) error {
	if r == nil {
		return errors.New("Reader can't be nil.")
	}
	// defer close(r)

	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  rect.Dy(),      // height in pixel or Y ratio(see Ratio Option below)
		Width:   rect.Dx(),      // width in pixel or X ratio
		Mode:    cutter.TopLeft, // Accepted Mode: TopLeft, Centered
		Anchor:  rect.Min,       // Position of the top left point
		Options: cutter.Copy,
	})
	if err != nil {
		return err
	}

	return imgEncode(w, cImg, imgType)
}

// 将图片缩放为指定宽和高。
// w 图片裁剪后的输出目的地
// r 源图片
// width 图片目标宽度
// height 图片目标高度
// imgType 支持的图片类型
func Resize(w io.Writer, r io.Reader, width, height uint, imgType SupportedImgType) error {
	if r == nil {
		return errors.New("Reader can't be nil.")
	}

	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	cImg := resize.Resize(width, height, img, resize.NearestNeighbor)

	return imgEncode(w, cImg, imgType)
}

func imgEncode(w io.Writer, cImg image.Image, imgType SupportedImgType) error {
	switch imgType {
	case IMG_PNG:
		return png.Encode(w, cImg)
	case IMG_JPEG:
		return jpeg.Encode(w, cImg, &jpeg.Options{Quality: 100})
	default:
		return errors.New("Unsupported format: " + string(imgType))
	}
}

// 裁剪图片指定区域并将该区域缩放为指定大小的宽和高。
func Thumbnails(w io.Writer, r io.Reader, rect image.Rectangle, width, height uint, imgType SupportedImgType) error {
	if r == nil {
		return errors.New("Reader can't be nil.")
	}

	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  rect.Dy(),      // height in pixel or Y ratio(see Ratio Option below)
		Width:   rect.Dx(),      // width in pixel or X ratio
		Mode:    cutter.TopLeft, // Accepted Mode: TopLeft, Centered
		Anchor:  rect.Min,       // Position of the top left point
		Options: cutter.Copy,
	})
	if err != nil {
		return err
	}

	cImg = resize.Resize(width, height, cImg, resize.NearestNeighbor)

	return imgEncode(w, cImg, imgType)
}
