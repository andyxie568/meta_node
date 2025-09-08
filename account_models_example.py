#!/usr/bin/env python3
"""
区块链账户模型示例代码
演示UTXO模型、以太坊账户生成和公私钥关系
"""

import hashlib
import secrets
import json
from typing import List, Dict, Any
from dataclasses import dataclass
from ecdsa import SigningKey, SECP256k1
from ecdsa.util import string_to_number, number_to_string

# 配置
SECP256k1_ORDER = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141

@dataclass
class UTXO:
    """UTXO数据结构"""
    txid: str
    vout: int
    value: int  # 以satoshi为单位
    script_pubkey: str
    address: str
    
    def to_dict(self) -> Dict[str, Any]:
        return {
            'txid': self.txid,
            'vout': self.vout,
            'value': self.value,
            'script_pubkey': self.script_pubkey,
            'address': self.address
        }

@dataclass
class TransactionInput:
    """交易输入"""
    txid: str
    vout: int
    script_sig: str
    sequence: int = 0xFFFFFFFF

@dataclass
class TransactionOutput:
    """交易输出"""
    value: int
    script_pubkey: str
    address: str

class BitcoinUTXOModel:
    """比特币UTXO模型实现"""
    
    def __init__(self):
        self.utxos: Dict[str, UTXO] = {}  # 全局UTXO集合
        self.transactions: List[Dict] = []  # 交易历史
    
    def create_utxo(self, txid: str, vout: int, value: int, address: str) -> UTXO:
        """创建UTXO"""
        utxo = UTXO(
            txid=txid,
            vout=vout,
            value=value,
            script_pubkey=f"OP_DUP OP_HASH160 {address} OP_EQUALVERIFY OP_CHECKSIG",
            address=address
        )
        self.utxos[f"{txid}:{vout}"] = utxo
        return utxo
    
    def find_utxos_by_address(self, address: str) -> List[UTXO]:
        """根据地址查找UTXO"""
        return [utxo for utxo in self.utxos.values() if utxo.address == address]
    
    def create_transaction(self, from_address: str, to_address: str, amount: int, fee_rate: int = 1000) -> Dict:
        """创建交易"""
        # 查找发送方的UTXO
        available_utxos = self.find_utxos_by_address(from_address)
        if not available_utxos:
            raise ValueError("No UTXOs found for address")
        
        # 选择UTXO (简化版本，实际需要更复杂的算法)
        selected_utxos = []
        total_input = 0
        
        for utxo in available_utxos:
            selected_utxos.append(utxo)
            total_input += utxo.value
            if total_input >= amount + fee_rate:
                break
        
        if total_input < amount + fee_rate:
            raise ValueError("Insufficient funds")
        
        # 创建交易输入
        inputs = []
        for utxo in selected_utxos:
            inputs.append(TransactionInput(
                txid=utxo.txid,
                vout=utxo.vout,
                script_sig=f"<signature> <pubkey>"
            ))
        
        # 创建交易输出
        outputs = []
        outputs.append(TransactionOutput(
            value=amount,
            script_pubkey=f"OP_DUP OP_HASH160 {to_address} OP_EQUALVERIFY OP_CHECKSIG",
            address=to_address
        ))
        
        # 计算找零
        change = total_input - amount - fee_rate
        if change > 0:
            outputs.append(TransactionOutput(
                value=change,
                script_pubkey=f"OP_DUP OP_HASH160 {from_address} OP_EQUALVERIFY OP_CHECKSIG",
                address=from_address
            ))
        
        # 创建交易
        transaction = {
            'version': 1,
            'inputs': [{'txid': inp.txid, 'vout': inp.vout, 'script_sig': inp.script_sig, 'sequence': inp.sequence} for inp in inputs],
            'outputs': [{'value': out.value, 'script_pubkey': out.script_pubkey} for out in outputs],
            'locktime': 0
        }
        
        # 消耗UTXO
        for utxo in selected_utxos:
            del self.utxos[f"{utxo.txid}:{utxo.vout}"]
        
        # 创建新的UTXO
        txid = hashlib.sha256(json.dumps(transaction, sort_keys=True).encode()).hexdigest()
        for i, output in enumerate(outputs):
            self.create_utxo(txid, i, output.value, output.address)
        
        self.transactions.append(transaction)
        return transaction

