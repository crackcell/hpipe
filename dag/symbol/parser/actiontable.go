
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
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(4),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(5),		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S3
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			shift(6),		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(1),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(8),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int */
			shift(17),		/* string */
			shift(18),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(19),		/* ${env: */
			
		},

	},
	actionRow{ // S7
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(20),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Stmt */
			reduce(3),		/* ;, reduce: Stmt */
			nil,		/* id */
			nil,		/* = */
			shift(21),		/* + */
			shift(22),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Expr */
			reduce(6),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(23),		/* * */
			shift(24),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S11
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S13
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S14
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S15
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S16
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: Factor */
			reduce(15),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(15),		/* +, reduce: Factor */
			reduce(15),		/* -, reduce: Factor */
			reduce(15),		/* *, reduce: Factor */
			reduce(15),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			shift(37),		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(38),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(16),		/* $, reduce: Id */
			reduce(16),		/* ;, reduce: Id */
			nil,		/* id */
			nil,		/* = */
			reduce(16),		/* +, reduce: Id */
			reduce(16),		/* -, reduce: Id */
			reduce(16),		/* *, reduce: Id */
			reduce(16),		/* /, reduce: Id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(8),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int */
			shift(17),		/* string */
			shift(18),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(19),		/* ${env: */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(8),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int */
			shift(17),		/* string */
			shift(18),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(19),		/* ${env: */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(8),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int */
			shift(17),		/* string */
			shift(18),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(19),		/* ${env: */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(8),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int */
			shift(17),		/* string */
			shift(18),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(19),		/* ${env: */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(43),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
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
			shift(44),		/* + */
			shift(45),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(46),		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(47),		/* * */
			shift(48),		/* / */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S28
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S30
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S31
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S32
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S33
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(15),		/* +, reduce: Factor */
			reduce(15),		/* -, reduce: Factor */
			reduce(15),		/* *, reduce: Factor */
			reduce(15),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(15),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			shift(50),		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(51),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
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
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(52),		/* } */
			nil,		/* ${env: */
			
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
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(53),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			reduce(4),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(23),		/* * */
			shift(24),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Expr */
			reduce(5),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(23),		/* * */
			shift(24),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S41
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S42
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(16),		/* +, reduce: Id */
			reduce(16),		/* -, reduce: Id */
			reduce(16),		/* *, reduce: Id */
			reduce(16),		/* /, reduce: Id */
			nil,		/* ( */
			reduce(16),		/* ), reduce: Id */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S46
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(25),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(29),		/* ( */
			nil,		/* ) */
			shift(30),		/* int */
			shift(34),		/* string */
			shift(35),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(36),		/* ${env: */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			shift(44),		/* + */
			shift(45),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(58),		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(59),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(60),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: Date */
			reduce(17),		/* ;, reduce: Date */
			nil,		/* id */
			nil,		/* = */
			reduce(17),		/* +, reduce: Date */
			reduce(17),		/* -, reduce: Date */
			reduce(17),		/* *, reduce: Date */
			reduce(17),		/* /, reduce: Date */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Env */
			reduce(18),		/* ;, reduce: Env */
			nil,		/* id */
			nil,		/* = */
			reduce(18),		/* +, reduce: Env */
			reduce(18),		/* -, reduce: Env */
			reduce(18),		/* *, reduce: Env */
			reduce(18),		/* /, reduce: Env */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(47),		/* * */
			shift(48),		/* / */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(47),		/* * */
			shift(48),		/* / */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S56
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S57
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S58
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
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(17),		/* +, reduce: Date */
			reduce(17),		/* -, reduce: Date */
			reduce(17),		/* *, reduce: Date */
			reduce(17),		/* /, reduce: Date */
			nil,		/* ( */
			reduce(17),		/* ), reduce: Date */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(18),		/* +, reduce: Env */
			reduce(18),		/* -, reduce: Env */
			reduce(18),		/* *, reduce: Env */
			reduce(18),		/* /, reduce: Env */
			nil,		/* ( */
			reduce(18),		/* ), reduce: Env */
			nil,		/* int */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	
}

