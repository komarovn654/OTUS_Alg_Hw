package hw03alg

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultuply(t *testing.T) {
	a := Matrix{
		x11: big.NewInt(234),
		x12: big.NewInt(23),
		x21: big.NewInt(115),
		x22: big.NewInt(467),
	}
	b := Matrix{
		x11: big.NewInt(1),
		x12: big.NewInt(2),
		x21: big.NewInt(3),
		x22: big.NewInt(4),
	}
	expect := Matrix{
		x11: big.NewInt(303),
		x12: big.NewInt(560),
		x21: big.NewInt(1516),
		x22: big.NewInt(2098),
	}

	t.Run("matrix multiplication", func(t *testing.T) {
		require.Equal(t, expect, a.Multiply(b))
	})

	b = IdentityMatrix()
	t.Run("identity matrix multiplication", func(t *testing.T) {
		require.Equal(t, expect, a.Multiply(b))
	})
}

func TestMatrixPwr(t *testing.T) {
	a := Matrix{
		x11: big.NewInt(2),
		x12: big.NewInt(2),
		x21: big.NewInt(2),
		x22: big.NewInt(2),
	}
	expect := Matrix{
		x11: big.NewInt(8),
		x12: big.NewInt(8),
		x21: big.NewInt(8),
		x22: big.NewInt(8),
	}

	t.Run("matrix multiplication", func(t *testing.T) {
		require.Equal(t, expect, a.Pwr(2))
	})
}
