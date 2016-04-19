package main

import (
	"fmt"
	"os"

	"code.google.com/p/biogo.bam"
)

var (
	ioBam     *os.File
	ioBamIdx  *os.File
	bamReader *bam.Reader
	bamIndex  *bam.Index
	err       error
	rid       int = 0
	beg       int = 25000
	end       int = 55000
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	bamFile := "/mnt/home/wei/biogo.bam/src/code.google.com/p/biogo.bam/test_fetch/data1.bam"
	check(err)
	bamIdx := "/mnt/home/wei/biogo.bam/src/code.google.com/p/biogo.bam/test_fetch/data1.bam.bai"
	check(err)
	ioBam, err = os.Open(bamFile)
	ioBamIdx, err = os.Open(bamIdx)
	defer ioBam.Close()
	defer ioBamIdx.Close()

	bamReader, err = bam.NewReader(ioBam)
	check(err)
	bamIndex, err = bam.ReadIndex(ioBamIdx)
	check(err)

	fi, ok := bamReader.Fetch(bamIndex, rid, beg, end)
	fmt.Println(ok)
	if ok {
		for fi.Next() {
			algn := fi.Get()
			if algn != nil {
				fmt.Println(algn.Pos)
			}
		}
	}

}
