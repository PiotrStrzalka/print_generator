package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

type batch struct {
	start int
	end   int
	data  *[]int
}

func (b batch) String() string {
	return fmt.Sprintf("Start: %d, End: %d, Data: %v", b.start, b.end, *b.data)
}

func (b batch) generate() {
	pages := b.end - b.start + 1
	sheets := int(math.Ceil(float64(pages) / 4.0))
	half := b.start + sheets*2

	for i := 0; i < sheets; i++ {
		first := (b.start + i*2)
		*b.data = append(*b.data, first)

		second := (half + i*2)
		if second <= b.end {
			*b.data = append(*b.data, second)
		} else {
			*b.data = append(*b.data, -1)
		}

		third := (half + i*2 + 1)
		if third <= b.end {
			*b.data = append(*b.data, third)
		} else {
			*b.data = append(*b.data, -1)
		}

		fourth := (b.start + i*2 + 1)
		if fourth <= b.end {
			*b.data = append(*b.data, fourth)
		}
	}
}

func GeneratePrintOrder(start, end, pagesOnSheet, sheetInBatch int) ([][]int, error) {
	if pagesOnSheet != 4 {
		return nil, fmt.Errorf("currently supported only 4 pages on sheet")
	}

	if start > end {
		return nil, fmt.Errorf("start page number cannot be bigger that end")
	}

	batches := []batch{}

	i := start
	for i < end {
		e := i + sheetInBatch*pagesOnSheet - 1
		if e > end {
			e = end
		}
		batches = append(batches, batch{
			start: i,
			end:   e,
			data:  &[]int{},
		})
		i += sheetInBatch * pagesOnSheet
	}

	for _, b := range batches {
		b.generate()
	}

	//some other checks?
	res := make([][]int, 0, len(batches))

	for _, b := range batches {
		res = append(res, *b.data)
	}

	return res, nil

}

func main() {
	start := flag.Int("start", 1, "Start page number")
	end := flag.Int("end", 0, "End page number")
	batch := flag.Int("batch", 0, "How many pages printed to be cut together")
	flag.Parse()

	res, err := GeneratePrintOrder(*start, *end, 4, *batch)
	if err != nil {
		log.Fatalf("Error during generation: %v", err)
	}

	for i, batch := range res {
		fmt.Printf("Batch %d:\n", i+1)

		for i, e := range batch[0 : len(batch)-1] {
			fmt.Printf("%d,", e)

			if (i+1)%8 == 0 {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("%d\n", batch[len(batch)-1])
	}
}
