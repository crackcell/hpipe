
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			shift(14),		/* float */
			shift(18),		/* string */
			shift(19),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(20),		/* ${env: */
			
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
			nil,		/* float */
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
			shift(21),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			shift(22),		/* + */
			shift(23),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			shift(24),		/* * */
			shift(25),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			nil,		/* float */
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
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
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
			reduce(16),		/* $, reduce: Factor */
			reduce(16),		/* ;, reduce: Factor */
			nil,		/* id */
			nil,		/* = */
			reduce(16),		/* +, reduce: Factor */
			reduce(16),		/* -, reduce: Factor */
			reduce(16),		/* *, reduce: Factor */
			reduce(16),		/* /, reduce: Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
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
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			shift(39),		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(40),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			reduce(17),		/* $, reduce: Id */
			reduce(17),		/* ;, reduce: Id */
			nil,		/* id */
			nil,		/* = */
			reduce(17),		/* +, reduce: Id */
			reduce(17),		/* -, reduce: Id */
			reduce(17),		/* *, reduce: Id */
			reduce(17),		/* /, reduce: Id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
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
			shift(14),		/* float */
			shift(18),		/* string */
			shift(19),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(20),		/* ${env: */
			
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
			shift(14),		/* float */
			shift(18),		/* string */
			shift(19),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(20),		/* ${env: */
			
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
			shift(14),		/* float */
			shift(18),		/* string */
			shift(19),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(20),		/* ${env: */
			
		},

	},
	actionRow{ // S25
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
			shift(14),		/* float */
			shift(18),		/* string */
			shift(19),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(20),		/* ${env: */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(45),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			shift(46),		/* + */
			shift(47),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(48),		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(49),		/* * */
			shift(50),		/* / */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* float */
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
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
			reduce(11),		/* +, reduce: Factor */
			reduce(11),		/* -, reduce: Factor */
			reduce(11),		/* *, reduce: Factor */
			reduce(11),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
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
			reduce(12),		/* +, reduce: Factor */
			reduce(12),		/* -, reduce: Factor */
			reduce(12),		/* *, reduce: Factor */
			reduce(12),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(12),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
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
			reduce(13),		/* +, reduce: Factor */
			reduce(13),		/* -, reduce: Factor */
			reduce(13),		/* *, reduce: Factor */
			reduce(13),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(13),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
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
			reduce(14),		/* +, reduce: Factor */
			reduce(14),		/* -, reduce: Factor */
			reduce(14),		/* *, reduce: Factor */
			reduce(14),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(14),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
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
			reduce(15),		/* +, reduce: Factor */
			reduce(15),		/* -, reduce: Factor */
			reduce(15),		/* *, reduce: Factor */
			reduce(15),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(15),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
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
			nil,		/* id */
			nil,		/* = */
			reduce(16),		/* +, reduce: Factor */
			reduce(16),		/* -, reduce: Factor */
			reduce(16),		/* *, reduce: Factor */
			reduce(16),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(16),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			shift(52),		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(53),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
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
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(54),		/* } */
			nil,		/* ${env: */
			
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
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(55),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			reduce(4),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(24),		/* * */
			shift(25),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			reduce(5),		/* $, reduce: Expr */
			reduce(5),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(24),		/* * */
			shift(25),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			nil,		/* float */
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(17),		/* +, reduce: Id */
			reduce(17),		/* -, reduce: Id */
			reduce(17),		/* *, reduce: Id */
			reduce(17),		/* /, reduce: Id */
			nil,		/* ( */
			reduce(17),		/* ), reduce: Id */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
		},

	},
	actionRow{ // S48
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			shift(26),		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(30),		/* ( */
			nil,		/* ) */
			shift(31),		/* int */
			shift(32),		/* float */
			shift(36),		/* string */
			shift(37),		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			shift(38),		/* ${env: */
			
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
			shift(46),		/* + */
			shift(47),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(60),		/* ) */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S52
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(61),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S53
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
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			shift(62),		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Date */
			reduce(18),		/* ;, reduce: Date */
			nil,		/* id */
			nil,		/* = */
			reduce(18),		/* +, reduce: Date */
			reduce(18),		/* -, reduce: Date */
			reduce(18),		/* *, reduce: Date */
			reduce(18),		/* /, reduce: Date */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			reduce(19),		/* $, reduce: Env */
			reduce(19),		/* ;, reduce: Env */
			nil,		/* id */
			nil,		/* = */
			reduce(19),		/* +, reduce: Env */
			reduce(19),		/* -, reduce: Env */
			reduce(19),		/* *, reduce: Env */
			reduce(19),		/* /, reduce: Env */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int */
			nil,		/* float */
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
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(49),		/* * */
			shift(50),		/* / */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* float */
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
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(49),		/* * */
			shift(50),		/* / */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Expr */
			nil,		/* int */
			nil,		/* float */
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
			reduce(7),		/* +, reduce: Term */
			reduce(7),		/* -, reduce: Term */
			reduce(7),		/* *, reduce: Term */
			reduce(7),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* float */
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
			reduce(8),		/* +, reduce: Term */
			reduce(8),		/* -, reduce: Term */
			reduce(8),		/* *, reduce: Term */
			reduce(8),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Term */
			nil,		/* int */
			nil,		/* float */
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
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Factor */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(18),		/* +, reduce: Date */
			reduce(18),		/* -, reduce: Date */
			reduce(18),		/* *, reduce: Date */
			reduce(18),		/* /, reduce: Date */
			nil,		/* ( */
			reduce(18),		/* ), reduce: Date */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			nil,		/* id */
			nil,		/* = */
			reduce(19),		/* +, reduce: Env */
			reduce(19),		/* -, reduce: Env */
			reduce(19),		/* *, reduce: Env */
			reduce(19),		/* /, reduce: Env */
			nil,		/* ( */
			reduce(19),		/* ), reduce: Env */
			nil,		/* int */
			nil,		/* float */
			nil,		/* string */
			nil,		/* ${date: */
			nil,		/* date */
			nil,		/* } */
			nil,		/* ${env: */
			
		},

	},
	
}

