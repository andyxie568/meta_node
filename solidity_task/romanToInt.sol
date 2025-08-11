// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

// 罗马数字转整数
contract RomanToInt {
    mapping(bytes1 => uint256) private romanValueMap;

    constructor() {
        romanValueMap["I"] = 1;
        romanValueMap["V"] = 5;
        romanValueMap["X"] = 10;
        romanValueMap["L"] = 50;
        romanValueMap["C"] = 100;
        romanValueMap["D"] = 500;
        romanValueMap["M"] = 1000;
    }

    function romanToInt(string memory s) public view returns(uint256) {
        bytes memory strBytes = bytes(s);
        uint256 total = 0;
        uint256 len = strBytes.length;

        for (uint256 i = 0; i < len; i++) {
            uint256 curValue = romanValueMap[strBytes[i]];
            if (i < len - 1 && curValue < romanValueMap[strBytes[i + 1]]) {
                total -= curValue;
            } else {
                total += curValue;
            }
        }
        return total;
    }
}