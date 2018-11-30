// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"
import "math/big"
import "github.com/shopspring/decimal"

// Fast Binary Encoding decimal field model
type FieldModelDecimal struct {
    // Field model buffer
    buffer *Buffer
    // Field model buffer offset
    offset int
}

// Create a new decimal field model
func NewFieldModelDecimal(buffer *Buffer, offset int) *FieldModelDecimal {
    return &FieldModelDecimal{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelDecimal) FBESize() int { return 16 }
// Get the field extra size
func (fm *FieldModelDecimal) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelDecimal) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelDecimal) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelDecimal) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelDecimal) FBEUnshift(size int) { fm.offset -= size }

// Check if the decimal value is valid
func (fm *FieldModelDecimal) Verify() bool { return true }

// Get the decimal value
func (fm *FieldModelDecimal) Get() (Decimal, error) {
    return fm.GetDefault(DecimalZero())
}

// Get the decimal value with provided default value
func (fm *FieldModelDecimal) GetDefault(defaults Decimal) (Decimal, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    // Read decimal parts
    low := ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())
    mid := ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 4)
    high := ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 8)
    flags := ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 12)

    // Calculate decimal value
    negative := (flags & 0x80000000) != 0
    scale := (flags & 0x7FFFFFFF) >> 16
    result := decimal.New(int64(high), 0).Mul(lowScaleField)
    result = result.Add(decimal.New(int64(mid), 0).Mul(midScaleField))
    result = result.Add(decimal.New(int64(low), 0))
    result = result.Shift(-int32(scale))
    if negative {
        result = result.Neg()
    }

    return Decimal{result}, nil
}

// Set the decimal value
func (fm *FieldModelDecimal) Set(value Decimal) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    // Extract decimal parts
    negative := value.IsNegative()
    number := value.Coefficient()
    scale := -value.Exponent()

    // Check for decimal number overflow
    bits := number.BitLen()
    if (bits < 0) || (bits > 96) {
        // Value too big for .NET Decimal (bit length is limited to [0, 96])
        WriteCount(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), 0, fm.FBESize())
        return errors.New("value too big for .NET Decimal (bit length is limited to [0, 96])")
    }

    // Check for decimal scale overflow
    if (scale < 0) || (scale > 28) {
        // Value scale exceeds .NET Decimal limit of [0, 28]
        WriteCount(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), 0, fm.FBESize())
        return errors.New("value scale exceeds .NET Decimal limit of [0, 28]")
    }

    // Write unscaled value to bytes 0-11
    bytes := number.Bytes()
    for i := range bytes {
        WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + i, bytes[len(bytes) - i - 1])
    }
    WriteCount(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + len(bytes), 0, 12 - len(bytes))

    // Write scale at byte 14
    WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 14, byte(scale))

    // Write signum at byte 15
    if negative {
        WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 15, 0x80)
    } else {
        WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset() + 15, 0)
    }
    return nil
}

var lowScaleField, midScaleField decimal.Decimal

func init()  {
    var low, mid big.Int
    low.SetString("18446744073709551616", 10)
    mid.SetString("4294967296", 10)
    lowScaleField = decimal.NewFromBigInt(&low, 0)
    midScaleField = decimal.NewFromBigInt(&mid, 0)
}
