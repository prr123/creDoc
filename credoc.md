<p style="font-size:18pt; text-align:center;">credoc</p>

# Description
     
   credoc   
   create doc documentation   
   usage: credoc file.go   
     
   author: prr azul software   
   date: 20 May 2022   
   copyright prr azul software   
     


# Types
no types defined!    

# Functions
## main    
func main()     

## credocfil    
func credocfil(inpfilnam string)(err error)     

 function that creates doc output file    
## IsTyp    
func IsTyp(buf []byte)(typNam string, res bool)     

 function that checks whether an input line is a type definition.    
 It returns the type name if the input line is a typpe definition.    
## IsFunc    
func IsFunc(buf []byte)(funcNam string, res bool)     

 function that checks whether an input line is a function.    
 It returns the function name if the input line is a  function.    
## IsMethod    
func IsMethod(buf []byte)(methNam string, typNam string, res bool)     

 function that detemines whether a input line is a  method.    
 if so, the function returns the method name and the name of the structure the method is associated with    
## IsComment    
func IsComment (buf []byte)(desc string, res bool)     

 function that determines whether the input line is a comment line.    
 If so, it returns the comment in the desc.    

# Methods
no methods defined!    
