# Tests should be run using "python test_modified_palindrome.py"
import unittest
from unittest import mock

from modified_palindrome import isPalindrome

m = mock.Mock()


class TestIsPalindrome(unittest.TestCase):
    """Tests for isPalindrome.py."""

    def test_eye_is_a_palindrome(self):
        """is eye a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('eye')

        self.assertEqual(result, 'eye')

    def test_madam_is_a_palindrome(self):
        """is madam a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('madam')

        self.assertEqual(result, 'madam')

    def test_racecar_is_a_palindrome(self):
        """is racecar a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('racecar')

        self.assertEqual(result, 'racecar')

    def test_noon_is_a_palindrome(self):
        """is noon a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('noon')

        self.assertEqual(result, 'noon')
    
    def test_with_nonalphanumeric_characters_is_a_palindrome(self):
        """is _!E#@y#e a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('_!E#@y#e')

        self.assertEqual(result, 'eye')

    def test_with_nonalphanumeric_characters2_is_a_palindrome(self):
        """is MaDaM!! a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('MaDaM!!')

        self.assertEqual(result, 'madam')

    def test_abc_is_not_a_palindrome(self):
        """is abc a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('abc')

        self.assertEqual(result, 'abc is not a palindrome')

    def test_anything_is_not_a_palindrome(self):
        """is anything a palindrome?"""
        m.some_guy.side_effect = isPalindrome
        result = m.some_guy.side_effect('anything')

        self.assertEqual(result, 'anything is not a palindrome')

    def test_wrong_user_input(self):
        """wrong user input should lead to system exit"""
        m.some_guy.side_effect = isPalindrome

        with self.assertRaises(SystemExit) as cm:
            result = m.some_guy.side_effect('m123')

        self.assertEqual(cm.exception.code, 1)


unittest.main()