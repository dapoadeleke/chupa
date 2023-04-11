package utility

import "testing"

func TestPrepareUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "should prepend http:// to url without scheme",
			url:  "website.com",
			want: "http://website.com",
		},
		{
			name: "should leave url with scheme as it is",
			url:  "http://website.com",
			want: "http://website.com",
		},
		{
			name: "should leave url with ssl scheme as it is",
			url:  "https://website.com",
			want: "https://website.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrepareUrl(tt.url); got != tt.want {
				t.Errorf("PrepareUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{
			name: "should return false for url without scheme",
			url:  "website.com",
			want: false,
		}, {
			name: "should return false for url without tdl",
			url:  "http://website",
			want: false,
		}, {
			name: "should return true for correct url without ssl",
			url:  "http://website.com",
			want: true,
		}, {
			name: "should return true for correct url with ssl",
			url:  "https://website.com",
			want: true,
		}, {
			name: "should return true for website with sub-domain",
			url:  "https://sub.website.com",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidUrl(tt.url); got != tt.want {
				t.Errorf("ValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
