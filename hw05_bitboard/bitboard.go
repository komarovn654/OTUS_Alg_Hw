package hw05bitboard

type King struct {
}

func (k *King) GetMoves(pos int) int64 {
	kp := int64(1 << pos) // king position
	noA := int64(0xfefefefefefefe)
	kpA := kp & noA
	mask := (kpA << 7) | (kp << 8) | (kp << 9) |
		(kpA >> 1) | (kp << 1) |
		(kpA >> 9) | (kp >> 8) | (kp >> 7)

	return mask
}
