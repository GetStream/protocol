package generator

import (
	"path"
)

func swiftFileName(name string) string {
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		base := path.Base(name)
		name = base[:len(base)-len(path.Ext(base))]
	}

	name += ".model.swift"

	return name
}
