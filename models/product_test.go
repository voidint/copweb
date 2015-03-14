package models

import "testing"

func getProductInstance() *Product {
	coverImg := &ProductImage{
		Path: "/image/acedd39e-f253-47e4-aeb9-0bc2bfd59ecb.jpg",
	}

	detailImg0 := &ProductImage{
		Path: "/image/e21fb8ed-2d46-46f0-a8e1-f03cf6bb86fc.jpg",
	}
	detailImg1 := &ProductImage{
		Path: "/image/c01d652d-7457-49e3-b353-2ae9031d6964.jpg",
	}
	detailImgs := []*ProductImage{detailImg0, detailImg1}

	prod := &Product{
		Title:           "unit test",
		Intro:           "unit test",
		Desc:            "unit test",
		DescUseMarkdown: 1,
		IsPublic:        ACCESSABLE_PRIVATE,
		SortNo:          0,
		CoverImg:        coverImg,
		DetailImgs:      detailImgs,
	}
	return prod
}

func TestAddProduct(t *testing.T) {
	err := AddProduct(getProductInstance())
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProducts(t *testing.T) {
	prods, err := GetProducts(nil, 1, 1, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(prods))
	for _, prod := range prods {
		t.Logf("%#v\n", prod)
	}
}

func TestCountProducts(t *testing.T) {
	count, err := CountProducts(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("count=%d\n", count)

	cond := &Product{IsPublic: 1}
	count, err = CountProducts(cond)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("count=%d\n", count)
}

func TestGetProductImages(t *testing.T) {
	imgs, err := GetProductImages(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Total prod imgs: %d\n", len(imgs))

	imgs, err = GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_COVER})
	if err != nil {
		t.Fatal(err)
	}
	for _, img := range imgs {
		t.Log(img.Path)
	}
}

func TestRmProducts(t *testing.T) {
	prod := getProductInstance()
	err := AddProduct(prod)
	if err != nil {
		t.Fatal(err)
	}

	// t.Logf("prodId:%s\n", prod.Id)
	affected, err := RmProducts(prod.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d records removed.\n", affected)
}

func TestModProduct(t *testing.T) {
	prod := getProductInstance()
	err := AddProduct(prod)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Before update: %#v\n", prod)
	t.Logf("Before update coverImg:%s\n", prod.CoverImg.Path)

	prod.Title = "unit test modified"
	prod.IsPublic = ACCESSABLE_PUBLIC
	coverImg := &ProductImage{
		Path: "/image/cd6e04ae-768d-48f2-88c2-3585c615cbab.jpg",
	}
	prod.CoverImg = coverImg

	affected, err := ModProduct(prod)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("affected: %d\n", affected)

	prods, err := GetProducts(&Product{Id: prod.Id}, 1, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(prods) < 1 {
		return
	}
	t.Logf("After update: %#v\n", prods[0])
	t.Logf("After update coverImg:%s\n", prods[0].CoverImg.Path)
}

func TestGetProductById(t *testing.T) {
	prods, err := GetProducts(nil, 1, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, prod := range prods {
		product, has, err := GetProductById(prod.Id, false)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%t %#v\n", has, product)
	}
}

func TestPushPinProduct(t *testing.T) {
	prod := getProductInstance()
	err := AddProduct(prod)
	if err != nil {
		t.Fatal(err)
	}

	affected, err := PushPinProduct(prod.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(affected)
}
