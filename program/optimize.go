package program

// Optimize optimizes a list of instruction in the following ways:
//
// 1. Consecutive instructions of InstMoveHead will be merged.
// 2. Consecutive instructions of InstAddToByte will be merged.
//
// Example:
//
// If there are 10 consecutive InstMoveHead{1} instructions
// these will be combined into 1 InstMoveHead{10} instruction.
func Optimize(insts []Instruction) []Instruction {
	opted := make([]Instruction, 0)
	for i := 0; i < len(insts); {
		ii := insts[i]
		switch ii.(type) {
		case InstMoveHead:
			i2 := ii.(InstMoveHead)
			j := i + 1
			for ; j < len(insts); j++ {
				ij := insts[j]
				if v, ok := ij.(InstMoveHead); ok {
					i2 = InstMoveHead{i2.V + v.V}
				} else {
					break
				}
			}
			opted = append(opted, i2)
			i = j
		case InstAddToByte:
			i2 := ii.(InstAddToByte)
			j := i + 1
			for ; j < len(insts); j++ {
				ij := insts[j]
				if v, ok := ij.(InstAddToByte); ok {
					i2 = InstAddToByte{i2.V + v.V}
				} else {
					break
				}
			}
			opted = append(opted, i2)
			i = j
		case InstLoop:
			v := ii.(InstLoop)
			opted = append(opted, InstLoop{Optimize(v.Insts)})
			i++
		default:
			opted = append(opted, ii)
			i++
		}
	}
	return opted
}
