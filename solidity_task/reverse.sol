// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 反转一个字符串。输入 "abcde"，输出 "edcba"
contract Reverse {
    function reverseString(string memory input) public  pure returns (string memory) {
        bytes memory inputBytes = bytes(input);
        uint len = inputBytes.length;
        bytes memory result = new bytes(len);
        for (uint i = 0; i < len; i++) {
            result[i] = inputBytes[len - i - 1];
        }

        return string(result);
    }
}