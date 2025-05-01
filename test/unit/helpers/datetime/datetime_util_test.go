package datetimeutil_test

import (
	datetimeutil "github.com/hbttundar/diabuddy-api-infra/helpers/datetime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNow(t *testing.T) {
	now := datetimeutil.Now()
	_, err := time.Parse(time.RFC3339, now)
	assert.NoError(t, err, "expected Now() to return a valid RFC3339 datetime")
}

func TestNowWithFormat(t *testing.T) {
	format := time.RFC1123
	formatted := datetimeutil.NowWithFormat(format)
	_, err := time.Parse(format, formatted)
	assert.NoError(t, err, "expected NowWithFormat to return a valid formatted time")

	defaultFormatted := datetimeutil.NowWithFormat("")
	_, err = time.Parse(time.RFC3339, defaultFormatted)
	assert.NoError(t, err, "expected NowWithFormat with empty format to default to RFC3339")
}

func TestConvert(t *testing.T) {
	d := time.Now()
	result := datetimeutil.Convert(d)
	expected := d.UTC().Format(time.RFC3339)
	assert.Equal(t, expected, result)
}

func TestConvertWithFormat(t *testing.T) {
	d := time.Now()
	customFormat := time.Kitchen
	formatted := datetimeutil.ConvertWithFormat(d, customFormat)
	expected := d.UTC().Format(customFormat)
	assert.Equal(t, expected, formatted)

	defaultFormatted := datetimeutil.ConvertWithFormat(d, "")
	assert.Equal(t, d.UTC().Format(time.RFC3339), defaultFormatted)
}

func TestParse(t *testing.T) {
	d := time.Now().UTC().Truncate(time.Second)
	str := d.Format(time.RFC3339)

	parsed, err := datetimeutil.Parse(str, time.RFC3339)
	require.NoError(t, err)
	assert.True(t, d.Equal(parsed))

	parsedDefault, err := datetimeutil.Parse(str, "")
	require.NoError(t, err)
	assert.True(t, d.Equal(parsedDefault))

	// Invalid date
	_, err = datetimeutil.Parse("not-a-date", time.RFC3339)
	assert.Error(t, err)
}

func TestDateBetween(t *testing.T) {
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	for i := 0; i < 10; i++ {
		randomDate := datetimeutil.DateBetween(start, end)
		assert.False(t, randomDate.Before(start), "Date should not be before start")
		assert.False(t, randomDate.After(end), "Date should not be after end")
	}

	// Also test when start is after end
	reversed := datetimeutil.DateBetween(end, start)
	assert.False(t, reversed.Before(start))
	assert.False(t, reversed.After(end))
}
