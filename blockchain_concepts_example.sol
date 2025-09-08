// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * 区块链核心概念示例合约
 * 演示区块、交易、Merkle Tree、共识机制和Gas费用的实际应用
 */
contract BlockchainConceptsExample {
    
    // 事件：用于记录重要操作，消耗较少Gas
    event TransactionExecuted(address indexed from, address indexed to, uint256 amount, uint256 gasUsed);
    event MerkleRootUpdated(bytes32 newRoot);
    
    // 存储结构：模拟区块头信息
    struct BlockHeader {
        bytes32 previousHash;    // 前一个区块的哈希
        bytes32 merkleRoot;      // Merkle根
        uint256 timestamp;       // 时间戳
        uint256 difficulty;      // 难度目标
        uint256 nonce;          // 随机数
        uint256 blockNumber;    // 区块号
    }
    
    // 交易结构
    struct Transaction {
        address from;
        address to;
        uint256 amount;
        uint256 nonce;
        bytes signature;
        uint256 gasPrice;
        uint256 gasLimit;
    }
    
    // 状态变量
    mapping(address => uint256) public balances;
    mapping(bytes32 => bool) public processedTransactions;
    
    BlockHeader public currentBlock;
    uint256 public totalTransactions;
    
    // 构造函数：初始化创世区块
    constructor() {
        currentBlock = BlockHeader({
            previousHash: bytes32(0),  // 创世区块没有父区块
            merkleRoot: bytes32(0),
            timestamp: block.timestamp,
            difficulty: 1000000,       // 初始难度
            nonce: 0,
            blockNumber: 0
        });
    }
    
    /**
     * 执行交易 - 演示Gas费用计算
     * @param to 接收方地址
     * @param amount 转账金额
     */
    function executeTransaction(address to, uint256 amount) external {
        uint256 gasStart = gasleft(); // 记录开始时的Gas
        
        require(balances[msg.sender] >= amount, "Insufficient balance");
        require(to != address(0), "Invalid recipient");
        
        // 执行转账
        balances[msg.sender] -= amount;
        balances[to] += amount;
        
        // 更新交易计数
        totalTransactions++;
        
        uint256 gasUsed = gasStart - gasleft(); // 计算实际使用的Gas
        
        // 发出事件（比存储更省Gas）
        emit TransactionExecuted(msg.sender, to, amount, gasUsed);
    }
    
    /**
     * 模拟Merkle根计算
     * 在实际区块链中，这由矿工完成
     */
    function calculateMerkleRoot(bytes32[] memory transactionHashes) external pure returns (bytes32) {
        require(transactionHashes.length > 0, "No transactions");
        
        if (transactionHashes.length == 1) {
            return transactionHashes[0];
        }
        
        // 简化的Merkle树计算
        bytes32[] memory currentLevel = transactionHashes;
        
        while (currentLevel.length > 1) {
            bytes32[] memory nextLevel = new bytes32[]((currentLevel.length + 1) / 2);
            
            for (uint256 i = 0; i < currentLevel.length; i += 2) {
                if (i + 1 < currentLevel.length) {
                    // 两个子节点都存在
                    nextLevel[i / 2] = keccak256(abi.encodePacked(currentLevel[i], currentLevel[i + 1]));
                } else {
                    // 只有一个子节点，复制它
                    nextLevel[i / 2] = currentLevel[i];
                }
            }
            
            currentLevel = nextLevel;
        }
        
        return currentLevel[0];
    }
    
    /**
     * 模拟PoW挖矿过程
     * 在实际区块链中，这需要大量计算
     */
    function mineBlock(bytes32 newMerkleRoot, uint256 difficulty) external {
        require(difficulty > 0, "Invalid difficulty");
        
        // 模拟寻找有效Nonce的过程
        uint256 nonce = 0;
        bytes32 blockHash;
        
        do {
            blockHash = keccak256(abi.encodePacked(
                currentBlock.previousHash,
                newMerkleRoot,
                block.timestamp,
                difficulty,
                nonce
            ));
            nonce++;
        } while (uint256(blockHash) >= (2**256 - 1) / difficulty && nonce < 1000); // 限制循环次数避免Gas耗尽
        
        // 更新区块头
        currentBlock = BlockHeader({
            previousHash: keccak256(abi.encodePacked(currentBlock)),
            merkleRoot: newMerkleRoot,
            timestamp: block.timestamp,
            difficulty: difficulty,
            nonce: nonce,
            blockNumber: currentBlock.blockNumber + 1
        });
        
        emit MerkleRootUpdated(newMerkleRoot);
    }
    
    /**
     * 模拟PoS验证者选择
     * 根据质押金额和持币时间选择验证者
     */
    function selectValidator(address[] memory validators, uint256[] memory stakes, uint256[] memory stakingTime) 
        external pure returns (address) {
        require(validators.length == stakes.length && stakes.length == stakingTime.length, "Array length mismatch");
        
        uint256 totalWeight = 0;
        uint256[] memory weights = new uint256[](validators.length);
        
        // 计算每个验证者的权重
        for (uint256 i = 0; i < validators.length; i++) {
            weights[i] = stakes[i] * stakingTime[i]; // 质押金额 × 持币时间
            totalWeight += weights[i];
        }
        
        // 简化的随机选择（实际应用中需要更复杂的随机性）
        uint256 randomValue = uint256(keccak256(abi.encodePacked(block.timestamp, block.difficulty))) % totalWeight;
        
        uint256 cumulativeWeight = 0;
        for (uint256 i = 0; i < validators.length; i++) {
            cumulativeWeight += weights[i];
            if (randomValue < cumulativeWeight) {
                return validators[i];
            }
        }
        
        return validators[validators.length - 1]; // 默认返回最后一个
    }
    
    /**
     * 获取当前区块信息
     */
    function getCurrentBlock() external view returns (BlockHeader memory) {
        return currentBlock;
    }
    
    /**
     * 获取账户余额
     */
    function getBalance(address account) external view returns (uint256) {
        return balances[account];
    }
    
    /**
     * 存款函数（用于测试）
     */
    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }
    
    /**
     * 估算Gas费用
     */
    function estimateGasCost(address to, uint256 amount) external view returns (uint256) {
        // 基础Gas费用
        uint256 baseGas = 21000; // 基础转账费用
        
        // 根据操作复杂度添加额外费用
        if (to == address(this)) {
            baseGas += 20000; // 合约调用额外费用
        }
        
        return baseGas * tx.gasprice;
    }
}
