# python3 solution.py < input.txt

# **************************************************************************** #
#                                SOLUTION                                      #
# **************************************************************************** #
Regex_Pattern = r"\S\S\s\S\S\s\S\S"	# Do not delete 'r'.

# **************************************************************************** #
#                         ̿ ̿ ̿'̿'\̵͇̿̿\з=(•_•)=ε/̵͇̿̿/'̿'̿ ̿                              #
# **************************************************************************** #


import re

print(str(bool(re.search(Regex_Pattern, input()))).lower())
