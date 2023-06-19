# Sliding Windows

Theory -> https://platzi.com/clases/4592-algoritmos-avanzados-arrays/56153-ejercicios-recomendados-de-ventana-deslizante/

## Max Consecutive Ones
[LeetCode Link](https://leetcode.com/problems/max-consecutive-ones/)
### Analysis
We are given an array of zeros and ones, we need to find the max count of
consequitive ones and return it as a number.

### Intuition
We know that we need to apply sliding window for this exercise, so the idea
would be to have a pointer that traverses the array while we have other pointer
marking the start of a one sequence.

so we start by checking 
if the current position is a one, then we need to check if the pointer marking the start is a one or
a zero, if is a zero...

[My solution](https://leetcode.com/problems/max-consecutive-ones/submissions/973656933/)

The easy way was using an acumulator for the ones found, and reseting thecounter if it finds a zero.

## Longest substring without repeating characters
[LeetCode Link](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
[My solution](https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/967786483/)

## Longest repeating character replacement (Pending go solution)
[LeetCode Link](https://leetcode.com/problems/longest-repeating-character-replacement/)

## Find all anagrams in a string
[LeetCode Link](https://leetcode.com/problems/find-all-anagrams-in-a-string/)
[[My solution|Python]](https://leetcode.com/problems/find-all-anagrams-in-a-string/submissions/975040558/)
[My Solution](https://leetcode.com/problems/find-all-anagrams-in-a-string/submissions/975082040/)

## Best Time to Buy And Sell Stock
[LeetCode Link](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

## Permutation In String
[LeetCode Link](https://leetcode.com/problems/permutation-in-string/)

## Minimum Window Substring
[LeetCode Link](https://leetcode.com/problems/minimum-window-substring/)

## Sliding Window Maximum
[LeetCode Link](https://leetcode.com/problems/sliding-window-maximum/)
