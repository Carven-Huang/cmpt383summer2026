### A1
lst = [i for i in range(1,11)]
print(lst)

### A2
lst = [i**2 for i in range(1,21)]
print(lst)

### A3
lst = [i**2 for i in range(1,31) if i % 2 == 0]
print(lst)

### A4
def odd_cube(n):
    return [i**2 for i in range(n) if i % 2 != 0]
print(odd_cube(10))

### A5
words = ["cat", "house", "a", "python", "hi"]
length_of_each = [len(_) for _ in words]
print(length_of_each)

### A6
names = ["bob", "", "alice", "  ", "charlie", "dean"]
cap_name = [ _.title() for _ in names if _.strip() != '']
print(cap_name)

### A7
words = ["bob", "", "mary", "carla", "sam", "dean", "ed"]
cap_word = [_.title() for _ in words
                if _.strip() != ''
                if len(_) % 2 == 0]
print(cap_word)

### A8
result = []
for x in range(1, 51):
    if x % 3 == 0:
        result.append(x * 10)
print(result)
result = [ _*10 for _ in range(1,51) if _ % 3 ==0]
print(result)

### A9
result = []
for x in range(1, 101):
    if x % 2 == 0:
        if x % 5 != 0:
            result.append(x)
print(result)
result = [_ for _ in range(1,101) if _ % 2 == 0 if _ % 5 != 0]
print(result)

### A10
def classify_nums(x):
    return ["positive" if x > 0 else "negative" if x < 0 else "neutral" for x in xs]
print(classify_nums([1,2,3,4,-1,0]))

### A11
print([x for x in range(10) if x % 2 == 0]) # a list of even number in range [1,9]
print(["even" if x % 2 == 0 else "odd" for x in range(10)])# a list of 'even','odd','even'...

### A12
matrix = [
    [1, 2, 3],
    [4, 5],
    [6, 7, 8, 9]
]
flatten = [x for row in matrix for x in row]
print(flatten)

### A13
matrix = [
    [1, 2, 3],
    [4, 5],
    [6, 7, 8, 9]
]
flatten = [x for row in matrix for x in row if x % 2 == 0]
print(flatten)
### A14
def clean_words(words):
    return [word.strip().lower() for word in words if word.strip() != '']
print(clean_words(["  hello ", "", "WORLD", "   ", "Python "]))

### A15
# [x if x % 2 == 0 for x in range(10)] 
# the list comprehension is not in right order
print([x for x in range(10) if x % 2 == 0])


### B1
matrix = [(1, 1), (1, 2), (1, 3),
 (2, 1), (2, 2), (2, 3),
 (3, 1), (3, 2), (3, 3)]

comp_matrix = [(x,y) for x in range(1,4) for y in range(1,4)]
print(matrix == comp_matrix)

### B2/B3
print([(x, y) for x in [1, 2] for y in ["a", "b"]])
print([(x, y) for y in ["a", "b"] for x in [1, 2]])

### B4
cartesian = [ (x,y) for x in range(1,6) for y in range(1,6) if x != y]
print(cartesian)
print((1,2) not in cartesian)

### B5
cartesian = [ (x,y) for x in range(1,6) for y in range(1,6) if x < y]
print(cartesian)
print((3,2) not in cartesian)

### B6
product = [(i,j,k) for i in range(2) for j in range(2) for k in range(2)]
print(product)

### B7
def bit_strings_4():
    return [(i,j,k,l) for i in range(2) 
               for j in range(2) 
               for k in range(2) 
               for l in range(2)]
print(bit_strings_4())

### B8
result = [ (x,y) for x in range(1,101) for y in range(1,101)
                    if x ** 2 + y == 80]
print(result)

### B9
def pythagorean_triples(n):
    return [ (a,b,c) for a in range(n) for b in range(n) for c in range(n)
                        if a ** 2 + b ** 2 == c ** 2
                        if 1 <= a <= b <= c <= n]
print(pythagorean_triples(10))

### B10
# 10^4 in total, quadracit growth

### B11
def reverse_sum_same_digits():
    solutions = [ (a,b,c) for a in range(1,10)
                        for b in range(10)
                        if a != b
                        for c in range(1,10)
                        if b != c and a != c and b != a
                        if len(str((abc := a*100+b*10+c) + (cba := c*100+b*10+a))) == 3
                        if set(str(abc+cba))  == 0
                        ]
    print(solutions)
    print(f"Number of solutions: {len(solutions)}")
reverse_sum_same_digits()

### B12
print(
    len([x + y for x in [1, 2, 3] for y in [10, 20, 30]])
)

### C - skip

#----- ZIP
### D1
def add_list(a,b):
    print(sum([ x+y for x,y in zip(a,b)]))
