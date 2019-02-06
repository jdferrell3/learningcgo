# Find the common elements given two sorted lists

Another challenge from Baltimore GoLang Meetup (https://gist.github.com/jboursiquot/a66237985bf7a30eae378cb17f17f953)

Write a function that returns a list of common elements between two sorted lists.

Example:

```go
commonElements([]int{1, 2, 3}, []int{2, 3, 4}) // should return []int{2, 3}
```

## Problem Set

```go
[]int{1, 3, 4, 6, 7, 9}
[]int{1, 2, 4, 5, 9, 10}
// should return []int{1, 4, 9}

[]int{1, 2, 9, 10, 11, 12}
[]int{0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15}
// should return []int{1, 2, 9, 10, 12}

[]int{0, 1, 2, 3, 4, 5}
[]int{6, 7, 8, 9, 10, 11}
// should return nil
```

## Questions

What is the time complexity of your solution and why?
