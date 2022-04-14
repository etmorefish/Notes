package main

import (
	"reflect"
	"testing"
)

var testPage = Page{
	Title: "hello",
	Body:  []byte("Hello, world!"),
}

func TestPage_save(t *testing.T) {
	type fields struct {
		Title string
		Body  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"save", fields(testPage), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title: tt.fields.Title,
				Body:  tt.fields.Body,
			}
			if err := p.save(); (err != nil) != tt.wantErr {
				t.Errorf("save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_loadPage(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    *Page
		wantErr bool
	}{
		{"load", args{title: testPage.Title}, &testPage, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadPage(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPage() got = %v, want %v", got, tt.want)
			}
			t.Logf("got body %#v", string(got.Body))
		})
	}
}
