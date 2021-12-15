package confredis

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewPackage(v interface{}) *Package {
	return &Package{
		Data: v,
	}
}

type Package struct {
	Data interface{}
}

func (p *Package) MarshalBinary() ([]byte, error) {
	// return msgpack.Marshal(p.Data)
	return json.Marshal(p)
}

func (p *Package) UnmarshalBinary(data []byte) error {
	// return msgpack.Unmarshal(data, p)
	return json.Unmarshal(data, p)
}
