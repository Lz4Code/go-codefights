/*
Challenge: Is isIPv4Address
https://codefights.com/arcade/intro/level-5/veW5xJednTy4qcjso/description

An IP address is a numerical label assigned to each device (e.g., computer, printer) participating in a computer network that uses the Internet Protocol for communication. There are two versions of the Internet protocol, and thus two versions of addresses. One of them is the IPv4 address.

IPv4 addresses are represented in dot-decimal notation, which consists of four decimal numbers, each ranging from 0 to 255, separated by dots, e.g., 172.16.254.1.

Given a string, find out if it satisfies the IPv4 address naming rules.

Example

For inputString = "172.16.254.1", the output should be
isIPv4Address(inputString) = true;

For inputString = "172.316.254.1", the output should be
isIPv4Address(inputString) = false.

316 is not in range [0, 255].

For inputString = ".254.255.0", the output should be
isIPv4Address(inputString) = false.

There is no first number.

Input/Output

[time limit] 4000ms (go)
[input] string inputString

A string consisting of digits, full stops and lowercase Latin letters.

Guaranteed constraints:
5 ≤ inputString.length ≤ 15.

[output] boolean

true if inputString satisfies the IPv4 address naming rules, false otherwise.
*/
import "strings"
func isIPv4Address(s string) bool {
   arr := strings.Split(s, ".")
   if len(arr) != 4 {
      return false
   }
   for _, v := range arr {
      i, err := strconv.Atoi(v)
      if err != nil { 
         return false
      }
      if i < 0 || i > 255 {
         return false
      }
   }
   return true
}
