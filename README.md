use 'godoc cmd/github.com/whytheplatypus/lights' for documentation on the github.com/whytheplatypus/lights command 









# Variables


```go
var (
    // BuildTime is supplied by the compiler as the time at which the binary was built
    BuildTime string
    // GitCommit is supplied by the compiler as the most recent commit the binary was built from
    GitCommit string
    // Version is supplied by the compiler as the most recent git tag the binary was built from
    // defaults to 0.0.1
    Version string
    // VersionDescription is a modifier to Version that describes the binary build
    VersionDescription = "dev"

    // Subcmds are the possible excutable sub-commands for this program
    Subcmds = CmdRegistry{}
)
```



# Functions


## [Description](main.go#L112)
```go
func Description() string
```
Description returns a string describing the binary build bartender
<version>(-<VersionDescription>) ( :: commit - <GitCommit> [ :: built @
<BuildTime> ] )





## [Short](main.go#L131)
```go
func Short() string
```






# Types


## [CmdRegistry](main.go#L55)

```go
type CmdRegistry map[string]Runnable
```







### [Register](main.go#L57)
```go
func (c CmdRegistry) Register(name string, cmd Runnable)
```





### [Run](main.go#L64)
```go
func (c CmdRegistry) Run(args []string) error
```







## [RunFunc](main.go#L49)

```go
type RunFunc func(args []string) error
```







### [Run](main.go#L51)
```go
func (r RunFunc) Run(args []string) error
```







## [Runnable](main.go#L45)

```go
type Runnable interface {
    Run(args []string) error
}
```











#Notes

##Todos

- this panics and only reports info about -v, need this and any info about
commands?


- clean this error text up


- dup of line 63-68


- this should go at either the bottom or top of the file for easy location





# Packages

	
[commands](commands)  
	
	
[razer](razer)  
	


