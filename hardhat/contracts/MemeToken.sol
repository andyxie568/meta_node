// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MemeToken is ERC20, Ownable {
    // 税率配置（百分比，以10000为基数，如500=5%）
    uint256 public buyTax = 500; // 买入税5%
    uint256 public sellTax = 1000; // 卖出税10%
    uint256 public constant TAX_DENOMINATOR = 10000;

    // 流动性池地址（如Uniswap V2 Pair）
    address public liquidityPool;
    // 国库地址（接收部分税费）
    address public treasury;

    // 交易限制：单笔最大交易额（占总量比例，100=1%）
    uint256 public maxTransactionPercent = 100; // 1%
    // 每日最大交易次数
    uint256 public dailyMaxTransactions = 10;
    // 记录地址每日交易次数
    mapping(address => uint256) public dailyTransactions;
    // 记录地址最后交易日期（防止跨天计数）
    mapping(address => uint256) public lastTransactionDay;

    event TaxUpdated(uint256 newBuyTax, uint256 newSellTax);
    event LiquidityPoolSet(address indexed pool);
    event TransactionLimitUpdated(uint256 maxPercent, uint256 dailyMax);

    constructor(
        string memory name,
        string memory symbol,
        uint256 totalSupply,
        address initialOwner,
        address _treasury
    ) ERC20(name, symbol) Ownable(initialOwner) {
        _mint(initialOwner, totalSupply * 10 ** decimals());
        treasury = _treasury;
    }

    // 设置流动性池地址（仅Owner）
    function setLiquidityPool(address _pool) external onlyOwner {
        liquidityPool = _pool;
        emit LiquidityPoolSet(_pool);
    }

    // 更新税率（仅Owner）
    function updateTaxes(uint256 _buyTax, uint256 _sellTax) external onlyOwner {
        require(_buyTax <= 2000 && _sellTax <= 3000, "Tax too high"); // 限制最高税率
        buyTax = _buyTax;
        sellTax = _sellTax;
        emit TaxUpdated(_buyTax, _sellTax);
    }

    // 更新交易限制（仅Owner）
    function updateTransactionLimits(uint256 _maxPercent, uint256 _dailyMax) external onlyOwner {
        require(_maxPercent <= 500, "Max percent too high"); // 最高5%
        maxTransactionPercent = _maxPercent;
        dailyMaxTransactions = _dailyMax;
        emit TransactionLimitUpdated(_maxPercent, _dailyMax);
    }

    // 检查交易限制
    function _checkTransactionLimits(address sender, uint256 amount) internal {
        // 单笔额度限制：不超过总量的maxTransactionPercent
        uint256 maxAmount = (totalSupply() * maxTransactionPercent) / TAX_DENOMINATOR;
        require(amount <= maxAmount, "Exceeds max transaction amount");

        // 每日交易次数限制
        uint256 currentDay = block.timestamp / 1 days;
        if (lastTransactionDay[sender] != currentDay) {
            dailyTransactions[sender] = 0;
            lastTransactionDay[sender] = currentDay;
        }
        require(dailyTransactions[sender] < dailyMaxTransactions, "Exceeds daily transaction limit");
        dailyTransactions[sender]++;
    }

    // 计算税费并分配
    function _calculateTaxes(address sender, address recipient, uint256 amount) internal view returns (uint256 tax) {
        // 区分买入/卖出：向流动性池卖出时征收卖出税，其他情况按买入税
        bool isSell = recipient == liquidityPool;
        bool isBuy = sender == liquidityPool;

        if (isSell) {
            tax = (amount * sellTax) / TAX_DENOMINATOR;
        } else if (isBuy) {
            tax = (amount * buyTax) / TAX_DENOMINATOR;
        } else {
            tax = 0; // 钱包间转账免税（可根据需求调整）
        }
    }

    // 分配税费（流向国库和流动性池）
    function _distributeTaxes(uint256 taxAmount, address sender) internal {
        if (taxAmount == 0) return;

        // 50%税费进入国库，50%注入流动性池
        uint256 toTreasury = (taxAmount * 5000) / TAX_DENOMINATOR;
        uint256 toLiquidity = taxAmount - toTreasury;

        if (toTreasury > 0 && treasury != address(0)) {
            _transfer(sender, treasury, toTreasury);
        }
        if (toLiquidity > 0 && liquidityPool != address(0)) {
            _transfer(sender, liquidityPool, toLiquidity);
        }
    }

    // 重写transfer函数，加入税费和限制逻辑
    function transfer(address recipient, uint256 amount) public override returns (bool) {
        _checkTransactionLimits(_msgSender(), amount);

        uint256 tax = _calculateTaxes(_msgSender(), recipient, amount);
        uint256 transferAmount = amount - tax;

        _transfer(_msgSender(), recipient, transferAmount);
        if (tax > 0) {
            _distributeTaxes(tax, _msgSender());
        }
        return true;
    }
}