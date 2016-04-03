package sunrpc

// XDR encoding
// implements: /usr/share/doc/RFC/standard/rfc4506.txt.gz

import "reflect"
import "encoding/binary"
import "math"

func Marshal(v interface{}, data []byte) error {
	var pos = 0
	var end = 0
	var val = reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		end = pos + 4
		binary.BigEndian.PutUint32(data[pos:end], uint32(val.Uint()))
	case reflect.Int8, reflect.Int16, reflect.Int32:
		end = pos + 4
		binary.BigEndian.PutUint32(data[pos:end], uint32(val.Uint()))
	case reflect.Int64:
		end = pos + 8
		binary.BigEndian.PutUint64(data[pos:end], val.Uint())
	case reflect.Uint64:
		end = pos + 8
		binary.BigEndian.PutUint64(data[pos:end], val.Uint())
	case reflect.Float32:
		end = pos + 4
		binary.BigEndian.PutUint32(data[pos:end], math.Float32bits(float32(val.Float())))
	case reflect.Float64:
		end = pos + 8
		binary.BigEndian.PutUint64(data[pos:end], math.Float64bits(val.Float()))
	default:
		break
	}
	pos = end
	return nil
}

func Unmarshal(v interface{}, data []byte) error {
	return nil
}
