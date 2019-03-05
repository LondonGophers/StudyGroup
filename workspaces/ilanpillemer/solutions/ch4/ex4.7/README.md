# Thought Experiment

Modify reverse to reverse the characters of a []byte slice that
represents a UTF-8-encoded string, in place. Can you do it without
allocating new memory?

## Logic

I need to find one counter example that shows this is not possible. In
order for it not to be possible I needed to show of the three possible
swaps for a carefully chosen three rune byte array all of them will
lose data if no new memory is allocated to prevent this..

Upon looking for such an example it became clear that although it is 
impossible to swap "over" an rune without needing temporary storage, 
it is always possible to swap adjacent runes.

Thus, if you swap iteratively each rune from right to left, you can reverse
a unicode byte array without allocating any new memory..

This is implemented in `main.go`.
