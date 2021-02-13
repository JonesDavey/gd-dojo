#include <iostream>
#include <iomanip>
#include <vector>
#include <string>

template<typename T>
T binary_search(T value, std::vector<T> collection) {
   if (collection.empty())
   { 
      return -1;
   }
   else
   {
      if (collection.size() == 1)
      {
         if(collection[0] != value ) return -1;
         else return 0;
      }
      if (collection[collection.size()/2] > value)
      {
         // recurse on the first half of the vector
         return binary_search(value, std::vector<T>(collection.begin(), collection.end() - (collection.size()/2)));
      }
      else if(collection[collection.size()/2] < value)
      {
         // recurse on the second half
         return collection.size()/2 + binary_search(value, std::vector<T>(collection.begin() + (collection.size()/2), collection.end()));
      }
      else
      {
         return collection.size()/2;
      }
   }
}

template<typename T, typename ... Args>
bool apply_test(T (*functionUnderTest)(Args...), T expectedOutput, Args ... testInput) {
   return expectedOutput == functionUnderTest(testInput...);
}

void print_test(std::string testDescription, bool testResult)
{
   std::cout << std::right << std::setw(50) << testDescription << std::boolalpha << testResult << std::endl;
}

void apply_tests() {
   // testing for integer values
   print_test("empty vector - ", apply_test(binary_search, -1, 1, std::vector<int>({})));
   print_test("value not found - ", apply_test(binary_search, -1, 5, std::vector<int>({4})));
   print_test("only element in vector - ", apply_test(binary_search, 0, 4, std::vector<int>{4}));

   std::cout << "\n";

   // odd length vectors
   print_test("value in the middle - odd - ", apply_test(binary_search, 1, 4, std::vector<int>{1,4,6}));
   print_test("value far left - odd - ", apply_test(binary_search, 0, 4, std::vector<int>{4,5,6}));
   print_test("value far right - odd - ", apply_test(binary_search, 4, 4, std::vector<int>{1,2,3,3,4}));

   std::cout << "\n";

   // even length vectors
   print_test("value close to middle - even - ", apply_test(binary_search, 1, 4, std::vector<int>{1,4,5,6}));
   print_test("value far left - even - ", apply_test(binary_search, 0, 4, std::vector<int>{4,5,6,7}));
   print_test("value far right - even - ", apply_test(binary_search, 3, 4, std::vector<int>{1,2,3,4}));
}

int main() {
   apply_tests();
   return 0;
}
