package trivialerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	err := New("msg %d %d", 1, 2)
	assert.IsType(t, trivialError{}, err)
	assert.Equal(t, "msg 1 2", err.Error())
}

func Test_Wrap(t *testing.T) {
	err := Wrap(errors.New("eww"))
	assert.IsType(t, trivialError{}, err)
	assert.Equal(t, "eww", err.Error())
}

func Test_WrapIf(t *testing.T) {
	var originalErr = errors.New("hello")

	err := WrapIf(false, nil)
	assert.Nil(t, err)
	err = WrapIf(true, nil)
	assert.Nil(t, err)

	err = WrapIf(true, originalErr)
	assert.Equal(t, err, originalErr)

	err = WrapIf(false, originalErr)
	assert.NotEqual(t, err, originalErr)
	assert.IsType(t, trivialError{}, err)
	assert.Equal(t, originalErr.Error(), err.Error())
}

func Test_IsTrivial(t *testing.T) {
	var (
		originalErr = errors.New("hello")
		trivialErr  = New("hello")
	)

	assert.False(t, IsTrivial(originalErr))
	assert.True(t, IsTrivial(trivialErr))
}
