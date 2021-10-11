package uns

import "testing"

func TestBytesStruct(t *testing.T) {
	bs := []byte{65, 66, 67, 68, 69, 70, 71, 72}
	bss := BytesStructOf(bs)
	t.Log(bss)
	t.Log(bss.Bytes())
	t.Log("=========================")
	t.Log(bss.ToString())
	t.Log(bss.ToStringStruct())
	t.Log("=========================")
	t.Log(bss.IntSlice())
	t.Log(bss.ToIntSliceStruct())
	t.Log("=========================")
	t.Log(bss.UintSlice())
	t.Log(bss.ToUintSliceStruct())
	t.Log("=========================")
	t.Log(bss.Int8Slice())
	t.Log(bss.ToUint8SliceStruct())
	t.Log("=========================")
	t.Log(bss.Uint8Slice())
	t.Log(bss.ToUint8SliceStruct())
	t.Log("=========================")
	t.Log(bss.Int16Slice())
	t.Log(bss.ToUint16SliceStruct())
	t.Log("=========================")
	t.Log(bss.Uint16Slice())
	t.Log(bss.ToUint16SliceStruct())
	t.Log("=========================")
	t.Log(bss.Int32Slice())
	t.Log(bss.ToUint32SliceStruct())
	t.Log("=========================")
	t.Log(bss.Uint32Slice())
	t.Log(bss.ToUint32SliceStruct())
	t.Log("=========================")
	t.Log(bss.Int64Slice())
	t.Log(bss.ToUint64SliceStruct())
	t.Log("=========================")
	t.Log(bss.Uint64Slice())
	t.Log(bss.ToUint64SliceStruct())
	t.Log("=========================")
	t.Log(bss.Float32Slice())
	t.Log(bss.ToFloat32SliceStruct())
	t.Log("=========================")
	t.Log(bss.Float64Slice())
	t.Log(bss.ToFloat64SliceStruct())
	t.Log("=========================")
	t.Log(BytesToString(bs))
	t.Log("=========================")
	t.Log(BytesToIntSlice(bs))
	t.Log(BytesToUintSlice(bs))
	t.Log("=========================")
	t.Log(BytesToInt8Slice(bs))
	t.Log(BytesToUint8Slice(bs))
	t.Log("=========================")
	t.Log(BytesToInt16Slice(bs))
	t.Log(BytesToUint16Slice(bs))
	t.Log("=========================")
	t.Log(BytesToInt32Slice(bs))
	t.Log(BytesToUint32Slice(bs))
	t.Log("=========================")
	t.Log(BytesToInt64Slice(bs))
	t.Log(BytesToUint64Slice(bs))
	t.Log("=========================")
	t.Log(BytesToFloat32Slice(bs))
	t.Log(BytesToFloat64Slice(bs))
}
