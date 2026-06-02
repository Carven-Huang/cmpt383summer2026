class my_string:
    def __init__(self, string):
        self.string = string + '\n'
    def __str__(self):
        return self.string
    def __iter__(self):
        return iter(self.string)
    def __len__(self):
        return len(self.string)
    def __next__(self):
        return self.string

my_str = my_string('Hello, World!')
print(my_str)  # Output: Hello, World!
print(len(my_str))  # Output: 14 (including the newline character)
for char in my_str:
    print(char)  # Output: H e l l o ,   W o r l d ! (each character on a new line)
iter_my_str = iter(my_str)
print(next(iter_my_str) for _ in range(len(my_str)))  # Output: Hello, World!

lst = [1, 2, 3, 4, 5]

# 1. 列表和它的迭代器：地址必然不同，返回 False
list_iterator = iter(lst) # 等价于 lst.__iter__()
print(f"列表地址:   {hex(id(lst))}")
print(f"迭代器地址: {hex(id(list_iterator))}")
print(f"它们一样吗? {lst is list_iterator}\n") # False

# 2. 迭代器调用自己的 __iter__：地址必然完全相同，返回 True
iterator_of_iterator = iter(list_iterator) # 等价于 list_iterator.__iter__()
print(f"迭代器地址:     {hex(id(list_iterator))}")
print(f"二次迭代器地址: {hex(id(iterator_of_iterator))}")
print(f"它们是同一个吗? {list_iterator is iterator_of_iterator}") # True

lst = [1, 2, 3, 4, 5]
addr = iter(lst)
print(f"address at: {addr}")
print(f"Type of address: {type(addr)}")
print(f"dereferenced value: {lst}")

