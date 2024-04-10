import re
import pyperclip

def to_snake_case(s):
    """
    将字符串转换为蛇形命名，处理多种分隔符和大小写变化。

    参数:
    - s: 要转换的字符串。

    返回:
    - 转换后的蛇形命名字符串。
    """
    # 使用正则表达式分割字符串：分割条件是非字母数字或者大写字母（考虑到驼峰命名）
    words = re.split(r'(?=[A-Z])|\W+', s)

    # 去除空字符串，可能出现在连续分隔符或字符串开头/结尾的情况
    words = [word for word in words if word]

    # 将所有单词转换为小写，并使用下划线连接它们
    snake_case = '_'.join(word.lower() for word in words)

    return snake_case

# 从剪贴板读取内容
content = pyperclip.paste()

# 转换内容
snake_case_content = to_snake_case(content)

# 打印转换后的内容（可选）
print(f"转换后的蛇形命名: {snake_case_content}")

# 将转换后的内容写回剪贴板
pyperclip.copy(snake_case_content)
