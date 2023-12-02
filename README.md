# Advent of Code 2023 - Go
My Advent of Code solutions for 2023. Each solution is done using Go and *hopefully* only the stdlib.

## Thoughts for Each Day
This section contains my thoughts on each of the days, including struggles I had solving each problem, and fun pieces of information I learned along the way.

### Day 1
The toughest part of this one for me was accounting for overlapping matches in Part 2. I was initially using RegEx for a solution, but I realized (after eventually going crazy) that finding a solution this way was going to be more tedious than just using the `slices` package.