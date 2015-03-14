package models

import "testing"

func TestAddFlagshipProducts(t *testing.T) {
	prod := getProductInstance()
	err := AddProduct(prod)
	if err != nil {
		t.Fatal(err)
	}

	fProd := &FlagshipProduct{
		ProductId: prod.Id,
	}

	affected, err := AddFlagshipProducts(fProd)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("AddFlagshipProducts affected:%d\n", affected)
}

func TestGetAllFlagshipProducts(t *testing.T) {
	fProds, err := GetAllFlagshipProducts(false)
	if err != nil {
		t.Fatal(err)
	}
	for _, fProd := range fProds {
		t.Logf("%#v\n", fProd)
		t.Logf("%#v\n", fProd.Product)
	}
}

func TestRmFlagshipProductById(t *testing.T) {
	fProds, err := GetAllFlagshipProducts(true)
	if err != nil {
		t.Fatal(err)
	}
	if len(fProds) > 0 {
		affected, err := RmFlagshipProductById(fProds[0].Id)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("RmFlagshipProductById affected:%d\n", affected)
	}

}

func TestPushPinFlagshipProduct(t *testing.T) {
	prod := getProductInstance()
	err := AddProduct(prod)
	if err != nil {
		t.Fatal(err)
	}

	fProd := &FlagshipProduct{
		ProductId: prod.Id,
	}

	_, err = AddFlagshipProducts(fProd)
	if err != nil {
		t.Fatal(err)
	}

	affected, err := PushPinFlagshipProduct(fProd.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(affected)
}
