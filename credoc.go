//
// credoc
// create doc documentation
// usage: credoc file.go
//
// author: prr azul software
// date: 20 May 2022
// copyright prr azul software
//
package main

import (
	"os"
	"fmt"
)

func main() {

	numArg := len(os.Args)

	switch numArg {
	case 0, 1:
		fmt.Printf("no input file provided\n")
		fmt.Printf("usage is: credoc file\n")
		os.Exit(1)

 	case 2:
		fmt.Printf("input file: %s\n", os.Args[1])

	default:
		fmt.Printf("in correct number of command line parameters: %d\n", numArg)
		fmt.Printf("usage is: credoc file\n")
		os.Exit(1)
	}

	err := credocfil(os.Args[1])
	if err != nil {
		fmt.Printf("error - credocfil: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("success")
}

func credocfil(inpfilnam string)(err error) {
// function that creates doc output file

	var outfilnam string
	var buf [4096]byte
	// check whether input file is valid
	inpfil, err := os.Open(os.Args[1])
	defer inpfil.Close()
	if err != nil {return fmt.Errorf("os.Open: %v\n", err)}

	// create output file name
//	found := false
	for i:=0; i< len(inpfilnam); i++ {
		if inpfilnam[i] == '.' {
//			found = true
			outfilnam = string(inpfilnam[:i])
			break
		}
	}

	fmt.Printf("output file: %s\n", outfilnam)

	outfil, err := os.Create(outfilnam + ".md")
	if err != nil { return fmt.Errorf("os.Create: %v\n", err)}
	defer outfil.Close()

	outstr:= fmt.Sprintf("%s\n", outfilnam)
	outfil.WriteString(outstr)

	bufp := buf[:]

	linSt := 0
	ilin := 1
	for iblock :=0; iblock < 10; iblock++ {
		offset := int64(iblock * 4096)
		nb, _ := inpfil.ReadAt(bufp, offset)
//		if err != nil {return fmt.Errorf("read: %d %v", nb, err)}
		for i:=0; i< nb; i++ {
			if bufp[i] == '\n' {
				linEnd := i
				fmt.Printf("line %d: %s\n", ilin, string(bufp[linSt:linEnd]))
				ilin++
				linSt = i+1
			}
		}
		if nb < 4096 {break}
	}

	outfil.WriteString(string(bufp[:linSt]))

	return nil
}
