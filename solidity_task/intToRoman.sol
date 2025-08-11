// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

// 整数转罗马数字
contract IntToRoman {
    struct RomanPair {
        uint256 value;
        string roman;
    }


   RomanPair[] romanPairList;
    constructor() {
        romanPairList.push(RomanPair(1000, "M"));
        romanPairList.push(RomanPair(900, "CM"));
        romanPairList.push(RomanPair(500, "D"));
        romanPairList.push(RomanPair(400, "CD"));
        romanPairList.push(RomanPair(100, "C"));
        romanPairList.push(RomanPair(90, "XC"));
        romanPairList.push(RomanPair(50, "L"));
        romanPairList.push(RomanPair(40, "XL"));
        romanPairList.push(RomanPair(10, "X"));
        romanPairList.push(RomanPair(9, "IX"));
        romanPairList.push(RomanPair(5, "V"));
        romanPairList.push(RomanPair(4, "IV"));
        romanPairList.push(RomanPair(1, "I"));
    }

    function intToRoman(uint256 num) public view returns (string memory) {
        require(num > 0, "Number must be positive");

        string memory result = "";

        for (uint256 i = 0; i < romanPairList.length; i++) {
            RomanPair memory pair = romanPairList[i];
            while (num >= pair.value) {
                result = string(abi.encodePacked(result, pair.roman));
                num -= pair.value;
            }
        }
        return result;
    }
}