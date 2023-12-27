# GolangLeetcode
AC leetcode with goLang

`using go test -v `

### 1. Two Sum

### 2. Add Two Numbers

### 3. Longest Substring Without Repeating Characters
### 4. Median of Two Sorted Arrays
### 5. Longest Palindromic Substring
### 6. ZigZag Conversion
### 7. Reverse Integer
### 8. String to Integer (atoi)
### 9. Palindrome Number
### 10. Regular Expression Matching
### 11. Container With Most Water
1. 思路:
### 12. Integer to Roman
1. 思路: 做一个映射，简单转换十进制转换即可
### 13. Roman to Integer
1. 思路： 倒着遍历字符串，如果发现当前字符比前一个字符小，则减去当前字符，否则加上当前字符，最后将结果加上
### 14. Longest Common Prefix
1. 思路： 遍历第一个字符串的，将每一个字符和在后续的字符串中遍历，当后续的字符串和该字符相等的时候添加到公共字符串中，否则就返回失败
### 15. 3Sum
1. 将输入排序，然后遍历，将第一个元素作为起始位置，第二个元素作为起始位置+1，第三个元素作为末尾位置，如果三个元素和大于0，则第三个位置向中间移动，小于0 表示数目不够大，移动第二个元素，直到便利出所有等于0的可能值，当第二个元素超过第三个元素时表示便利结束，开始移动第一个位置
1. 注意： 第一个元素的移动要加入重复元素判断，提升运行效率
### 16. 3Sum Closest
1. ，将输入排序，三个位置开始搜索，第一个为起始位置，第二个为起始位置+1，第三个为末尾位置，如果三个位置的和大于目标值，则第三个位置向中间移动，
如果小于目标值，则第二个位置向中间移动，如果等于目标值，则直接返回结果
1. 注意： 需要设定一个初始值保证结果的正确性
### 17. Letter Combinations of a Phone Number
1. 思路： 模拟一个进位加法，超过含有的字母的数目就自动进位，检查是否所有的遍历完成了就结束，将满足条件的字符添加即可
1. 注意： 字符顺序可能和进位加法有冲突，下标的时候要注意
### 18. 4Sum
1. 思路：递归调用3SUM 的方法，
1. 注意： 处理好重复的数字和设定提前退出条件
### 19. Remove Nth Node From End of List
1. 思路： 取一个指针指向N 个之后的Node ，遍历链表，当前指针到达末尾的时候，去除掉当前的指针即可完成任务
1. 注意：删除第一个节点的情况和最后一个节点的情况，删除 第一个节点的情况是，N个之后个前指针是nil ，删除最后一个节点的情况是长度为1 时单独处理
### 20. Valid Parentheses
1. 思路： 建立一个待匹配的队列，当出现结束的字符的时候，于队列中的字符出队，直到队列为空，则匹配成功。
1. 注意： 单独出现结束字符和结束的时候待匹配队列中有字符，这两种情况都算匹配失败