class EthereumAccount:
    """以太坊账户实现"""
    
    def __init__(self, private_key: int = None):
        if private_key is None:
            self.private_key = self._generate_private_key()
        else:
            self.private_key = private_key
        
        self.public_key = self._private_key_to_public_key(self.private_key)
        self.address = self._public_key_to_address(self.public_key)
        self.nonce = 0
        self.balance = 0
    
    def _generate_private_key(self) -> int:
        """生成私钥"""
        while True:
            private_key = secrets.randbits(256)
            if 0 < private_key < SECP256k1_ORDER:
                return private_key
    
    def _private_key_to_public_key(self, private_key: int) -> bytes:
        """从私钥生成公钥"""
        signing_key = SigningKey.from_secret_exponent(private_key, curve=SECP256k1)
        return signing_key.get_verifying_key().to_string()
    
    def _public_key_to_address(self, public_key: bytes) -> str:
        """从公钥生成地址"""
        # 以太坊使用Keccak256而不是SHA256
        public_key_hash = hashlib.sha3_256(public_key).digest()
        return "0x" + public_key_hash[-20:].hex()
    
    def sign_transaction(self, to_address: str, value: int, data: bytes = b'', gas_limit: int = 21000, gas_price: int = 20) -> Dict:
        """签名交易"""
        transaction = {
            'nonce': self.nonce,
            'gasPrice': gas_price,
            'gasLimit': gas_limit,
            'to': to_address,
            'value': value,
            'data': data.hex() if data else '0x'
        }
        
        # 序列化交易
        transaction_data = self._serialize_transaction(transaction)
        
        # 签名
        signing_key = SigningKey.from_secret_exponent(self.private_key, curve=SECP256k1)
        signature = signing_key.sign(transaction_data)
        
        # 解析签名
        r, s, v = self._parse_signature(signature)
        
        # 添加签名到交易
        transaction['r'] = hex(r)
        transaction['s'] = hex(s)
        transaction['v'] = v
        
        self.nonce += 1
        return transaction
    
    def _serialize_transaction(self, transaction: Dict) -> bytes:
        """序列化交易（简化版本）"""
        # 实际实现需要RLP编码
        return json.dumps(transaction, sort_keys=True).encode()
    
    def _parse_signature(self, signature: bytes) -> tuple:
        """解析签名"""
        r = int.from_bytes(signature[:32], 'big')
        s = int.from_bytes(signature[32:64], 'big')
        v = 27  # 简化版本
        return r, s, v

class AccountManager:
    """账户管理器"""
    
    def __init__(self):
        self.accounts: Dict[str, EthereumAccount] = {}
        self.utxo_model = BitcoinUTXOModel()
    
    def create_ethereum_account(self) -> EthereumAccount:
        """创建以太坊账户"""
        account = EthereumAccount()
        self.accounts[account.address] = account
        return account
    
    def create_bitcoin_utxo(self, txid: str, vout: int, value: int, address: str) -> UTXO:
        """创建比特币UTXO"""
        return self.utxo_model.create_utxo(txid, vout, value, address)
    
    def get_account_balance(self, address: str) -> int:
        """获取账户余额"""
        if address in self.accounts:
            return self.accounts[address].balance
        return 0
    
    def get_utxo_balance(self, address: str) -> int:
        """获取UTXO余额"""
        utxos = self.utxo_model.find_utxos_by_address(address)
        return sum(utxo.value for utxo in utxos)

