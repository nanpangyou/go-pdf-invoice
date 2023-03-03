package structs

type config struct {
	Pdf_File_Path string
}
type configRoot struct {
	Basic config
}

func NewConfig() *configRoot {
	return new(configRoot)
}
