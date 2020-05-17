package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// 2021-05-06 07:09:08
	tm := time.Unix(1620256148, 0)
	date := Format("YYYY-MM-DD hh:mm:ss", tm)
	assert.Equal(t, "2021-05-06 07:09:08", date)

	date = Format("YYYY", tm)
	assert.Equal(t, "2021", date)
	date = Format("YY", tm)
	assert.Equal(t, "21", date)

	date = Format("MM", tm)
	assert.Equal(t, "05", date)

	date = Format("M", tm)
	assert.Equal(t, "5", date)

	date = Format("DD", tm)
	assert.Equal(t, "06", date)
	date = Format("D", tm)
	assert.Equal(t, "6", date)

	date = Format("hh", tm)
	assert.Equal(t, "07", date)

	date = Format("mm", tm)
	assert.Equal(t, "09", date)

	date = Format("ss", tm)
	assert.Equal(t, "08", date)
}