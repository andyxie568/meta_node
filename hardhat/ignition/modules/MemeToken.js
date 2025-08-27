const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("MemeTokenModule", (m) => {
    // 部署参数配置（根据需求调整）
    const name = "MemeToken";
    const symbol = "MTK";
    const totalSupply = 0.001;
    const initialOwner = m.getAccount(0); // 部署者地址自动获取
    const treasury = "0xdCee1d1D79461c7f87b1C19bfacce0e986038451"; // 替换为实际国库地址

    // 部署主合约
    const memeToken = m.contract("MemeToken", [
        name,
        symbol,
        totalSupply * (10 ** 18), // 转换为带精度单位
        initialOwner,
        treasury
    ]);

    // 部署后操作：设置初始流动性池（可选）
    m.call(memeToken, "setLiquidityPool", ["0xdCee1d1D79461c7f87b1C19bfacce0e986038451"], {
        id: "setLiquidityPool",
        after: [memeToken] // 确保在合约部署后执行
    });

    return { memeToken };
});