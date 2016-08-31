// PointerTo project PointerTo.go
package ValueReference

import (
	"reflect"
	"testing"
	"time"
)

func TestValueReferenceImplementsValueReferencer(t *testing.T) {
	x := new(ValueReference)
	_ = ValueReferencer(x)
}

func TestNewAsValueReference(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Expected ValueReference
	}{
		{
			"Int",
			&xInt,
			ValueReference{
				reference:                 &xInt,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentIntValue,
				getReferentValue:          getReferentIntValue,
				reinitializeReferentValue: nil,
			},
		},
		{
			"IntPtr",
			&xIntPtr,
			ValueReference{
				reference:                 &xIntPtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentIntPtrValue,
				getReferentValue:          getReferentIntPtrValue,
				reinitializeReferentValue: reinitializesReferentIntPtrValue,
			},
		},
		{
			"String",
			&xString,
			ValueReference{
				reference:                 &xString,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"StringPtr",
			&xStringPtr,
			ValueReference{
				reference:                 &xStringPtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Int64",
			&xInt64,
			ValueReference{
				reference:                 &xInt64,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Int64Ptr",
			&xInt64Ptr,
			ValueReference{
				reference:                 &xInt64Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Time",
			&xTime,
			ValueReference{
				reference:                 &xTime,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"TimePtr",
			&xTimePtr,
			ValueReference{
				reference:                 &xTimePtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float32",
			&xFloat32,
			ValueReference{
				reference:                 &xFloat32,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float32Ptr",
			&xFloat32Ptr,
			ValueReference{
				reference:                 &xFloat32Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float64",
			&xFloat64,
			ValueReference{
				reference:                 &xFloat64,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float64Ptr",
			&xFloat64Ptr,
			ValueReference{
				reference:                 &xFloat64Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Byte",
			&xByte,
			ValueReference{
				reference:                 &xByte,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"BytePtr",
			&xBytePtr,
			ValueReference{
				reference:                 &xBytePtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		actual := *(New(x)).(*ValueReference)

		if actual.reference != expected.reference {
			t.Errorf("Testing x as %v: New(&x)  returned reference %#v \n\t\t    while expected %#v", name, actual.reference, expected.reference)
		}
		if actual.referentType != expected.referentType {
			t.Errorf("Testing x as %v: New(&x)  returned refrentType %#v \n\t\t    while expected %#v", name, actual.referentType, expected.referentType)
		}
		if reflect.ValueOf(actual.setReferentValue).Pointer() != reflect.ValueOf(expected.setReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned setReferentValue %#v \n\t\t    while expected %#v", name, actual.setReferentValue, expected.setReferentValue)
		}
		if reflect.ValueOf(actual.getReferentValue).Pointer() != reflect.ValueOf(expected.getReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned getReferentValue %#v \n\t\t    while expected %#v", name, actual.getReferentValue, expected.getReferentValue)
		}
		if reflect.ValueOf(actual.reinitializeReferentValue).Pointer() != reflect.ValueOf(expected.reinitializeReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned reinitializeReferentValue %#v \n\t\t    while expected %#v", name, actual.reinitializeReferentValue, expected.reinitializeReferentValue)
		}

	}

}

func TestSetReferentValue(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Actual   func() interface{}
		Expected interface{}
	}{
		{
			"xInt",
			&xInt,
			func() interface{} { return xInt },
			interface{}(5),
		},
		{
			"xIntPtr",
			&xIntPtr,
			func() interface{} { return xIntPtr },
			interface{}(new(int)),
		},
		{
			"xString",
			&xString,
			func() interface{} { return xString },
			interface{}("Test string"),
		},
		{
			"xStringPtr",
			&xStringPtr,
			func() interface{} { return xStringPtr },
			interface{}(new(string)),
		},
		{
			"xInt64",
			&xInt64,
			func() interface{} { return xInt64 },
			interface{}(int64(9999999999999)),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			func() interface{} { return xInt64Ptr },
			interface{}(new(int64)),
		},
		{
			"xTime",
			&xTime,
			func() interface{} { return xTime },
			interface{}(time.Now()),
		},
		{
			"xTimePtr",
			&xTimePtr,
			func() interface{} { return xTimePtr },
			interface{}(new(time.Time)),
		},
		{
			"xFloat32",
			&xFloat32,
			func() interface{} { return xFloat32 },
			interface{}(float32(567.987651012)),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			func() interface{} { return xFloat32Ptr },
			interface{}(new(float32)),
		},
		{
			"xFloat64",
			&xFloat64,
			func() interface{} { return xFloat64 },
			interface{}(float64(56789101112.987651012)),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			func() interface{} { return xFloat64Ptr },
			interface{}(new(float64)),
		},
		{
			"xByte",
			&xByte,
			func() interface{} { return xByte },
			interface{}([]byte(`Testing set refernet value for []byte`)),
		},
		{
			"xBytePtr",
			&xBytePtr,
			func() interface{} { return xBytePtr },
			interface{}(new([]byte)),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		vr.SetReferentValue(expected)
		actual := testCase.Actual()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Testing x as %v: SetReferentValue() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}
}

func TestGetReferentValue(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Assign   func(interface{})
		Expected interface{}
	}{
		{
			"xInt",
			&xInt,
			func(i interface{}) { xInt = i.(int) },
			interface{}(5),
		},
		{
			"xIntPtr",
			&xIntPtr,
			func(i interface{}) { xIntPtr = i.(*int) },
			interface{}(new(int)),
		},
		{
			"xString",
			&xString,
			func(i interface{}) { xString = i.(string) },
			interface{}("Test string"),
		},
		{
			"xStringPtr",
			&xStringPtr,
			func(i interface{}) { xStringPtr = i.(*string) },
			interface{}(new(string)),
		},
		{
			"xInt64",
			&xInt64,
			func(i interface{}) { xInt64 = i.(int64) },
			interface{}(int64(9999999999999)),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			func(i interface{}) { xInt64Ptr = i.(*int64) },
			interface{}(new(int64)),
		},
		{
			"xTime",
			&xTime,
			func(i interface{}) { xTime = i.(time.Time) },
			interface{}(time.Now()),
		},
		{
			"xTimePtr",
			&xTimePtr,
			func(i interface{}) { xTimePtr = i.(*time.Time) },
			interface{}(new(time.Time)),
		},
		{
			"xFloat32",
			&xFloat32,
			func(i interface{}) { xFloat32 = i.(float32) },
			interface{}(float32(567.987651012)),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			func(i interface{}) { xFloat32Ptr = i.(*float32) },
			interface{}(new(float32)),
		},
		{
			"xFloat64",
			&xFloat64,
			func(i interface{}) { xFloat64 = i.(float64) },
			interface{}(float64(56789101112.987651012)),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			func(i interface{}) { xFloat64Ptr = i.(*float64) },
			interface{}(new(float64)),
		},
		{
			"xByte",
			&xByte,
			func(i interface{}) { xByte = i.([]byte) },
			interface{}([]byte(`Testing set refernet value for []byte`)),
		},
		{
			"xBytePtr",
			&xBytePtr,
			func(i interface{}) { xBytePtr = i.(*[]byte) },
			interface{}(new([]byte)),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		testCase.Assign(expected)
		actual := vr.GetReferentValue()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Testing x as %v: GetReferentValue() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}
}

func TestGetReferentType(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Expected reflect.Type
	}{
		{
			"xInt",
			&xInt,
			reflect.TypeOf(int(0)),
		},
		{
			"xIntPtr",
			&xIntPtr,
			reflect.TypeOf((*int)(nil)),
		},
		{
			"xString",
			&xString,
			reflect.TypeOf(string("")),
		},
		{
			"xStringPtr",
			&xStringPtr,
			reflect.TypeOf((*string)(nil)),
		},
		{
			"xInt64",
			&xInt64,
			reflect.TypeOf(int64(0)),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			reflect.TypeOf((*int64)(nil)),
		},
		{
			"xTime",
			&xTime,
			reflect.TypeOf(time.Time{}),
		},
		{
			"xTimePtr",
			&xTimePtr,
			reflect.TypeOf((*time.Time)(nil)),
		},
		{
			"xFloat32",
			&xFloat32,
			reflect.TypeOf(float32(0)),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			reflect.TypeOf((*float32)(nil)),
		},
		{
			"xFloat64",
			&xFloat64,
			reflect.TypeOf(float64(0)),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			reflect.TypeOf((*float64)(nil)),
		},
		{
			"xByte",
			&xByte,
			reflect.TypeOf([]byte(nil)),
		},
		{
			"xBytePtr",
			&xBytePtr,
			reflect.TypeOf((*[]byte)(nil)),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		actual := vr.GetReferentType()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Testing x as %v: GetReferentType() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}
}

func TestIsReferentPtr(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Expected bool
	}{
		{
			"xInt",
			&xInt,
			false,
		},
		{
			"xIntPtr",
			&xIntPtr,
			true,
		},
		{
			"xString",
			&xString,
			false,
		},
		{
			"xStringPtr",
			&xStringPtr,
			true,
		},
		{
			"xInt64",
			&xInt64,
			false,
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			true,
		},
		{
			"xTime",
			&xTime,
			false,
		},
		{
			"xTimePtr",
			&xTimePtr,
			true,
		},
		{
			"xFloat32",
			&xFloat32,
			false,
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			true,
		},
		{
			"xFloat64",
			&xFloat64,
			false,
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			true,
		},
		{
			"xByte",
			&xByte,
			false,
		},
		{
			"xBytePtr",
			&xBytePtr,
			true,
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		actual := vr.IsReferentPtr()
		if actual != expected {
			t.Errorf("Testing x as %v: IsReferentPtr() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}
}

func TestReInitializeReferentValueNil(t *testing.T) {
	var xIntPtr *int
	var xStringPtr *string
	var xInt64Ptr *int64
	var xTimePtr *time.Time
	var xFloat32Ptr *float32
	var xFloat64Ptr *float64
	var xBytePtr *[]byte

	testCases := []struct {
		Name        string
		X           interface{}
		Actual      func() interface{}
		ExpectedNot interface{}
	}{
		{
			"xIntPtr",
			&xIntPtr,
			func() interface{} { return xIntPtr },
			(*int)(nil),
		},
		{
			"xStringPtr",
			&xStringPtr,
			func() interface{} { return xStringPtr },
			(*string)(nil),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			func() interface{} { return xInt64Ptr },
			(*int64)(nil),
		},
		{
			"xTimePtr",
			&xTimePtr,
			func() interface{} { return xTimePtr },
			(*time.Time)(nil),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			func() interface{} { return xFloat32Ptr },
			(*float32)(nil),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			func() interface{} { return xFloat64Ptr },
			(*float64)(nil),
		},
		{
			"xBytePtr",
			&xBytePtr,
			func() interface{} { return xBytePtr },
			(*[]byte)(nil),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expectedNot := testCase.ExpectedNot

		vr := New(x)
		// Check that it was initialy nil
		actual := testCase.Actual()
		if actual != expectedNot {
			t.Errorf("Testing x as %v: initial value is %#v \n\t\t\twhile expected %#v", name, actual, expectedNot)
		}

		vr.ReInitializeReferentValue()
		actual = testCase.Actual()
		if actual == expectedNot {
			t.Errorf("Testing x as %v: ReInitializeReferentValue() assigned %#v \n\t\t\twhile expected NOT %#v", name, actual, expectedNot)
		}
	}
}

func TestReInitializeReferentValueNotNil(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Assign   func(interface{})
		Actual   func() interface{}
		Expected interface{}
	}{
		{
			"xInt",
			&xInt,
			func(i interface{}) { xInt = i.(int) },
			func() interface{} { return xInt },
			interface{}(5),
		},
		{
			"xIntPtr",
			&xIntPtr,
			func(i interface{}) { xIntPtr = i.(*int) },
			func() interface{} { return xIntPtr },
			interface{}(new(int)),
		},
		{
			"xString",
			&xString,
			func(i interface{}) { xString = i.(string) },
			func() interface{} { return xString },
			interface{}("Test string"),
		},
		{
			"xStringPtr",
			&xStringPtr,
			func(i interface{}) { xStringPtr = i.(*string) },
			func() interface{} { return xStringPtr },
			interface{}(new(string)),
		},
		{
			"xInt64",
			&xInt64,
			func(i interface{}) { xInt64 = i.(int64) },
			func() interface{} { return xInt64 },
			interface{}(int64(9999999999999)),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			func(i interface{}) { xInt64Ptr = i.(*int64) },
			func() interface{} { return xInt64Ptr },
			interface{}(new(int64)),
		},
		{
			"xTime",
			&xTime,
			func(i interface{}) { xTime = i.(time.Time) },
			func() interface{} { return xTime },
			interface{}(time.Now()),
		},
		{
			"xTimePtr",
			&xTimePtr,
			func(i interface{}) { xTimePtr = i.(*time.Time) },
			func() interface{} { return xTimePtr },
			interface{}(new(time.Time)),
		},
		{
			"xFloat32",
			&xFloat32,
			func(i interface{}) { xFloat32 = i.(float32) },
			func() interface{} { return xFloat32 },
			interface{}(float32(567.987651012)),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			func(i interface{}) { xFloat32Ptr = i.(*float32) },
			func() interface{} { return xFloat32Ptr },
			interface{}(new(float32)),
		},
		{
			"xFloat64",
			&xFloat64,
			func(i interface{}) { xFloat64 = i.(float64) },
			func() interface{} { return xFloat64 },
			interface{}(float64(56789101112.987651012)),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			func(i interface{}) { xFloat64Ptr = i.(*float64) },
			func() interface{} { return xFloat64Ptr },
			interface{}(new(float64)),
		},
		{
			"xByte",
			&xByte,
			func(i interface{}) { xByte = i.([]byte) },
			func() interface{} { return xByte },
			interface{}([]byte(`Testing set refernet value for []byte`)),
		},
		{
			"xBytePtr",
			&xBytePtr,
			func(i interface{}) { xBytePtr = i.(*[]byte) },
			func() interface{} { return xBytePtr },
			interface{}(new([]byte)),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		testCase.Assign(expected)
		vr.ReInitializeReferentValue()
		actual := testCase.Actual()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Testing x as %v: ReInitializeReferentValue() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
	}
}

func TestReinitializeReferentValueByReflect(t *testing.T) {

	x := (*int)(nil)
	expected := 0

	vr := *(New(&x)).(*ValueReference)

	initialPtr := getReferentIntPtrValue(&vr).(*int)
	if initialPtr != (*int)(nil) {
		t.Logf("Initial value is %#v while expected %#v", initialPtr, (*int)(nil))
	}

	reinitializeReferentValueByReflect(&vr)
	actualPtr := getReferentIntPtrValue(&vr).(*int)
	if actualPtr == (*int)(nil) {
		t.Errorf("actualPtr value is %#v while expected NOT NIL", actualPtr)
	}
	actual := *(actualPtr)
	if actual != expected {
		t.Errorf("reinitializeReferentValueByReflect returned %#v while expected %#v", actual, expected)
	}
}

/*
func Test(t *testing.T) {
    expected :=

    actual :=
	if actual != expected {
		t.Errorf(" returned %#v while expected %#v", actual, expected)
	}
}
*/

//--------------------------------------//
//             Benchmarks               //
//--------------------------------------//

/*
func Benchmark(b *testing.B) {
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		s :=
		s += ""
	}
	b.ReportAllocs()
}
*/
