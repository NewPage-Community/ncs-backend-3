package model

import "testing"

func TestUser_Level(t *testing.T) {
	type fields struct {
		UID      int64
		PassType int32
		Point    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			"level 0",
			fields{
				UID:      1,
				PassType: 0,
				Point:    0,
			},
			0,
		},
		{
			"level 1",
			fields{
				UID:      1,
				PassType: 0,
				Point:    1,
			},
			1,
		},
		{
			"level 2",
			fields{
				UID:      1,
				PassType: 0,
				Point:    PassLevelPoint,
			},
			2,
		},
		{
			"level 101",
			fields{
				UID:      1,
				PassType: 0,
				Point:    PassLevelPoint*100 + 7201,
			},
			101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				UID:      tt.fields.UID,
				PassType: tt.fields.PassType,
				Point:    tt.fields.Point,
			}
			if got := u.Level(); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}
}
