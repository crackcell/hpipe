
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
			shift(6),		/* int64 */
			nil,		/* id */
			shift(8),		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			shift(9),		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			shift(22),		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S9
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			shift(6),		/* int64 */
			nil,		/* id */
			shift(8),		/* ${ */
			nil,		/* } */
			
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
			shift(6),		/* int64 */
			nil,		/* id */
			shift(8),		/* ${ */
			nil,		/* } */
			
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
			shift(6),		/* int64 */
			nil,		/* id */
			shift(8),		/* ${ */
			nil,		/* } */
			
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
			shift(6),		/* int64 */
			nil,		/* id */
			shift(8),		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			shift(27),		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(28),		/* + */
			shift(29),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(30),		/* ) */
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(3),		/* +, reduce: Expr */
			reduce(3),		/* -, reduce: Expr */
			shift(31),		/* * */
			shift(32),		/* / */
			nil,		/* ( */
			reduce(3),		/* ), reduce: Expr */
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			shift(34),		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			shift(35),		/* } */
			
		},

	},
	actionRow{ // S23
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S24
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S25
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S26
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S27
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S28
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S30
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S31
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
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
			shift(19),		/* int64 */
			nil,		/* id */
			shift(21),		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(28),		/* + */
			shift(29),		/* - */
			nil,		/* * */
			nil,		/* / */
			nil,		/* ( */
			shift(40),		/* ) */
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S34
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			shift(41),		/* } */
			
		},

	},
	actionRow{ // S35
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(1),		/* +, reduce: Expr */
			reduce(1),		/* -, reduce: Expr */
			shift(31),		/* * */
			shift(32),		/* / */
			nil,		/* ( */
			reduce(1),		/* ), reduce: Expr */
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(2),		/* +, reduce: Expr */
			reduce(2),		/* -, reduce: Expr */
			shift(31),		/* * */
			shift(32),		/* / */
			nil,		/* ( */
			reduce(2),		/* ), reduce: Expr */
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S38
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S39
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S40
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	actionRow{ // S41
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
			nil,		/* int64 */
			nil,		/* id */
			nil,		/* ${ */
			nil,		/* } */
			
		},

	},
	
}

