```-- pc[i] is the population count of i

    var pc [256] byte -- array of 256 bytes, indexed by i where each
                         byte is a count of the population count of i.

    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }

===========

All 256 values of the array are initialised to 0

i = 0 -> pc[0] = pc[0/2] + byte(00000000 & 00000001) -> pc[0] = pc[0] + 0 -> pc[0] = 0 + 0 = 0
i = 1 -> pc[1] = pc[1/2] + byte(00000001 & 00000001) -> pc[1] = pc[0] + 1 -> pc[1] = 0 + 1 = 1
i = 2 -> pc[2] = pc[2/2] + byte(00000010 & 00000001) -> pc[2] = pc[1] + 0 -> pc[2] = 1 + 0 = 1
i = 3 -> pc[3] = pc[3/2] + byte(00000011 & 00000001) -> pc[3] = pc[1] + 1 -> pc[3] = 1 + 1 = 2
i = 4 -> pc[4] = pc[4/2] + byte(00000100 & 00000001) -> pc[4] = pc[2] + 0 -> pc[4] = 1 + 0 = 1
i = 5 -> pc[5] = pc[5/2] + byte(00000101 & 00000001) -> pc[5] = pc[2] + 1 -> pc[5] = 1 + 1 = 2
i = 6 -> pc[6] = pc[6/2] + byte(00000110 & 00000001) -> pc[6] = pc[3] + 0 -> pc[6] = 2 + 0 = 2
i = 7 -> pc[7] = pc[7/2] + byte(00000111 & 00000001) -> pc[7] = pc[3] + 1 -> pc[7] = 2 + 1 = 3



[1] This shows us that for all odd numbers the right most bit has to be 1
... so the `byte(i&1)` is always the value 1 for odd numbers
... as all odd number have to a population count of at least one and all even numbers of at least 0.

[2] i / 2 effectively is a reference to the population count that
must have already been calculated at this point as its an earlier population
count and the population is counted monotonically forward. So you can always
find out the value of a previous count.

Also it refers to the population of the current number shifted once
to the right, as its a division by two, and in binary that would be
shift of values to the right, dropping off the rightmost bit.

But, you know what the right bit would be as that is determined by if
its odd or even by [1]

Hence [2] + [1] will equal the population count. Feels like a proof by induction written out... ```