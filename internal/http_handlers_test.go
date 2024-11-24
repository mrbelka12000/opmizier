package internal

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mrbelka12000/optimizer/internal/models"
)

type mockValue struct {
	Input1  any
	Input2  any
	Output1 any
	Output2 any
	Calls   int
}

func TestService_makeListHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := NewMockadapter(ctrl)

	cases := []struct {
		name string

		query     url.Values
		mockValue mockValue

		expectedReq    models.Request
		expectedStatus int
	}{
		{
			name: "ok",

			query: url.Values{
				"id":         []string{"15"},
				"last_name":  []string{"teka"},
				"first_name": []string{"beka"},
				"website":    []string{"https://google.com"},
			},

			mockValue: mockValue{
				Input1: gomock.Any(),
				Input2: models.Request{
					ID:        15,
					FirstName: "beka",
					LastName:  "teka",
					Website:   "https://google.com",
				},
				Output1: nil,
				Calls:   1,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "error cannot decode query",
			query: url.Values{
				"id":         []string{"15b"},
				"last_name":  []string{"teka"},
				"first_name": []string{"beka"},
				"website":    []string{"https://google.com"},
			},

			mockValue: mockValue{
				Calls: 0,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error cannot handle request",

			query: url.Values{
				"id":                []string{"15"},
				"subscription_date": []string{"15-02-2002"},
				"email":             []string{"test@gmail.com"},
				"website":           []string{"https://google.com"},
			},

			mockValue: mockValue{
				Input1: gomock.Any(),
				Input2: models.Request{
					ID:               15,
					SubscriptionDate: "15-02-2002",
					Email:            "test@gmail.com",
					Website:          "https://google.com",
				},
				Output1: assert.AnError,
				Calls:   1,
			},

			expectedStatus: http.StatusInternalServerError,
		},
	}

	s := NewService(adapter, slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			adapter.EXPECT().
				List(tc.mockValue.Input1, tc.mockValue.Input2).
				Return(tc.mockValue.Output1).
				Times(tc.mockValue.Calls)

			req := httptest.NewRequest(http.MethodGet, "/list?"+tc.query.Encode(), nil)
			w := httptest.NewRecorder()

			s.makeListHandler()(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
