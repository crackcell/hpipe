
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
			shift(1),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int */
			shift(8),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			shift(9),		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(10),		/* + */
			shift(11),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Expr */
			reduce(3),		/* +, reduce: Expr */
			reduce(3),		/* -, reduce: Expr */
			shift(12),		/* * */
			shift(13),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S4
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S6
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S7
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			shift(23),		/* date */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Variable */
			reduce(11),		/* +, reduce: Variable */
			reduce(11),		/* -, reduce: Variable */
			reduce(11),		/* *, reduce: Variable */
			reduce(11),		/* /, reduce: Variable */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(1),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int */
			shift(8),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(1),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int */
			shift(8),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(1),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int */
			shift(8),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(1),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int */
			shift(8),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			shift(28),		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* + */
			shift(30),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(31),		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(3),		/* +, reduce: Expr */
			reduce(3),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(3),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S17
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S19
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S20
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			shift(23),		/* date */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			shift(36),		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			reduce(12),		/* }, reduce: Builtins */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Expr */
			reduce(1),		/* +, reduce: Expr */
			reduce(1),		/* -, reduce: Expr */
			shift(12),		/* * */
			shift(13),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Expr */
			reduce(2),		/* +, reduce: Expr */
			reduce(2),		/* -, reduce: Expr */
			shift(12),		/* * */
			shift(13),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S26
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S27
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(11),		/* +, reduce: Variable */
			reduce(11),		/* -, reduce: Variable */
			reduce(11),		/* *, reduce: Variable */
			reduce(11),		/* /, reduce: Variable */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Variable */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S31
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(14),		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(18),		/* ( */
			nil,		/* ) */
			shift(19),		/* int */
			shift(21),		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(29),		/* + */
			shift(30),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(41),		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			shift(42),		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Variable */
			reduce(10),		/* +, reduce: Variable */
			reduce(10),		/* -, reduce: Variable */
			reduce(10),		/* *, reduce: Variable */
			reduce(10),		/* /, reduce: Variable */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(1),		/* +, reduce: Expr */
			reduce(1),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(1),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(2),		/* +, reduce: Expr */
			reduce(2),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(2),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S39
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S40
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S41
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
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(10),		/* +, reduce: Variable */
			reduce(10),		/* -, reduce: Variable */
			reduce(10),		/* *, reduce: Variable */
			reduce(10),		/* /, reduce: Variable */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Variable */
			nil,		/* int */
			nil,		/* ${ */
			nil,		/* } */
			nil,		/* id */
			nil,		/* date */
			
		},

	},
	
}

