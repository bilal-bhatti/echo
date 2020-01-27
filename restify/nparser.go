package restify

import (
	"bytes"
	"context"
	"flag"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

func NewStdPackageParser() *StdPackageParser {
	return &StdPackageParser{}
}

func (p *StdPackageParser) Start() {
	pkgs := load()

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, d := range file.Decls {
				switch t := d.(type) {
				case *ast.FuncDecl:
					if isBindingSpecNode(pkg.TypesInfo, t) {

						// fset := token.NewFileSet()
						// var buf bytes.Buffer
						// printer.Fprint(&buf, fset, t)
						// log.Println("body", string(buf.Bytes()))

						p.Packets = append(p.Packets, &StdGenPacket{
							Pkg:      pkg,
							Bindings: t,
						})
					}
				}
			}
		}
	}

	for _, packet := range p.Packets {
		// 	wr := astutil.Apply(packet.Bindings, func(c *astutil.Cursor) bool {
		// 		log.Println("pre", c.Node())

		// 		return true
		// 	}, func(c *astutil.Cursor) bool {
		// 		return true
		// 	})

		// 	fset := token.NewFileSet()
		// 	var buf bytes.Buffer
		// 	printer.Fprint(&buf, fset, wr)
		// 	log.Println("body", string(buf.Bytes()))
		// rewrite(packet.Pkg.TypesInfo, packet.Bindings)
		prepare(packet)
		parseSpecExpression(packet)

	}

	// for _, packet := range p.Packets {
	// 	log.Println("found router", packet.Bindings)
	// 	// log.Println("package info", packet.Pkg.TypesInfo)

	// 	prepare(packet)
	// 	// log.Println(packet.Preamble)
	// 	bindings := parseSpecExpression(packet)

	// 	for _, binding := range bindings {
	// 		log.Println("bindings", binding)
	// 		for _, p := range binding.Params {
	// 			log.Println("param", p)
	// 		}
	// 		for _, r := range binding.Returns {
	// 			log.Println("return", r)
	// 		}
	// 	}
	// }

	// tmpl, err := template.New("bindings").Parse(templates.Header)
	// if err != nil {
	// 	panic(err)
	// }

	// metadata := &GenMetadata{
	// 	PackageName: packet.Pkg.Name,
	// }

	// err = tmpl.Execute(os.Stdout, metadata)
	// if err != nil {
	// 	panic(err)
	// }
}

// mux.Get("/things/{id}", zipline.Get(things.GetOne))
// func rewrite(info *types.Info, funk *ast.FuncDecl) {
// 	f := astcopy.FuncDecl(funk)
// 	fset := token.NewFileSet()
// 	// ast.Print(fset, f)

// 	wr := astutil.Apply(f, func(c *astutil.Cursor) bool {
// 		return true
// 	}, func(c *astutil.Cursor) bool {
// 		switch t := c.Node().(type) {
// 		case *ast.CallExpr:
// 			for i := 0; i < len(t.Args); i++ {

// 				arg := t.Args[i]
// 				if call, ok := arg.(*ast.CallExpr); ok {
// 					if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
// 						if id, ok := sel.X.(*ast.Ident); ok {
// 							if id.Name == "zipline" {
// 								t.Args[i] = ast.NewIdent("hi")
// 							}
// 						}
// 					}
// 				}
// 			}

// 			return true
// 		default:
// 			return true
// 		}
// 	})

// 	// fset := token.NewFileSet()
// 	var buf bytes.Buffer
// 	printer.Fprint(&buf, fset, wr)
// 	log.Println("body", string(buf.Bytes()))
// }

func rewrite(info *types.Info, funk *ast.FuncDecl) {
	// copy := astcopy.FuncDecl(funk)

	// fset := token.NewFileSet()
	// ast.Print(fset, f)

	for _, stmt := range funk.Body.List {
		switch stmtType := stmt.(type) {
		case *ast.ExprStmt:
			// log.Println("stmt", stmtType, stmtType.X, reflect.TypeOf(stmtType.X))

			switch expType := stmtType.X.(type) {
			case *ast.CallExpr:

				for i := 0; i < len(expType.Args); i++ {
					arg := expType.Args[i]
					// log.Println("arg", arg, reflect.TypeOf(arg))

					if call, ok := arg.(*ast.CallExpr); ok {
						if isBindingSpecNode(info, call) {
							// log.Println("call fun", reflect.TypeOf(call.Fun))
							if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
								if id, ok := sel.X.(*ast.Ident); ok {
									if id.Name == "zipline" {

										generate(info, sel)
										// expType.Args[i] = ast.NewIdent(generate(info, sel))
									}
								}
							} else {
								log.Println("Somethinger other than call to zipling found")
							}
						}
					}
				}
			default:
				log.Println("Unhandled exp type", reflect.TypeOf(expType))
			}

		default:
			// log.Println("Unhandled stmt type", reflect.TypeOf(stmtType))
		}
	}

	fset := token.NewFileSet()
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, funk)
	log.Println("func body \n", string(buf.Bytes()))

}

