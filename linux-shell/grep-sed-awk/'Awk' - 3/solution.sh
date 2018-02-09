awk -F" " '{
  # compute the average mark of the student
  average = (( $2 + $3 + $4) / 3)

  # default to FAIL
  rank = "FAIL"
  if (average >= 80)
    rank = "A"
  else if (average >= 60)
    rank = "B"
  else if (average > 50)
    rank = "C"

  # print the rank next to original line
  print $0 " : " rank
}'
