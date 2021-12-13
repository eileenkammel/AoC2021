package aoc

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want BingoInput
	}{
		{"1,2,3", args{s: "1,2,3"}, BingoInput{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBingoCard(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want BingoCard
	}{
		{"sample input 1", args{s: []string{
			"22 13 17 11  0",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10  3 18  5",
			" 1 12 20 15 19",
		}}, NewBingoCard(
			[5][5]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBingoCard(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBingoCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

var sampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestParseChallange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  BingoInput
		want1 []BingoCard
	}{
		{
			"Sample Input",
			args{s: sampleInput},
			BingoInput{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			[]BingoCard{
				{
					[5][5]int{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					[5][5]bool{},
				},
				{
					[5][5]int{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					[5][5]bool{},
				},
				{
					[5][5]int{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
					[5][5]bool{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseChallange(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseChallange() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseChallange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
