
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
		String: `StmtList : Stmt	<< ast.NewStmtList(X[0].(*ast.Stmt)), nil >>`,
		Id: "StmtList",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0].(*ast.Stmt)), nil
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList ";" Stmt	<< ast.AppendStmtList(X[0].([]*ast.Stmt), X[2].(*ast.Stmt)), nil >>`,
		Id: "StmtList",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmtList(X[0].([]*ast.Stmt), X[2].(*ast.Stmt)), nil
		},
	},
	ProdTabEntry{
		String: `Stmt : "$" id "=" Expr	<< ast.NewOperatorFromParser(ast.NewLeftID(string(X[1].(*token.Token).Lit)), "=", X[3].(*ast.Stmt)) >>`,
		Id: "Stmt",
		NTType: 2,
		Index: 3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(ast.NewLeftID(string(X[1].(*token.Token).Lit)), "=", X[3].(*ast.Stmt))
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "+" Term	<< ast.NewOperatorFromParser(X[0].(*ast.Stmt), "+", X[2].(*ast.Stmt)) >>`,
		Id: "Expr",
		NTType: 3,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Stmt), "+", X[2].(*ast.Stmt))
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "-" Term	<< ast.NewOperatorFromParser(X[0].(*ast.Stmt), "-", X[2].(*ast.Stmt)) >>`,
		Id: "Expr",
		NTType: 3,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Stmt), "-", X[2].(*ast.Stmt))
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
		String: `Term : Term "*" Factor	<< ast.NewOperatorFromParser(X[0].(*ast.Stmt), "*", X[2].(*ast.Stmt)) >>`,
		Id: "Term",
		NTType: 4,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Stmt), "*", X[2].(*ast.Stmt))
		},
	},
	ProdTabEntry{
		String: `Term : Term "/" Factor	<< ast.NewOperatorFromParser(X[0].(*ast.Stmt), "/", X[2].(*ast.Stmt)) >>`,
		Id: "Term",
		NTType: 4,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperatorFromParser(X[0].(*ast.Stmt), "/", X[2].(*ast.Stmt))
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
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
		String: `Factor : int	<< ast.NewIntFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 5,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : float	<< ast.NewFloatFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFloatFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Factor : Date	<< X[0], nil >>`,
		Id: "Factor",
		NTType: 5,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : Env	<< X[0], nil >>`,
		Id: "Factor",
		NTType: 5,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : Id	<< X[0], nil >>`,
		Id: "Factor",
		NTType: 5,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : string	<< ast.NewStringFromParser(string(X[0].(*token.Token).Lit)) >>`,
		Id: "Factor",
		NTType: 5,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringFromParser(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Id : "$" id	<< ast.NewRightIDFromParser(string(X[1].(*token.Token).Lit)) >>`,
		Id: "Id",
		NTType: 6,
		Index: 17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRightIDFromParser(string(X[1].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Date : "${date:" date "}"	<< ast.NewDateFromParser(string(X[1].(*token.Token).Lit)) >>`,
		Id: "Date",
		NTType: 7,
		Index: 18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDateFromParser(string(X[1].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Env : "${env:" id "}"	<< ast.NewEnvFromParser(string(X[1].(*token.Token).Lit)) >>`,
		Id: "Env",
		NTType: 8,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEnvFromParser(string(X[1].(*token.Token).Lit))
		},
	},
	
}
