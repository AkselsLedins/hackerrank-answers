# python3 solution.py < input.txt

# **************************************************************************** #
#                                SOLUTION                                      #
# **************************************************************************** #
Regex_Pattern = r'^[123][120][xs0][30aA][xsu][.,]$'	# Do not delete 'r'.
# **************************************************************************** #
#                         ̿ ̿ ̿'̿'\̵͇̿̿\з=(•_•)=ε/̵͇̿̿/'̿'̿ ̿                              #
# **************************************************************************** #

import re

print(str(bool(re.search(Regex_Pattern, input()))).lower())
