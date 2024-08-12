package api

import (
	"reflect"
	"sort"
	"testing"
)

func Test_parsePutColumns(t *testing.T) {
	type args struct {
		reqBody []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "default test case",
			args: args{
				reqBody: []byte(`
				{
					"name": "Tom",
					"age": 34,
					"occupation": "plumber"
				}
				`),
			},
			want: []string{"age", "name", "occupation"},
		},
		{
			name: "empty json test case",
			args: args{
				reqBody: []byte("{}"),
			},
			want: []string{},
		},
		{
			name: "invalid json test case",
			args: args{
				reqBody: []byte("{"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePutColumns(tt.args.reqBody)

			if (err != nil) != tt.wantErr {
				t.Errorf("parsePutColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sortedGot := got.Cols
			sort.Strings(sortedGot)

			if !reflect.DeepEqual(sortedGot, tt.want) {
				t.Errorf("parsePutColumns() = %v, want %v", sortedGot, tt.want)
			}
		})
	}
}
