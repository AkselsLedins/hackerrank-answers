# ./solution.sh < input.txt

array=()

while read line
do
  array=("${array[@]}" $line)
done

echo "${array[@]}"
