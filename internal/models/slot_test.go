package models

import (
	"reflect"
	"testing"
)

func TestNewSlot(t *testing.T) {
	type args struct {
		number uint
	}
	tests := []struct {
		name string
		args args
		want *Slot
	}{
		{
			"Test 1",
			args{1},
			&Slot{
				nil,
				1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewSlot(test.args.number)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewSlot() = %+v, want %+v", got, test.want)
			}
		})
	}
}
