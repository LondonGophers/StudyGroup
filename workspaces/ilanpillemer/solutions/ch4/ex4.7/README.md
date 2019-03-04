# Thought Experiment

Modify reverse to reverse the characters of a []byte slice that
represents a UTF-8-encoded string, in place. Can you do it without
allocating new memory?

## Logic

I need to find one counter example that shows this is not possible.
In order for it not be possible I need to show of the three possible
swaps for a carefully chosen three rune byte array all of them will lose data if no
new memory is allocated to prevent this..

Which I am still looking for.
