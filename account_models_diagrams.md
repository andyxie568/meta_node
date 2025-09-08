# 区块链账户模型可视化图表

## 1. UTXO模型结构图

```mermaid
graph TD
    A[UTXO模型] --> B[交易输入 Inputs]
    A --> C[交易输出 Outputs]
    
    B --> D[引用UTXO]
    B --> E[解锁脚本]
    
    C --> F[创建新UTXO]
    C --> G[锁定脚本]
    
    D --> H[UTXO1: 1 BTC]
    D --> I[UTXO2: 0.5 BTC]
    
    F --> J[给Bob: 0.3 BTC]
    F --> K[找零: 1.2 BTC]
    
    H --> L[被消费]
    I --> L
    J --> M[新UTXO]
    K --> N[新UTXO]
```

## 2. 以太坊账户生成流程图

```mermaid
graph LR
    A[随机数生成] --> B[私钥 Private Key]
    B --> C[椭圆曲线计算]
    C --> D[公钥 Public Key]
    D --> E[Keccak256哈希]
    E --> F[取后20字节]
    F --> G[以太坊地址]
    
    B --> H[数字签名]
    H --> I[交易验证]
    
    G --> J[EOA账户]
    G --> K[合约账户]
```

## 3. 公私钥关系图

```mermaid
graph TD
    A[私钥 Private Key] --> B[椭圆曲线乘法]
    B --> C[公钥 Public Key]
    C --> D[哈希函数]
    D --> E[地址 Address]
    
    A --> F[数字签名]
    F --> G[交易签名]
    
    C --> H[签名验证]
    H --> I[身份验证]
    
    E --> J[账户标识]
    J --> K[余额查询]
    J --> L[交易发送]
```

## 4. UTXO vs 账户模型对比

```mermaid
graph TB
    A[区块链账户模型] --> B[UTXO模型]
    A --> C[账户模型]
    
    B --> D[比特币]
    B --> E[Litecoin]
    B --> F[Bitcoin Cash]
    
    C --> G[以太坊]
    C --> H[EOS]
    C --> I[Tron]
    
    B --> J[特点]
    J --> K[并行处理]
    J --> L[高隐私性]
    J --> M[防双花]
    
    C --> N[特点]
    N --> O[智能合约]
    N --> P[状态管理]
    N --> Q[开发者友好]
```

## 5. 交易生命周期对比

### UTXO交易生命周期
```mermaid
graph LR
    A[选择UTXO] --> B[创建交易]
    B --> C[签名交易]
    C --> D[广播交易]
    D --> E[验证UTXO]
    E --> F[消费UTXO]
    F --> G[创建新UTXO]
    G --> H[交易确认]
```

### 以太坊交易生命周期
```mermaid
graph LR
    A[检查余额] --> B[创建交易]
    B --> C[签名交易]
    C --> D[广播交易]
    D --> E[验证签名]
    E --> F[更新账户状态]
    F --> G[执行智能合约]
    G --> H[交易确认]
```

## 6. 地址生成过程对比

### 比特币地址生成
```mermaid
graph TD
    A[私钥] --> B[公钥]
    B --> C[SHA256哈希]
    C --> D[RIPEMD160哈希]
    D --> E[添加版本字节]
    E --> F[双重SHA256校验]
    F --> G[Base58编码]
    G --> H[比特币地址]
```

### 以太坊地址生成
```mermaid
graph TD
    A[私钥] --> B[公钥]
    B --> C[Keccak256哈希]
    C --> D[取后20字节]
    D --> E[添加0x前缀]
    E --> F[以太坊地址]
```

## 7. 安全机制对比

```mermaid
graph TB
    A[安全机制] --> B[UTXO安全]
    A --> C[账户安全]
    
    B --> D[UTXO唯一性]
    B --> E[脚本验证]
    B --> F[防双花]
    
    C --> G[Nonce防重放]
    C --> H[Gas费用]
    C --> I[状态验证]
    
    D --> J[每个UTXO只能使用一次]
    E --> K[脚本语言验证]
    F --> L[交易排序]
    
    G --> M[递增序号]
    H --> N[防止垃圾交易]
    I --> O[状态一致性]
```

## 8. 智能合约支持对比

```mermaid
graph TD
    A[智能合约支持] --> B[UTXO模型]
    A --> C[账户模型]
    
    B --> D[脚本语言]
    B --> E[有限功能]
    B --> F[复杂实现]
    
    C --> G[图灵完整]
    C --> H[丰富功能]
    C --> I[简单开发]
    
    D --> J[Bitcoin Script]
    E --> K[基础操作]
    F --> L[需要复杂设计]
    
    G --> M[Solidity/Vyper]
    H --> N[状态管理]
    I --> O[开发者友好]
```

## 9. 性能对比

```mermaid
graph TB
    A[性能指标] --> B[UTXO模型]
    A --> C[账户模型]
    
    B --> D[并行处理]
    B --> E[状态查询]
    B --> F[交易大小]
    
    C --> G[串行处理]
    C --> H[状态查询]
    C --> I[交易大小]
    
    D --> J[高并发]
    E --> K[复杂查询]
    F --> L[可变大小]
    
    G --> M[有限并发]
    H --> N[简单查询]
    I --> O[固定大小]
```

## 10. 实际应用场景

```mermaid
graph TD
    A[应用场景] --> B[UTXO适用]
    A --> C[账户适用]
    
    B --> D[简单转账]
    B --> E[高隐私需求]
    B --> F[并行处理]
    
    C --> G[智能合约]
    C --> H[复杂业务逻辑]
    C --> I[状态管理]
    
    D --> J[比特币转账]
    E --> K[隐私币]
    F --> L[高TPS需求]
    
    G --> M[DeFi应用]
    H --> N[游戏应用]
    I --> O[企业应用]
```

这些图表展示了UTXO模型和账户模型的核心差异，以及它们在区块链系统中的不同应用场景。理解这些概念对于选择合适的区块链平台和开发策略至关重要。
