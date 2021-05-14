# 使用golang重写python小脚本

## sushu.go

python版本

```python
a = []
for i in range(0,8):
	if i > 1:
		num = i // 2
		while num > 1:
			if i % num == 0:
				break
			num -= 1
		else:
			a.append(i)
print(a)
```