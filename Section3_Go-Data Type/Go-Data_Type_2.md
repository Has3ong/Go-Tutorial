# Go-Data_Type_2

#### Numeric Type
* Integer, Real Number, Complex Number
* 32bit, 64bit, unsigned

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

https://golang.org/ref/spec
```

##### Integer
* Octal(0) Decimal Hexadecimal(0x)
```
var num1 int = 1
var num2 int = -1
var num3 int = 012  // 10
var num4 int = 0x12 // 18
```

* ASCII, UNICODE
```
var ascii1 byte = 65    // 65, A
var ascii2 byte = 0102  // 66, B
var ascii3 byte = 0x43  // 67, C

var u1 rune = 44256     // 고
var u2 rune = 0126340   // 고
var u3 run = 0xACE0    // 고
```

##### Real Number
* float32, float64
```
var num1 float32 = 3.14159267
var num2 float32 = .1415926

var test float32 = 10.0
test - 0.1 = 9.9
float32(test - (0.1)) = 9.9
float64(test - (0.1)) = 9.899999618530273
```
* exponential notation
```
var num3 float32 = 1e-1         // 0.1
var num4 float32 = 15e3         // 15000
var num5 float32 = 0.234E+3     // 234
```

##### Complex Number
```
var num1 complex64 = 2 + 1i
num2 := complex64(2 + 1i)
num3 := complex(2, 1) //complex128
var num4 complex128 = 2 + 1i

real()  // real part
imag()  // imaginary part
```
