
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
				return 2
			case r == 34 : // ['"','"']
				return 3
			case r == 36 : // ['$','$']
				return 4
			case r == 40 : // ['(','(']
				return 5
			case r == 41 : // [')',')']
				return 6
			case r == 42 : // ['*','*']
				return 7
			case r == 43 : // ['+','+']
				return 8
			case r == 45 : // ['-','-']
				return 9
			case r == 46 : // ['.','.']
				return 10
			case r == 47 : // ['/','/']
				return 11
			case r == 48 : // ['0','0']
				return 12
			case 49 <= r && r <= 57 : // ['1','9']
				return 13
			case r == 58 : // [':',':']
				return 14
			case r == 59 : // [';',';']
				return 15
			case r == 61 : // ['=','=']
				return 16
			case 65 <= r && r <= 67 : // ['A','C']
				return 17
			case r == 68 : // ['D','D']
				return 18
			case 69 <= r && r <= 76 : // ['E','L']
				return 17
			case r == 77 : // ['M','M']
				return 18
			case 78 <= r && r <= 88 : // ['N','X']
				return 17
			case r == 89 : // ['Y','Y']
				return 18
			case r == 90 : // ['Z','Z']
				return 17
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 103 : // ['a','g']
				return 17
			case r == 104 : // ['h','h']
				return 18
			case 105 <= r && r <= 108 : // ['i','l']
				return 17
			case r == 109 : // ['m','m']
				return 18
			case 110 <= r && r <= 114 : // ['n','r']
				return 17
			case r == 115 : // ['s','s']
				return 18
			case 116 <= r && r <= 122 : // ['t','z']
				return 17
			case r == 125 : // ['}','}']
				return 20
			
			
			
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
			case r == 32 : // [' ',' ']
				return 14
			case r == 45 : // ['-','-']
				return 14
			case r == 58 : // [':',':']
				return 14
			case r == 68 : // ['D','D']
				return 14
			case r == 77 : // ['M','M']
				return 14
			case r == 89 : // ['Y','Y']
				return 14
			case r == 104 : // ['h','h']
				return 14
			case r == 109 : // ['m','m']
				return 14
			case r == 115 : // ['s','s']
				return 14
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 21
			case r == 46 : // ['.','.']
				return 22
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			case 65 <= r && r <= 90 : // ['A','Z']
				return 24
			case r == 95 : // ['_','_']
				return 22
			case 97 <= r && r <= 122 : // ['a','z']
				return 24
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 123 : // ['{','{']
				return 25
			
			
			
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
			case r == 32 : // [' ',' ']
				return 14
			case r == 45 : // ['-','-']
				return 14
			case r == 58 : // [':',':']
				return 14
			case r == 68 : // ['D','D']
				return 14
			case r == 77 : // ['M','M']
				return 14
			case r == 89 : // ['Y','Y']
				return 14
			case r == 104 : // ['h','h']
				return 14
			case r == 109 : // ['m','m']
				return 14
			case r == 115 : // ['s','s']
				return 14
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 26
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 26
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 13
			case 65 <= r && r <= 90 : // ['A','Z']
				return 26
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 32 : // [' ',' ']
				return 14
			case r == 45 : // ['-','-']
				return 14
			case r == 58 : // [':',':']
				return 14
			case r == 68 : // ['D','D']
				return 14
			case r == 77 : // ['M','M']
				return 14
			case r == 89 : // ['Y','Y']
				return 14
			case r == 104 : // ['h','h']
				return 14
			case r == 109 : // ['m','m']
				return 14
			case r == 115 : // ['s','s']
				return 14
			
			
			
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
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 17
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 32 : // [' ',' ']
				return 14
			case r == 45 : // ['-','-']
				return 14
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case r == 58 : // [':',':']
				return 14
			case 65 <= r && r <= 67 : // ['A','C']
				return 17
			case r == 68 : // ['D','D']
				return 18
			case 69 <= r && r <= 76 : // ['E','L']
				return 17
			case r == 77 : // ['M','M']
				return 18
			case 78 <= r && r <= 88 : // ['N','X']
				return 17
			case r == 89 : // ['Y','Y']
				return 18
			case r == 90 : // ['Z','Z']
				return 17
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 103 : // ['a','g']
				return 17
			case r == 104 : // ['h','h']
				return 18
			case 105 <= r && r <= 108 : // ['i','l']
				return 17
			case r == 109 : // ['m','m']
				return 18
			case 110 <= r && r <= 114 : // ['n','r']
				return 17
			case r == 115 : // ['s','s']
				return 18
			case 116 <= r && r <= 122 : // ['t','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 17
			case r == 95 : // ['_','_']
				return 19
			case 97 <= r && r <= 122 : // ['a','z']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 21
			case r == 46 : // ['.','.']
				return 22
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			case 65 <= r && r <= 90 : // ['A','Z']
				return 24
			case r == 95 : // ['_','_']
				return 22
			case 97 <= r && r <= 122 : // ['a','z']
				return 24
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 21
			case r == 46 : // ['.','.']
				return 22
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			case 65 <= r && r <= 90 : // ['A','Z']
				return 24
			case r == 95 : // ['_','_']
				return 22
			case 97 <= r && r <= 122 : // ['a','z']
				return 24
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 21
			case r == 46 : // ['.','.']
				return 22
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			case 65 <= r && r <= 90 : // ['A','Z']
				return 24
			case r == 95 : // ['_','_']
				return 22
			case 97 <= r && r <= 122 : // ['a','z']
				return 24
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 100 : // ['d','d']
				return 27
			case r == 101 : // ['e','e']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 26
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 97 : // ['a','a']
				return 29
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 30
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 116 : // ['t','t']
				return 31
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 118 : // ['v','v']
				return 32
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case r == 58 : // [':',':']
				return 34
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 58 : // [':',':']
				return 35
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
}
