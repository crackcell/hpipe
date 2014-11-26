
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 34 : // ['"','"']
				return 2
			case r == 40 : // ['(','(']
				return 3
			case r == 41 : // [')',')']
				return 4
			case r == 42 : // ['*','*']
				return 5
			case r == 43 : // ['+','+']
				return 6
			case r == 45 : // ['-','-']
				return 7
			case r == 46 : // ['.','.']
				return 8
			case r == 47 : // ['/','/']
				return 9
			case r == 48 : // ['0','0']
				return 10
			case 49 <= r && r <= 57 : // ['1','9']
				return 11
			case r == 59 : // [';',';']
				return 12
			case r == 61 : // ['=','=']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 14
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 15
			case r == 46 : // ['.','.']
				return 16
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 8
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 14
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 19
			case r == 95 : // ['_','_']
				return 20
			case 97 <= r && r <= 122 : // ['a','z']
				return 19
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 11
			case 65 <= r && r <= 90 : // ['A','Z']
				return 19
			case r == 95 : // ['_','_']
				return 20
			case 97 <= r && r <= 122 : // ['a','z']
				return 19
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 8
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 8
			case r == 95 : // ['_','_']
				return 14
			case 97 <= r && r <= 122 : // ['a','z']
				return 8
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 15
			case r == 46 : // ['.','.']
				return 16
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 15
			case r == 46 : // ['.','.']
				return 16
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 15
			case r == 46 : // ['.','.']
				return 16
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 19
			case r == 95 : // ['_','_']
				return 20
			case 97 <= r && r <= 122 : // ['a','z']
				return 19
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 19
			case r == 95 : // ['_','_']
				return 20
			case 97 <= r && r <= 122 : // ['a','z']
				return 19
			
			
			
			}
			return NoState
			
		},
	
}
