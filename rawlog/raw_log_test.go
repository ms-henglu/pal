package rawlog_test

import (
	"testing"
	"time"

	"github.com/ms-henglu/pal/rawlog"
)

func Test_NewRawLog(t *testing.T) {
	timestampInCase1, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2019-01-01T00:00:00.000+0800")
	timestampInCase2, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
	ntimestampInCase3, _ := time.Parse("2006/01/02 15:04:05", "2023/07/28 14:03:36")
	testcases := []struct {
		input   string
		wantErr bool
		want    *rawlog.RawLog
	}{
		{
			input:   "2019-01-01T00:00:00.000+0800 [INFO] hello world",
			wantErr: false,
			want: &rawlog.RawLog{
				TimeStamp: timestampInCase1,
				Level:     "INFO",
				Message:   "hello world",
			},
		},
		{
			input:   "2019-01-01 00:00:00 [INFO] hello world",
			wantErr: false,
			want: &rawlog.RawLog{
				TimeStamp: timestampInCase2,
				Level:     "INFO",
				Message:   "hello world",
			},
		},
		{
			input:   "2023/07/28 14:03:36 [DEBUG] AzureRM Request: ",
			wantErr: false,
			want: &rawlog.RawLog{
				TimeStamp: ntimestampInCase3,
				Level:     "DEBUG",
				Message:   "AzureRM Request:",
			},
		},
	}
	for _, tc := range testcases {
		t.Logf("[DEBUG] input: %s", tc.input)
		got, err := rawlog.NewRawLog(tc.input)
		if tc.wantErr {
			if err == nil {
				t.Errorf("want error, got nil")
				continue
			}
		}
		if err != nil {
			if !tc.wantErr {
				t.Errorf("want no error, got %v", err)
				continue
			}
		}
		if got == nil {
			t.Errorf("want not nil, got nil")
			continue
		}
		if got.TimeStamp != tc.want.TimeStamp {
			t.Errorf("want timestamp %v, got %v", tc.want.TimeStamp, got.TimeStamp)
		}
		if got.Level != tc.want.Level {
			t.Errorf("want level %s, got %s", tc.want.Level, got.Level)
		}
		if got.Message != tc.want.Message {
			t.Errorf("want message `%s`, got `%s`", tc.want.Message, got.Message)
		}
	}
}
