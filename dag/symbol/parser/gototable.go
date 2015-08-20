
/*
*/
package parser

const numNTSymbols = 5
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
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S5
		
		-1, // S'
		15, // Expr
		16, // Term
		17, // Factor
		20, // Variable
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Expr
		23, // Term
		4, // Factor
		7, // Variable
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Expr
		24, // Term
		4, // Factor
		7, // Variable
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Expr
		-1, // Term
		25, // Factor
		7, // Variable
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Expr
		-1, // Term
		26, // Factor
		7, // Variable
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S18
		
		-1, // S'
		33, // Expr
		16, // Term
		17, // Factor
		20, // Variable
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Expr
		36, // Term
		17, // Factor
		20, // Variable
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Expr
		37, // Term
		17, // Factor
		20, // Variable
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Expr
		-1, // Term
		38, // Factor
		20, // Variable
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Expr
		-1, // Term
		39, // Factor
		20, // Variable
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Expr
		-1, // Term
		-1, // Factor
		-1, // Variable
		

	},
	
}
