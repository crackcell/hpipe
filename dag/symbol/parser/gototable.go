
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
		2, // Expr
		3, // Term
		4, // Factor
		7, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S5
		
		-1, // S'
		15, // Expr
		16, // Term
		17, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		22, // Builtins
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Expr
		24, // Term
		4, // Factor
		7, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Expr
		25, // Term
		4, // Factor
		7, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Expr
		-1, // Term
		26, // Factor
		7, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Expr
		-1, // Term
		27, // Factor
		7, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S18
		
		-1, // S'
		34, // Expr
		16, // Term
		17, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		35, // Builtins
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Expr
		37, // Term
		17, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Expr
		38, // Term
		17, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Expr
		-1, // Term
		39, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Expr
		-1, // Term
		40, // Factor
		20, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		-1, // Builtins
		

	},
	
}
