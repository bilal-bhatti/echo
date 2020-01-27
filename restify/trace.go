package restify

import (
	"go/ast"
	"log"
)

func Inspect(node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		// log.Println("node", n)
		switch t := n.(type) {
		case *ast.File:
			log.Println("file", t)
		case *ast.GenDecl:
			log.Println("gen decl", t)
		case *ast.FieldList:
			log.Println("file list", t)
		case *ast.FuncType:
			log.Println("func", t)
		case *ast.BasicLit:
			log.Println("basic lit", t)
		case *ast.ImportSpec:
			log.Printf("import spec %v", t)
		case *ast.Ident:
			log.Println("ident", t)
		case *ast.Comment:
			log.Println("comment", t)
		case *ast.CommentGroup:
			log.Println("comment group", t)
		case *ast.FuncDecl:
			log.Println("function", t)
		case *ast.BlockStmt:
			log.Println("block statement", t)
		case *ast.AssignStmt:
			log.Println("assignment statement", t)
		case *ast.CallExpr:
			log.Println("call expression", t)
		case *ast.ExprStmt:
			log.Println("expression statement", t)
		case *ast.RangeStmt:
			log.Println("range statement", t)
		case *ast.Field:
			log.Println("field", t)
		case *ast.SelectorExpr:
			log.Println("selector expression", t)
		case *ast.IfStmt:
			log.Println("if statement", t)
		case *ast.BinaryExpr:
			log.Println("binary expression", t)
		case *ast.FuncLit:
			log.Println("func lit", t)
		case *ast.TypeSwitchStmt:
			log.Println("type switch statement", t)
		case *ast.TypeAssertExpr:
			log.Println("type assert statement", t)
		case *ast.CaseClause:
			log.Println("case clause", t)
		case *ast.StarExpr:
			log.Println("star expression", t)
		case *ast.ReturnStmt:
			log.Println("return statement", t)
		case *ast.UnaryExpr:
			log.Println("unary expression", t)
		case *ast.CompositeLit:
			log.Println("composite lit", t)
		case *ast.KeyValueExpr:
			log.Println("key value expression", t)
		case *ast.ArrayType:
			log.Println("array type", t)
		case *ast.DeclStmt:
			log.Println("decl statement", t)
		case *ast.ValueSpec:
			log.Println("value spec", t)
		case *ast.TypeSpec:
			log.Println("type spec", t)
		case *ast.StructType:
			log.Println("struct type", t)
		case *ast.MapType:
			log.Println("map type", t)
		case *ast.ForStmt:
			log.Println("for statement", t)
		case *ast.IncDecStmt:
			log.Println("inc dec statement", t)
		case *ast.ParenExpr:
			log.Println("paren expression", t)
		case *ast.IndexExpr:
			log.Println("index expression", t)
		case *ast.SliceExpr:
			log.Println("slice expression", t)
		case *ast.InterfaceType:
			log.Println("interface type", t)
		case *ast.ChanType:
			log.Println("chan type", t)
		case *ast.GoStmt:
			log.Println("go statement", t)
		case *ast.DeferStmt:
			log.Println("defer statement", t)
		default:
			// log.Println("Not handling declaration type ", reflect.TypeOf(t))
			PrintP(t)
		}
		return true
	})
}
