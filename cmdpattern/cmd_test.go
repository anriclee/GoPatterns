package cmdpattern

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	flags := stripFlags([]string{"new", "theme", "-force=true"})
	fmt.Println(flags)
}
