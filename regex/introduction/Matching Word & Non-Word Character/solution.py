# python3 solution.py < input.txt

# **************************************************************************** #
#                                SOLUTION                                      #
# **************************************************************************** #
Regex_Pattern = r"\w\w\w\W\w\w\w\w\w\w\w\w\w\w\W\w\w\w"	# Do not delete 'r'.

# **************************************************************************** #
#                         ̿ ̿ ̿'̿'\̵͇̿̿\з=(•_•)=ε/̵͇̿̿/'̿'̿ ̿                              #
# **************************************************************************** #


import re

print(str(bool(re.search(Regex_Pattern, input()))).lower())