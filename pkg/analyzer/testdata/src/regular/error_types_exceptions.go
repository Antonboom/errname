package regular

import "strings"

type ValidationErrors []string

func (ve ValidationErrors) Error() string {
	return strings.Join(ve, "\n")
}

type TenErrors [10]string

func (te TenErrors) Error() string {
	return strings.Join(te[:], "\n")
}
