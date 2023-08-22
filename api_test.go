package danger_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResults(t *testing.T) {
	d := danger.New()
	r, err := d.Results()
	require.Nil(t, err)
	require.Equal(t, `{"fails":[],"warnings":[],"messages":[],"markdowns":[]}`, r)

	d.Message("test", "", 0)
	r, err = d.Results()
	require.Nil(t, err)
	require.Equal(t, `{"fails":[],"warnings":[],"messages":[{"message":"test"}],"markdowns":[]}`, r)
}
