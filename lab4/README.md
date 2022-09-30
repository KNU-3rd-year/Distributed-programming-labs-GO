# Lab 4
**Topic**: Read-write lock

## Part 2:
### Task:
**Topic:**
Implement using Golang

**Task:**
A question about the garden. Create a multithreaded application that works with a generic two-dimensional array. To protect operations with a common array, use a read-write lock. A two-dimensional array describes a garden. The following threads should work in the application:
1) stream-gardener watches over the garden and waters the wilted plants;
2) flow-nature can arbitrarily change the state of plants;
3) stream-monitor1 periodically outputs the state of the garden to a file (without erasing the previous state);
4) flow-monitor2 displays the state of the garden on the screen.

### Solution:
[Lab4_part2.go](Lab4_part2.go)