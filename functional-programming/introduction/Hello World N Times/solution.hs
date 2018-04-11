import Control.Applicative
import Control.Monad
import System.IO


--  Print "Hello World" on a new line 'n' times.
hello_worlds :: Int -> IO ()
hello_worlds n = replicateM_ n $ putStrLn "Hello World"

main :: IO ()
main = do
    n_temp <- getLine
    let n = read n_temp :: Int
    hello_worlds n