add_list([1,2,3],[3,2,1])

### D2
print(list(zip([1, 2, 3], ["a", "b", "c"])))

### D3
print(list(zip([1, 2, 3, 4], ["a", "b"])))

### D4
# [x + y for x, y in zip([1, 2, 3], [10, 20, 30])]  -> 3 in total
# [x + y for x in [1, 2, 3] for y in [10, 20, 30]] -> 3x3 = 9

### D5
def dot_product(a, b):
    return sum(list(x*y for x,y in zip(a,b)))
print(dot_product([1,2,3],[4,5,6]))

### D6
def passing_students(names, scores):
    return [name for name,score in zip(names,scores) if score >= 50]

### D7
def adjacent_chars(s):
    return [(x,y) for x,y in zip(s,s[1:])]
print(adjacent_chars("hour"))

### D8
def is_sorted(nums):
    return all(x <= y for x,y in zip(nums,nums[1:]))
print(is_sorted("12345"))
print(is_sorted("123465"))

### D9
def is_sorted(nums):
    return all(x < y for x,y in zip(nums,nums[1:]))
print(is_sorted("123445"))
print(is_sorted("12345"))

### D10
def first_drop_index(nums):
    prev_value = nums[0]
    for i, value in enumerate(nums):
        if prev_value > value:
            return i - 1
        else:
            prev_value = value
    return None
print(first_drop_index([1,4,3,4,5]))
def first_drop_index_2(nums):
    for i,value in enumerate(zip(nums,nums[1:])):
        if value[0] > value[1]:
            return i

    return None
print(first_drop_index_2([1,4,3,4,5]))
### D11
print(all([])) # True, vacuously true

### D12
def combine_records(names, ages, cities):
    return [ (name,age,city) for name,age,city in zip(names,ages,cities)]
print(combine_records(["Alice", "Bob"], [30, 25], ["NY", "LA"]))

### D13
z = zip([1, 2, 3], ["a", "b", "c"])
print(list(z))
print(list(z))

### D14
def print_pairs(a,b):
    for x,y in zip(a,b):
         print((x,y))
print_pairs([1,2,3],[4,5,6])


### E1
def f(a, b, c):
    return a + b + c
values = [1,2,3]
# f(values) no output, need 3 args but given 1
print(f(*values)) # act like passing 3 args

### E2
def call_with_three_args(f, values):
    return f(*values)

### E3
matrix = [[1, 2],
 [3, 4],
 [5, 6]]
def transpose(matrix):
    return[list(row) for row in zip(*matrix)]
    return list(zip(*matrix))
print(transpose(matrix))


### E4/E5
[list(row) for row in zip(*matrix)] # a list of lists (mutable)
list(zip(*matrix)) # a list of immutable tuples

### E6
matrix = [[1, 2, 3], [4, 5], [6, 7, 8]]
# depend on the elem with fewest numbers
print(list(zip(*matrix))) #output [(1,4,6),(2,5,7)]

### E7/E8
def transpose_manual(matrix):
    transpose = list()
    num_rows = len(matrix)
    num_cols = len(matrix[0])
    for c in range(num_cols):
        new_row = []
        for r in range(num_rows):
            new_row.append(matrix[r][c])
        transpose.append(new_row)
    return transpose
matrix = [[1, 2],
 [3, 4],
 [5, 6]]
print(transpose_manual(matrix)) # less obivious version

### F1
def print_index_value(lst):
    for i,value in enumerate(lst):
        print(i,value)
print_index_value(["a",'b','c'])

### F2
def numbered_lines(lines):
    for i,value in enumerate(lines):
        print(f"{i+1}. {value}")
numbered_lines(["hello", "world"])

### F3
def get_min(lst):
    min_index = 0
    min_value = lst[0]
    for i,value in enumerate(lst):
        if value < min_value:
            min_value = value
            min_index = i
    return min_index,min_value
print(get_min([5,4,3,2,1]))

### F4
def get_last_max(lst):
    max_index = 0
    max_value = lst[0]
    for i,value in enumerate(lst):
        if value >= max_value:
            max_value = value
            max_index = i
    return max_index,max_value
print(get_last_max([5,4,3,2,1,5]))

### F5
for i, value in enumerate(["x", "y"], start=5):
    print(i, value) # 5 x 6 y, just the index changed

### F6
#返回两个列表中所有不相等的位置
def mismatches(a, b):
    result = list()
    for i,value in enumerate(zip(a,b)):
        if value[0] != value[1]:
            result.append((i, value[0], value[1]))
    return result
print(mismatches("abcde", "abXde")) # output [(2,'c','X')]

### F7
nums = [1, 2, 3, 4]

