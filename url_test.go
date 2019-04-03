package url2struct

import "testing"

func Test_field(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int type in string value.",
			args: args{key: "key", val: "0"},
			want: "Key int`url:\"key\"`;",
		},
		{
			name: "float type in string value.",
			args: args{key: "key", val: "1.1"},
			want: "Key float64`url:\"key\"`;",
		},
		{
			name: "bool type in string value.",
			args: args{key: "key", val: "true"},
			want: "Key bool`url:\"key\"`;",
		},
		{
			name: "string type",
			args: args{key: "key", val: "hoge"},
			want: "Key string`url:\"key\"`;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := field(tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("field() = %v, want %v", got, tt.want)
			}
		})
	}
}
