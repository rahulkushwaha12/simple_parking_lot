package models

import (
	"reflect"
	"testing"
)

func TestNewCar(t *testing.T) {
	type args struct {
		color, number string
	}
	tests := []struct {
		name string
		args args
		want *Car
	}{
		{
			name: "Test 1",
			args: args{
				color:  "Red",
				number: "KA-01-HH-7777",
			},
			want: &Car{
				number: "KA-01-HH-7777",
				color:  "Red",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewCar(test.args.number, test.args.color)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewCar() = %+v, want %+v", got, test.want)
			}
		})
	}
}
