# EasySwap 合约交互时序图

## 1. 创建 List 订单时序图

```
用户    OrderBook    Vault    NFT合约
 │         │         │         │
 │         │         │         │
 │ 1. makeOrders([listOrder])  │
 ├────────►│         │         │
 │         │         │         │
 │         │ 2. 验证订单       │
 │         │         │         │
 │         │ 3. depositNFT(orderKey, maker, collection, tokenId)
 │         ├────────►│         │
 │         │         │         │
 │         │         │ 4. safeTransferFrom(maker, vault, tokenId)
 │         │         ├────────►│
 │         │         │         │
 │         │         │ 5. 返回成功
 │         │         │◄────────┤
 │         │         │         │
 │         │ 6. _addOrder(order)
 │         │         │         │
 │         │ 7. LogMake 事件   │
 │         │         │         │
 │ 8. 返回订单Key    │         │
 │◄────────┤         │         │
 │         │         │         │
```

## 2. 创建 Bid 订单时序图

```
用户    OrderBook    Vault
 │         │         │
 │         │         │
 │ 1. makeOrders([bidOrder]) + ETH
 ├────────►│         │
 │         │         │
 │         │ 2. 验证订单       │
 │         │         │
 │         │ 3. depositETH{value: price}(orderKey, price)
 │         ├────────►│
 │         │         │
 │         │ 4. ETHBalance[orderKey] += price
 │         │         │
 │         │ 5. _addOrder(order)
 │         │         │
 │         │ 6. LogMake 事件   │
 │         │         │
 │ 7. 返回订单Key    │
 │◄────────┤         │
 │         │         │
```

## 3. 订单匹配时序图

```
买家    OrderBook    Vault    卖家
 │         │         │        │
 │         │         │        │
 │ 1. matchOrders([matchDetail])
 ├────────►│         │        │
 │         │         │        │
 │         │ 2. 验证订单匹配条件
 │         │         │        │
 │         │ 3. withdrawETH(buyOrderKey, price, orderBook)
 │         ├────────►│        │
 │         │         │        │
 │         │ 4. 计算协议费用   │
 │         │         │        │
 │         │ 5. 转移 ETH 给卖家
 │         ├─────────┼────────►│
 │         │         │        │
 │         │ 6. withdrawNFT(sellOrderKey, buyer, collection, tokenId)
 │         ├────────►│        │
 │         │         │        │
 │         │ 7. 转移 NFT 给买家
 │         ├─────────┼────────►│
 │         │         │        │
 │         │ 8. 更新订单状态   │
 │         │         │        │
 │         │ 9. LogMatch 事件 │
 │         │         │        │
 │ 10. 返回成功      │        │
 │◄────────┤         │        │
 │         │         │        │
```

## 4. 订单取消时序图

```
用户    OrderBook    Vault
 │         │         │
 │         │         │
 │ 1. cancelOrders([orderKeys])
 ├────────►│         │
 │         │         │
 │         │ 2. 验证订单所有权
 │         │         │
 │         │ 3. 根据订单类型提取资产
 │         │         │
 │         │ 3a. 如果是 List 订单：
 │         │     withdrawNFT(orderKey, maker, collection, tokenId)
 │         ├────────►│
 │         │         │
 │         │ 3b. 如果是 Bid 订单：
 │         │     withdrawETH(orderKey, amount, maker)
 │         ├────────►│
 │         │         │
 │         │ 4. 资产返还给用户
 │         ├─────────┼────────►│
 │         │         │        │
 │         │ 5. 从订单簿移除订单
 │         │         │        │
 │         │ 6. LogCancel 事件
 │         │         │        │
 │ 7. 返回成功       │        │
 │◄────────┤         │        │
 │         │         │        │
```

## 5. 订单编辑时序图

```
用户    OrderBook    Vault
 │         │         │
 │         │         │
 │ 1. editOrders([editDetails])
 ├────────►│         │
 │         │         │
 │         │ 2. 验证订单可编辑性
 │         │         │
 │         │ 3. 取消旧订单
 │         │ 3a. _removeOrder(oldOrder)
 │         │ 3b. LogCancel 事件
 │         │         │
 │         │ 4. 处理资产差异
 │         │ 4a. 如果是 List 订单：
 │         │     editNFT(oldOrderKey, newOrderKey)
 │         ├────────►│
 │         │         │
 │         │ 4b. 如果是 Bid 订单：
 │         │     editETH(oldOrderKey, newOrderKey, oldAmount, newAmount, maker)
 │         ├────────►│
 │         │         │
 │         │ 5. 创建新订单
 │         │ 5a. _addOrder(newOrder)
 │         │ 5b. LogMake 事件
 │         │         │
 │ 6. 返回新订单Key  │
 │◄────────┤         │
 │         │         │
```

## 6. 批量匹配时序图