for i, x in enumerate(nums):
    if x % 2 == 0:
        nums.pop(i) # i - 1 after pop, the list becomes [1,3,4], but the index will move to 2, which is out of bound, so the loop ends here
    # not safe to modify the list while iterating, because the index will be messed up after pop
print(nums)

### G1
a = [1, 2, 3] # a list
b = iter(a) # iterator
c = "cat" # string
d = iter("cat") #iterator
e = range(5) # iterable

### G2/G3/G4
it = iter([10, 20, 30])
print(next(it)) # 10
print(next(it)) # 20
it = iter([1, 2, 3])
print(list(it))
print(list(it)) # output [], raise StopIteration

### G5/G6
class Letters:
    def __init__(self,str):
        self.data = str
        self.index = 0
    def __iter__(self):
        return self
    def __next__(self):
        if self.index < len(self.data):
            self.index += 1
            return self.data[self.index - 1]
        else:
            raise StopIteration
        
    def next(self):
        return self.__next__()
letters = Letters("hello")
for _ in letters:
    print(_)
# print(letters.next()) # StopIteration raised
letters = Letters("hello")
print(letters.next()) # h

### G7/G8
class MyEnumerate:
    def __init__(self, iterable,start=0):
        self.iterable = iterable
        self.start = start
        self.index = 0
    def __iter__(self):
        return self
    def __next__(self):
        if self.index < len(self.iterable):
            self.index += 1
            return (self.index - 1+self.start, self.iterable[self.index - 1])
        else:
            raise StopIteration

for index, value in MyEnumerate(["a", "b", "c"], start=1):
    print(index, value) # 1 a, 2 b, 3 c

### G9
class MyReserve:
    def __init__(self, iterable):
        self.iterable = iterable
        self.index = len(iterable)
    def __iter__(self):
        return self
    def __next__(self):
        if self.index > 0:
            self.index -= 1
            return self.iterable[self.index]
        else:
            raise StopIteration
for value in MyReserve("hello"):
    print(value) # o l l e h

### G10
### the __iter__ method shall return the iterator object itself,
#  which is the object that has the __next__ method implemented.
#  In this case, the __iter__ method returns a list, 
# which is not an iterator, 
# so it will cause an error when we try to call next() on it. 
# The correct implementation should return self instead of a list.
class BadCounter:
    def __init__(self):
        self.n = 0

    def __iter__(self):
        # return [1, 2, 3]
        return self

    def __next__(self):
        self.n += 1
        return self.n
counter = BadCounter()
iterator = iter(counter)
print(next(iterator)) # TypeError: 'list' object is not an iterator

### G11/G12
class GeneralCounter:
    def __init__(self, upper_bound=None):
        self.n = 0
        self.upper_bound = upper_bound
    def __iter__(self):
        return self
    def __next__(self):
        if self.upper_bound is not None and self.n >= self.upper_bound:
            raise StopIteration
        self.n += 1
        return self.n
counter = GeneralCounter(5)
for _ in counter:
    print(_) # 1 2 3 4 5

### G13/G14
class Primes:
    def __init__(self, upper_bound=None):
        self.n = 1
        self.upper_bound = upper_bound
    def __iter__(self):
        return self
    def __next__(self):
        while True:
            if self.n > self.upper_bound:
                raise StopIteration
            self.n += 1
            if self.is_prime(self.n):
                return self.n
    def is_prime(self, n):
        if n < 2:
            return False
        for i in range(2, int(n**0.5) + 1):
            if n % i == 0:
                return False
        return True
    def next_prime__(self):
        return self.__next__()
primes = Primes(5)
for _ in primes:
    print(_) # 2 3 5

### G15
it = iter([1, 2, 3])
 # True, __iter_ of an iterator should return itself
print(f"{iter(it) is it}")

it = iter([1, 2, 3])
for x in it:
    for y in it:
        print(x, y)
    # output: 1 2   1 3
    # the iterator is exhausted after the first loop

### H1/H2
def simple_steps():
    yield "Step 1"
    yield "Step 2"
    yield "Step 3"
for step in simple_steps():
    print(step)
gen = simple_steps()
# print(next(gen)) # Step 1
# print(next(gen)) # Step 2
# print(next(gen)) # Step 3
# print(next(gen)) # StopIteration raised

### H3/H4/H5/H6
def my_enumerate(lst):
    for i in range(len(lst)):
        yield i, lst[i]
print(list(my_enumerate(["a", "b", "c"]))) # [(0, 'a'), (1, 'b'), (2, 'c')]

def my_reversed(lst):
    for i in range(len(lst)-1, -1, -1):
        yield i, lst[i]
print(list(my_reversed(["a", "b", "c"]))) 

def evens_up_to(n):
    for i in range(n):
        if i % 2 == 0:
            yield i
        else:
            continue
print(list(evens_up_to(10)))

