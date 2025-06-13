package concurrency

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

// creates random files with random texts s is number of strings to add ,n is no of files
func CreateFiles(n, nt int, s string) error {

	if n == 0 {
		return fmt.Errorf("files can't be zero ")
	}
	err := os.Mkdir(s, 0777)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		basePath, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(basePath, s, RandString("file_"))
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		for j := 0; j < nt; j++ {
			if _, err := f.Write([]byte(s)); err != nil {
				return err
			}
		}
	}
	return nil

}

// pure cpu bound work (simply adding slice items)
func AddSLiceItems(s []int64) int64 {

	var sum int64 = 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// concurrent version of addSliceItems
func AddSliceItemsC(goroutines int64, s []int64) int64 {
	//divide slices equally among go routines
	sp := SlicesToProcess(goroutines, s)
	// launch routines to process these slices
	var res int64 = 0
	var wg sync.WaitGroup
	for i := range goroutines {
		wg.Add(1)
		go func(idx int64) {
			defer wg.Done()
			var sum int64 = 0
			for _, v := range sp[idx] {
				sum += v
			}
			atomic.AddInt64(&res, sum)
		}(i)

	}
	wg.Wait()

	return res
}
func addSliceItemsCChannels(goroutines int64, s []int64) int64 {
	//divide slices equally among go routines
	sp := SlicesToProcess(goroutines, s)
	// launch routines to process these slices
	resChan := make(chan int64, goroutines)
	var wg sync.WaitGroup

	for i := range goroutines {
		wg.Add(1)
		go func(idx int64) {
			defer wg.Done()
			var sum int64 = 0
			for _, v := range sp[idx] {
				sum += v
			}
			resChan <- sum
		}(i)

	}
	wg.Wait()
	close(resChan)
	var res int64
	for v := range resChan {
		res += v
	}

	return res

}

// utils fn for addSliceC returns slice of slices for go routines
func SlicesToProcess(g int64, s []int64) [][]int64 {

	elementsPerGSlice := []int64{}
	r := int64(len(s) % int(g))
	q := int64((len(s) - int(r))) / g
	res := [][]int64{}
	var i, j int64 = 0, 0
	for i < g {
		if j < r {
			elementsPerGSlice = append(elementsPerGSlice, q+1)
		} else {
			elementsPerGSlice = append(elementsPerGSlice, q)
		}
		j++
		i++
	}
	i, j = 0, 0
	for v := range g {
		j += elementsPerGSlice[v]
		res = append(res, s[i:j])
		i = j
	}
	return res
}
