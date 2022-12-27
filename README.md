Go-reloaded
A simple text completion/editing/auto-correction tool.

List of possibilities:

1.  Replace the word with the decimal version of the word. (hex)
2.  Replace the word before with the decimal version of the word. (bin)
3.  Converts the word before with the Uppercase version of it. (up)
4.  converts the word before with the Lowercase version of it. (low)
5.  Converts the word before with the capitalized version of it. (cap)
6.  For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.
7.  Every instance of the punctuations ., ,, !, ?, : and ; will be close to the previous word and with space apart from the next one. Except if there are groups of punctuation like: ... or !?.
8.  The punctuation mark ' will always be found with another instance of it and they will be placed to the right and left of the word in the middle of them.
9.  If there are more than one word between the two ' ' marks, the program will place the marks next to the corresponding words.
10. Every instance of a will be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h.

HOW TO USE the Project.
In a file called sample.txt, place the following text:

1. If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
2. I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
3. Don not be sad ,because sad backwards is das . And das not good .
4. harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '

In result.txt you might see:

1. If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
2. I have to pack 5 outfits. Packed 26 just to be sure
3. Don not be sad, because sad backwards is das. And das not good.
4. Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
