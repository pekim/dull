package widget

type editActionPos struct {
	cursor    int
	selection int
}

type textEditAction struct {
	//actionType editActionType
	beforePos  editActionPos
	afterPos   editActionPos
	deleteText []rune
	insertText []rune
}