func generate(info *types.Info, sel *ast.SelectorExpr) string {
	// obj := qualifiedIdentObject(info, call.Fun)
	obj := qualifiedIdentObject(info, sel)
	log.Println("obj", obj)

	sig := obj.Type().(*types.Signature)

	for i := 0; i < sig.Params().Len(); i++ {
		p := sig.Params().At(i)
		log.Println("p", p.Name(), p.Type().String())
	}

	for i := 0; i < sig.Results().Len(); i++ {
		r := sig.Results().At(i)
		log.Println("r", r.Name(), r.Type().String())
	}

	// log.Println("sel.Sel", sel.Sel, reflect.TypeOf(sel.Sel))
	// log.Println("sel.X", sel.X, reflect.TypeOf(sel.X))

	return "boogey"
}

func pkgs(f *flag.FlagSet) []string {
	pkgs := f.Args()
	if len(pkgs) == 0 {
		pkgs = []string{"."}
	}
	return pkgs
}

func load() []*packages.Package {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("failed to get working directory: ", err)
		panic(err)
	}

	log.Println("Working directory", wd)

	cfg := &packages.Config{
		Context:    context.Background(),
		Mode:       packages.LoadAllSyntax,
		Dir:        wd,
		Env:        os.Environ(),
		BuildFlags: []string{"-tags=restified"},
		// TODO(light): Use ParseFile to skip function bodies and comments in indirect packages.
	}

	// Package pattern to search
	ps := []string{"./..."}
	// ps := pkgs(nil)

	pkgs, err := packages.Load(cfg, ps...)
	if err != nil {
		panic(err)
	}
	var errs []error
	for _, p := range pkgs {
		for _, e := range p.Errors {
			errs = append(errs, e)
		}
	}
	if len(errs) > 0 {
		log.Println(errs)
		panic(errs)
	}

	return pkgs
}

func prepare(packet *StdGenPacket) {
	dfunc := packet.Bindings
	body := bytes.Buffer{}

	for _, stmt := range dfunc.Body.List {
		switch stmtType := stmt.(type) {
		case *ast.AssignStmt:
			printNode(&body, stmtType)
			body.WriteString("\n")
		case *ast.ExprStmt:
			if isBindingSpecNode(packet.Pkg.TypesInfo, stmtType) {
				// quit, we have the preamble for the func
				packet.Specs = append(packet.Specs, stmtType)

				// found = false
				break // out of switch
			}
			printNode(&body, stmtType)
			body.WriteString("\n")
		case *ast.ReturnStmt:
			printNode(&body, stmtType)

		default:
			log.Println("Unhandled stmt type", reflect.TypeOf(stmtType))
		}

	}
	packet.Preamble = string(body.Bytes())
}

func parseSpecExpression(packet *StdGenPacket) []*Binding {
	bindings := []*Binding{}

	for _, spec := range packet.Specs {
		call, ok := spec.X.(*ast.CallExpr)
		if !ok {
			panic("spec invalid")
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			panic("spec invalid")
		}

		path := strings.Trim(call.Args[0].(*ast.BasicLit).Value, "\"")

		binding := NewBinding(sel.Sel.Name, path)

		zipline, ok := call.Args[1].(*ast.CallExpr)
		if !ok {
			panic("invalid expression")
		}

		handler := zipline.Args[0].(*ast.SelectorExpr)
		// log.Println("handler", handler.X, handler.Sel)
		obj := qualifiedIdentObject(packet.Pkg.TypesInfo, handler.Sel)
		log.Println("obj", obj)

		sig := obj.Type().(*types.Signature)

		for i := 0; i < sig.Params().Len(); i++ {
			p := sig.Params().At(i)
			binding.AddParam(p.Name(), p.Type().String())
		}

		for i := 0; i < sig.Results().Len(); i++ {
			r := sig.Results().At(i)
			binding.AddReturn(r.Name(), r.Type().String())
		}

		bindings = append(bindings, binding)
	}

	fset := token.NewFileSet()
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, packet.Bindings)
	log.Println("function body\n", string(buf.Bytes()))

	return bindings
}

func isBindingSpecNode(info *types.Info, fn ast.Node) bool {
	foundit := false

	ast.Inspect(fn, func(n ast.Node) bool {
		// skip everything other than call expression

		callExp, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExp, ok := callExp.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		// check the receiver is the package spec
		if x, ok := selExp.X.(*ast.Ident); !ok {
			return true
		} else {
			if x.Name != "zipline" {
				return true
			}
		}

		// // verify arguments
		if len(callExp.Args) != 1 {
			return true
		}

		buildObj := qualifiedIdentObject(info, callExp.Fun)

		if buildObj == nil || buildObj.Pkg() == nil || !isRestifyImport(buildObj.Pkg().Path()) || !isSupportedMethod(buildObj) {
			return true
		}

		// found what we were looking for
		foundit = true
		if foundit {
			return false
		}
		return true
	})
	return foundit
}
