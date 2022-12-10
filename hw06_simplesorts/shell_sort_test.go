package hw06simplesorts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShellSort(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := ShellSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := ShellSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := ShellSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := ShellSort(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}

func TestShellSortFrankLazarus(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := ShellSortFrankLazarus(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := ShellSortFrankLazarus(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := ShellSortFrankLazarus(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := ShellSortFrankLazarus(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}

func TestShellSortHibbard(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		ar := RandArray(100)
		done := ShellSortHibbard(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("random digits", func(t *testing.T) {
		ar := RandDigits(100)
		done := ShellSortHibbard(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("sorted array", func(t *testing.T) {
		ar := SortedArray(100)
		done := ShellSortHibbard(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
	t.Run("reverse array", func(t *testing.T) {
		ar := ReversArray(100)
		done := ShellSortHibbard(&ar)
		<-done
		require.True(t, IsSorted(&ar))
	})
}
