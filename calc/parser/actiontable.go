
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			shift(12),		/* string */
			shift(13),		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			shift(14),		/* + */
			shift(15),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
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
			shift(16),		/* * */
			shift(17),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S14
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
			shift(12),		/* string */
			shift(13),		/* int64 */
			
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
			shift(12),		/* string */
			shift(13),		/* int64 */
			
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
			shift(12),		/* string */
			shift(13),		/* int64 */
			
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
			shift(12),		/* string */
			shift(13),		/* int64 */
			
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
			reduce(13),		/* +, reduce: Factor */
			reduce(13),		/* -, reduce: Factor */
			reduce(13),		/* *, reduce: Factor */
			reduce(13),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(13),		/* ), reduce: Factor */
			nil,		/* string */
			nil,		/* int64 */
			
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
			shift(29),		/* + */
			shift(30),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(31),		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(6),		/* +, reduce: Expr */
			reduce(6),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Expr */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(9),		/* +, reduce: Term */
			reduce(9),		/* -, reduce: Term */
			reduce(9),		/* *, reduce: Term */
			reduce(9),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Term */
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
		},

	},
	actionRow{ // S23
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
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(12),		/* +, reduce: Factor */
			reduce(12),		/* -, reduce: Factor */
			reduce(12),		/* *, reduce: Factor */
			reduce(12),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(12),		/* ), reduce: Factor */
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Expr */
			reduce(4),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(16),		/* * */
			shift(17),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Expr */
			reduce(5),		/* ;, reduce: Expr */
			nil,		/* id */
			nil,		/* = */
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(16),		/* * */
			shift(17),		/* / */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S27
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
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S28
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
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
		},

	},
	actionRow{ // S31
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
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* ; */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* + */
			nil,		/* - */
			nil,		/* * */
			nil,		/* / */
			shift(22),		/* ( */
			nil,		/* ) */
			shift(23),		/* string */
			shift(24),		/* int64 */
			
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
			shift(29),		/* + */
			shift(30),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(39),		/* ) */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(4),		/* +, reduce: Expr */
			reduce(4),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Expr */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(5),		/* +, reduce: Expr */
			reduce(5),		/* -, reduce: Expr */
			shift(32),		/* * */
			shift(33),		/* / */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Expr */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(7),		/* +, reduce: Term */
			reduce(7),		/* -, reduce: Term */
			reduce(7),		/* *, reduce: Term */
			reduce(7),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Term */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(8),		/* +, reduce: Term */
			reduce(8),		/* -, reduce: Term */
			reduce(8),		/* *, reduce: Term */
			reduce(8),		/* /, reduce: Term */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Term */
			nil,		/* string */
			nil,		/* int64 */
			
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
			reduce(10),		/* +, reduce: Factor */
			reduce(10),		/* -, reduce: Factor */
			reduce(10),		/* *, reduce: Factor */
			reduce(10),		/* /, reduce: Factor */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Factor */
			nil,		/* string */
			nil,		/* int64 */
			
		},

	},
	
}

