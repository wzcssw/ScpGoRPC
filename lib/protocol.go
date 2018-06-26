package lib

type Package struct {
	Header Header
	Body   []byte
}

type Header struct {
	Filename     string
	Filesize     int64
	PackageSize  int64
	PackageIndex int
	PackageCount int
}
