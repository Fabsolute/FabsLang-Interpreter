# Fabs Lang Interpreter

Easiest programming language you can see

| Command       | Explanation                                                                                    | 
| ------------- | ---------------------------------------------------------------------------------------------- |
| f             | output the byte at the data pointer.                                                           |
| fa            | increment (increase by one) the byte at the data pointer.                                      |
| fab           | decrement (decrease by one) the byte at the data pointer.                                      |
| fabs          | increment the data pointer (to point to the next cell to the right).                           |
| fabsl         | decrement the data pointer (to point to the next cell to the left).                            |
| fabsla        | accept one byte of input, storing its value in the byte at the data pointer.                   |
| fabslan       | if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching fabslang command.|
| fabslang      | if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching fabslan command                     |


| Command       | C Equivalent                                                                                   | 
| ------------- | ---------------------------------------------------------------------------------------------- |
| f             | putchar(*ptr);                                                           |
| fa            | ++*ptr;                                      |
| fab           | --*ptr;                                      |
| fabs          | ++ptr;                           |
| fabsl         | --ptr;                           |
| fabsla        | *ptr=getchar();                   |
| fabslan       | while (*ptr) { |
| fabslang      | }                    |


