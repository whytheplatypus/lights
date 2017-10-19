use 'godoc cmd/github.com/whytheplatypus/lights/razer' for documentation on the github.com/whytheplatypus/lights/razer command 







`import github.com/whytheplatypus/lights/razer`



# Variables


```go
var Conn *dbus.Conn
```


```go
var DeviceList []string
```


```go
var Keyboards = map[string]map[string][]struct{ Col, Row uint8 }{
    "Razer Blade Stealth (Late 2016)": {
        "Escape": {{1, 0}},
        "F1":     {{2, 0}},
        "F2":     {{3, 0}},
        "F3":     {{4, 0}},
        "F4":     {{5, 0}},
        "F5":     {{6, 0}},
        "F6":     {{7, 0}},
        "F7":     {{8, 0}},
        "F8":     {{9, 0}},
        "F9":     {{10, 0}},
        "F10":    {{11, 0}},
        "F11":    {{12, 0}},
        "F12":    {{13, 0}},
        "Insert": {{14, 0}},
        "Delete": {{15, 0}},

        "`":         {{1, 1}},
        "1":         {{2, 1}},
        "2":         {{3, 1}},
        "3":         {{4, 1}},
        "4":         {{5, 1}},
        "5":         {{6, 1}},
        "6":         {{7, 1}},
        "7":         {{8, 1}},
        "8":         {{9, 1}},
        "9":         {{10, 1}},
        "0":         {{11, 1}},
        "-":         {{12, 1}},
        "=":         {{13, 1}},
        "BackSpace": {{15, 1}},

        "Tab": {{1, 2}},
        "q":   {{2, 2}},
        "w":   {{3, 2}},
        "e":   {{4, 2}},
        "r":   {{5, 2}},
        "t":   {{6, 2}},
        "y":   {{7, 2}},
        "u":   {{8, 2}},
        "i":   {{9, 2}},
        "o":   {{10, 2}},
        "p":   {{11, 2}},
        "[":   {{12, 2}},
        "]":   {{13, 2}},
        "\\":  {{15, 2}},

        "Caps_Lock": {{1, 3}},
        "a":         {{2, 3}},
        "s":         {{3, 3}},
        "d":         {{4, 3}},
        "f":         {{5, 3}},
        "g":         {{6, 3}},
        "h":         {{7, 3}},
        "j":         {{8, 3}},
        "k":         {{9, 3}},
        "l":         {{10, 3}},
        ";":         {{11, 3}},
        "'":         {{12, 3}},
        "Return":    {{15, 3}},

        "Shift_L": {
            {1, 4},
            {2, 4},
        },
        "z": {{3, 4}},
        "x": {{4, 4}},
        "c": {{5, 4}},
        "v": {{6, 4}},
        "b": {{7, 4}},
        "n": {{8, 4}},
        "m": {{9, 4}},
        ",": {{10, 4}},
        ".": {{11, 4}},
        "/": {{12, 4}},
        "Shift_R": {
            {13, 4},
            {14, 4},
            {15, 4},
        },

        "Control_L": {{1, 5}},
        "Super_L":   {{3, 5}},
        "Alt_L":     {{4, 5}},
        "Alt_R":     {{9, 5}},
        "Control_R": {{11, 5}},
        "Left":      {{12, 5}},
        "Up":        {{13, 5}},
        "Down":      {{14, 5}},
        "Right":     {{15, 5}},
    },
}
```


```go
var Keys = []string{
    "Escape",
    "F1",
    "F2",
    "F3",
    "F4",
    "F5",
    "F6",
    "F7",
    "F8",
    "F9",
    "F10",
    "F11",
    "F12",
    "Insert",
    "Delete",

    "`",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "0",
    "-",
    "=",
    "BackSpace",

    "Tab",
    "q",
    "w",
    "e",
    "r",
    "t",
    "y",
    "u",
    "i",
    "o",
    "p",
    "[",
    "]",
    "\\",

    "Caps_Lock",
    "a",
    "s",
    "d",
    "f",
    "g",
    "h",
    "j",
    "k",
    "l",
    ";",
    "'",
    "Return",

    "Shift_L",
    "z",
    "x",
    "c",
    "v",
    "b",
    "n",
    "m",
    ",",
    ".",
    "/",
    "Shift_R",

    "Control_L",
    "Super_L",
    "Alt_L",
    "Alt_R",
    "Control_R",
    "Left",
    "Up",
    "Down",
    "Right",
}
```



# Functions


## [Apply](razer.go#L103)
```go
func Apply(name string, conn *dbus.Conn, s *Set)
```





## [ClearCustom](razer.go#L132)
```go
func ClearCustom(name string, conn *dbus.Conn, color Color)
```





## [GetDeviceName](razer.go#L166)
```go
func GetDeviceName(name string, conn *dbus.Conn) (string, error)
```





## [GetMatrixDimensions](razer.go#L173)
```go
func GetMatrixDimensions(name string, conn *dbus.Conn) ([]int32, error)
```





## [SetBreathRandom](razer.go#L158)
```go
func SetBreathRandom(name string, conn *dbus.Conn)
```





## [SetCustom](razer.go#L116)
```go
func SetCustom(name string, conn *dbus.Conn)
```





## [SetReactive](razer.go#L124)
```go
func SetReactive(name string, conn *dbus.Conn, color Color)
```





## [ValidDevice](razer.go#L33)
```go
func ValidDevice(d string) bool
```






# Types


## [Color](razer.go#L54)

```go
type Color interface {
    RGB() []uint8
}
```









## [RGB](razer.go#L46)

```go
type RGB struct {
    // contains filtered or unexported fields
}
```







### [RGB](razer.go#L50)
```go
func (rgb *RGB) RGB() []uint8
```







## [RGBA](razer.go#L38)

```go
type RGBA struct {
    color.RGBA
}
```







### [RGB](razer.go#L42)
```go
func (rgba *RGBA) RGB() []uint8
```







## [Row](razer.go#L59)
TODO rename

```go
type Row struct {
    Num uint8
    // can make this matter but requires updates to the driver I think?
    Start  uint8
    Colors []Color
}
```







### [Encode](razer.go#L66)
```go
func (r *Row) Encode(w io.Writer) error
```







## [Set](razer.go#L90)
TODO rename

```go
type Set struct {
    Rows []*Row
}
```







### [Encode](razer.go#L94)
```go
func (s *Set) Encode(w io.Writer) error
```










