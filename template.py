import argparse

parser = argparse.ArgumentParser()

parser.add_argument("--inputFile", default = "input.txt", help = "Relative path to the input file")
parser.add_argument("--lineDelimiter", default = "\n", help = "The end of line delimiter")

if __name__ == "__main__":
    args = parser.parse_args()

    with open(args.inputFile, 'r') as f:

        lineIdx = -1
        for line in f:
            lineIdx += 1
            token = line.split(args.lineDelimiter)
            
            if len(token) > 2:
                print(f"[WARN] Skipping line with index = {lineIdx} and value = {token}")
            
            firstToken = token[0]
            print(f"The line contents are {firstToken}")
    