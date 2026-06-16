import time

# def make_counter2():
#     n = 0
#     def set_n(x):
#         nonlocal n
#         n = x
#     def get_n():
#         nonlocal n
#         return n
#     def increment():
#         nonlocal n
#         n += 1
#     return set_n, get_n, increment
# set_n, get_n, increment = make_counter2()
# print(get_n()) # 0
# increment()
# print(get_n()) # 1
# increment()
# print(get_n()) # 2
# set_n(10)
# print(get_n()) # 10

# def do_laundry():
#     print("Doing ....")
#     time.sleep(2)
#     print("Done!")

def make_timed_function(f):
    def func(*args, **kwargs):
        s_time = time.time()
        result = f(*args, **kwargs)
        f_time = time.time()
        print(f"Time taken: {f_time - s_time:.4f} seconds")
        return result
    return func

@make_timed_function
def do_laundry(*args, **kwargs):
    if args:
        print(f"{args[0]} is doing laundry...")
    print("Doing ...")
    time.sleep(kwargs.get('duration', 2))
    print("Done!")
do_laundry("Alice", duration=1)




# class MyIterator:
#     def __init__(self, data):
#         self.data = data
#         self.index = 0
    
#     def __iter__(self):
#         return self
    
#     def __next__(self):
#         if self.index < len(self.data):
#             result = self.data[self.index]
#             self.index += 1
#             return result
#         else:
#             raise StopIteration
# my_iter = MyIterator("Hello, World!")
# iter_of_my_iter = iter(my_iter)
# print(f"my_iterator_type: {type(my_iter)}")
# print(f"iter_of_my_iterator_type: {type(iter_of_my_iter)}")
# print(f"my_iter is iter_of_my_iter? {my_iter is iter_of_my_iter}") # True

# def test_yield():
#     yield 1
#     yield 2
#     yield 3
# gen = test_yield()
# for index, value in enumerate(gen):
#     print(f"Value: {index, value}")


# def make_adder(n):
#     def adder(x):
#         return x + n
#     return adder
# add3 = make_adder(3)
# print(add3(4))

