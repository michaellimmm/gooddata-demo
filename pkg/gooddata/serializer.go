package gooddata

type Serializer interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
