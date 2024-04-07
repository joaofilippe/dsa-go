# Bubble sort

* Simple Sort Algorithm;
* Element by element;
* Compare `v[i]` with `v[i+1]`;
* Swap if necessary;

![an example](https://upload.wikimedia.org/wikipedia/commons/c/c8/Bubble-sort-example-300px.gif)

## Performance

* Worst-case performance:
  * O(n<sup>2</sup>) compariosons;
  * O(n<sup>2</sup>) swaps;
* Best-case performance:
  * O(n) comparisons;
  * O(1) swaps;
* Average performance:
  * O(n<sup>2</sup>) compariosons;
  * O(n<sup>2</sup>) swaps;
* Worst-case space complexity:
  * O(n) total;
  * O(1) auxiliary;

## Example

Take an array of numbers `"5 1 4 2 8"`, and sort the array from lowest number to greatest number using bubble sort. 

In each step, elements written in bold are being compared. Three passes will be required;

**First Pass:**

( 5 1 4 2 8 ) → ( 1 5 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.

( 1 5 4 2 8 ) → ( 1 4 5 2 8 ), Swap since 5 > 4

( 1 4 5 2 8 ) → ( 1 4 2 5 8 ), Swap since 5 > 2

( 1 4 2 5 8 ) → ( 1 4 2 5 8 ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.

**Second Pass:**

( 1 4 2 5 8 ) → ( 1 4 2 5 8 )

( 1 4 2 5 8 ) → ( 1 2 4 5 8 ), Swap since 4 > 2

( 1 2 4 5 8 ) → ( 1 2 4 5 8 )

( 1 2 4 5 8 ) → ( 1 2 4 5 8 )

Now, the array is already sorted, but the algorithm *does not know if it is completed*.

The algorithm needs one additional whole pass without any swap to know it is sorted.

**Third Pass:**
( 1 2 4 5 8 ) → ( 1 2 4 5 8 )

( 1 2 4 5 8 ) → ( 1 2 4 5 8 )

( 1 2 4 5 8 ) → ( 1 2 4 5 8 )

( 1 2 4 5 8 ) → ( 1 2 4 5 8 )