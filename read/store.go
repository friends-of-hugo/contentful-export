package read

type Store interface {
	ReadFromFile(path string) (result []byte, err error)
}
