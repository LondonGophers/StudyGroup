# Thought Experiment

Modify reverse to reverse the characters of a []byte slice that
represents a UTF-8-encoded string, in place. Can you do it without
allocating new memory?

## Logic

### Step 1

Suppose you have a 3 rune unicode with
+ `x` a 4 byte rune
+ `y` a 3 byte rune
+ `z` a 1 byte rune

`xxxxyyyz`

### Step 2

Now you reverse the first and last runes... getting
+ `zxxxxxxx`

Now you have *lost* the second rune as it was overwritten in the swap.

Hence you cannot swap these two runes without first storing something in memory.

# In Conclusion

You cannot reverse a unicode string without allocating new memory for temporary storage.
If I am wrong, please let me know by raising an issue and explaining how.