def demonstrate_utxo_model():
    """演示UTXO模型"""
    print("=== UTXO模型演示 ===")
    
    manager = AccountManager()
    
    # 创建一些UTXO
    utxo1 = manager.create_bitcoin_utxo("tx1", 0, 100000, "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")  # 0.001 BTC
    utxo2 = manager.create_bitcoin_utxo("tx2", 0, 200000, "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")  # 0.002 BTC
    
    print(f"创建UTXO1: {utxo1.value} satoshis")
    print(f"创建UTXO2: {utxo2.value} satoshis")
    
    # 查询余额
    balance = manager.get_utxo_balance("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
    print(f"总余额: {balance} satoshis")
    
    # 创建交易
    try:
        transaction = manager.utxo_model.create_transaction(
            "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
            "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
            150000,  # 0.0015 BTC
            1000     # 0.00001 BTC 手续费
        )
        print(f"交易创建成功: {json.dumps(transaction, indent=2)}")
    except ValueError as e:
        print(f"交易创建失败: {e}")

def demonstrate_ethereum_accounts():
    """演示以太坊账户"""
    print("\n=== 以太坊账户演示 ===")
    
    manager = AccountManager()
    
    # 创建账户
    account1 = manager.create_ethereum_account()
    account2 = manager.create_ethereum_account()
    
    print(f"账户1地址: {account1.address}")
    print(f"账户1公钥: {account1.public_key.hex()}")
    print(f"账户1私钥: {hex(account1.private_key)}")
    
    print(f"账户2地址: {account2.address}")
    print(f"账户2公钥: {account2.public_key.hex()}")
    print(f"账户2私钥: {hex(account2.private_key)}")
    
    # 签名交易
    transaction = account1.sign_transaction(account2.address, 1000000000000000000)  # 1 ETH
    print(f"签名交易: {json.dumps(transaction, indent=2)}")

def demonstrate_key_relationship():
    """演示公私钥关系"""
    print("\n=== 公私钥关系演示 ===")
    
    # 生成私钥
    private_key = secrets.randbits(256)
    print(f"私钥: {hex(private_key)}")
    
    # 从私钥生成公钥
    signing_key = SigningKey.from_secret_exponent(private_key, curve=SECP256k1)
    public_key = signing_key.get_verifying_key().to_string()
    print(f"公钥: {public_key.hex()}")
    
    # 从公钥生成地址
    public_key_hash = hashlib.sha3_256(public_key).digest()
    address = "0x" + public_key_hash[-20:].hex()
    print(f"地址: {address}")
    
    # 验证关系
    print(f"私钥长度: {private_key.bit_length()} bits")
    print(f"公钥长度: {len(public_key)} bytes")
    print(f"地址长度: {len(address)} characters")

def demonstrate_mnemonic():
    """演示助记词生成"""
    print("\n=== 助记词演示 ===")
    
    # 生成熵
    entropy = secrets.randbits(128)
    print(f"熵: {hex(entropy)}")
    
    # 简化的助记词生成（实际需要BIP39标准）
    word_list = [
        "abandon", "ability", "able", "about", "above", "absent", "absorb", "abstract", "absurd", "abuse",
        "access", "accident", "account", "accuse", "achieve", "acid", "acoustic", "acquire", "across", "act"
    ]
    
    # 将熵映射到单词
    mnemonic_words = []
    temp_entropy = entropy
    for _ in range(12):
        word_index = temp_entropy % len(word_list)
        mnemonic_words.append(word_list[word_index])
        temp_entropy //= len(word_list)
    
    print(f"助记词: {' '.join(mnemonic_words)}")

if __name__ == "__main__":
    print("区块链账户模型示例")
    print("=" * 50)
    
    # 演示UTXO模型
    demonstrate_utxo_model()
    
    # 演示以太坊账户
    demonstrate_ethereum_accounts()
    
    # 演示公私钥关系
    demonstrate_key_relationship()
    
    # 演示助记词
    demonstrate_mnemonic()
    
    print("\n演示完成！")
