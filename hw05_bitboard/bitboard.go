package hw05bitboard

var (
	noA = uint64(0xfefefefefefefefe)
	noB = uint64(0xfdfdfdfdfdfdfdfd)
	noH = uint64(0x7f7f7f7f7f7f7f7f)
	noG = uint64(0xbfbfbfbfbfbfbfbf)

	column1 = uint64(0xff)
	lineA   = uint64(0x101010101010101)
)

type ChessPiece interface {
	GetMoves(pos uint64) uint64
}

type King struct {
}

func (k *King) GetMoves(pos uint64) uint64 {
	kp := uint64(1 << pos)
	kpA := kp & noA
	kpH := kp & noH

	mask := (kpA << 7) | (kp << 8) | (kpH << 9) |
		(kpA >> 1) | (kpH << 1) |
		(kpA >> 9) | (kp >> 8) | (kpH >> 7)

	return mask
}

type Knight struct {
}

func (k *Knight) GetMoves(pos uint64) uint64 {
	kp := uint64(1 << pos)

	kpA := kp & noA
	kpAB := kp & noA & noB
	kpH := kp & noH
	kpHG := kp & noH & noG

	mask := (kpHG << 17) | (kpH << 10) | (kpHG >> 6) | (kpH >> 15) |
		(kpAB >> 17) | (kpA >> 10) | (kpAB << 6) | (kpA << 15)

	return mask
}

type Rook struct {
}

func (r *Rook) GetMoves(pos uint64) uint64 {
	rp := uint64(1 << pos)

	line := lineA << (pos % 8)
	column := column1 << (8 * (pos / 8))

	return (line | column) ^ rp
}

type Bishop struct {
}

func (r *Bishop) GetMoves(pos uint64) uint64 {
	bp := uint64(1 << pos)
	posDiag := bp
	negDiag := bp

	for i := uint64(1); i <= pos%8; i++ {
		posDiag |= (bp >> ((8 * i) + i))
		negDiag |= (bp << ((8 * i) - i))
	}
	for i := uint64(1); i < (8 - pos%8); i++ {
		posDiag |= (bp << ((8 * i) + i))
		negDiag |= (bp >> ((8 * i) - i))
	}

	return (posDiag | negDiag) ^ bp
}

type Queen struct {
}

func (q *Queen) GetMoves(pos uint64) uint64 {
	b := Bishop{}
	r := Rook{}

	return (b.GetMoves(pos) | r.GetMoves(pos))
}

func BitsCountShift(bitField uint64) (count uint64) {
	for ; bitField > 0; bitField = bitField >> 1 {
		if bitField&1 == 1 {
			count++
		}
	}

	return count
}

func BitsCountSubtract(bitField uint64) (count uint64) {
	for ; bitField > 0; bitField &= bitField - 1 {
		count++
	}

	return count
}

type Cache struct {
	bits [256]uint8
}

func (c *Cache) Init() {
	for i := uint64(0); i < 256; i++ {
		c.bits[i] = uint8(BitsCountSubtract(i))
	}
}

func BitsCountCache(bitField uint64, c *Cache) (count uint64) {
	for i := 0; i < 8; i++ {
		mask := bitField & 0xff
		count += uint64(c.bits[mask])
		bitField >>= 8
	}
	return count
}
