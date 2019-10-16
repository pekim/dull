package dull

import (
	"fmt"
	"unicode"
)

var keysByName = make(map[string]Key)
var modifierKeysByName = make(map[string]ModifierKey)

func init() {
	keysByName["Space"] = KeySpace
	keysByName["Apostrophe"] = KeyApostrophe
	keysByName["Comma"] = KeyComma
	keysByName["Minus"] = KeyMinus
	keysByName["Period"] = KeyPeriod
	keysByName["Slash"] = KeySlash
	keysByName["0"] = Key0
	keysByName["1"] = Key1
	keysByName["2"] = Key2
	keysByName["3"] = Key3
	keysByName["4"] = Key4
	keysByName["5"] = Key5
	keysByName["6"] = Key6
	keysByName["7"] = Key7
	keysByName["8"] = Key8
	keysByName["9"] = Key9
	keysByName["Semicolon"] = KeySemicolon
	keysByName["Equal"] = KeyEqual
	keysByName["A"] = KeyA
	keysByName["B"] = KeyB
	keysByName["C"] = KeyC
	keysByName["D"] = KeyD
	keysByName["E"] = KeyE
	keysByName["F"] = KeyF
	keysByName["G"] = KeyG
	keysByName["H"] = KeyH
	keysByName["I"] = KeyI
	keysByName["J"] = KeyJ
	keysByName["K"] = KeyK
	keysByName["L"] = KeyL
	keysByName["M"] = KeyM
	keysByName["N"] = KeyN
	keysByName["O"] = KeyO
	keysByName["P"] = KeyP
	keysByName["Q"] = KeyQ
	keysByName["R"] = KeyR
	keysByName["S"] = KeyS
	keysByName["T"] = KeyT
	keysByName["U"] = KeyU
	keysByName["V"] = KeyV
	keysByName["W"] = KeyW
	keysByName["X"] = KeyX
	keysByName["Y"] = KeyY
	keysByName["Z"] = KeyZ
	keysByName["LeftBracket"] = KeyLeftBracket
	keysByName["Backslash"] = KeyBackslash
	keysByName["RightBracket"] = KeyRightBracket
	keysByName["GraveAccent"] = KeyGraveAccent
	keysByName["World1"] = KeyWorld1
	keysByName["World2"] = KeyWorld2
	keysByName["Escape"] = KeyEscape
	keysByName["Enter"] = KeyEnter
	keysByName["Tab"] = KeyTab
	keysByName["Backspace"] = KeyBackspace
	keysByName["Insert"] = KeyInsert
	keysByName["Delete"] = KeyDelete
	keysByName["Right"] = KeyRight
	keysByName["Left"] = KeyLeft
	keysByName["Down"] = KeyDown
	keysByName["Up"] = KeyUp
	keysByName["PageUp"] = KeyPageUp
	keysByName["PageDown"] = KeyPageDown
	keysByName["Home"] = KeyHome
	keysByName["End"] = KeyEnd
	keysByName["CapsLock"] = KeyCapsLock
	keysByName["ScrollLock"] = KeyScrollLock
	keysByName["NumLock"] = KeyNumLock
	keysByName["PrintScreen"] = KeyPrintScreen
	keysByName["Pause"] = KeyPause
	keysByName["F1"] = KeyF1
	keysByName["F2"] = KeyF2
	keysByName["F3"] = KeyF3
	keysByName["F4"] = KeyF4
	keysByName["F5"] = KeyF5
	keysByName["F6"] = KeyF6
	keysByName["F7"] = KeyF7
	keysByName["F8"] = KeyF8
	keysByName["F9"] = KeyF9
	keysByName["F10"] = KeyF10
	keysByName["F11"] = KeyF11
	keysByName["F12"] = KeyF12
	keysByName["F13"] = KeyF13
	keysByName["F14"] = KeyF14
	keysByName["F15"] = KeyF15
	keysByName["F16"] = KeyF16
	keysByName["F17"] = KeyF17
	keysByName["F18"] = KeyF18
	keysByName["F19"] = KeyF19
	keysByName["F20"] = KeyF20
	keysByName["F21"] = KeyF21
	keysByName["F22"] = KeyF22
	keysByName["F23"] = KeyF23
	keysByName["F24"] = KeyF24
	keysByName["F25"] = KeyF25
	keysByName["KP0"] = KeyKP0
	keysByName["KP1"] = KeyKP1
	keysByName["KP2"] = KeyKP2
	keysByName["KP3"] = KeyKP3
	keysByName["KP4"] = KeyKP4
	keysByName["KP5"] = KeyKP5
	keysByName["KP6"] = KeyKP6
	keysByName["KP7"] = KeyKP7
	keysByName["KP8"] = KeyKP8
	keysByName["KP9"] = KeyKP9
	keysByName["KPDecimal"] = KeyKPDecimal
	keysByName["KPDivide"] = KeyKPDivide
	keysByName["KPMultiply"] = KeyKPMultiply
	keysByName["KPSubtract"] = KeyKPSubtract
	keysByName["KPAdd"] = KeyKPAdd
	keysByName["KPEnter"] = KeyKPEnter
	keysByName["KPEqual"] = KeyKPEqual
	keysByName["LeftShift"] = KeyLeftShift
	keysByName["LeftControl"] = KeyLeftControl
	keysByName["LeftAlt"] = KeyLeftAlt
	keysByName["LeftSuper"] = KeyLeftSuper
	keysByName["RightShift"] = KeyRightShift
	keysByName["RightControl"] = KeyRightControl
	keysByName["RightAlt"] = KeyRightAlt
	keysByName["RightSuper"] = KeyRightSuper
	keysByName["Menu"] = KeyMenu

	modifierKeysByName["Shift"] = ModShift
	modifierKeysByName["Control"] = ModControl
	modifierKeysByName["Alt"] = ModAlt
	modifierKeysByName["Super"] = ModSuper
}

