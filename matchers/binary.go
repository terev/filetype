package matchers

import (
  "bytes"
  "debug/macho"
  "encoding/binary"
  "unsafe"
)

var NativeEndian binary.ByteOrder

func init() {
  if getEndian() {
    NativeEndian = binary.BigEndian
  } else {
    NativeEndian = binary.LittleEndian
  }
}

const IntSize int = int(unsafe.Sizeof(0))

func getEndian() (ret bool) {
  var i int = 0x1
  bs := (*[IntSize]byte)(unsafe.Pointer(&i))
  if bs[0] == 0 {
    return true
  } else {
    return false
  }
}

var (
  TypeMachO32  = newType("Mach-O 32-bit", "")
  TypeMachO64  = newType("Mach-O 64-bit", "")
  TypeMachOFat = newType("Mach-O Fat", "")
)

var Binary = Map{
  TypeMachO32:  MachO32,
  TypeMachO64:  MachO64,
  TypeMachOFat: MachOFat,
}

func MachO32(buf []byte) bool {
  var header macho.FileHeader

  binary.Read(bytes.NewReader(buf), NativeEndian, &header)

  return header.Magic == macho.Magic32
}

func MachO64(buf []byte) bool {
  var header macho.FileHeader

  binary.Read(bytes.NewReader(buf), NativeEndian, &header)

  return header.Magic == macho.Magic64
}

func MachOFat(buf []byte) bool {
  var magic uint32

  binary.Read(bytes.NewReader(buf), binary.BigEndian, &magic)

  return magic == macho.MagicFat
}
