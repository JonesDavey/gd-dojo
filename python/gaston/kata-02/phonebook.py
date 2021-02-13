import sys

def read_phonebook():
    phonebook = {}

    num_entries = int(input())
    for _ in range(num_entries):
        entry = str.split(input())
        phonebook[entry[0]] = entry[1]

    return phonebook

# Read the entries
phonebook = read_phonebook()

for line in sys.stdin:

    name = line.strip()
    if name in phonebook:
        number = phonebook[name]
        print(f"{name}={number}")
    else:
        print("Not found")
