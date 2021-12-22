package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneratePageOrderBasicOK(t *testing.T) {
	res, err := GeneratePrintOrder(1, 8, 4, 10)

	require.NoError(t, err)

	exp := [][]int{{1, 5, 6, 2, 3, 7, 8, 4}}

	require.Equal(t, exp, res)
}

func TestGeneratePageOrderBoundaries(t *testing.T) {
	_, err := GeneratePrintOrder(1, 0, 4, 10)
	assert.Error(t, err)

	_, err = GeneratePrintOrder(1, 5, 3, 10)
	assert.Error(t, err)
}

func TestGeneratePageOrderBatches(t *testing.T) {
	res, err := GeneratePrintOrder(1, 40, 4, 5)
	require.NoError(t, err)

	exp := [][]int{
		{1, 11, 12, 2, 3, 13, 14, 4, 5, 15, 16, 6, 7, 17, 18, 8, 9, 19, 20, 10},
		{21, 31, 32, 22, 23, 33, 34, 24, 25, 35, 36, 26, 27, 37, 38, 28, 29, 39, 40, 30},
	}

	require.Equal(t, exp, res)
}

func TestGeneratePageOrderBatchNotFull(t *testing.T) {
	res, err := GeneratePrintOrder(3, 33, 4, 3)

	require.NoError(t, err)

	//1: 3..14 -> 12
	//2: 15..26 -> 12
	//3: 27..33 -> 7
	exp := [][]int{
		{3, 9, 10, 4, 5, 11, 12, 6, 7, 13, 14, 8},
		{15, 21, 22, 16, 17, 23, 24, 18, 19, 25, 26, 20},
		{27, 31, 32, 28, 29, 33, -1, 30},
	}

	require.Equal(t, exp, res)
}
