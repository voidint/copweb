package models

import "testing"

func getBlog() *Blog {
	return &Blog{
		Title:     "unit test",
		Body:      "### write in markdown",
		BodyUseMd: USE_MD_YES,
		IsPublic:  ACCESSABLE_PRIVATE,
	}
}

func TestAddBlog(t *testing.T) {
	blog := getBlog()
	err := AddBlog(blog)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBlogById(t *testing.T) {
	blog := getBlog()
	err := AddBlog(blog)
	if err != nil {
		t.Fatal(err)
	}
	b, has, err := GetBlogById(blog.Id, false)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%t\n", has)
	t.Logf("%#v\n", b)
}

func TestRmBlog(t *testing.T) {
	blog := getBlog()
	err := AddBlog(blog)
	if err != nil {
		t.Fatal(err)
	}
	affected, err := RmBlog(&Blog{Id: blog.Id})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d records is deleted.\n", affected)
}

func TestPushPinBlog(t *testing.T) {
	blog := getBlog()
	err := AddBlog(blog)
	if err != nil {
		t.Fatal(err)
	}
	affected, err := PushPinBlog(blog.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("affected:%d\n", affected)
}

func TestGetBlogImagesByBlogId(t *testing.T) {
	blog := getBlog()
	blog.Cover = &BlogImage{
		Path: "/image/e35ce8bf-d84a-4efe-9c8a-a803f333d464.jpg",
	}
	err := AddBlog(blog)
	if err != nil {
		t.Fatal(err)
	}

	imgs, err := GetBlogImagesByBlogId(blog.Id)
	if err != nil {
		t.Fatal(err)
	}
	for _, img := range imgs {
		t.Logf("%#v\n", img)
	}

}
