package fsquota

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLimit_GetHard(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		l := &Limit{
			mu:   new(sync.Mutex),
			hard: nil,
		}
		assert.EqualValues(t, 0, l.GetHard())
	})

	t.Run("NotNil", func(t *testing.T) {
		l := &Limit{
			mu:   new(sync.Mutex),
			hard: new(uint64),
		}
		*l.hard = 16
		assert.EqualValues(t, 16, l.GetHard())
	})
}

func TestLimit_GetSoft(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		l := &Limit{
			mu:   new(sync.Mutex),
			soft: nil,
		}
		assert.EqualValues(t, 0, l.GetSoft())
	})

	t.Run("NotNil", func(t *testing.T) {
		l := &Limit{
			mu:   new(sync.Mutex),
			soft: new(uint64),
		}
		*l.soft = 16
		assert.EqualValues(t, 16, l.GetSoft())
	})
}

func TestLimit_SetHard(t *testing.T) {
	l := &Limit{mu: new(sync.Mutex)}
	l.SetHard(64)
	require.NotNil(t, l.hard)
	assert.EqualValues(t, 64, *l.hard)
}

func TestLimit_SetSoft(t *testing.T) {
	l := &Limit{mu: new(sync.Mutex)}
	l.SetSoft(64)
	require.NotNil(t, l.soft)
	assert.EqualValues(t, 64, *l.soft)
}
