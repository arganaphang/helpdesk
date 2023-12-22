package pagination

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

// func TestTransform(t *testing.T) {
// 	tests := []struct {
// 		got     url.Values
// 		want    *Pagination
// 		wantErr bool
// 	}{
// 		{
// 			got: url.Values{
// 				"page":     []string{"a"},
// 				"per_page": []string{"10"},
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			got: url.Values{
// 				"page":     []string{"1"},
// 				"per_page": []string{"a"},
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			got: url.Values{
// 				"page":     []string{""},
// 				"per_page": []string{""},
// 			},
// 			want: &Pagination{
// 				Page:      1,
// 				PerPage:   10,
// 				Total:     0,
// 				TotalPage: 0,
// 				Limit:     10,
// 				Offset:    0,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for idx, tt := range tests {
// 		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
// 			got, err := Transform(tt.got)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Transform() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestPagination_Finish(t *testing.T) {
	type fields struct {
		PerPage int
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Pagination
	}{
		{
			name:   "Test Finish Pagination",
			fields: fields{PerPage: 10},
			args:   args{count: 20},
			want: &Pagination{
				Page:      0,
				PerPage:   10,
				Total:     20,
				TotalPage: 2,
				Limit:     0,
				Offset:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				PerPage: tt.fields.PerPage,
			}
			p.Finish(tt.args.count)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Finish() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPagination_Transform(t *testing.T) {
	type fields struct {
		Page      int
		PerPage   int
		Limit     int
		Offset    int
		Total     int
		TotalPage int
	}
	type args struct {
		query url.Values
	}
	tests := []struct {
		fields  fields
		args    args
		want    Pagination
		wantErr bool
	}{
		{
			fields: fields{},
			args: args{
				query: url.Values{
					"page":     []string{"a"},
					"per_page": []string{"10"},
				},
			},
			want:    Pagination{},
			wantErr: true,
		},
		{
			fields: fields{},
			args: args{
				query: url.Values{
					"page":     []string{"1"},
					"per_page": []string{"a"},
				},
			},
			want:    Pagination{},
			wantErr: true,
		},
		{
			fields: fields{},
			args: args{
				query: url.Values{
					"page":     []string{""},
					"per_page": []string{""},
				},
			},
			want: Pagination{
				Page:      1,
				PerPage:   10,
				Total:     0,
				TotalPage: 0,
				Limit:     10,
				Offset:    0,
			},
			wantErr: false,
		},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			p := &Pagination{
				Page:      tt.fields.Page,
				PerPage:   tt.fields.PerPage,
				Limit:     tt.fields.Limit,
				Offset:    tt.fields.Offset,
				Total:     tt.fields.Total,
				TotalPage: tt.fields.TotalPage,
			}
			if err := p.Transform(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("Pagination.Transform() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*p, tt.want) {
				t.Errorf("Transform() = %v, want %v", *p, tt.want)
			}
		})
	}
}
