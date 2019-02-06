# Most frequently occurring element in a list

Another challenge from Baltimore GoLang Meetup (https://gist.github.com/jboursiquot/c3b8bec7dcf7589e2107dd7d72eb22e0)

Find the most frequently occurring element in a given list.

For example, the most frequent element in `[1, 2, 3, 1]` is `1`.

Assume a single unique element that appears more than once. In other words, you won't be given something like `[1, 1, 2, 2]`.

## Problem Set

```go
[]int{1, 3, 1, 3, 2, 1}  // should return 1
[]int{3, 3, 1, 3, 2, 1} // 3
[]int{0} // 0
[]int{0, -1, 10, 10, -1, 10, -1, -1, -1, 1} // -1
```

## Questions

What is the time complexity of your solution and why?