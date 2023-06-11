# Verifying Alien Dictionary

## Description

In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different  `order`. The  `order`  of the alphabet is some permutation of lowercase letters.

Given a sequence of  `words`  written in the alien language, and the  `order`  of the alphabet, return  `true`  if and only if the given  `words`  are sorted lexicographically in this alien language. 

The problem can be found and resolved [here](https://leetcode.com/problems/verifying-an-alien-dictionary/).

## Analysis
### Inputs and outputs
We're gonna have two entries, an array of strings that has the words and a string which is an alphabet that determines the order of the words.

As an output we are going to have a bool that is true if the words are in ascending order, according to the alphabet given or false otherwise. Lexicographically means that the characters have a hierarchy based on the alphabet, so that we can look at it in the dictionary.

### Intuition for logical attempt
For the first word we can do much it is just our starting word,
we need to remember the word and check the next word in the sequence.

Compare the word with the previous one so that we know if the order is correct.

For that we can compare the characters taking into account the following cases:

If the words have the same length, compare each character in the same index until they are different, check the index of the character in the alphabet: If the index in the alphabet for the previous word is bigger than the character for the current one return false.

If the words have different length, traverse the shortest word and compare each character in the same index until they are different, check the index of the character in the alphabet.

#### Summary
We have a k*O(n) where n is the number of words in the array and the k is a variable that accounts for the length of the words when comparing each character.
1. Two pointers: One points to the current word and other points the previous word.
2. Check the length of each word
3. Comparison of the characters
4. Position of the caracters in the alphabet (use the alphabet as a hash table).

#### Reflections
First I need to improve on the way i approach the problem
Desmitify break keyword
think of indexes
Prioritize eficiency over readability
The thought process I saw in more eficient solutions was
sort of backwards to my process. I think that is because
I thought more on the posible cases than the worst case.

Thinking on the worst case leeds to figure a solution for
the comparison of two letters from two words.

Then adjacent cases like what if one word is longer than the other

Then what words I need to compare.

# Merge Two Sorted Arrays

## Description
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

## Analysis
### inputs and outputs
We receive two integer arrays in non-decreasing order like 1,2,3..
We have to return an array with all the elements in non-decreasing order.
The result has to be the length of the first argument. We also have the length of each
valid array item. In the case of the first array the rest is ordered with zeros.
Example:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

### intuition
We need to compare each item so that we can determine where they have to be in the array. 

Since the arrays are ordered and the result has to be of length of the first array, we can doit in place of nums1. We can compare the last valid item of each array and the biggest of the two can be placed at the end of the first array.

## Summary
This is a O(n+m) algorithm there m and n are the number of items in the array.

## Reflection
In the go implementation I made the following mistakes

1. Used a for clause when I could've use the condition, like a while loop.
2. I didn't use the n as a pointer and condition.

In the python implementation we almost got the sense for the pointer but miss the second consideration in the above list.

# Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

## Analysis

### inputs and outputs
We are given an array of size n where its items represent a height. The idea is getting the biggest area taking into account that the water can't be spilled.
Example:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

### intuition

We need to controll for the distance between the lines and the height, we can start by having one pointer in the beggining of the array and another at the end, store the current covered area.

We have to move the pointer with less height so that we can obtain a bigger area, if the next line is higher we should take the area and if is bigger than the current biggest area replace it. Then change compare again until the two pointers meet.

Let's write it in steps

Locate pointer at the beginning and the end of the array

Take the current area by multiplying the distance of the two index by the pointer with the shortest length. set the biggest area to the number if is bigger than the current area.

Move the pointer with the lowest or equall height.

## summary
The container with most water now seems pretty easy, I could've use the built in min max functions from go tho. That made my solution a little bit slower but I feel like i learned a ton with this.


# Trapping Rain Water

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

## Analysis

### Inputs and outputs
We're given an array of integers and we need to return an integer that saids how much water is trapped
for example:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

### Intuition

First i need to identify when the water is trappped based on the height of the arrays
Then I need to have some sort of acumulator that sums every time the water is trapped
It seems like with two pointers, i can determine the position where the posible trapping begins
and with the second when it ends.

## Summary
Bad intuition, this was one of the cases where there were two ways of going arround the problem

Check the contribution from the left and check the contribution from the right

Or check if theres some sort of criteria to move left or right, in this case we could've thouhgt 
what ensures the water is going to be trapped, for example if the value in the left is bigger than
the value in the right if i find one step higher in the right the water would be trap

the second thought is not necessarily correct! this was a mess cause i had to see the solution. It was a hard problem tho. the second is actually correct!

## Reflections

the most important thing i can take from this is that inside the 2 pointers patterns I identify 3

A criteria to move one pointer in respect to the other

A pointer that references a position that we want to alter an another that traverses the array that creates a condition for the modification.

A pointer that checks the previous or next value and another that checks the current value.

# Move Zeros

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

## Analysis
### inputs and outputs
As an input we get an array of ints with zeros between them and as an output is got to be the same array. The modification has to be in place.

### Intuition
After a visualization analysis:
Pointer1 has the responsability of holding the position of the left most zero.
Pointer2 has the responsability of traversing the array.
when p2 encounters a number different to zero and p1 has position they interchange values and p1 increases
if p2 encounters a zero and p1 doesn't have a position, p1 = p2, y p1 has a position nothing happens.

