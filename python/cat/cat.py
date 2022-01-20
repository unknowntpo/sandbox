import sys

def cat(fname: str):
    try:
        f = open(fname, 'rb')
    except OSError:
        print("failed to open file:", fname)
        sys.exit()
    with f:
        print(f.read())

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 cat.py <file>\n")
        sys.exit(1)
    fname = sys.argv[1]
    # TODO: Accept input from IO-Redirection
    cat(fname)

if __name__ == "__main__":
    main()
