# ComboCrafter

**ComboCrafter** — hashcat-slyle mask wordlist generator.

---

## Install

1. Clone repo:

```bash
git clone https://github.com/yourusername/maskgen.git
cd ComboCrafter
```

2. Build:

```
bash
go build .
```

## Usage

Mask syntax:

```
?l — lowercase letter (a-z).
?d — dight (0-9).
?a — lowercase letter or dight.
?s — hyphen (-).
?A — all symbols from domain characters set.
?w — word from wordlist.
```

Flags:

```
-m — Path to the file containing masks.
-w — Path to the file containing the wordlist.
-o — Path to the output file.
-stdout — Print output to stdout instead of a file.
```

### Example

**masks.txt**:

```
?l?d
?w?d?s
```

**wordlist.txt**:

```
word1
word2
```

Run:

```
bash
./ComboCrafter -m masks.txt -w wordlist.txt -o output.txt
```

**output.txt**:

```
a0
a1
...
z9
word10-
word11-
...
word29-
word20-
word21-
...
word29-
```
