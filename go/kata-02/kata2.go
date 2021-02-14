package main

import(
   "bufio"
   "io"
   "errors"
   "flag"
   "fmt"
   "strconv"
   "strings"
   "os"
)

var PhoneBook = make(map[string]string)

func validNumOfEntries(n int) bool {
  return (n > 0) && (n <= 100000)
}

// expects "string number" pairs from reader, delimited by newline
// populates the global phonebook map & returns a slice of string queries
func fileInput(r io.Reader) ([]string, error) {
   scanner := bufio.NewScanner(r)
   scanner.Split(bufio.ScanWords)

   var input[]string
   for scanner.Scan() {
      input = append(input, scanner.Text())
   }

   num, err := strconv.Atoi(input[0])
   if err != nil {
      return nil, err
   }

   // First entry in the file is expected to be some valid integer with bounds
   // 1 <= n <= 10^5
   if validNumOfEntries(num) != true {
      return nil, errors.New("Amount of entries not within valid range 0-100,000")
   }

   // Each entry consists of a name - number pair
   entryAmount := (num*2)
   if entryAmount+1 > len(input) {
      return nil, errors.New("Invalid phonebook file detected")
   }

   entries := input[1:entryAmount+1]

   for i := 0; i < len(entries); i+=2 {
      PhoneBook[entries[i]] = entries[i+1]
   }

   var queries []string
   // If there are remaining lines of the input to be parsed, then they are
   // considered top be queries to the phonebook.
   if entryAmount+1 < len(input) {
      queries = input[entryAmount+1:]
   }

   return queries, nil
}

// expects "string number" pairs from stdin, delimited by newline
// populates the global phonebook map & a slice of string queries
func readFromStdIn() ([]string, error) {
   scanner := bufio.NewScanner(os.Stdin)

   if scanner.Scan() {
      input := scanner.Text()

      num, err := strconv.Atoi(input)
      if err != nil {
         return nil, err
      }

      // First entry in the file is expected to be some valid integer with bounds
      // 1 <= n <= 10^5
      if validNumOfEntries(num) != true {
         return nil, errors.New("Amount of entries not within valid range 0-100,000")
      }
      for i := 0; i < num; i++ {
         scanner.Scan()
         line := strings.Fields(scanner.Text())
         if len(line) > 2 {
            return nil, errors.New("Invalid input, exiting.")
            os.Exit(2)
         }

         PhoneBook[line[0]] = line[1]
      }
   }

   // input should now be queries
   var queries []string
   for scanner.Scan() {
      input := scanner.Text()
      if input != "" {
         queries = append(queries, input)
      } else { // assume empty query means stop parsing
         break
      }
   }

   return queries, nil
}

func handleQueries(queries []string) {
   for _, key := range queries {
      value, ok := PhoneBook[key]
      if !ok {
         fmt.Println("Not found")
      } else {
         fmt.Println(key+"="+value)
      }
   }
}

func main() {
   filePtr := flag.String("file", "", "Input file")
   flag.Parse()

   if *filePtr != "" { // read input from file if provided
      r, err := os.Open(*filePtr)

      if err != nil {
         fmt.Println(err)
         os.Exit(2)
      }
      defer r.Close()

      queries, err := fileInput(r)
      if err != nil {
         fmt.Println("Failed to initialize: ", err)
         os.Exit(2)
      }
      handleQueries(queries)

   } else { // we're in stdin mode
      queries, err := readFromStdIn()
      if err != nil {
         fmt.Println("Error while reading stdin: ", err)
         os.Exit(2)
      }
      handleQueries(queries)
   }
}
