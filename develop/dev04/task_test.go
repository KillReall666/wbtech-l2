package main

import (
	"reflect"
	"testing"
)

func TestSetOfAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *map[string][]string
	}{
		{
			name: "test_1",
			args: []string{"Пятак", "пЯтка", "тяпкА", "Листок", "слИток", "стОлик", "одувАн"},
			want: &map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name: "test_2",
			args: []string{"Anagram", "nAgaram", "cAr", "Arc", "Love"},
			want: &map[string][]string{
				"anagram": {"anagram", "nagaram"},
				"car":     {"car", "arc"},
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetsOfAnagrams(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOfAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}