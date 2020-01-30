package imui

type Id string

const emptyId = ""

func (id Id) appendPath(newId Id) Id {
	if newId == "" {
		return id
	}

	return id + "/" + newId
}
