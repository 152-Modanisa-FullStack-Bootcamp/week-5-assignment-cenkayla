package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	type args struct {
		x uint32
		y uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{name: "0", args: args{x: math.MaxUint32, y: 1}, want: 0, want1: true},
		{name: "1", args: args{x: 1, y: 1}, want: 2, want1: false},
		{name: "2", args: args{x: 42, y: 2701}, want: 2743, want1: false},
		{name: "3", args: args{x: 42, y: math.MaxUint32}, want: 41, want1: true},
		{name: "4", args: args{x: 4294967290, y: 5}, want: 4294967295, want1: false},
		{name: "5", args: args{x: 4294967290, y: 6}, want: 0, want1: true},
		{name: "6", args: args{x: 4294967290, y: 10}, want: 4, want1: true},
		{name: "7", args: args{x: 4294967295}, want: 4294967295, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint32(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestVariadicSet(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want []interface{}
	}{
		{name: "1", args: []interface{}{4, 2, 5, 4, 2, 4}, want: []interface{}{4, 2, 5}},
		{name: "2", args: []interface{}{"bootcamp", "rocks!", "really", "rocks!"}, want: []interface{}{"bootcamp", "rocks!", "really"}},
		{name: "3", args: []interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, want: []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VariadicSet(tt.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{f: 42.42}, want: 42.50},
		{name: "2", args: args{f: 42}, want: 42},
		{name: "3", args: args{f: 42.01}, want: 42.25},
		{name: "4", args: args{f: 42.24}, want: 42.25},
		{name: "5", args: args{f: 42.25}, want: 42.25},
		{name: "6", args: args{f: 42.26}, want: 42.50},
		{name: "7", args: args{f: 42.55}, want: 42.75},
		{name: "8", args: args{f: 42.75}, want: 42.75},
		{name: "9", args: args{f: 42.76}, want: 43},
		{name: "10", args: args{f: 42.99}, want: 43},
		{name: "11", args: args{f: 43.13}, want: 43.25},
	}

	for _, tt := range tests {
		got := CeilNumber(tt.args.f)
		assert.Equal(t, tt.want, got)
	}
}

func TestAlphabetSoup(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "hello"}, want: "ehllo"},
		{name: "2", args: args{s: ""}, want: ""},
		{name: "3", args: args{s: "h"}, want: "h"},
		{name: "4", args: args{s: "ab"}, want: "ab"},
		{name: "5", args: args{s: "ba"}, want: "ab"},
		{name: "6", args: args{s: "bac"}, want: "abc"},
		{name: "7", args: args{s: "cba"}, want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlphabetSoup(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStringMask(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "!mysecret*", n: 2}, want: "!m********"},
		{name: "2", args: args{s: "", n: 1}, want: "*"},
		{name: "3", args: args{s: "a", n: 1}, want: "*"},
		{name: "4", args: args{s: "string", n: 0}, want: "******"},
		{name: "5", args: args{s: "string", n: 3}, want: "str***"},
		{name: "6", args: args{s: "string", n: 5}, want: "strin*"},
		{name: "7", args: args{s: "string", n: 6}, want: "******"},
		{name: "8", args: args{s: "string", n: 7}, want: "******"},
		{name: "9", args: args{s: "s*r*n*", n: 3}, want: "s*r***"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringMaskWithBuffer(tt.args.s, tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	//words := "apple,bat,hello,goodbye,cat,yellow,why"

	type args struct {
		arr [2]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{arr: [2]string{"hellocat", words}}, want: "hello,cat"},
		{name: "2", args: args{arr: [2]string{"catbat", words}}, want: "cat,bat"},
		{name: "3", args: args{arr: [2]string{"yellowapple", words}}, want: "yellow,apple"},
		{name: "4", args: args{arr: [2]string{"", words}}, want: "not possible"},
		{name: "5", args: args{arr: [2]string{"notcat", words}}, want: "not possible"},
		{name: "6", args: args{arr: [2]string{"bootcamprocks!", words}}, want: "not possible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordSplit(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkStringMaskConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringMaskConcat("s*r*n*", 3)
	}
}

func BenchmarkStringMaskWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringMaskWithBuffer("s*r*n*", 3)
	}
}
