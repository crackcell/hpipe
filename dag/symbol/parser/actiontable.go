
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* int */
			shift(6),		/* date */
			shift(7),		/* id */
			shift(8),		/* string */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(9),		/* + */
			shift(10),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Expr */
			reduce(3),		/* +, reduce: Expr */
			reduce(3),		/* -, reduce: Expr */
			shift(11),		/* * */
			shift(12),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Term */
			reduce(6),		/* +, reduce: Term */
			reduce(6),		/* -, reduce: Term */
			reduce(6),		/* *, reduce: Term */
			reduce(6),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Factor */
			reduce(8),		/* +, reduce: Factor */
			reduce(8),		/* -, reduce: Factor */
			reduce(8),		/* *, reduce: Factor */
			reduce(8),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Factor */
			reduce(9),		/* +, reduce: Factor */
			reduce(9),		/* -, reduce: Factor */
			reduce(9),		/* *, reduce: Factor */
			reduce(9),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Factor */
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Factor */
			reduce(11),		/* +, reduce: Factor */
			reduce(11),		/* -, reduce: Factor */
			reduce(11),		/* *, reduce: Factor */
			reduce(11),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* int */
			shift(6),		/* date */
			shift(7),		/* id */
			shift(8),		/* string */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* int */
			shift(6),		/* date */
			shift(7),		/* id */
			shift(8),		/* string */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* int */
			shift(6),		/* date */
			shift(7),		/* id */
			shift(8),		/* string */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* int */
			shift(6),		/* date */
			shift(7),		/* id */
			shift(8),		/* string */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(25),		/* + */
			shift(26),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(27),		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(3),		/* +, reduce: Expr */
			reduce(3),		/* -, reduce: Expr */
			shift(28),		/* * */
			shift(29),		/* / */
			nil,		/* ( */
			reduce(3),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(6),		/* +, reduce: Term */
			reduce(6),		/* -, reduce: Term */
			reduce(6),		/* *, reduce: Term */
			reduce(6),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(8),		/* +, reduce: Factor */
			reduce(8),		/* -, reduce: Factor */
			reduce(8),		/* *, reduce: Factor */
			reduce(8),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(9),		/* +, reduce: Factor */
			reduce(9),		/* -, reduce: Factor */
			reduce(9),		/* *, reduce: Factor */
			reduce(9),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(11),		/* +, reduce: Factor */
			reduce(11),		/* -, reduce: Factor */
			reduce(11),		/* *, reduce: Factor */
			reduce(11),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Expr */
			reduce(1),		/* +, reduce: Expr */
			reduce(1),		/* -, reduce: Expr */
			shift(11),		/* * */
			shift(12),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Expr */
			reduce(2),		/* +, reduce: Expr */
			reduce(2),		/* -, reduce: Expr */
			shift(11),		/* * */
			shift(12),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Term */
			reduce(4),		/* +, reduce: Term */
			reduce(4),		/* -, reduce: Term */
			reduce(4),		/* *, reduce: Term */
			reduce(4),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Term */
			reduce(5),		/* +, reduce: Term */
			reduce(5),		/* -, reduce: Term */
			reduce(5),		/* *, reduce: Term */
			reduce(5),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Factor */
			reduce(7),		/* +, reduce: Factor */
			reduce(7),		/* -, reduce: Factor */
			reduce(7),		/* *, reduce: Factor */
			reduce(7),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* int */
			shift(18),		/* date */
			shift(19),		/* id */
			shift(20),		/* string */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(25),		/* + */
			shift(26),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(35),		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(1),		/* +, reduce: Expr */
			reduce(1),		/* -, reduce: Expr */
			shift(28),		/* * */
			shift(29),		/* / */
			nil,		/* ( */
			reduce(1),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(2),		/* +, reduce: Expr */
			reduce(2),		/* -, reduce: Expr */
			shift(28),		/* * */
			shift(29),		/* / */
			nil,		/* ( */
			reduce(2),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(4),		/* +, reduce: Term */
			reduce(4),		/* -, reduce: Term */
			reduce(4),		/* *, reduce: Term */
			reduce(4),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(5),		/* +, reduce: Term */
			reduce(5),		/* -, reduce: Term */
			reduce(5),		/* *, reduce: Term */
			reduce(5),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(7),		/* +, reduce: Factor */
			reduce(7),		/* -, reduce: Factor */
			reduce(7),		/* *, reduce: Factor */
			reduce(7),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* id */
			nil,		/* string */
			
		},

	},
	
}

