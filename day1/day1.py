import re

def spelled_out_to_int(spelled_number):
    number_dict = {
        'zero': 0, 'one': 1, 'two': 2, 'three': 3, 'four': 4,
        'five': 5, 'six': 6, 'seven': 7, 'eight': 8, 'nine': 9
    }

    # Convert the spelled-out number to lowercase for case-insensitive matching
    spelled_number_lower = spelled_number.lower()

    # Check if the spelled-out number is in the dictionary
    if spelled_number_lower in number_dict:
        return number_dict[spelled_number_lower]
    else:
        return spelled_number

with open("input.txt") as f:
    day1_in = f.read()
    output = 0
    for line in day1_in.splitlines():
        match = re.search(r'\d', line)
        match2 = re.search(r'\d', line[::-1])
        output += int(f'{match.group()}{match2.group()}') 

    # print(output)

#part2


regex_new ='(?:zero|one|two|three|four|five|six|seven|eight|nine|\d)' 
pattern_rev = 'zero|one|two|three|four|five|six|seven|eight|nine'[::-1]
regex_new_b =f'(?:{pattern_rev}|\d)' 
# print(pattern_rev, regex_new_b)

with open("input.txt") as f:
    day1_in = f.read()
    output = 0
    count = 0
    for line in day1_in.splitlines():
        match_f = re.search(regex_new,line)
        match_b = re.search(regex_new_b,line[::-1])

        # print(line,'\n',match_f,match_b)
        match_f_int = spelled_out_to_int(match_f.group())
        match_b_int = spelled_out_to_int(match_b.group()[::-1])
        # print(match_f_int, match_b_int)
        output += int(f'{match_f_int}{match_b_int}')

    print(output)