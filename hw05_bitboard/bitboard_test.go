package hw05bitboard

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKing(t *testing.T) {
	tests := []struct {
		pos     uint64
		expMask uint64
	}{
		{
			pos:     0,
			expMask: 770,
		},
		{
			pos:     7,
			expMask: 49216,
		},
		{
			pos:     56,
			expMask: 144959613005987840,
		},
		{
			pos:     63,
			expMask: 4665729213955833856,
		},
		{
			pos:     36,
			expMask: 61745389371392,
		},
	}

	k := King{}
	for _, tc := range tests {
		t.Run("pos "+strconv.FormatUint(tc.pos, 10), func(t *testing.T) {
			require.Equal(t, tc.expMask, k.GetMoves(tc.pos))
		})
	}
}

func TestKnight(t *testing.T) {
	tests := []struct {
		pos     uint64
		expMask uint64
	}{
		{
			pos:     0,
			expMask: 132096,
		},
		{
			pos:     7,
			expMask: 4202496,
		},
		{
			pos:     56,
			expMask: 1128098930098176,
		},
		{
			pos:     63,
			expMask: 9077567998918656,
		},
		{
			pos:     36,
			expMask: 11333767002587136,
		},
	}

	k := Knight{}
	for _, tc := range tests {
		t.Run("pos "+strconv.FormatUint(tc.pos, 10), func(t *testing.T) {
			require.Equal(t, tc.expMask, k.GetMoves(tc.pos))
		})
	}
}

func TestRook(t *testing.T) {
	tests := []struct {
		pos     uint64
		expMask uint64
	}{
		{
			pos:     0,
			expMask: 72340172838076926,
		},
		{
			pos:     7,
			expMask: 9259542123273814143,
		},
		{
			pos:     56,
			expMask: 18302911464433844481,
		},
		{
			pos:     63,
			expMask: 9187484529235886208,
		},
		{
			pos:     36,
			expMask: 1157443723186933776,
		},
	}

	r := Rook{}
	for _, tc := range tests {
		t.Run("pos "+strconv.FormatUint(tc.pos, 10), func(t *testing.T) {
			require.Equal(t, tc.expMask, r.GetMoves(tc.pos))
		})
	}
}

func TestBishop(t *testing.T) {
	tests := []struct {
		pos     uint64
		expMask uint64
	}{
		{
			pos:     0,
			expMask: 9241421688590303744,
		},
		{
			pos:     7,
			expMask: 72624976668147712,
		},
		{
			pos:     56,
			expMask: 567382630219904,
		},
		{
			pos:     63,
			expMask: 18049651735527937,
		},
		{
			pos:     36,
			expMask: 9386671504487645697,
		},
	}

	b := Bishop{}
	for _, tc := range tests {
		t.Run("pos "+strconv.FormatUint(tc.pos, 10), func(t *testing.T) {
			require.Equal(t, tc.expMask, b.GetMoves(tc.pos))
		})
	}
}

func TestQueen(t *testing.T) {
	tests := []struct {
		pos     uint64
		expMask uint64
	}{
		{
			pos:     0,
			expMask: 9313761861428380670,
		},
		{
			pos:     7,
			expMask: 9332167099941961855,
		},
		{
			pos:     56,
			expMask: 18303478847064064385,
		},
		{
			pos:     63,
			expMask: 9205534180971414145,
		},
		{
			pos:     36,
			expMask: 10544115227674579473,
		},
	}

	q := Queen{}
	for _, tc := range tests {
		t.Run("pos "+strconv.FormatUint(tc.pos, 10), func(t *testing.T) {
			require.Equal(t, tc.expMask, q.GetMoves(tc.pos))
		})
	}
}

func TestBitsCount(t *testing.T) {
	tests := []struct {
		mask  uint64
		count uint64
	}{
		{
			mask:  0xFFFFFFFFFFFFFFFF,
			count: 64,
		},
		{
			mask:  0x0000000000000000,
			count: 0,
		},
		{
			mask:  0xFAD0E3245DFA3001,
			count: 30,
		},
	}

	c := Cache{}
	c.Init()

	for _, tc := range tests {
		t.Run("cache method", func(t *testing.T) {
			require.Equal(t, tc.count, BitsCountCache(tc.mask, &c))
		})
		t.Run("subtract method", func(t *testing.T) {
			require.Equal(t, tc.count, BitsCountSubtract(tc.mask))
		})
		t.Run("shift method", func(t *testing.T) {
			require.Equal(t, tc.count, BitsCountShift(tc.mask))
		})
	}
}
