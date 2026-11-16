Suppose:
nums := []int{1, 2, 3}


Initially, your slice looks conceptually like this:
Initial slice
1

2

3


Length: 3

Capacity: at least 3


Now execute:
append(nums, nums...)


What does this mean?
- nums is the original slice.
- nums... expands the slice into individual elements.
- Those elements are appended to the original slice.
Conceptually, it is similar to:
append(nums, 1, 2, 3)


The resulting slice is:
Resulting slice
1

2

3

1

2

3


Original elements

Appended elements

Length = 6

2. What happens internally when append executes?
This is the important part.

A Go slice has three components:
1. Pointer: Points to the underlying array.
2. Length: Number of elements in the slice.
3. Capacity: Maximum number of elements available from the slice's starting position before reaching the end of its underlying array.
Let's consider two scenarios.
Scenario A: Sufficient capacity
Suppose:
nums := make([]int, 3, 6)

nums[0] = 1
nums[1] = 2
nums[2] = 3


Here:
- Length = 3
- Capacity = 6
The underlying array has enough room for another three elements.

When you execute:
result := append(nums, nums...)


Go reuses the existing underlying array.
Underlying array (capacity = 6)
1


2


3


1


2


3



3 existing elements + 3 appended elements. No new underlying array needed.

No allocation of a new underlying array is required.
Scenario B: Insufficient capacity
Now suppose:
nums := []int{1, 2, 3}


Assume its capacity is 3.

When you execute:
result := append(nums, nums...)


There isn't enough capacity to accommodate three additional elements.

Go must:
1. Allocate a larger underlying array.
2. Copy the existing elements into it.
3. Append the additional elements.
4. Return a new slice header pointing to the new underlying array.
Before append
1

2

3


Old underlying array

After allocation and copying
1


2


3


1


2


3



New underlying array containing all six elements

The exact new capacity is determined by Go's runtime allocation and growth rules. Don't assume it will always double.
3. Now calculate time complexity
Let n be the length of the original slice.

Your operation is:
append(nums, nums...)


We have two possible scenarios:
Operation	Time complexity
Enough capacity	O(n)
Insufficient capacity	O(n)

Why is the first scenario O(n) even without allocating a new array?

Because Go still needs to copy the n elements from nums into the available positions.

In the second scenario, it needs to copy the original n elements and append another n elements.

Therefore:
\[
T(n) = O(n+n) = O(2n) = O(n)
\]

Time complexity: O(n)
4. Space complexity
Your resulting slice contains 2n elements.

If a new underlying array is allocated, it needs space for those elements.
\[
S(n) = O(2n) = O(n)
\]

Auxiliary space: O(n)
One subtle distinction:
- If capacity is sufficient, no additional underlying array allocation is necessary.
- If capacity is insufficient, a new underlying array is allocated.
- Regardless, the resulting slice has 2n elements.
For algorithm analysis, we generally report O(n) auxiliary space for this operation when considering the additional storage required for the result.
5. Can we make this more efficient?
You can explicitly allocate the required capacity:
func getConcatenation(nums []int) []int {
    n := len(nums)

    result := make([]int, 0, 2*n)

    result = append(result, nums...)
    result = append(result, nums...)

    return result
}


This avoids repeated growth allocations. But the time and space complexities remain O(n).

For this particular problem, your original solution is already perfectly reasonable:
func getConcatenation(nums []int) []int {
    return append(nums, nums...)
}


One important interview detail: This solution may reuse and modify the original slice's underlying array. If you want to preserve the input, allocate a new slice as shown above.