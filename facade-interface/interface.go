package main

type CPU interface {
	Freeze()
	Jump(position int)
	Execute()
}

type Memory interface {
	Load(position int, data string)
}

// this will return string so the string outside
type Harddisk interface {
	Read(position int) string
}
