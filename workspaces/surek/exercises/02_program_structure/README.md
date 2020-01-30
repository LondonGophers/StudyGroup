# Chapter 2 - Program Structure

## Packages and Files

### Exercise 2.1
Add types, constants, and functions to `tempconv` for processing temperatures
in the Kelvin scale, where Kelvin is -273.15°C and a difference of 1K
has the same magnitude as 1°C.

## Imports

### Exercise 2.2
Write a general purpose unit-conversion program analogous to `cf` that reads
numbers from its command-line arguments or from the standard imput if there are
no arguments, and converts each nunber into units like temperature in Celsius
and Fahrenheit, length in feet and meters, weight in pounds, kilograms, and the
like.

## Package Initialization

### Exercise 2.3
Rewrite `PopCount` to use a loop instead of a single expression. Compare the
performance of the two versions. (Section 11.4 shows how to compare the
performance of different implementations systematically.)

### Exercise 2.4
Write a version of `PopCount` that counts bits by shifting its argument through
64 bit positions, testing the rightmost bit each time. Compare its performance
to the table-lookup version.

### Exercise 2.5
The expression `x&(x-1)` clears the rightmost non-zero bit of x. Write a version
of `PopCount` that counts bits by using this fact, and assess its performance.
