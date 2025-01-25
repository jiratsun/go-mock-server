package strings

import (
	"fmt"
	"strings"
)

type StringBuilder struct {
	builder strings.Builder
}

func (wrapper *StringBuilder) String() string {
	return wrapper.builder.String()
}

func (wrapper *StringBuilder) WriteString(s string) *StringBuilder {
	wrapper.builder.WriteString(s)
	return wrapper
}

func (wrapper *StringBuilder) WriteStringf(s string, a ...any) *StringBuilder {
	wrapper.builder.WriteString(fmt.Sprintf(s, a...))
	return wrapper
}

func (wrapper *StringBuilder) WriteStringln(s string) *StringBuilder {
	wrapper.builder.WriteString(s + "\n")
	return wrapper
}

func (wrapper *StringBuilder) WriteStringRepeat(s string, count int) *StringBuilder {
	wrapper.builder.WriteString(strings.Repeat(s, count))
	return wrapper
}

func (wrapper *StringBuilder) WriteStringlnRepeat(s string, count int) *StringBuilder {
	wrapper.builder.WriteString(strings.Repeat(s+"\n", count))
	return wrapper
}
