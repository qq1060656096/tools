package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// 2021-05-06 07:09:08
	tm := time.Unix(1620256148, 0)
	date := Format("Y-m-d h:i:s", tm)
	assert.Equal(t, "2021-05-06 07:09:08", date)

	date = Format("Y", tm)
	assert.Equal(t, "2021", date)
	date = Format("y", tm)
	assert.Equal(t, "21", date)

	date = Format("m", tm)
	assert.Equal(t, "05", date)

	date = Format("n", tm)
	assert.Equal(t, "5", date)

	date = Format("d", tm)
	assert.Equal(t, "06", date)

	date = Format("H", tm)
	assert.Equal(t, "07", date)

	date = Format("i", tm)
	assert.Equal(t, "09", date)

	date = Format("s", tm)
	assert.Equal(t, "08", date)
}

func TestDays(t *testing.T) {
	// 2021-05-06 07:09:08
	tm := time.Unix(1620256148, 0)
	days := Days(tm)
	assert.Equal(t, 25, days)
}