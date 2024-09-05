package gooddata

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	u, err := url.ParseRequestURI("https://marvel.cloud.gooddata.com")
	assert.Nil(t, err)

	gd := &gooddataAPI{
		baseUrl: u,
	}

	type args struct {
		endpoint   string
		queryParam interface{}
	}

	type option struct {
		Query string `url:"q"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Given endpoints with no queryparam",
			args: args{
				endpoint: "/api/v1/entities/users",
			},
			want:    "https://marvel.cloud.gooddata.com/api/v1/entities/users",
			wantErr: false,
		},
		{
			name: "Given endpoint with query param",
			args: args{
				endpoint: "/api/v1/entities/users",
				queryParam: option{
					Query: "wonderwoman",
				},
			},
			want:    "https://marvel.cloud.gooddata.com/api/v1/entities/users?q=wonderwoman",
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url, err := gd.url(test.args.endpoint, test.args.queryParam)
			assert.Equal(t, url, test.want)
			assert.Equal(t, err != nil, test.wantErr)
		})
	}
}