# def naturals(n):
#     i = 0
#     while True:
#         yield i
#         i += 1
# print(list(naturals(10)))

### H7
def take(n,iterable):
    iterable = iter(iterable)
    for _ in range(n):
        try:
            yield next(iterable)
        except StopIteration:
            return

### H8/H9

def is_prime(n):
    if n < 2:
        return False
    for i in range(2, int(n**0.5)+1):
        if n % i == 0:
            return False
    return True

def primes_gen(n = 100):
    i = 1
    for i in range(n):
        if is_prime(i):
            yield i
print(list(primes_gen(50)))

### H10
def f():
    yield 1
    return 2 # this is not a yield, it will cause StopIteration(2) when the generator is exhausted
    yield 3 # never run, return 2 causes the f() generator to be exhausted
# print(list(f()))

### H11/H12

def evens_up_to(n):
    for i in range(n):
        if i % 2 == 0:
            yield i
        else:
            continue
gen = evens_up_to(6)
print(gen)
print(list(gen)) 
print(list(gen)) # this line output [], because the generator is exhausted
print(type(evens_up_to)) # <class 'function'>
print(type(evens_up_to(6))) # <class 'generator'>

### H13
lines = ["INFO start", "ERROR bad", "INFO end", "ERROR worse"]
def error_lines(lines):
    for line in lines:
        if line.startswith("ERROR"):
            yield line

### H14
lines = ["INFO start", "ERROR bad", "INFO end", "ERROR worse"]

def strip_lines(lines):
    for line in lines:
        yield line.strip()

def non_empty(lines):
    for line in lines:
        if line.strip() != '':
            yield line

def uppercase(lines):
    for line in lines:
        yield line.upper()
result = list(uppercase(non_empty(strip_lines(lines))))
print(result)


### I1
def make_adder(n):
    def adder(v):
        return v + n
    return adder
add3 = make_adder(3)
add5 = make_adder(5)
print(add3(3)) # 6
print(add5(5)) # 10

### I2
def make_multiplier(n):
    def multiplier(v):
        return v * n
    return multiplier
mul2 = make_multiplier(2)
mul3 = make_multiplier(3)
print(mul2(4)) # 8
print(mul3(4)) # 12


### I3
def make_power(n):
    def multiplier(v):
        return v ** n
    return multiplier
mul2 = make_power(2)
mul3 = make_power(3)
print(mul2(4))
print(mul3(4))


### I4
def make_counter():
    count = 0
    def counter():
        nonlocal count
        count += 1
        return count
    return counter
counter = make_counter()
for i in range(10):
    print(counter())

### I5/I6
# def make_counter():
#     n = 0

#     def counter():
#         # n is local and undefined but we want to modify it,
#         #  so we need to declare it as nonlocal
#         # nonlocal n
#         n += 1
#         return n

#     return counter

### I7
def make_bounds_checker(low, high):
    def bounds_checker(x):
        nonlocal low,high
        return low <= x <= high
    return bounds_checker
check_0_10 = make_bounds_checker(0, 10)
print(check_0_10(5)) # True

### I8
def make_prefixer(prefix):
    def prefixer(x):
        nonlocal prefix
        return str(prefix) + str(x)
    return prefixer
addHello = make_prefixer("Hello")
print(addHello(" World!"))


### I12
def make_funcs():
    funcs = []
    for i in range(3):
        def f():
            return i
        funcs.append(f)
    return funcs

f0, f1, f2 = make_funcs()
print(type(f0))
print(f0(), f1(), f2()) # 2,2,2, each fx is a function class



### J1/J2/J3
def make_hello_goodbye(f):
    def wrapper():
        print("starting")
        result = f()
        print("done")
        return result
    return wrapper
# test_func = make_hello_goodbye(lambda: "Hello, World!")
# print(test_func()) # starting, done, Hello, World!
@make_hello_goodbye
def say_hi():
    print("hi")
say_hi()

### J4/J5
def print_before_after(f):
    def wrapper():
        print('before')
        result = f()
        print('after')
        return result
    return wrapper
def add():
    return 1 + 2
adder = print_before_after(add)
result = adder()
print(result)

### J6
def deco1(f):
    def wrapper():
        print("deco1 before")
        result = f()
        print("deco1 after")
        return result
    return wrapper

def deco2(f):
    def wrapper():
        print("deco2 before")
        result = f()
        print("deco2 after")
        return result
    return wrapper

@deco1
@deco2
def hello():
    print("hello")

hello()
# ---------------output the same as below----------------
def deco3(f):
    def wrapper():
        print("deco1 before")
        print("deco2 before")
        result = f()
        print("deco2 after")
        print("deco1 after")
        return result
    return wrapper
@deco3
def hello():
    print("hello")
hello()

