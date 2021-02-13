module Main where

import System.IO
import Control.Monad
import qualified Data.Map as M

main :: IO ()
main = do
  phonebook <- readEntries
  runUntilEOF (getNumber phonebook)

getNumber :: M.Map String String -> IO ()
getNumber phonebook = do
  name <- getLine
  let number = M.lookup name phonebook
  case number of (Just num) -> putStrLn (name ++ "=" ++ show num)
                 _          -> putStrLn "Not found"


readEntries :: IO (M.Map String String)
readEntries = do
  entryCount <- getLine
  entries <- replicateM (read entryCount) readEntry
  return $ M.fromList entries

readEntry :: IO (String, String)
readEntry = do
  _ln <- getLine
  let [name, number] = words _ln
  return (name, number)

-- Util
runUntilEOF :: IO () -> IO ()
runUntilEOF f = do
  eof <- isEOF
  if eof
    then return ()
    else f >> runUntilEOF f
