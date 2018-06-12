# ThePuzzler
A go project that creates an iterative (Onion) hash.

The program allows you to solve and create puzzle using:

- Md5Hash
- Sha1Hash
- Sha256Hash
- Sha512Hash
- Sha3256

When the goal is to solve a puzzle, it will put the solution in a JSON file of the form:

- Layer   : Which layer of the algo the entry points to
- Algorithm  : which hashing algorithm was used for said layer
- Time     : how long it took to parse out the layer
- Question  : what plain text was found in the layer
- Body      : re reminder of the puzzle

The program uses flags to differentiate between solving and creation modes as well as to get desired file paths.

### Mode Flag
The mode flag, denoted by a `-m` in the command line takes two possible values: s & c (solve and create)
### Puzzle Path Flag
The puzzle path flag, denoted by a `-pp`, has two different effect depending on the mode. If the puzzle is in `solve` mode, the program will solve the puzzle found in that path, if the puzzle is in `create` mode, the program will put the created puzzle in that path. it is defaulted to look for `puzzle.txt` in the current directory. 

## Solution File Path Flag
The solution file path flag, denoted by `-sp` allows you to specify where you want the program to write the solution to (when in solve mode).
    
### Max Layer Flag
The program will stop solving the puzzle if the number of layers exceed the number specified by the max layer flag, denoted by `-ln`.   

### Configuration Flag
This flag, denoted by `-config`, is the file path to to configuration used in creation mode. There is an example in the Create a Puzzle section later.

## Solve a Puzzle
To solve a puzzle, all you need is to set the program to `solve` mode and give it a path

Using the executable and you are in it's directory:
```bash
./main -m=s -pp=example_puzzle.txt 
 ```
To save the solution in a specific path,
```bash
./main -m=s -pp=example_puzzle.txt -sp=ILikePie.json
 ```

 ## Create a Puzzle

To create a puzzle, you need to first create a configuration file json file that mimics the json below
```javascript
{"input" : [{"layer"   : 1,
    "algorithm" : "sha-1",
    "message": "to be or not to be"},
        {"layer"   : 2,
    "algorithm" : "md5",
    "message": "I like pie"}]
        }
```

Assuming you are using the executable, and your files are in the same directory 
```bash
./main -m=c -pp=example_created_puzzle.txt -config=example_puzzle_config.json
```

The maximum number of layers you can create is 20.

## 