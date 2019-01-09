<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    Openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

## 753. Cracking the Safe (Hard)

<p>
There is a box protected by a password.  The password is <code>n</code> digits, where each letter can be one of the first <code>k</code> digits <code>0, 1, ..., k-1</code>.
</p><p>
You can keep inputting the password, the password will automatically be matched against the last <code>n</code> digits entered.
</p><p>
For example, assuming the password is <code>"345"</code>, I can open it when I type <code>"012345"</code>, but I enter a total of 6 digits.
</p><p>
Please return any string of minimum length that is guaranteed to open the box after the entire string is inputted.
</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> n = 1, k = 2
<b>Output:</b> "01"
<b>Note:</b> "10" will be accepted too.
</pre>
</p>

<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> n = 2, k = 2
<b>Output:</b> "00110"
<b>Note:</b> "01100", "10011", "11001" will be accepted too.
</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li><code>n</code> will be in the range <code>[1, 4]</code>.</li>
<li><code>k</code> will be in the range <code>[1, 10]</code>.</li>
<li><code>k^n</code> will be at most <code>4096</code>.</li>
</ol>
</p>

### Related Topics
  [[Depth-first Search](https://github.com/openset/leetcode/tree/master/tag/depth-first-search/README.md)]
  [[Math](https://github.com/openset/leetcode/tree/master/tag/math/README.md)]

### Hints
  1. We can think of this problem as the problem of finding an Euler path (a path visiting every edge exactly once) on the following graph: there are $$k^{n-1}$$ nodes with each node having $$k$$ edges.  It turns out this graph always has an Eulerian circuit (path starting where it ends.)

We should visit each node in "post-order" so as to not get stuck in the graph prematurely.