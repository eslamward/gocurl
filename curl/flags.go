package curl

type Flags struct {
	SaveInFile bool
}

func NewFlags(saveInFile bool) *Flags {

	return &Flags{
		SaveInFile: saveInFile,
	}
}
