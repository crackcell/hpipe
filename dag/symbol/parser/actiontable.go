
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
			nil,		/* ; */
			shift(3),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(4),		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: StmtList */
			reduce(1),		/* ;, reduce: StmtList */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			shift(5),		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(3),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(7),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(11),		/* ( */
			nil,		/* ) */
			shift(12),		/* int */
			shift(13),		/* date */
			shift(14),		/* string */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: StmtList */
			reduce(2),		/* ;, reduce: StmtList */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: Factor */
			reduce(13),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(13),		/* +, reduce: Factor */
			reduce(13),		/* -, reduce: Factor */
			reduce(13),		/* *, reduce: Factor */
			reduce(13),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Stmt */
			reduce(3),		/* ;, reduce: Stmt */
			nil,		/* id */
			nil,		/* = */
			shift(15),		/* + */
			shift(16),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Expr */
			reduce(6),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(17),		/* * */
			shift(18),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Term */
			reduce(9),		/* ;, reduce: Term */
			nil,		/* id */
			nil,		/* = */
			reduce(9),		/* +, reduce: Term */
			reduce(9),		/* -, reduce: Term */
			reduce(9),		/* *, reduce: Term */
			reduce(9),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Factor */
			reduce(11),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(11),		/* +, reduce: Factor */
			reduce(11),		/* -, reduce: Factor */
			reduce(11),		/* *, reduce: Factor */
			reduce(11),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Factor */
			reduce(12),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(12),		/* +, reduce: Factor */
			reduce(12),		/* -, reduce: Factor */
			reduce(12),		/* *, reduce: Factor */
			reduce(12),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: Factor */
			reduce(14),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(14),		/* +, reduce: Factor */
			reduce(14),		/* -, reduce: Factor */
			reduce(14),		/* *, reduce: Factor */
			reduce(14),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(7),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(11),		/* ( */
			nil,		/* ) */
			shift(12),		/* int */
			shift(13),		/* date */
			shift(14),		/* string */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(7),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(11),		/* ( */
			nil,		/* ) */
			shift(12),		/* int */
			shift(13),		/* date */
			shift(14),		/* string */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(7),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(11),		/* ( */
			nil,		/* ) */
			shift(12),		/* int */
			shift(13),		/* date */
			shift(14),		/* string */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(7),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(11),		/* ( */
			nil,		/* ) */
			shift(12),		/* int */
			shift(13),		/* date */
			shift(14),		/* string */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(13),		/* +, reduce: Factor */
			reduce(13),		/* -, reduce: Factor */
			reduce(13),		/* *, reduce: Factor */
			reduce(13),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(13),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			shift(31),		/* + */
			shift(32),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(33),		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(34),		/* * */
			shift(35),		/* / */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(9),		/* +, reduce: Term */
			reduce(9),		/* -, reduce: Term */
			reduce(9),		/* *, reduce: Term */
			reduce(9),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(11),		/* +, reduce: Factor */
			reduce(11),		/* -, reduce: Factor */
			reduce(11),		/* *, reduce: Factor */
			reduce(11),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(12),		/* +, reduce: Factor */
			reduce(12),		/* -, reduce: Factor */
			reduce(12),		/* *, reduce: Factor */
			reduce(12),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(12),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(14),		/* +, reduce: Factor */
			reduce(14),		/* -, reduce: Factor */
			reduce(14),		/* *, reduce: Factor */
			reduce(14),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(14),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			reduce(4),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(17),		/* * */
			shift(18),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Expr */
			reduce(5),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(17),		/* * */
			shift(18),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Term */
			reduce(7),		/* ;, reduce: Term */
			nil,		/* id */
			nil,		/* = */
			reduce(7),		/* +, reduce: Term */
			reduce(7),		/* -, reduce: Term */
			reduce(7),		/* *, reduce: Term */
			reduce(7),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Term */
			reduce(8),		/* ;, reduce: Term */
			nil,		/* id */
			nil,		/* = */
			reduce(8),		/* +, reduce: Term */
			reduce(8),		/* -, reduce: Term */
			reduce(8),		/* *, reduce: Term */
			reduce(8),		/* /, reduce: Term */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Factor */
			reduce(10),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(19),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(23),		/* ( */
			nil,		/* ) */
			shift(24),		/* int */
			shift(25),		/* date */
			shift(26),		/* string */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			shift(31),		/* + */
			shift(32),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(41),		/* ) */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(34),		/* * */
			shift(35),		/* / */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(34),		/* * */
			shift(35),		/* / */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(7),		/* +, reduce: Term */
			reduce(7),		/* -, reduce: Term */
			reduce(7),		/* *, reduce: Term */
			reduce(7),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(8),		/* +, reduce: Term */
			reduce(8),		/* -, reduce: Term */
			reduce(8),		/* *, reduce: Term */
			reduce(8),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* date */
			nil,		/* string */
			
		},

	},
	
}

