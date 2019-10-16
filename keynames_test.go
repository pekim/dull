package dull

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseKeyCombination_Empty(t *testing.T) {
	keyCombination, err := ParseKeyCombination("")
	assert.NotNil(t, err)
	assert.Equal(t, KeyCombination{}, keyCombination)
}

func TestParseKeyCombination_KeyOnly(t *testing.T) {
	keyCombination, err := ParseKeyCombination("F1")
	assert.Nil(t, err)
	assert.Equal(t, KeyCombination{
		key:  KeyF1,
		mods: ModNone,
	}, keyCombination)
}

func TestParseKeyCombination_UnrecognisedKey(t *testing.T) {
	keyCombination, err := ParseKeyCombination("bad")
	assert.NotNil(t, err)
	assert.Equal(t, KeyCombination{}, keyCombination)
}

func TestParseKeyCombination_MultipleKeys(t *testing.T) {
	keyCombination, err := ParseKeyCombination("F1 F2")
	assert.NotNil(t, err)
	assert.Equal(t, KeyCombination{}, keyCombination)
}

func TestParseKeyCombination_KeyAndMod(t *testing.T) {
	keyCombination, err := ParseKeyCombination("F1<Control>")
	assert.Nil(t, err)
	assert.Equal(t, KeyCombination{
		key:  KeyF1,
		mods: ModControl,
	}, keyCombination)
}

func TestParseKeyCombination_ModAndKey(t *testing.T) {
	keyCombination, err := ParseKeyCombination("<Control>F1")
	assert.Nil(t, err)
	assert.Equal(t, KeyCombination{
		key:  KeyF1,
		mods: ModControl,
	}, keyCombination)
}

func TestParseKeyCombination_KeyAndMultipleMods(t *testing.T) {
	keyCombination, err := ParseKeyCombination("<Control>F1<Alt><Shift>")
	assert.Nil(t, err)
	assert.Equal(t, KeyCombination{
		key:  KeyF1,
		mods: ModControl | ModAlt | ModShift,
	}, keyCombination)
}

func TestParseKeyCombination_UnreconisedMod(t *testing.T) {
	keyCombination, err := ParseKeyCombination("F1<bad>")
	assert.NotNil(t, err)
	assert.Equal(t, KeyCombination{}, keyCombination)
}
