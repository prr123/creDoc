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
	"credoc/util"
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
	fmt.Println("*** success ***")
}

func credocfil(inpfilnam string)(err error) {
// function that creates doc output file

	var outfilnam string
	// check whether input file is valid
	inpfil, err := os.Open(os.Args[1])
	defer inpfil.Close()
	if err != nil {return fmt.Errorf("os.Open: %v\n", err)}

	inpfilInfo,_ := inpfil.Stat()
	inpSize := inpfilInfo.Size()

	fmt.Printf("*** input file:  %s Size: %d\n", inpfilnam, inpSize)

	// create output file name
//	found := false
	for i:=0; i< len(inpfilnam); i++ {
		if inpfilnam[i] == '.' {
//			found = true
			outfilnam = string(inpfilnam[:i])
			break
		}
	}

	fmt.Printf("*** output file: %s\n", outfilnam + ".md")

	outfil, err := os.Create(outfilnam + ".md")
	if err != nil { return fmt.Errorf("os.Create: %v\n", err)}
	defer outfil.Close()

	prefix := "<p style=\"font-size:18pt; text-align:center;\">"
	suffix := "</p>\n\n"

	outfil.WriteString(prefix + outfilnam + suffix)

	bufp := make([]byte, inpSize)

//	for iblock :=0; iblock < 10; iblock++ {
	nb, _ := inpfil.Read(bufp)
//		if err != nil {return fmt.Errorf("read: %d %v", nb, err)}

	// top comments
	linSt := 0
	ilin := 1
	introlin := 0
	for i:=0; i< nb; i++ {
		if bufp[i] == '\n' {
			linEnd := i
//			fmt.Printf("line %d: %s\n", ilin, string(bufp[linSt:linEnd]))
			if (bufp[linSt] == '/') && (bufp[linSt+1] == '/') {
				outfil.WriteString("  " + string(bufp[linSt+2: linEnd]) + "   \n")
				introlin = ilin
			}
			if (introlin >0) && (introlin < ilin) {
				outfil.WriteString("\n\n")
				break
			}
			ilin++
			linSt = i+1
		}
	}

	//type definitions

	outfil.WriteString("# Types\n\n")


	outfil.WriteString("# Functions\n\n")
	linSt = 0
	ilin = 1
	for i:=0; i< nb; i++ {
		if bufp[i] == '\n' {
			linEnd := i
			fnam, res := Isfunc(bufp[linSt: linEnd])
			if res {
				outfil.WriteString(fnam+"\n")
			}
			ilin++
			linSt = i+1
		}
	}
	outfil.WriteString("# Methods\n\n")

	return nil
}

func Isfunc(buf []byte)(fnam string, res bool) {

	if len(buf) < 5 { return "", false }

	if buf[0] != 'f' {return "", false}
	if buf[1] != 'u' {return "", false}
	if buf[2] != 'n' {return "", false}
	if buf[3] != 'c' {return "", false}
	if buf[4] != ' ' {return "", false}

	for i:= 5; i<len(buf); i++ {
		if utilLib.IsAlpha(buf[i]) {continue}
//		if buf[i]
	}
	fnam = "xyz"
	return fnam, true
}
