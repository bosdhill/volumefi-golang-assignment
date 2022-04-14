# volumefi-golang-assignment


My approach:

0. Made a copy of the document to annotate with my thoughts as a read through it and revisited it.

1. Walked through the examples to make sure I understood the assignment. Whiteboarded some ideas of what algorithm I would need to use.

2. Soon realized it was a topological sorting problem so googled around to review topological sort (found [this article](https://efficientcodeblog.wordpress.com/2017/11/28/topological-sort-dfs-bfs-and-dag/) particularly helpful)

3. Wrote a rough implementation of topological sort and ran it to make sure it passed the provided examples.

4. Wrote some unit tests to test edge cases, and so made a few passes over the sorting code to update it (a bit of TDD), mainly testing for cycles, loops, disconnectedness.

5. Wrapped up with parsing json input from stdin (since it wasn't specified, but made end to end testing easier).

6. Wrote a Makefile to build a binary and run the unit and end to end tests.