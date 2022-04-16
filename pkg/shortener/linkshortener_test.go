package shortener

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

//hz cho delat' s .env
func TestBitlyShortener(t *testing.T) {
	testCases := []struct {
		name    string
		link    string
		wantErr bool
	}{
		{
			name:    "normal link",
			link:    "https://github.com/EternalQ/go-coding-challenge-jr",
			wantErr: false,
		},
		{
			name:    "not a link",
			link:    "",
			wantErr: true,
		},
		{
			name:    "wrong link",
			link:    "http://giaasd.co/something/unreal",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shorten, err := GetBitlyShorten(tc.link)
			if err != nil {
				t.Error(err)
			}

			req, err := http.NewRequest("GET", shorten, nil)
			if err != nil {
				t.Error(err)
			}

			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				t.Error(err)
			}
			defer res.Body.Close()

			if !tc.wantErr {
				assert.Equal(t, res.Request.URL.String(), tc.link)
			}
		})
	}
}
