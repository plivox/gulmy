package shell

const (
	Windows = "windows"
	Darwin  = "darwin"
	Linux   = "linux"
)

var GlobalMakeStyle = false

func MakeStyle() {
	GlobalMakeStyle = !GlobalMakeStyle
}
