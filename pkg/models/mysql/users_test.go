/*
MIT License
Copyright (c) 2022 Lee Choon Siong
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mysql

import (
	"github.com/choonsiong/snippetbox/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestUserModelGet(t *testing.T) {
	// Skip the test if the `-short` flag is provided when running the test.
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	// Set up a suite of table-driven tests and expected results.
	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Alice Jones",
				Email:   "alice@example.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
				Active:  true,
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "Non-existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := UserModel{db}

			user, err := m.Get(tt.userID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}
