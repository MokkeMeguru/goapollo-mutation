package pkg_test

import (
	"testing"

	"github.com/MokkeMeguru/goapollo-mutation/pkg"
)

func TestBasicRandomGenerator_RandInt(t *testing.T) {
	type fields struct {
		Seed int64
	}
	type args struct {
		closedLeft  int
		openedRight int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLess int
		wantMore int
		wantErr  bool
	}{
		{
			name:     "normal",
			fields:   fields{Seed: 100},
			args:     args{closedLeft: 0, openedRight: 100},
			wantLess: 100,
			wantMore: -1,
			wantErr:  false,
		},
		{
			name:     "invalid args",
			fields:   fields{Seed: 100},
			args:     args{closedLeft: 0, openedRight: 0},
			wantLess: 100,
			wantMore: -1,
			wantErr:  true,
		},
		{
			name:     "accept negative",
			fields:   fields{Seed: 100},
			args:     args{closedLeft: -100, openedRight: 0},
			wantLess: -1,
			wantMore: -101,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := pkg.NewBasicRandomGenerator(tt.fields.Seed)
			got, err := m.RandInt(tt.args.closedLeft, tt.args.openedRight)
			if (err != nil) != tt.wantErr {
				t.Errorf("BasicRandomGenerator.RandInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantMore >= got || got >= tt.wantLess {
				t.Errorf("BasicRandomGenerator.RandInt(): %v < %v < %v", tt.wantMore, got, tt.wantLess)
			}
		})
	}
}
