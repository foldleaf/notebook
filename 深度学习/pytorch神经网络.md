学自李宏毅老师

# 神经网络训练设置
```py
# 从数据集中读取数据
dataset = MyDataset(file)
# 将数据放到数据加载器，分批
tr_set = DataLoader(dataset, 16, shuffle=True)
# 构造模型并指定设备(cpu/cuda(gpu))
model = MyModel().to(device)
# 设置损失函数
criterion = nn.MSELoss()
# 设置优化器
optimizer = torch.optim.SGD(model.parameters(), 0.1)

```
# 神经网络训练回路
感觉这里用回路比循环好听一点
```py
# 迭代 n_epochs
for epoch in range(n_epochs):
    # 将模型设置为训练模式
    model.train()
    # 通过数据加载器迭代
    for x, y in tr_set:
        # 梯度置为 0
        optimizer.zero_grad()
        # cpu/cuda
        x, y = x.to(device), y.to(device)
        # 正向传播（计算输出）
        pred = model(x)
        # 计算损失
        loss = criterion(pred, y)
        # 反向传播（计算梯度）
        loss.backward()
        # 使用优化器更新模型
        optimizer.step()
```

# 神经网络验证回路
```py
# 将模型设置为评估模式
model.eval()
total_loss = 0
# 通过数据加载器迭代
for x, y in dv_set:
    # cpu/cuda(gpu)
    x, y = x.to(device), y.to(device)
    # 禁用梯度计算
    with torch.no_grad():
        # 前向传播（计算输出）
        pred = model(x)
        # 计算损失
        loss = criterion(pred, y)
    # 累积损失
    total_loss += loss.cpu().item() * len(x)
    # 计算平均损失
    avg_loss = total_loss / len(dv_set.dataset)
```
# 神经网络测试回路
```py
# # 将模型设置为评估模式
model.eval()
preds = []
for x in tt_set:
    x = x.to(device)
    with torch.no_grad():
        pred = model(x)
        # 收集预测结果
        preds.append(pred.cpu())
```

