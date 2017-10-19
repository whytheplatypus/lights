use 'godoc cmd/github.com/whytheplatypus/lights/commands' for documentation on the github.com/whytheplatypus/lights/commands command 







`import github.com/whytheplatypus/lights/commands`




# Functions


## [Hex](render.go#L162)
```go
func Hex(scol string, contrast uint8) (color.RGBA, error)
```






# Types


## [Clear](clear.go#L8)

```go
type Clear struct{}
```







### [Run](clear.go#L10)
```go
func (r *Clear) Run(args []string) error
```







## [Renderer](render.go#L20)

```go
type Renderer struct{}
```







### [Run](render.go#L22)
```go
func (r *Renderer) Run(args []string) error
```









#Notes

##Todos

- same as before why 0? should be something like GetKeyboard


- what does this mean that it's the first device?


- why set it back to this? what's going on here?


- better for this to be a named function so it's clear what's happening


- this should be the first thing that happens


- named function to be clear we're making a fifo


- can we not use these?


- I've seen ths before, that's no good


- encapsulate to be clear we're getting a keyboard


- what if there's no keyboard


- what is this? a render loop?


- these 3 or 4 statements feel jumbled


- this is a parser/unmarshler, whtever


- it's rendering into s and sending s to a pipe, this should be made clear via
code and function names


- return don't panic




