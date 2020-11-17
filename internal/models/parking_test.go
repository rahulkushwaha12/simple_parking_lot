package models

import (
	"reflect"
	"testing"
)

func TestNewParking(t *testing.T) {

	type args struct {
		capacity uint
	}
	tests := []struct {
		name string
		args args
		want *Parking
	}{
		{
			"test 1",
			args{capacity: 0},
			&Parking{
				capacity: 0,
				slots:    make([]*Slot, 0),
			},
		},
		{
			"test 2",
			args{capacity: 2},
			&Parking{
				capacity: 2,
				slots:    []*Slot{{number: 1}, {number: 2}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewParking(test.args.capacity)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %+v, want %+v", got, test.want)
			}
		})
	}

}

func TestParking_Slots(t *testing.T) {
	type args struct {
		parking *Parking
	}
	tests := []struct {
		name string
		args args
		want []*Slot
	}{
		{
			"Test 1",
			args{&Parking{
				capacity: 0,
				slots:    []*Slot{},
			}},
			[]*Slot{},
		},
		{
			"Test 2",
			args{},
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.args.parking.Slots()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %+v, want %+v", got, test.want)
			}
		})
	}
}

func TestParking_GetSlotByIndex(t *testing.T) {
	type args struct {
		parking *Parking
		index   uint
	}
	tests := []struct {
		name string
		args args
		want *Slot
	}{
		{
			"Test 1",
			args{&Parking{
				capacity: 0,
				slots:    []*Slot{},
			},
				0,
			},
			nil,
		},
		{
			"Test 2",
			args{nil, 1},
			nil,
		},
		{
			"Test 3",
			args{&Parking{
				capacity: 1,
				slots: []*Slot{
					{
						nil,
						uint(1),
					},
				},
			},
				0,
			},
			&Slot{nil, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := test.args.parking.GetSlotByIndex(test.args.index)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %+v, want %+v", got, test.want)
			}
		})
	}
}
