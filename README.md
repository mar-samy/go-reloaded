Go Reloaded Overview

Go Reloaded is a command-line text processing tool written in Go.

The program reads an input text file, applies a series of transformations based on embedded commands, corrects punctuation and grammar formatting, and writes the processed result to a new output file.

Although the functionality is relatively small, the project focuses on implementing the logic manually rather than relying on external libraries.

Why I Built This?
------------------
This is the very first software project I have completed entirely on my own.

Before starting, I knew the basics of Go syntax, but I had never built a complete application from scratch.

The biggest challenge wasn't writing Go code.

It was learning how to convert ideas into software.

Taking a problem description...

Breaking it into functions...

Designing the program flow...

Handling unexpected cases...

Debugging...

Testing...

and finally connecting everything together.

Every function in this project was written manually as part of my learning journey.

Features

The application supports:

-> Number Conversion
 - Hexadecimal → Decimal
 - Binary → Decimal

-> Text Modifiers
 - (up)
 - (low)
 - (cap)

including

 - (up, n)
 - (low, n)
 - (cap, n)

which modify multiple previous words.

Punctuation Formatting

Automatically fixes spacing around

- .
- ,
- !
- ?
- :
- ;

and grouped punctuation such as
- ...
- !!
- !?
- ?!

  
Quote Formatting

Formats quoted phrases correctly by removing unnecessary spaces inside quotation marks.

Article Correction
-
Automatically converts
 
 - a → an

when followed by a word beginning with

 - a
 - e
 - i
 - o
 - u
 - h
   
Concepts Practiced

Throughout this project I practiced:

 - Go command-line applications
 - File handling
 - Reading and writing files
 - String manipulation
- Slice manipulation
- Parsing
- Number conversion
- Error handling
- Modular programming
- Control structures
- Algorithm design
- Debugging
- Testing edge cases

Technologies
- Go
- Standard Library only

No external packages were used.

Biggest Lesson

The biggest lesson from this project wasn't learning Go syntax.

It was learning how to think through a programming problem.

Many times I knew exactly what I wanted the program to do...

but translating that idea into code was difficult.

Debugging forced me to understand my own logic instead of simply writing code that "looked right."

That experience made this project far more valuable than simply completing another tutorial.

Future Improvements

Possible future enhancements include:

 - Unit testing
 - Performance optimizations
 - Cleaner parser architecture
 - Better error reporting
 - Additional text transformations
 - Refactoring repeated modifier logic


A Personal Note

If someone reading this is also starting their programming journey, I'd like to say one thing:

Programming can be frustrating.

There were many moments during this project where I felt completely stuck.

But every bug I fixed taught me something that no tutorial could.

This project isn't perfect.

It's simply the first milestone in what I hope will be a long career in software engineering.