type KeyCombination struct {
	key  Key
	mods ModifierKey
}

func ParseKeyCombination(text string) (KeyCombination, error) {
	key := KeyUnknown
	var foundMods = make(map[ModifierKey]bool)

	start := 0
	end := 0
	inKey := false
	inMod := false
	for n, ch := range text {
		switch ch {
		case '<':
			if inMod {
				return KeyCombination{}, fmt.Errorf("failed to parse '%s', nested keys at %d and %d",
					text, start, end)
			}
			if inKey {
				if key != KeyUnknown {
					return KeyCombination{}, fmt.Errorf("failed to parse '%s', more than one key found",
						text)
				}
				k, found := keysByName[text[start:n]]
				if !found {
					return KeyCombination{}, fmt.Errorf("failed to parse '%s', unrecognised key name %s",
						text, text[start:n])
				}
				key = k

				inKey = false
				start = n + 1
			}

			inMod = true
			start = n + 1
		case '>':
			if inMod {
				mod, ok := modifierKeysByName[text[start:n]]
				if !ok {
					return KeyCombination{}, fmt.Errorf("failed to parse '%s', unrecognised mod name %s",
						text, text[start:n])
				}
				foundMods[mod] = true

				inMod = false
				start = n + 1
			} else if inKey {
				return KeyCombination{}, fmt.Errorf("failed to parse '%s', unexpected '>' at %d",
					text, n)
			}
		default:
			if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
				return KeyCombination{}, fmt.Errorf("failed to parse '%s', unexpected character '%s' at %d",
					text, string(ch), n)
			}

			if !inKey && !inMod {
				inKey = true
			}
		}

		end = n
	}

	if inKey {
		if key != KeyUnknown {
			return KeyCombination{}, fmt.Errorf("failed to parse '%s', more than one key found",
				text)
		}
		k, found := keysByName[text[start:end+1]]
		if !found {
			return KeyCombination{}, fmt.Errorf("failed to parse '%s', unreconised key name %s",
				text, text[start:end+1])
		}
		key = k
	}

	if key == KeyUnknown {
		return KeyCombination{}, fmt.Errorf("failed to parse '%s', no key found",
			text)
	}

	mods := ModNone
	for mod, have := range foundMods {
		if have {
			mods |= mod
		}
	}

	return KeyCombination{
		key:  key,
		mods: mods,
	}, nil
}
