# creDoc
## Program Description
credoc is a program to facilitate the documentation of golang code files.

credoc requires an go file as second argument in the command line.
The program reads the go file as input and generates a markdown file as output. The markdown file contains the code documentation of the go input file. 

The program first reads the top comment lines to generate a descriptiont of the go file being documented. The program then iterates through the rest of the input file three times to find all the type, function and method definitions.

The markdown file currently has four sections:
- Description
- Types
- Functions
- Methods

### Description
The first sections contains a description of the code. The description is generated from the top comment lines in the input file.

### Types
The type section lists each type definition contained in the file. If a type definition is followed by one or more comment lines, the comment lines are read and used to provide a description of the type.

### Functions
The function section lists each function definition. Similarily to the type definition, if a function definition is immediatly followed by one or more comment lines, these comment lines are assumed to provide a description of that function.

### Methods
The method section operates similarily to the function section.


## Usage
The program greatly facilitates the production of a file's documentation.

The author simply needs to add a few comment lines at the beginning of the file to describe the code in general and its functionality.
The code author then only needs to add one or more comment lines directly following each type, function and method definition to describe the respective type, function or method.

## Installation
download credoc.go and util/utilLib.go

compile file:     
*go build credoc.go*

add alias in .bashrc

## Output Sample
A sample documentatin of the credoc.go can be found in the markdown file credoc.md.
