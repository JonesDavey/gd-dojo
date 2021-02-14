module Main where

main :: IO ()
main = runTests

binarySearch :: Ord a => a -> [a] -> Int
binarySearch target list = _binarySearch target (zip list [0..])

-- Helper method that takes associative list of tuples
-- of the form (element, index), so that the index can be easily retrieved.
_binarySearch :: Ord a => a -> [(a, Int)] -> Int
_binarySearch _ [] = (-1)
_binarySearch target [(x, index)] = if x == target then index else (-1)
_binarySearch target list =
  let midpoint = (length list) `quot` 2
      (leftList, rightList) = splitAt midpoint list

      list' = if (fst . head) rightList <= target then rightList else leftList
  in _binarySearch target list'

applyTest :: Int -> Int -> [Int] -> IO ()
applyTest expected target list =
  if result == expected
  then putStrLn "Test Passed"
  else  putStrLn $
        "Test Failed. Expected: " ++ show expected
        ++ " Actual: " ++ show result
  where result = binarySearch target list

runTests :: IO ()
runTests = do
  applyTest (-1) 3 []
  applyTest (-1) 3 []
  applyTest 0 1 [1]
  applyTest 0 1 [1]

  applyTest 0 1 [1, 3, 5]
  applyTest 1 3 [1, 3, 5]
  applyTest 2 5 [1, 3, 5]
  applyTest (-1) 0 [1, 3, 5]
  applyTest (-1) 2 [1, 3, 5]
  applyTest (-1) 4 [1, 3, 5]
  applyTest (-1) 6 [1, 3, 5]

  applyTest 0 1 [1, 3, 5, 7]
  applyTest 1 3 [1, 3, 5, 7]
  applyTest 2 5 [1, 3, 5, 7]
  applyTest 3 7 [1, 3, 5, 7]
  applyTest (-1) 0 [1, 3, 5, 7]
  applyTest (-1) 2 [1, 3, 5, 7]
  applyTest (-1) 4 [1, 3, 5, 7]
  applyTest (-1) 6 [1, 3, 5, 7]
  applyTest (-1) 8 [1, 3, 5, 7]


