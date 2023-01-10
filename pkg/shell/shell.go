package shell

var GlobalMakeStyle = false

func MakeStyle() {
	GlobalMakeStyle = !GlobalMakeStyle
}
