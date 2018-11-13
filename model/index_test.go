package model

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	expected := Index{
		"foo": {
			"foo.txt": 1,
			"bar.txt": 1,
		},
	}
	got := NewIndex()
	got.Update("foo.txt", "foo")
	got.Update("bar.txt", "foo")

	if !reflect.DeepEqual(expected, got) {
		t.Error("expected", expected, "got", got)
	}
}

func TestIndexMerge(t *testing.T) {
	// не работает, если два индекса не согласованы, то есть
	// задают различное содержимое одних и тех же файлов
	firstIndex := Index{
		"foo": {
			"foo.txt": 1,
		},
		"bar": {
			"bar.txt": 1,
		},
	}
	secondIndex := Index{
		"foo": {
			"bar.txt": 1,
		},
		"bar": {
			"foo.txt": 1,
		},
	}
	got := NewIndex()
	IndexMerge(got, firstIndex)
	IndexMerge(got, secondIndex)
	expected := Index{
		"foo": {
			"foo.txt": 1,
			"bar.txt": 1,
		},
		"bar": {
			"foo.txt": 1,
			"bar.txt": 1,
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Error("expected", expected, "got", got)
	}
}

func TestFilesMerge(t *testing.T) {
	firstFiles := Files{
		"foo.txt": 1,
		"bar.txt": 1,
	}
	secondFiles := Files{
		"oof.txt": 1,
		"rab.txt": 1,
	}
	got := FilesMerge(firstFiles, secondFiles)
	expected := Files{
		"foo.txt": 1,
		"bar.txt": 1,
		"oof.txt": 1,
		"rab.txt": 1,
	}

	if !reflect.DeepEqual(expected, got) {
		t.Error("expected", expected, "got", got)
	}
}