```
用户    OrderBook    Vault
 │         │         │
 │         │         │
 │ 1. matchOrders([matchDetails])
 ├────────►│         │
 │         │         │
 │         │ 2. 遍历匹配详情
 │         │         │
 │         │ 3. 对每个匹配详情：
 │         │ 3a. delegatecall matchOrderWithoutPayback
 │         │ 3b. 验证订单匹配条件
 │         │ 3c. 提取和转移资产
 │         │ 3d. 更新订单状态
 │         │ 3e. 记录操作结果
 │         │         │
 │         │ 4. 处理失败的操作
 │         │ 4a. 记录错误信息
 │         │ 4b. 发出错误事件
 │         │         │
 │         │ 5. 返还多余的 ETH
 │         ├─────────┼────────►│
 │         │         │        │
 │ 6. 返回操作结果   │        │
 │◄────────┤         │        │
 │         │         │        │
```

## 7. 资产托管时序图

```
OrderBook    Vault    NFT合约
 │           │         │
 │           │         │
 │ 1. depositNFT(orderKey, maker, collection, tokenId)
 ├──────────►│         │
 │           │         │
 │           │ 2. safeTransferFrom(maker, vault, tokenId)
 │           ├────────►│
 │           │         │
 │           │ 3. NFTBalance[orderKey] = tokenId
 │           │         │
 │           │ 4. 返回成功
 │           │◄────────┤
 │           │         │
 │ 5. 订单创建成功    │
 │◄──────────┤         │
 │           │         │
```

## 8. 费用计算时序图

```
OrderBook    Vault    卖家    协议
 │           │        │      │
 │           │        │      │
 │ 1. 订单匹配成功     │      │
 │           │        │      │
 │ 2. 计算协议费用     │      │
 │ 2a. protocolFee = (fillPrice * protocolShare) / TOTAL_SHARE
 │ 2b. sellerAmount = fillPrice - protocolFee
 │           │        │      │
 │ 3. 分配资金        │      │
 │ 3a. 转移给卖家：sellerAmount
 │ ├─────────┼────────►│      │
 │           │        │      │
 │ 3b. 协议费用累积    │      │
 │           │        │      │
 │ 4. 更新状态        │      │
 │           │        │      │
```

## 9. 错误处理时序图

```
用户    OrderBook    Vault
 │         │         │
 │         │         │
 │ 1. 发起交易       │
 ├────────►│         │
 │         │         │
 │         │ 2. 验证失败
 │         │ 2a. 参数无效
 │         │ 2b. 权限不足
 │         │ 2c. 资产不足
 │         │         │
 │         │ 3. 回滚交易
 │         │ 3a. 恢复状态
 │         │ 3b. 返还资产
 │         │ 3c. 发出错误事件
 │         │         │
 │ 4. 返回错误信息   │
 │◄────────┤         │
 │         │         │
```

## 10. 事件监听时序图

```
前端    OrderBook    区块链
 │         │         │
 │         │         │
 │ 1. 监听合约事件   │
 ├────────►│         │
 │         │         │
 │         │ 2. 发出事件
 │         ├────────►│
 │         │         │
 │ 3. 接收事件       │
 │◄────────┼─────────┤
 │         │         │
 │ 4. 解析事件数据   │
 │ 4a. 提取订单信息
 │ 4b. 提取交易信息
 │ 4c. 提取价格信息
 │         │         │
 │ 5. 更新前端状态   │
 │ 5a. 更新订单列表
 │ 5b. 更新用户余额
 │ 5c. 更新交易历史
 │         │         │
 │ 6. 通知用户      │
 │         │         │
```

## 11. 升级时序图

```
管理员   代理合约    实现合约
 │         │         │
 │         │         │
 │ 1. 部署新实现合约 │
 │         │         │
 │ 2. 调用 upgradeTo(newImplementation)
 ├────────►│         │
 │         │         │
 │         │ 3. 更新实现地址
 │         ├────────►│
 │         │         │
 │         │ 4. 返回成功
 │         │◄────────┤
 │         │         │
 │ 5. 升级完成      │
 │◄────────┤         │
 │         │         │
```

## 12. 暂停/恢复时序图

```
管理员   OrderBook
 │         │
 │         │
 │ 1. pause()
 ├────────►│
 │         │
 │         │ 2. 设置暂停状态
 │         │
 │         │ 3. 发出 Paused 事件
 │         │
 │ 4. 暂停完成      │
 │◄────────┤
 │         │
 │         │
 │ 5. unpause()
 ├────────►│
 │         │
 │         │ 6. 清除暂停状态
 │         │
 │         │ 7. 发出 Unpaused 事件
 │         │
 │ 8. 恢复完成      │
 │◄────────┤
 │         │
```

这些时序图详细展示了 EasySwap 系统中各个组件之间的交互过程，包括正常流程、错误处理、事件监听等，帮助理解整个系统的工作机制。
