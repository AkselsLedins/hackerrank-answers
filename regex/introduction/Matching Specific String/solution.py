# python3 solution.py < input.txt

# **************************************************************************** #
#                                SOLUTION                                      #
# **************************************************************************** #
Regex_Pattern = r'hackerrank'
# **************************************************************************** #
#                         ̿ ̿ ̿'̿'\̵͇̿̿\з=(•_•)=ε/̵͇̿̿/'̿'̿ ̿                              #
# **************************************************************************** #

import re

Test_String = input()

match = re.findall(Regex_Pattern, Test_String)

print("Number of matches :", len(match))