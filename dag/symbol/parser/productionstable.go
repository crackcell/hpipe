
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
		String: `Expr : Expr "+" Term	<< ast.NewOperatorFromParser(X[0].(*ast.Expr), "+", X[2].(*ast.Expr)) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Expr), "+", X[2].(*ast.Expr))
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "-" Term	<< ast.NewOperatorFromParser(X[0].(*ast.Expr), "-", X[2].(*ast.Expr)) >>`,
		Id: "Expr",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Expr), "-", X[2].(*ast.Expr))
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
		String: `Term : Term "*" Factor	<< ast.NewOperatorFromParser(X[0].(*ast.Expr), "*", X[2].(*ast.Expr)) >>`,
		Id: "Term",
		NTType: 2,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Expr), "*", X[2].(*ast.Expr))
		},
	},
	ProdTabEntry{
		String: `Term : Term "/" Factor	<< ast.NewOperatorFromParser(X[0].(*ast.Expr), "/", X[2].(*ast.Expr)) >>`,
		Id: "Term",
		NTType: 2,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Expr), "/", X[2].(*ast.Expr))
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
		String: `Factor : int	<< ast.NewIntFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : date	<< ast.NewDateFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDateFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : id	<< ast.NewVarFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : string	<< ast.NewStringFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	
}
