
/*
*/
package parser

const numNTSymbols = 6
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // StmtList
		2, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // StmtList
		6, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		8, // Expr
		9, // Term
		10, // Factor
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		20, // Expr
		21, // Term
		22, // Factor
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		27, // Term
		10, // Factor
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		28, // Term
		10, // Factor
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		29, // Factor
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		30, // Factor
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		36, // Expr
		21, // Term
		22, // Factor
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		37, // Term
		22, // Factor
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		38, // Term
		22, // Factor
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		39, // Factor
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		40, // Factor
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // StmtList
		-1, // Stmt
		-1, // Expr
		-1, // Term
		-1, // Factor
		

	},
	
}
