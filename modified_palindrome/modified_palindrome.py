# program to Implement a modified palindrome, where it is case insensitive and ignore
# any non-alphanumeric characters.

# Assumption
# Output the palindrome when the word is a palindrome AND 
# output the string "is not a palindrome" when the word isn't a palindrome

import re
import sys


def isPalindrome(word):
    # remove special characters and numbers
    word = re.sub('[^A-Za-z]', '', word)
    
    # eliminate words with less than 2 valid characters
    if len(word) < 2:
        print("Kindly enter a valid word with more than 2 letters of the alphabet to try again")
        sys.exit(1)
    else:
        # reverse entered word
        pal = word.lower()[::-1]

        if word.lower() == pal:
            w = pal
        else:
            w = word + " is not a palindrome"

        return w


def main():
    print("Enter your desired word to test for a palindrome and press the enter key \n")
    word = input(">>> ")

    result = isPalindrome(word)
    print(result)
    
 
if __name__ == '__main__':
    main()









