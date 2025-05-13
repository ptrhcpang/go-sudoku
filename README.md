## Command line Sudoku solver

### introduction

This is a simple Sudoku solver written in Go 
primarily to help me get acclimatised to the 
language back in July of 2024.

### how to use it 

It does the following:

On running the program, it takes 9 inputs.

Each input represents a row of a Sudoku board,
starting from the top. Each input consists of 
a string of integers `0` through `9` separated 
only by commas. A `0` represents 
an empty space on an unsolved board.

If there are fewer than nine numbers in any input, 
the program populates that row of the unsolved board by 
the numbers in the input starting from the left, 
with the missing numbers implicitly understood as 
trailing noughts (i.e., spaces).

If there are more than nine numbers in any input, 
only the first nine will be used.

If there are fewer than nine inputs, the uninitiated 
rows are implicitly understood as being populated 
by noughts only.

### what it does

The program outputs the unsolved board, and, 
if the problem is neither underdetermined nor 
(inconsistently) over-determined, it outputs 
the solved board.
