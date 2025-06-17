-- SQL语句练习
-- 题目1：基本CRUD操作
-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
--  age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
-- 要求 ：
-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
    insert into students(name, age, grade) values('张三', 20, '三年级');
-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
    select * from students where age > 18;
-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
    update students set grade = '四年级' where name = '张三';
-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
    delete from students where age < 15;

-- 题目2：事务语句
-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
-- from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- 要求 ：
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
-- 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
-- 如果余额不足，则回滚事务。
BEGIN;

-- 声明要转账的金额
SET @amount = 100;
SET @from_account_id = A;
SET @to_account_id = B;

-- 1. 检查账户 A 的余额是否足够
SELECT balance INTO @from_balance
FROM accounts
WHERE id = @from_account_id
    FOR UPDATE; -- 锁定该行以防止其他事务并发修改

-- 2. 如果余额足够，进行转账操作
IF @from_balance >= @amount THEN
    -- 从账户 A 中扣除金额
UPDATE accounts
SET balance = balance - @amount
WHERE id = @from_account_id;

-- 向账户 B 中增加金额
UPDATE accounts
SET balance = balance + @amount
WHERE id = @to_account_id;

-- 记录转账信息到 transactions 表
INSERT INTO transactions (from_account_id, to_account_id, amount)
VALUES (@from_account_id, @to_account_id, @amount);

-- 提交事务
COMMIT;
ELSE
    -- 如果余额不足，回滚事务
    ROLLBACK;
END IF;
