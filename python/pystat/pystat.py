from sys import argv
import os

# in main
# parse file name from flag

if len(argv) < 2:
    print("Usage: python pystat.py <filename>")
    exit(1)

filename = argv[1]
print(f"filename: {filename}")
print(argv)

# use api to call syscall:stat
# print stat to stdout
# Ref: https://docs.python.org/3/library/os.html#os.stat
# What does parameter * means ?
statinfo = os.stat(filename)
print(statinfo)
