
package parser

import(
  "github.com/crackcell/hpipe/dag/symbol/token"
  "github.com/crackcell/hpipe/dag/symbol/ast"
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
		String: `S' : Expr	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "+" Term	<< ast.NewOperator(X[0], "+", X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "+", X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "-" Term	<< ast.NewOperator(X[0], "-", X[2]) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "-", X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id: "Expr",
		NTType: 1,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term "*" Factor	<< ast.NewOperator(X[0], "*", X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "*", X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Term "/" Factor	<< ast.NewOperator(X[0], "/", X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperator(X[0], "/", X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 2,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id: "Factor",
		NTType: 3,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int64	<< ast.NewInt64FromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInt64FromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : Variable	<<  >>`,
		Id: "Factor",
		NTType: 3,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Variable : "${" date "}"	<< ast.NewBuiltinVarFromParser(string(X[1].(*token.Token).Lit)) >>`,
		Id: "Variable",
		NTType: 4,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBuiltinVarFromParser(string(X[1].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Variable : "$" id	<< ast.NewVarFromParser(string(X[1].(*token.Token).Lit)) >>`,
		Id: "Variable",
		NTType: 4,
		Index: 11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarFromParser(string(X[1].(*token.Token).Lit))
		},
	},
	
}
