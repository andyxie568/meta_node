// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

/*
创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    // 存储候选人的得票数
    mapping(address => uint) public votes;
    // 存储所有候选人的地址
    address[] public addresses;
 
    function vote(address _candidate) external {
        votes[_candidate]++;
    }

    function getVotes(address _candidate) external view returns (uint) {
      return votes[_candidate];  
    }

    function resetVotes() external {
        for (uint i = 0; i < addresses.length; ++i) {
            votes[addresses[i]] = 0;  
        }
    }
}

/*
反转字符串 (Reverse String)
题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
*/
contract ReverseString {
     function reverse(string memory str) public pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        uint length = strBytes.length;
        bytes memory reversed = new bytes(length);

        for (uint i = 0; i < strBytes.length; i++) {
            reversed[i] = strBytes[length - 1 - i];
        }

        return string(reversed);
    }
}

/*
 用 solidity 实现整数转罗马数字
题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
*/
contract IntegerToRoman {
    function intToRoman(uint num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Input must be between 1 and 3999");
        
        // 定义可供匹配的整数值和对应的罗马数字符号
        uint[] memory values = new uint[](13);
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        
        values[0] = 1000; symbols[0] = "M";  // 1000 -> M
        values[1] = 900;  symbols[1] = "CM"; // 900 -> CM
        values[2] = 500;  symbols[2] = "D";  // 500 -> D
        values[3] = 400;  symbols[3] = "CD"; // 400 -> CD
        values[4] = 100;  symbols[4] = "C";  // 100 -> C
        values[5] = 90;   symbols[5] = "XC"; // 90 -> XC
        values[6] = 50;   symbols[6] = "L";  // 50 -> L
        values[7] = 40;   symbols[7] = "XL"; // 40 -> XL
        values[8] = 10;   symbols[8] = "X";  // 10 -> X
        values[9] = 9;    symbols[9] = "IX"; // 9 -> IX
        values[10] = 5;   symbols[10]= "V";  // 5 -> V
        values[11] = 4;   symbols[11]= "IV"; // 4 -> IV
        values[12] = 1;   symbols[12]= "I";  // 1 -> I
        
        string memory roman;
    
        for (uint i = 0; i < values.length; i++) {
            // 将当前的符号添加到结果中直到num小于该符号对应的值
            while (num >= values[i]) {
                roman = string(abi.encodePacked(roman, symbols[i]));
                num -= values[i];
            }
        }
        
        return roman;
    }
}

/*
用 solidity 实现罗马数字转数整数
题目描述在 https://leetcode.cn/problems/roman-to-integer/description/
*/
contract RomanToInt {
    mapping(bytes1 => uint16) romanValues;

    constructor() {
        // 初始化单个罗马字符映射到数值
        romanValues["I"] = 1;
        romanValues["V"] = 5;
        romanValues["X"] = 10;
        romanValues["L"] = 50;
        romanValues["C"] = 100;
        romanValues["D"] = 500;
        romanValues["M"] = 1000;
    }

    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory strBytes = bytes(s);
        uint256 sum = 0;
        
        for (uint i = 0; i < strBytes.length; i++) {
            uint16 currentValue = romanValues[strBytes[i]];
            uint16 nextValue;
            if (i < strBytes.length - 1) {
                nextValue = romanValues[strBytes[i+1]];
            }

            if (currentValue > nextValue) {
                sum += currentValue;
            } else {
                sum -= currentValue;
            }
        }
        
        return sum;
    }
}

/*
 合并两个有序数组 (Merge Sorted Array)
题目描述：将两个有序数组合并为一个有序数组。
*/
contract MergeSortedArrays {
    function mergeSortedArrays(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory mergedArray = new uint[](len1 + len2);

        uint i = 0;
        uint j = 0;
        uint k = 0;
        
        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                mergedArray[k] = arr1[i];
                i++;
            } else {
                mergedArray[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 把剩余的元素加到新数组后面
        while (i < len1) {
            mergedArray[k] = arr1[i];
            i++;
            k++;
        }

        while (j < len2) {
            mergedArray[k] = arr2[j];
            j++;
            k++;
        }

        return mergedArray;
    }
}

/*
二分查找 (Binary Search)
题目描述：在一个有序数组中查找目标值。
*/
contract BinarySearch {
    function binarySearch(uint[] memory sortedArray, uint target) public pure returns (int) {
        uint left = 0;
        uint right = sortedArray.length - 1;

        while (left <= right) {
            uint mid = left + (right - left) / 2;

            // 检查中间元素是否为目标值
            if (sortedArray[mid] == target) {
                return int(mid);
            }

            // 决定继续查找的半边
            if (sortedArray[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        // 如果未找到目标值，返回 -1
        return -1;
    }
}