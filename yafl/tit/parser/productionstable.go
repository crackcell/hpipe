
package parser

import (
	"../token"
	"../ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : StmtList	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt	<< ast.NewStmtList(X[0]) >>`,
		Id: "StmtList",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0])
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList ";" Stmt	<< ast.AppendStmtList(X[0], X[2]) >>`,
		Id: "StmtList",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmtList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Stmt : id "=" Expr	<< ast.NewOperator(ast.NewLeftID(string(X[0].(*token.Token).Lit)), "=", X[2]) >>`,
		Id: "Stmt",
		NTType: 2,
		Index: 3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(ast.NewLeftID(string(X[0].(*token.Token).Lit)), "=", X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "+" Term	<< ast.NewOperator(X[0], "+", X[2]) >>`,
		Id: "Expr",
		NTType: 3,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "+", X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "-" Term	<< ast.NewOperator(X[0], "-", X[2]) >>`,
		Id: "Expr",
		NTType: 3,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "-", X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id: "Expr",
		NTType: 3,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term "*" Factor	<< ast.NewOperator(X[0], "*", X[2]) >>`,
		Id: "Term",
		NTType: 4,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "*", X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Term "/" Factor	<< ast.NewOperator(X[0], "/", X[2]) >>`,
		Id: "Term",
		NTType: 4,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "/", X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<< X[0], nil >>`,
		Id: "Term",
		NTType: 4,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id: "Factor",
		NTType: 5,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : string	<< ast.NewValueFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 5,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewValueFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : int64	<< ast.NewInt64FromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInt64FromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : id	<< ast.NewRightID(string(X[0].(*token.Token).Lit)), nil >>`,
		Id: "Factor",
		NTType: 5,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRightID(string(X[0].(*token.Token).Lit)), nil
		},
	},
	
}
