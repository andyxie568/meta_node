// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


// 创建一个名为Voting的合约，包含以下功能：
// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
contract Vote {
    mapping (address => uint32) voteMap;
    address []candidateList;

    function vote(address candidate) public  {
        if (voteMap[candidate] == 0) {
            candidateList.push(candidate);
        }
        voteMap[candidate]++;
    }

    function getVotes(address candidate) public view returns(uint32) {
        return voteMap[candidate];
    }

    function resetVotes() public {
        for (uint i = 0; i < candidateList.length; i++) {
            voteMap[candidateList[i]] = 0;
        }
        delete candidateList;
    }
}
