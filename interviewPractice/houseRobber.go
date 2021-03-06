/*
Challenge: houseRobber
https://codefights.com/interview-practice/task/XBWN6DYRqPrKdMZs8/description

You are planning to rob houses on a specific street, and you know that every house on the street has a certain amount of money hidden. The only thing stopping you from robbing all of them in one night is that adjacent houses on the street have a connected security system. The system will automatically trigger an alarm if two adjacent houses are broken into on the same night.

Given a list of non-negative integers nums representing the amount of money hidden in each house, determine the maximum amount of money you can rob in one night without triggering an alarm.

Example

For nums = [1, 1, 1], the output should be
houseRobber(nums) = 2.

The optimal way to get the most money in one night is to rob the first and the third houses for a total of 2.

Input/Output

[time limit] 4000ms (go)
[input] array.integer nums

An array representing the amount of money that each house on the street has.

Guaranteed constraints:
0 ≤ nums.length ≤ 100,
0 ≤ nums[i] ≤ 500.

[output] integer
*/
func houseRobber(nums []int) int {
    
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if nums[1] < nums[0] {
        nums[1] = nums[0]
    }
    
    for i := 2; i < len(nums); i++ {
        if nums[i-2] + nums[i] > nums[i-1] {
            nums[i] = nums[i-2]+nums[i]
        } else {
            nums[i] = nums[i-1]
        }
    }
        
    return nums[len(nums)-1]
}
