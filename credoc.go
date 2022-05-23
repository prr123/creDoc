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
	outfil.WriteString("# Description\n")
	linSt := 0
	ilin := 1
	endTop := 0
	istate := 1
	linEnd := 0
	endTopLin := 1
	for i:=0; i< nb; i++ {
		if bufp[i] != '\n' {continue}
		ilin++
		switch istate {
		case 1:
			linEnd = i
//			fmt.Printf("line %d: %s\n", ilin, string(bufp[linSt:linEnd]))
			desc, res := IsComment(bufp[linSt: linEnd])
			if res {
				outfil.WriteString("  " + desc + "   \n")
				istate = 2
			}
			linSt = i+1
		case 2:
			linEnd = i
//			fmt.Printf("line %d: %s\n", ilin, string(bufp[linSt:linEnd]))
			desc, res := IsComment(bufp[linSt: linEnd])
			if res {
				outfil.WriteString("  " + desc + "   \n")
				istate = 2
				linSt = i+1
				endTop = i
				endTopLin = ilin
			}
		}
	}

	if endTop == 0 {
		outfil.WriteString("no description!    \n")
	}

	outfil.WriteString("\n\n")

	//type definitions
	outfil.WriteString("# Types\n")
	istate = 1
	linEnd = 0
	ilin = endTopLin
	linSt = 0
	found := false
//fmt.Println("nb: ", nb, " endTop: ", endTop)
	for i:=endTop; i< nb; i++ {
		if bufp[i] != '\n' {continue}
		ilin++
		linEnd = i
		switch istate {
		case 1:
			typNam, res := IsTyp(bufp[linSt: linEnd])
			if res {
				outfil.WriteString("## " + typNam + "    \n")
				outfil.WriteString(string(bufp[linSt: linEnd-1]) + "    \n\n")
				istate = 2
				found = true
			}
			linSt = i+1
		case 2:
			desc, res := IsComment(bufp[linSt: linEnd])
			if res {
				outfil.WriteString(desc + "    \n")
			}
			linSt = i+1
			if !res {istate = 1; break}
		}
	}

	if !found {
		outfil.WriteString("no types defined!    \n")
	}

	// function
	outfil.WriteString("\n# Functions\n")
	istate = 1
	linEnd = 0
	ilin = endTopLin
	linSt = 0
	found = false
//fmt.Println("nb: ", nb, " endTop: ", endTop)
	for i:=endTop; i< nb; i++ {
		if bufp[i] != '\n' {continue}
		ilin++
		linEnd = i
		switch istate {
		case 1:
			funcNam, res := IsFunc(bufp[linSt: linEnd])
			if res {
				outfil.WriteString("## " + funcNam + "    \n")
				outfil.WriteString(string(bufp[linSt: linEnd-1]) + "    \n\n")
				istate = 2
				found = true
			}
			linSt = i+1
		case 2:
			desc, res := IsComment(bufp[linSt: linEnd])
			if res {
				outfil.WriteString(desc + "    \n")
			}
			linSt = i+1
			if !res {istate = 1; break}
		}
	}

	if !found {
		outfil.WriteString("no functions defined!    \n")
	}

	// method
	outfil.WriteString("\n# Methods\n")
	istate = 1
	linEnd = 0
	ilin = endTopLin
	linSt = 0
	found = false
	for i:=endTop; i< nb; i++ {
		if bufp[i] != '\n' {continue}
		ilin++
		linEnd = i
		switch istate {
		case 1:
			methNam, methTyp, res := IsMethod(bufp[linSt: linEnd])
			if res {
				outfil.WriteString("## " + methTyp +": " + methNam + "    \n")
				outfil.WriteString(string(bufp[linSt: linEnd-1]) + "    \n\n")
				istate = 2
				found = true
			}
			linSt = i+1
		case 2:
			desc, res := IsComment(bufp[linSt: linEnd])
			if res {
				outfil.WriteString(desc + "    \n")
			}
			linSt = i+1
			if !res {istate = 1; break}
		}
	}
	if !found {
		outfil.WriteString("no methods defined!    \n")
	}

	return nil
}

