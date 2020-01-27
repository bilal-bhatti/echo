package restify

import (
	"go/ast"

	parseutil "gopkg.in/src-d/go-parse-utils.v1"
)

type GenPacket struct {
	Pkg      *ast.Package
	Bindings *ast.FuncDecl
}

type Putil struct {
	Packets []*GenPacket
}

func (pu *Putil) ParsePackage(p string) {
	pkg, err := parseutil.PackageAST(p)
	if err != nil {
		panic(err)
	}

	// bindings := []ast.Node{}
	// packets := []*genpacket{}

	ast.Inspect(pkg, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if HasSpecBindings(t) {
				pu.Packets = append(pu.Packets, &GenPacket{
					Pkg:      pkg,
					Bindings: t,
				})

				return false
			}
		}
		return true
	})
}
