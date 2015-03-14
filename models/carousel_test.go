package models

import (
	"testing"
)

func getDemoCarousel() *Carousel {
	return &Carousel{
		ImgPath: "/image/fc47f788-e815-4103-86b3-682f513017a6.jpg",
		Caption: "无论是什么任务，配备 Intel HD Graphics 5000 图形处理器的第四代 Intel Core 处理器都能应对自如。",
		SortNo:  0,
	}
}

func TestAddCarousel(t *testing.T) {
	affected, err := AddCarousel(getDemoCarousel())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("affected: %d\n", affected)
}

func TestRmCarousel(t *testing.T) {
	caro := getDemoCarousel()

	_, err := AddCarousel(caro)
	if err != nil {
		t.Fatal(err)
	}

	affected, err := RmCarousel(caro.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d records removed", affected)
}

func TestGetCarousels(t *testing.T) {
	lim := &Limiter{
		Limit:  10,
		Offset: 0,
	}
	carousels, err := GetCarousels(nil, lim)
	if err != nil {
		t.Fatal(err)
	}
	for _, car := range carousels {
		t.Logf("%#v\n", car)
	}
}

func TestGetCarouselById(t *testing.T) {
	car := getDemoCarousel()

	_, err := AddCarousel(car)
	if err != nil {
		t.Fatal(err)
	}

	carousel, _, err := GetCarouselById(car.Id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", carousel)

}

func TestModCarousel(t *testing.T) {
	car := getDemoCarousel()

	_, err := AddCarousel(car)
	if err != nil {
		t.Fatal(err)
	}
	carousel, _, err := GetCarouselById(car.Id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Before update: %#v\n", carousel)

	car.Caption = "modified caption"
	_, err = ModCarousel(car)
	if err != nil {
		t.Fatal(err)
	}

	carousel, _, err = GetCarouselById(car.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("After update: %#v\n", carousel)
}

func TestGetCarouselPage(t *testing.T) {
	page, err := GetCarouselPage(nil, 1, 3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", page)
	for _, row := range page.Rows {
		if val, ok := row.(*Carousel); ok {
			t.Logf("%#v", val)
		}
	}
}

func TestPushPinCarousel(t *testing.T) {
	caro := getDemoCarousel()
	_, err := AddCarousel(caro)
	if err != nil {
		t.Fatal(err)
	}

	affected, err := PushPinCarousel(caro.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(affected)
}