func IsTyp(buf []byte)(typNam string, res bool) {
// function that checks whether an input line is a type definition.
// It returns the type name if the input line is a typpe definition.

	if len(buf) < 5 { return "", false }

	if buf[0] != 't' {return "", false}
	if buf[1] != 'y' {return "", false}
	if buf[2] != 'p' {return "", false}
	if buf[3] != 'e' {return "", false}
	if buf[4] != ' ' {return "", false}

	fnamSt := 0
	for i:= 5; i<len(buf); i++ {
		if buf[i] == ' ' {continue}
		if utilLib.IsAlpha(buf[i]) {
			fnamSt = i
			break
		}
	}

	if fnamSt == 0 {return "", false}

	fnamEnd := 0
	for i:= fnamSt; i<len(buf); i++ {
		if (buf[i] == ' ') {
			fnamEnd = i
			break
		}
	}

	if fnamEnd == 0 {return "", false}

	typNam = string(buf[fnamSt:fnamEnd])
	return typNam, true
}

func IsFunc(buf []byte)(funcNam string, res bool) {
// function that checks whether an input line is a function.
// It returns the function name if the input line is a  function.

	if len(buf) < 5 { return "", false }

	if buf[0] != 'f' {return "", false}
	if buf[1] != 'u' {return "", false}
	if buf[2] != 'n' {return "", false}
	if buf[3] != 'c' {return "", false}
	if buf[4] != ' ' {return "", false}

	fnamSt := 0
	for i:= 5; i<len(buf); i++ {
		if buf[i] == ' ' {continue}
		if buf[i] == '(' {return "", false}
		if utilLib.IsAlpha(buf[i]) {
			fnamSt = i
			break
		}
	}

	if fnamSt == 0 {return "", false}

	fnamEnd := 0
	for i:= fnamSt; i<len(buf); i++ {
		if (buf[i] == ' ') || (buf[i] == '(') {
			fnamEnd = i
			break
		}
	}

	if fnamEnd == 0 {return "", false}

	funcNam = string(buf[fnamSt:fnamEnd])
	return funcNam, true
}

func IsMethod(buf []byte)(methNam string, typNam string, res bool) {
// function that detemines whether a input line is a  method.
// if so, the function returns the method name and the name of the structure the method is associated with

	if len(buf) < 5 { return "","", false }

	if buf[0] != 'f' {return "", "", false}
	if buf[1] != 'u' {return "", "", false}
	if buf[2] != 'n' {return "", "", false}
	if buf[3] != 'c' {return "", "", false}
	if buf[4] != ' ' {return "", "", false}

	typSt := 0
	for i:= 5; i<len(buf); i++ {
		if buf[i] == ' ' {continue}
		if utilLib.IsAlpha(buf[i]) {return "", "", false}
		if buf[i] == '(' {
			typSt = i
			break
		}
	}

	if typSt == 0 {return "", "", false}

	typEnd := 0
	for i:= typSt; i<len(buf); i++ {
		if buf[i] == ')' {
			typEnd = i
			break
		}
	}

	if typEnd == 0 {return "", "", false}

	typNam = string(buf[typSt+1:typEnd-1])
	methSt := 0
	for i:= typEnd + 1; i<len(buf); i++ {
		if utilLib.IsAlpha(buf[i]) {
			methSt = i
			break
		}
	}

	if methSt == 0 {return "", "", false}

	methEnd := 0
	for i:= methSt; i<len(buf); i++ {
		if (buf[i] == ' ') || (buf[i] == '(') {
			methEnd = i
			break
		}
	}

	if methEnd == 0 {return "", "", false}

	methNam = string(buf[methSt:methEnd])

	return methNam, typNam, true
}

func IsComment (buf []byte)(desc string, res bool) {
// function that determines whether the input line is a comment line.
// If so, it returns the comment in the desc.

	if len(buf) < 2 {return "", false}
	if buf[0] != '/' {return "", false}
	if buf[1] != '/' {return "",  false}

	desc = string(buf[2:])

	return desc, true
}
