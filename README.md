# MergeFiles

This is a very simple program that takes an arbitrary number of text files and
merges them into one, removing all duplicate lines.
More precisely, it joins all files and then, for each line, it removes all
further duplicates of that line.

*Note:* This program is in no way optimal, especially for large files.
Everything is done in memory, duplicates are removed even if they occur in the same
input file, and the outcome depends on the order of the input files.
Thus, for general use, some further tweaks may be needed.
