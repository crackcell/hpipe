
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
			case r == 36 : // ['$','$']
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
			case r == 47 : // ['/','/']
				return 8
			case r == 48 : // ['0','0']
				return 9
			case 49 <= r && r <= 57 : // ['1','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
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
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 14
			case r == 95 : // ['_','_']
				return 15
			case 97 <= r && r <= 122 : // ['a','z']
				return 14
			case r == 123 : // ['{','{']
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
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 9
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 9
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 9
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 17
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 14
			case r == 95 : // ['_','_']
				return 15
			case 97 <= r && r <= 122 : // ['a','z']
				return 14
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 14
			case r == 95 : // ['_','_']
				return 15
			case 97 <= r && r <= 122 : // ['a','z']
				return 14
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 32 : // [' ',' ']
				return 16
			case r == 45 : // ['-','-']
				return 16
			case r == 58 : // [':',':']
				return 16
			case r == 68 : // ['D','D']
				return 19
			case r == 77 : // ['M','M']
				return 19
			case r == 89 : // ['Y','Y']
				return 19
			case r == 104 : // ['h','h']
				return 19
			case r == 109 : // ['m','m']
				return 19
			case r == 115 : // ['s','s']
				return 19
			case r == 125 : // ['}','}']
				return 20
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 17
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 17
			case r == 95 : // ['_','_']
				return 18
			case 97 <= r && r <= 122 : // ['a','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 32 : // [' ',' ']
				return 16
			case r == 45 : // ['-','-']
				return 16
			case r == 58 : // [':',':']
				return 16
			case r == 68 : // ['D','D']
				return 19
			case r == 77 : // ['M','M']
				return 19
			case r == 89 : // ['Y','Y']
				return 19
			case r == 104 : // ['h','h']
				return 19
			case r == 109 : // ['m','m']
				return 19
			case r == 115 : // ['s','s']
				return 19
			case r == 125 : // ['}','}']
				return 20
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
}
