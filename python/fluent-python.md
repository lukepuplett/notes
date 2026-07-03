# Fluent Python – Complete Notes

**Book:** *Fluent Python* by Luciano Ramalho  
**Notes taken:** May 2026 (in preparation for Pacific Asset Management role)

---

## Chapter 1: Python's Data Model

Python has a "data model" which uses "dunder" (double underscore) functions on classes to implement logic for code such as `len(listything)`, indexation, string representation, addition, multiplication, etc. (See p.13)

**There are 83 dunder functions!!**

The data model lets you integrate seamlessly with built-in functions and operators by implementing special methods like:
- `__len__()` enables `len(obj)`
- `__getitem__()` enables `obj[i]` indexation
- `__repr__()`, `__str__()` for string representation
- `__add__()`, `__mul__()` for arithmetic operations

---

## Chapter 2: An Array of Sequences

### Container vs. Flat Sequences

**Container sequences** hold references to objects they contain (can contain mixed types):
- `list`
- `tuple`
- `collections.deque`

**Flat sequences** hold the values as copies (aren't mixed types):
- `str`
- `bytes`
- `bytearray`
- `memoryview`
- `array.array`

Immutable types: `tuple`, `str`, `bytes`

### List Comprehensions

```python
dummy = [ord(x) for x in x]  # List Comprehension syntax
cartesian = [(color, size) for color in colors for size in sizes]
```

Code in `[]` can be split over many lines for readability.

### Generator Expressions & Unpacking

**"Genexp" (generator expressions)** use similar syntax but iterate one-by-one, which can save memory (prevent creating the whole series):

```python
array.array('I', (ord(symbol) for symbol in symbols))  # Uses () instead of []
```

Generator expressions can be used to feed a for loop which may short-circuit iteration, saving memory/time.

**Unpacking:**
```python
coords = (33.9425, -118.408056)  # Tuple
for field1, _ in tuples:  # for will (unpack) destruct a tuple into named variables (p.29)
    print(field1)

lat, long = coords  # Tuples and any iterable object can be unpacked
b, a = a, b  # Swapping values works!
send_missile(*coords)  # * unpacks into args
lat, long = m.current_location()  # Unpack result
a, b, *rest = range(5)  # (0, 1, [2, 3, 4])
```

### Nested Unpacking & Named Tuples

```python
a, *body, c, d = range(5)  # The * "grab" can appear anywhere! (0, [1, 2], 3, 4)
```

Tuples can be nested, and unpacking works as expected so long as the brackets are in the right places (See p.32).

**Named Tuples:**
```python
City = collections.namedtuple('City', 'name country')  # Note: space instead of comma in string
tokyo = City('Tokyo', 'Japan')
print(tokyo.country)  # "Japan"
print(tokyo[0])  # "Tokyo"
```

The resulting tuple has some members: `City._fields`, `delhi = City._make(...)`, `delhi._asdict()`, also `City(*iter_list)`.

A tuple is almost the same as an immutable list (see p.35).

### Slicing & Mutation

```python
slicing: range(3) and my_list[:3] exclude the item at index 3!
```

This has some benefits, including easy code like `p1=list[:3]` and `p2=list[3:]`.

```python
[start:stop:stride]  # Stride can be negative, which reverses the sequence
UNIT_PRICE = slice(0, 6)
print(text[UNIT_PRICE])  # With []
[i, j]  # Is for multidimensional
[m:n, k:l]  # Will reference multidimensional slices!
```

You can mutate the contents of the slice if it's over a mutable list (like a window):

```python
nums = list(range(10))  # 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
nums[2:5] = [20, 30]  # 0, 1, 20, 30, 5, 6, 7...
del nums[2:5]  # 0, 1, 6, 7
```

You can use `sequence * 4` to make a repeating sequence.

### List Sorting & Arrays

```python
tic_tac_toe_board = [['_'] * 3 for i in range(3)]
```

This listcomp creates three lots of `_` in a list, times three in another list e.g. `[['_', '_', '_'], ['_', '_', '_'], ['_', '_', '_']]`.

```python
ttt_board[1][2] = 'X'
```

Mutable iterables can be extended/appended with `a += b` and `a *= b`, but mutables will get a new instance created (though with the contents you'd expect).

```python
list.sort()  # Sorts in-place
sorted(items, key=str.lower, reverse=True)  # Makes a new instance
```

**bisect and insort:** `haystack.insort(bisect(haystack, needle), needle)`

**Arrays** are much more efficient than lists, especially for numbers. They're as lean as a C array, so you need to initialize with a type code to determine the underlying C type (e.g., `'b'` is for signed char). Python will not let you put in the wrong type.

There's also a `set` which is optimized for duplicate/containment checks.

### Memoryview, NumPy, & Deque

There's a "shared-memory" `memoryview` which is like `Span<T>` in C# .NET. Shared-memory means "not copied":

```python
nums = array.array('h', [-2, -1, 0, 1, 2])
memv = memoryview(nums)
memv2 = memv.cast('B')
```

**NumPy** is built with C and Fortran. SciPy is built on NumPy and have powerful array types:

```python
a = numpy.arange(12)  # Creates a numpy.ndarray
a.shape = 3, 4  # Reshapes in-place
a.transpose()  # Flips the shape/rows/cols
```

Python also has a double-ended queue (`deque`). `append` and `pop` make it usable as a FIFO queue or a stack, but removing from the top is costly. It can have a `maxlen` and will discard the "oldest" items:

```python
dq = deque(range(10), maxlen=10)
```

`append` and `popleft` are atomic; no locks needed.

The `queue` package has `Queue`, `LifoQueue`, and `PriorityQueue` which block when full and are designed for thread-synchronous code.

There's also `multiprocessing.Queue` and `JoinableQueue`, and `asyncio` which has all the aforementioned queue types (new in 3.4 for managing tasks in async coding). Also, `heapq` has methods `heappush` and `heappop` which let you use a mutable sequence.

---

## Chapter 3: Dictionaries & Sets

### Creating Dicts

```python
a = dict(one=1, two=2, three=3)
b = {'one': 1, 'two': 2, 'three': 3}
c = dict(zip(['one', 'two', 'three'], [1, 2, 3]))
d = dict([('two', 2), ('one', 1), ('three', 3)])  # Note order
e = dict({'three': 3, 'one': 1, 'two': 2})  # Note order
a == b == c == d == e  # True!!
```

### Hashable Keys & Dict Implementation

`collections.abc` module contains `Mapping` and `MutableMapping`. Special map types often extend `dict` or `collections.UserDict`.

All standard lib maps use `dict` which is backed by a hash table, so they all need hashable keys. Hashable means they must have `__hash__()` dunder and have an `__eq__()` for equality comparison.

### Dict Comprehensions

```python
DIAL_CODES = [(86, 'China'), (...), ...]
country_code = {country: code for code, country in DIAL_CODES}
{code: country.upper() for country, code in country_code.items() if code < 66}
```

### Mapping Methods Overview

Page 71 has a table of the `dict`, `defaultdict`, and `OrderedDict` mapping types and which methods they implement, including any dunders.

**Duck typing!** `update(m, [kwargs])` looks for a `keys` method on `m`. If found, it assumes it's an iterable mapping. Otherwise, it iterates over `m` assuming its items are `(key, value)` pairs.

```python
d.setdefault(idx, []).append(item)  # Set the value at idx if not exist or get it if it exists
```

```python
d = defaultdict(list)  # Creates a list where missing keys automatically create a new key value of a new list
```

The arg must be "callable" and this only works for `d[missing]`, not `d.get(missing)`.

Subclass `collections.UserDict` and implement `__missing__` for custom logic (P77/78 shows a subclass of dict with custom logic).

### Dict Variations

Variations of dict are:
- `collections.OrderedDict` 
- `collections.ChainMap` (couldn't understand what it was saying)
- `collections.Counter` – holds an int counter for each key for tallies and stuff
- `collections.UserDict` – pure Python, designed to be subclassed

**UserDict** is actually a wrapper over a dict. You can see an example subclass on page 81. It subclasses `MutableMapping` and gains some useful methods.

```python
readonly_proxy = MappingProxyType({2: 'A'})
```

`MappingProxyType` returns a read-only proxy instance when given a mapping instance.

### Set Theory

```python
l = ['ham', 'egg', 'chips', 'ham', 'ham']
set(l)  # {'chips', 'egg', 'ham'}
```

Set of unique hashable values, with infix operators:
- `a | b` is union
- `a & b` is intersection
- `a - b` is difference

Set literal syntax `{1}`, `{1, 2}` looks like math but there's no empty set literal, so have to use `set()` as `{}` is an empty dict.

**Frozen Sets:**
```python
frozenset(range(10))
```

Set literals like `{1, 2, 3}` are much faster than `set([1, 2, 3])` which goes via a list (p.86).

But `frozenset` has no special syntax: `frozenset(range(10))`.

Python 2.7 added "setcomps" which look like listcomps but within `{}` braces (P.87).

**Warning:** Don't update a dict or set while iterating it—you won't reliably iterate all items (not sure if Python throws like .NET).

---

## Chapter 4: Text vs. Bytes

### Unicode & Code Points

The items you get out of a Python 3 `str` are Unicode characters; "code points" of each is a regular base 10 number from 0 to 1,114,111 as hex like `U+0041` (for 'A').

The actual bytes that represent a character depend on the encoding, e.g., single byte `\x41` for 'A' in UTF-8.

```python
str.encode('utf8')  # Returns those bytes
```

The type `bytes` (immutable) is not the same as `bytearray` (mutable).

You can use string methods like `endswith`, `replace`, etc. with binary sequences.

### Creating Bytes

```python
bytes.fromhex('31 4B CE A9')
bytes('foo', encoding='utf-8')
bytes([255, 0, 1, 255, etc.])
bytes(obj_implements_buffer_protocol)  # bytes, bytearray, memoryview, array.array
```

### Binary Operations & Memory

`memoryview` lets you share memory between binary data structures.

The `struct` module lets you binary serialize and deserialize structs/tuples, like .NET binary serializer.

Page 106 shows how to use `struct.unpack(format, data)` to pull some header info from a GIF file.

Interestingly, the `memoryview` instances must be deleted with `del var_name` to release memory!

Can even open the file with `mmap` module which would memory map it!!

### Regex on Strings vs. Bytes

Regex in the `re` module works on binary sequences when it's compiled from a binary sequence, not a string.

Regex matches differently depending on whether you use `str` or `bytes`, with the bytes version being more strict and the str version probably doing what you actually intended (matching way more than just ASCII).

### Encodings & Codecs

**% operator support** has been lost in Python 3.0 but should return in 3.5.

Python bundles >100 codecs for text-to-byte conversion, including cp1252 (Windows), utf-8, cp437, etc.

Most non-UTF codecs handle a small subset of Unicode and will throw a `UnicodeEncodeError` unless you pass an `errors` argument:
- `errors='ignore'` (skips unsupported chars)
- `errors='replace'` (replace with '?')

Pages 110 through 121 detail many pitfalls of Windows using/defaulting to cp1252. Basically, the motto is: **"Always specify the encoding"**.

### Unicode Normalization

Interesting discussion and demo on p.122 about normalization of characters with accents (and others).

P.124 discusses `.casefold()` which is similar to `.lower()` but converts some special Unicode chars like µ and ß, and is best used for case-insensitive comparisons.

P.125 demos writing a `nfc_equal()` function for string comparison, and `fold_equal()`.

P.126 demos shaving accents and combining chars for better searches.

### Sorting Unicode

```python
locale.strxfrm()  # Transforms a string into one that can be used in locale-aware comparisons
locale.setlocale(locale.LC_COLLATE, 'pt_BR.UTF-8')  # Set the locale to (e.g., Portuguese) before using
sorted_fruits = sorted(fruits, key=locale.strxfrm)
```

Manually setting `locale.setlocale` is not recommended. Your app or framework should set it at startup.

The syntax for the locale IDs is OS-dependent and more complex on Windows than *nix, like `"English_United States.850"`.

**NOTE:** Setting the locale may appear to work without error but the OS may not have a proper locale setup and your sorted strings may still be wrong.

There's a simpler solution: **PyUCA lib** by James Tauber, which is a pure Python implementation of the Unicode Collation Algorithm! (See p.131)

`unicodedata` module is basically the Unicode Database as a bunch of useful tools.

### Filenames

The Linux kernel is not Unicode savvy and filenames may not convert to `str`. So all `os` module functions accept paths as `str` or `bytes`.

`sys.getfilesystemencoding()` is used internally to convert between `str` and `bytes` filenames.

```python
os.listdir(b'.')  # Produces [b'abc.txt', b'x.txt']
```

In other words, it returns `bytes` if given `bytes`. There's `os.fsencode(filename)` and `os.fsdecode(filename)` to assist (p. 136).

---

## Chapter 5: First-Class Functions

### Functions as Objects

```python
def factorial(n): ...
fact = factorial  # Assign factorial to var fact
map(factorial, range(11))  # map "applies" one to the other and returns an iterable
```

`sorted` takes a function as a key which it applies to each element: `sorted(fruits, key=len)`. 

Others include `map`, `filter`, `reduce`, and `apply` (deprecated), but better alternatives are available for all their use-cases.

### Listcomp vs. map/filter

A listcomp or a genexp does the job of map and filter combined:

```python
[fact(n) for n in range(6)]  # More readable
[factorial(n) for n in range(6) if n % 2]  # And filter (p. 150)
```

### Lambda

The `lambda` keyword creates anonymous functions, but the body is limited to pure expressions:

```python
sorted(fruits, key=lambda word: word[::-1])  # Reverse chars
```

### Callable Objects

```python
callable(obj)  # Check if an object is callable
```

Callable objects include functions, methods, etc., but also a class instance with a `__call__` dunder can be called. Also, "generator" functions using `yield` (See chap 14 and 16).

```python
dir(func)  # Built-in that lists out all its dunders
# There's a __dict__ which can contain arbitrary values against the object
```

Example shows using `set(dir(a)) - set(dir(b))` to get the difference of a set! The function has several extra dunder funcs including `__code__`.

### Parameters: Positional, Keyword, Variadic

Can use `*` and `**` to explode iterables and mappings into separate args when calling a function:

```python
def tag(name, *content, cls=None, **attrs):
    # name – positional
    # *content – remaining positional args (iterable like params array in C#)
    # cls – keyword-only (must pass as name=value)
    # **attrs – keyword args (map of KV pairs) and also bind any that have matching names/keys to the named params
```

### Bobo Framework Example

```python
import bobo
@bobo.query('/')
def hello(person):
    return 'Hello %s!' % person
```

Note: bobo introspects the hello func to see person and extract it from the HTTP request. This is read from the code object found in `__code__`.

### The inspect Module

```python
from inspect import signature
sig = signature(x)  # where x is your function
```

The `sig` object has properties and methods to analyze and access the definition like .NET reflection.

```python
sig.bind(values)  # Will try to bind a map to the signature parameters using Python's rules
```

### Function Annotations (Python 3)

```python
def clip(text: str, max_len: int = 80) -> str:
    ...
```

⚠️ **No processing is done with the annotations!!** They are merely stored in `__annotations__` and may be useful for designers and IDEs, linters.

### Functional Programming Modules

**operator** and **functools** packages for functional programming:

```python
operator.mul  # Takes a and b and multiplies them! Woah!
itemgetter  # Just lambda fields: fields[1] (should be called getelement if you ask me)
# Though it can "destructure" many values into a tuple: tuple_data = itemgetter(1, 0, 4) (p. 165)
attrgetter('name', 'coord.lat')
```

`methodcaller` is cool and takes the name of the method to call and its args:

```python
hiphenate = methodcaller('replace', ' ', '-')
hiphenate('some value')  # 'some-value'
```

The `functools` module contains "higher-order" functions like `reduce`, but also `partial` and `partialmethod`:

```python
functools.partial(mul, 3)  # Partial application of a function: the arg 3 is fixed
triple = partial(mul, 3)
triple(7)  # 21
```

`partial` takes a callable first arg and any number of positional and keyword args. Python 3.4's `partialmethod` is designed to work "on methods".

---

## Chapter 6: Design Patterns with First-Class Functions

Page 183 shows a strategy pattern except it's stupid because the problem of deciding which strategy to apply is out of scope. But the strategies chose whether to apply themselves, which is bizarre as it spreads responsibility all over the place.

Then shows how each strategy class is really just a function runner and can be replaced by a function, which is obvious IMO.

It does eventually redress this by example of a meta strategy to pick the best available discount.

### Registration via globals()

```python
globals()  # Returns a dictionary representing the "current global symbol table"
promos = [globals()[name] for name in globals()
          if name.endswith('_promo')
          and name != 'best_promo']
```

Can also `import inspect` and run `inspect.getmembers(mymodule, inspect.isfunction)`.

---

## Chapter 7: Function Decorators & Closures

### What Are Decorators?

Function decorators let us 'mark' functions in source.

A decorator is a callable that takes another function as its argument: "the decorated function". The decorator may perform processing on the function and returns it or replaces it with another func or callable.

```python
@decorate
def target():
    print('running')

# Equivalent to:
target = decorate(target)
```

It simply wraps the func:

```python
def deco(func):
    def inner():
        print('foo')
    return inner
```

(Notes: defines a function inner inside the function deco and returns that.)

**Decorators are executed at "import time"** when a module is loaded!! This means they're cool for making "registration" systems, like route handles for web frameworks.

### Variable Scope Rules

Python doesn't require variables to be declared but assumes a var assigned in the body of a func to be local. Really odd:

```python
b = 6
def f2(a):
    print(a)
    print(b)
    b = 9  # This assignment makes b local!
```

This code works unless you remove the assignment causing an assignment in b to become local!! And this is "better than JS" because it prevents accidentally clobbering global vars!

### Closures & nonlocal

Closures are technically not the same as anonymous functions since closures mean the capture of referenced variables.

But because Python 2 assumes assignments to vars are local (makes them local), you cannot write to captured variables!! You must use `nonlocal` keyword:

```python
def outer():
    count = 0
    def inner():
        nonlocal count
        count += 1
```

Closures are essential for asynchronous programming with callbacks and functional-style coding.

### Standard Library Decorators

Three built-in: `property`, `classmethod`, `staticmethod` and `functools.wraps` is used for coding your own well-behaved decorators.

`lru_cache` and `singledispatch` are interesting.

### Memoization via lru_cache

```python
@functools.lru_cache()  # Note the () – it can take arguments
```

Memoization via `functools.lru_cache` which you simply decorate an expensive function with to get caching.

Since `lru_cache` can take arguments, it must have empty `()` parentheses to indicate you want to pass nothing to it. Weird!

The demo of `lru_cache` on recursive Fibonacci went from 17.7s to 0.0005s!!

### singledispatch (Function Overloading)

There's no method or function overloading in Python.

A common solution is to use verbose if/elif and dispatch to a set of specialized functions.

Not clear yet but it says the new Py 3.4 `functools.singledispatch` lets you decorate a plain function with `@singledispatch` and turn it into a generic function — a group of functions to perform the same operation in different ways depending on the type of the first argument.

Note – There's a "back-port" of singledispatch for 2.6-3.3 via a package. The book or AI should be referred to since the code sample is quite long but Gemini AI says this is the Pythonic way to do function overloading.

```python
from functools import singledispatch

@singledispatch
def htmlize(obj):
    # The base func is marked then its used as a decorator 
    # as the specialized func gets using special underscore syntax
    ...

@htmlize.register(str)
def _(text):
    # specialized implementation for strings
    ...
```

### Parameterized Decorators

Advice: use the base class e.g. `numbers.Integral` and `abc.MutableSequence` instead of concrete impls like `int` and `list`. The book recommends reading PEP 443.

To add args to a decorator your decorator func becomes a decorator function factory, returning a function definition:

```python
def myDecorator(args):
    def actualDecorator(func):
        # use args
        def wrapper(*vargs, **kwargs):
            return func(*vargs, **kwargs)
        return wrapper
    return actualDecorator

# Then use it:
@myDecorator(X)  # or @myDecorator()
def something():
    # etc.
    ...
```

---

## Chapter 8: Object References, Mutability & Recycling

Starts by reminding the reader that variables are labels, not boxes.

Python has `is` and `is not`.

The `is` operator compares identities while `==` compares values.

Mentions checking if a variable is bound to None but doesn't explain what this means.

`==` operator is sugar for `a.__eq__(b)`.

The default object impl of `__eq__` compares object IDs but you can override this.

Tuples are immutable but obviously their referenced values may not be.

Obviously, creating a new list copy via the ctor by passing in another list just shallow copies the values and refs.

There's a `copy` package with a `copy.deepcopy` func.

Function arguments seem to work like Java and C# and create new variables/labels in scope.

However, Tuples are copies (safe to change), but lists are not and are like reference types.

### Mutable Default Arguments (Pitfall!)

Do not use mutable default param/arg values like:

```python
def __init__(self, passengers=[]):
```

Weirdly this is legal but `[]` refers to "the default list" which seems to be a singleton attached to the class or something!!

**Technical Note:** In Python, `None` is a singleton. When you check if `x is None`, you are checking if the label `x` points to that one specific memory address where the None object lives.

Also, your intuition about the "singleton" behavior of `passengers=[]` is spot on—Python evaluates default arguments once at definition time, not every time the function is called. This means every instance of your class ends up sharing that exact same list object!

### del & Garbage Collection

The `del` statement deletes labels/refs, not objects. It frees the reference so Garbage Collection can kick in.

The `__del__` dunder method is called by Python just before GC and is tricky to implement properly.

**CPython uses reference counting.** CPython 2.0 has a generational GC algo somewhat like modern langs.

### Weak References

Weak references exist via `weakref` package:

```python
wref = weakref.ref(var_here)  # Returns a func/callable called wref which can access the referent via calling
wref()  # e.g. wref()
```

Obviously, if you reassign `var_here` to some other thing then `wref()` still points to the original thing but it may be GC'd so you need to expect `wref()` to point to None.

**Collections:**
- `WeakValueDictionary` – has its keys removed whenever the object is GC'd and is used for caching
- `WeakKeyDictionary` – similar
- `WeakSet` – useful if you make a class that must keep track of its own instances via a class attribute that's using a weakset so it doesn't block the GCing of its instances

**Note:** You cannot use a list or dictionary as a referent, but subclass them and you can.

---

## Chapter 9: A Pythonic Object

### String Representation Dunders

In addition to `__repr__` and `__str__`, there are two additional formatting dunder functions:
- `__bytes__()` – obviously byte sequence
- `__format__()` – for use with special formatting codes

The book will code up a Vector2D class and it implements dunder functions for init, iter, repr, str, bytes, eq, abs, bool.

### @classmethod & Alternative Constructors

Page 260 shows a strange function decorated with `@classmethod` which appears to be a way to attach static methods and is used here to create "an alternative constructor" or what I call a factory function.

It doesn't take-in `self` as an instance as its first parameter but takes in `cls` which seems to represent the "class itself" which is weird because the class is called Vector2D, so how does it know?!

The other weird thing is that it takes in a byte sequence `octets` but reads the first byte as a character defining the typecode (or encoding) which it uses in a cast expression:

```python
@classmethod
def frombytes(cls, octets):
    typecode = chr(octets[0])
    memv = memoryview(octets[1:]).cast(typecode)
    return cls(*memv)
```

Note that this func is defined within the indented class definition - so that's how it knows.

### @staticmethod

The `@staticmethod` decorator doesn't change how functions are called, so funcs should not take in `cls` as their first argument. This `@staticmethod` is simply used to define ordinary static functions.

The book doesn't explain what exactly `cls` is and why it takes in memoryview of bytes and somehow produces an instance of your class.

The author notes how he can't see the point of `@staticmethod` and sees static methods defined 'on' a class as pointless when you can just define a 'floating' function in the module. I think it's to help convey strong cohesion.

### __format__ Dunder

The `__format__` dunder func takes in a `format_spec` which is e.g. `'1 BRL = {rate:0.2f} USD'.format(rate=brl)` where `rate` is the class inst and `0.2f` in the curlies is the format_spec.

The book says there's a Format Specification Mini Language but obviously your actual impl of `__format__` can do whatever it likes.

### Private Attributes & Name Mangling

Page 266, as part of implementing hashing shows that to make a property private you prefix it with two underscores and refer to your private vars from a defined func which act serves as a getter function:

```python
@property
def x(self):
    return self.__x
```

The book calls out that for classes which have a sensible scalar value, you can implement `__int__` and `__float__` which are used by `int()` and `float()` constructors during type coercion.

P272: There's no way to create truly private variables. All there is is a way to prevent accidental overwrites of a "private" attribute in a subclass (I think he's referring to protected).

If you name your instance attribute with two underscore prefix then it'll be stored in the instance `__dict__` with a leading underscore then the class name, so `_Dog__mood` and a subclass with its own mood will be `_Beagle__mood` so the subclass one will not clobber the original one.

He says many Python coders prefer to use single underscore prefixes which do not get auto "mangled" and use a 'fully qualified' name like `_Dog__mood` which is a popular convention.

My own thoughts are to use double and rely on mangling but it's another example of how Python is a really shit language.

### __slots__

Because Python stores all instance attributes in the classes `__dict__` each class has an overhead of that hashtable! When dealing with millions of instances with few instance attribs then consider using `__slots__`:

```python
class Vector2D:
    __slots__ = ('__x', '__y')
```

Book notes that anyone dealing with millions of objects of numeric data should be using Numpy arrays. The Vector2d class is just a tangible demo, use Numpy or Scipy.

He shows the `__dict__` and `__slots__` versions with 10 million instances and the first one consumes 1.5Gb of RAM! The slots version consumes "only" 0.65Gb of RAM!!

---

## Chapter 12: Inheritance

First fascinating thing to learn is that Python supports **multiple inheritance**!!

Chapter starts with the pitfalls of subclassing 'built-ins'. Since the code of 'builtins' like 'list' or 'dict' is written in C your special methods overridden are NOT called.

E.g., an overridden `__getitem__` is never called by the built-in `get()`.

Interestingly in the example on page 362 the `[]` operator does call overidden `__setitem__` for some weird and random reason.

### The Built-in Subclassing Problem

Your notes perfectly capture the core frustrating reality of subclassing C-level built-ins in Python. Here is the breakdown of why those "weird and random" behaviors happen and how to structurally fix them.

**The C-implemented built-ins (dict, list, str) bypass overridden special methods for speed.** They look for methods directly on their own C-structs rather than looking up the Python method resolution order (MRO).

Why `get()` ignores `__getitem__`: The `dict.get` method is optimized in C. It does not call `self[key]` (which triggers `__getitem__`); it searches the internal hash table directly.

Why `[]` worked for `__setitem__`: If you were using an operator like `obj[key] = value`, Python's evaluation loop forces the call to the underlying slot handler. However, if a C-level method inside dict (like `update()`) is called, your overridden `__setitem__` will be ignored there too.

### The Solution: UserDict, UserList, UserString

To fix this behavior completely, you should not subclass `dict` or `list` directly. Instead, Python provides the `collections` module specifically to allow safe subclassing with expected Python-level behaviors.

Instead of `dict`, subclass `collections.UserDict`  
Instead of `list`, subclass `collections.UserList`  
Instead of `str`, subclass `collections.UserString`

**The Broken Way (Direct Subclassing):**
```python
class BadDict(dict):
    def __setitem__(self, key, value):
        super().__setitem__(key, [value] * 2)  # Duplicate values

d = BadDict(a=1)  # __init__ ignores your __setitem__!
print(d)  # {'a': 1}

d.update(b=2)  # update() ignores your __setitem__!
print(d)  # {'a': 1, 'b': 2}
```

**The Correct Way (Using UserDict):**
```python
from collections import UserDict
class GoodDict(UserDict):
    def __setitem__(self, key, value):
        super().__setitem__(key, [value] * 2)

d = GoodDict(a=1)  # Works!
print(d)  # {'a': [1, 1]}

d.update(b=2)  # Works!
print(d)  # {'a': [1, 1], 'b': [2, 2]}
```

### MRO & Multiple Inheritance

Fascinating: From the REPL call `bool.__mro__` and you'll see the `bool -> int -> object`

Author says that multiple inheritance is rare and mostly use with abstract base classes in `collections.abc` because ABCs are like interfaces and multiple inheritance of interfaces is fine and dandy.

Interestingly, it talks about the standard stuff for when to correctly use inheritance and talks about "Mixins" as a thing, but neglects to explain what they are exactly, as if the reader ought to know. Sounds to me like they're "static" classes of "static" functions?!

**Note - a Mixin is a bundle of functionality on a base class which is not intended to stand alone or be instantiated but offer a way for multiple-inheritance languages like Python to allow coders to "pull-in" sets of behaviours.**

Mixins would not leave methods abstract; for the subclass dev to "fill-in".

Recommends not to subclass from more than a single concrete class.

The remainder of the chapter discusses frameworks and the pros and cons of inheritance. All standard OO stuff.

---

## Chapter 13: Operator Overloading

Interoperate with "infix" operators like `+` and `*` or `-` and `~`.

Note that `()` calling and `.` and `[]` are also operators in Python but chapter covers unary & infix.

**Limitations:**
- Cannot overload operators for built-in types.
- Cannot create new operators, only overload.
- Cannot overload `is`, `and`, `or`, `not`.

### Unary Operators

It's easy to do unary operators; implement the special `__neg__` or `__pos__` or `__invert__` etc.

### Binary Operators

Book notes that 'sequence types' should `+` and `*` for concat and repetition but e.g. Vector type despite being sequence-like makes most sense to overload these ops as mathematical ops.

During sample code of `__add__(self, other)` the code makes use of a very interesting looking `itertools` package which looks like it contains helpful stuff somewhat like C# Linq.

Even more interesting is that the book says of a `itertools.zip_longest(...)` that it produces a "generator" which implies, like with Linq, that it is not a materialized sequence perhaps.

The above code on page 390 has `__add__` return a new instance of the type, which makes sense.

### Reflected Operators & NotImplemented

Python has a special dispatch system for calling overloads. It'll first look for `__add__` and then try `__radd__`, effectively trying the reverse/switched operands!

So I guess all the operator overload dunder functions get a "right handed" version.

You can return a special `NotImplemented` to indicate to the dispatcher, but don't get confused with `NotImplementedError`.

Implementing the right hand operator is often just delegating to the existing operator, which in the book is done via e.g. `+` rather than calling the dunder "manually":

```python
def __radd__(self, other):
    return self + other
```

Book mentions some ugly, unhelpful stock error messages when e.g. types in the iterable cannot be added to the receiving type.

Do implement the overload and return `NotImplemented` if you cannot handle the inbound type(s) because it "leaves the door open" to others to implement it using the reverse version.

To do this you can simply place your left-handed implementation inside a try, catch and return `NotImplemented` from the catch.

For e.g. scalar multiplication `__mul__` and `__rmul__` you can use `isinstance(..)` and use the `numbers.Real` ABC to check for a scalar (scalable, multiplyable).

Page 397 has a table of operators and their dunders.

Book says Python 3.4 will have a new `@` op for matrix multiplication! See PEP 465.

### Rich Comparison Operators

Handling of "Rich Comparison Operators" such as `==`, `!=`, `>`, `>=`, etc. is the same except that the reverse or right-hand version is often just the logical reverse, e.g., `__gt__` is `__lt__`.

Final Vector `__eq__` implementation is rather like what I'd do in C# and returns `NotImplemented` for anything non-Vector and then compares the length and does a traditional long-winded for loop over each value using `all(...)`.

### Augmented Assignment

"Augmented Assignment" operators like `+=` will work "for free" but you can overload them with e.g. `__iadd__` but the special method must return `self`.

In the further reading section he says it's a better design to use the try-catch approach and allow for dynamic typing which seems to contradict what he said about his Vector comparing true against a tuple of ints!

Notes how `__ne__` usually does not need to be implemented; only `__eq__`.

---

## Chapter 14: Iterables, Iterators & Generators

The `yield` keyword was added in v2.2 and allows the construction of generators which work as iterators.

First example says you want a list from a `range()` you must build one: `list(range())` which tells me generators are like .NETs enumerable.

First demo class Sentence implements the "sequence protocol" which is iterable.

Sentence implements `__getitem__` which is actually a fallback which does work but you should actually implement `__iter__`.

When your class implements `__iter__` it becomes iterable and weirdly will become a "goose" type `abc.Iterable` i.e. `issubclass(x, abc.Iterable)` will return true!!

Book recommends not bothering to check for iterable because the error thrown if it is not is clear enough.

Any object that returns an iterator from its `__iter__` function is iterable. Having `__getitem__` means it is a sequence, not strictly iterable.

Page 420 shows how to manually create and use a iterable using `iter(...)` and a while loop and guarding for (try-except) `StopIteration`.

The standard "interface" for an iterator has two methods: `__next__` and `__iter__`.

The `__iter__` function simply returns `self` and so an iterator can be used whenever a iterable is required.

To manually get the next item use `next(iterable)`.

It seems the weird "goose typing" is some trick of implementing `__subclasshook__` which I assume looks specifically for the `__next__` and `__iter__` dunders.

Weird. The book doesn't show how to implement `__next__` so I have no clue where the current pos state machine is defined!

Oh okay now it's going into detail and shows a classic stateful iterator pattern from GoF which it points out is not idiomatic Python. p424

**Woah! Important clarification:** iterables have a `__iter__` method that instantiates a new iterator every time while iterators implement `__next__` and `__iter__` which returns self.

Next section appears to discuss the final proper way to do it.

### Generator Functions (The Pythonic Way)

A Pythonic implementation uses a generator function to make an iterator. This uses the special `yield` keyword:

```python
def __iter__(self):
    for word in self.words:
        yield word
    return
```

It's essentially a factory which produces a generator object which wraps the function body and has `__next__` and `__iter__`.

Book mentions lazy and eager evaluations of the items, the usual stuff.

We can further advance by using a "generator expression" which is a lazy version of a "list comprehension". This is just like in C# .NET Linq:

```python
def __iter__(self):
    return (match.group() for match in RE.finditer(self.text))
```

### itertools & Generators

The `itertools.*` module has 19 generator functions for all manner of funky shit.

These appear composable like .NET Linq, for example:

```python
itertools.takewhile(
    lambda n: n < 3, 
    itertools.count(1, .5)
)
```

Page 439 begins a section on generator functions in the standard lib for filtering, selecting, ordering etc., again I assume like .NET Linq.

Some are "builtin" and need no qualification, like `filter(predicate, it)` and `enumerate(...)` and `map(func, it1, [it2, ...])` while the others seem all to be in the `itertools` module.

The book shows two tables listing them all and is not worth writing down.

### Reducing Functions

Page 450 lists `all(it)`, `any(it)` and then some more "reducing" funcs:
- `max(it, [key=], [default=])`
- `min(...)`
- `functools.reduce(func, it, [initial])`
- `sum(...)` (use `math.fsum` for floats)

**Cool thing:** the `iter()` function can take a callable like a function and a sentinel value which when encountered stops iteration. One obvious use is to read lines from a file until a blank line.

---

## Chapter 15: Context Managers & Else Blocks

Seems quite interesting and about a powerful `with` statement which appears to "do teardown".

Also seems you can use an `else` clause in a `for`, `while`, `try` block.

The `else` feature is not related to `with` at all but the author had to put it somewhere.

### else Blocks

The term `else` may hinder your understanding of the feature:

**for/else** - the else block runs upon loop completion/exhaustion, but not on break.

**while/else** - runs when the while condition becomes false, but not on break.

**try/else** - okay, weird, the else block runs if there's no exception, and is outside try so any errors go uncaught. Presumably the context of the block is maintained.

`return`, `break`, `continue` or exception causing an early loop exit will not trigger `else`.

Perhaps exceptions aren't all that expensive in Python so they use try/except more often as a kind of control flow, optimistic attempt.

### Context Managers

It's moved on to `with` now and is talking about a "context manager" object with special lifetime dunder funcs `__enter__` and `__exit__`.

Okay so I think it's a bit like using an `IDisposable` in C#, except the context manager arrives as a result of calling `with (x)` but the object bound via `as` is the result of the `__enter__` function!

The `as` clause of the `with` statement is optional. The `open` statement will always need an `as` clause in order to get a reference to the opened file, but some context managers return `None` because they have nothing useful to return.

The `__enter__` dunder takes in nothing other than `self`, but the `__exit__` function takes in `self`, `exc_type`, `exc_value`, `traceback` where `exc_type` is an optional exception class, `exc_value` is an optional exception instance, and `traceback` is a traceback object. `None` is passed when the block exits successfully.

Book says these context managers are new and that uses are (obviously) for holding locks, transactions, etc. Same as C# using `IDisposable`.

### contextlib

The `contextlib` module contains useful ones including:
- `closing` – a function that'll make a context manager from any object that has a `close()` method
- `suppress` – for ignoring specific exceptions
- `ExitStack` – lets you "stack" contexts and exit them all at the end

---

## Chapter 16: Coroutines

Also uses `yield`, but on the right side of an expression:

```python
datum = yield
```

It may or may not produce a value and yields `None` if there's no expression after the `yield` keyword.

The coroutine may receive data from the caller! Usually, a caller pushes data/values into the coroutine, so it's a control flow device often used for cooperative multitasking.

Coroutines evolved from generators. The book is hard to follow, discussing the history of this and discussing concepts not yet introduced.

### Simple Coroutine Example

Page 481 finally shows:

```python
def simple_coroutine():
    print('started!')
    x = yield
    print('received:', x)

my_coro = simple_coroutine()
next(my_coro)  # 'started!'
my_coro.send(42)  # 'received: 42'
# Exception! StopIteration
```

The exception is normal generator behavior for reaching the end of the flow/body.

### Coroutine States & Priming

Coroutines can be in one of four states: `GEN_CREATED`, `GEN_RUNNING`, `GEN_SUSPENDED`, `GEN_CLOSED`.

You must 'start' the coroutine with either `next(c)` or `c.send(None)` before sending data into it.

### Bi-directional Communication

It seems (p.483) that the coroutine is bi-directional as a value after `yield` is returned to the caller, then it waits for a value to be sent back.

The coroutine may have more than one `yield`.

P.485 shows a clever looping coroutine which computes a continuous running average: value in, average out. You must call `c.close()`.

P.486 shows how to write a decorator which auto-calls `next()` to prime the coroutine, then you define your coroutine and decorate it.

### Throw & Close

You can throw an exception into the coroutine from the outside?

Interestingly, sending a string into the "average" coroutine causes it to throw a `TypeError` but it doesn't seem to crash the app (unless because the demo is within the REPL) but it does close the coroutine so it cannot accept new values.

To discuss: `throw` and `close`. It flips back to generators, e.g., `generator.throw(exc_type[, exc_value[, traceback]])`.

The coroutine could handle the exception, so it'd then advance to the next `yield` and any value sent will be returned/accessible via `.throw()` result!

`close()` injects a `GeneratorExit` but you don't seem to need to handle it and it doesn't seem to actually raise/bubble it up?!

### Cleanup & Early Exit

If you need to run some mandatory cleanup, then place it in a try-finally block, naturally.

P. 492 shows a version of an averager which doesn't yield a running value, but rather returns it at the "end"—the end being the coroutine flow ending "naturally." Thus, it has a break in the loop if the caller sends in `None`.

However, this design causes the original `StopIteration` to be thrown due to the "generator" flow/body ending. However, the value attribute of the exception contains the return value.

### yield from & Sub-generators

`yield from` seems to individually yield the items in an iterable. For example: `yield from range(1,3)`.

It seems `yield from` can also yield from a sub-generator, delegating down through the current generator/coroutine and doing I/O with a coroutine within a coroutine!

P. 496 shows some wild code using a simple "grouper" delegating generator which delegates to the averager via a grouping function and is almost unintelligible. Subsequent pages are devoted to explaining the logic.

The PEP 380 inventor notes that the effect of sub-generators is as if the sub-generator were inlined into the code where the `yield from` statement occurs. (P. 10)

### Discrete Event Simulation

P. 506 - I've skipped over a ton of complex discussion around generators and sub-generators because I think it's tedious and my notes will never fully capture it.

Apparently, running simulations is a classic use case for coroutines. The first OO language, "Simula," used them.

To explain "Discrete Event Simulations," the book draws an analogy from games: a turn-based game has a discontinuous "clock" and jumps forward in time when there is an event, while continuous-clock simulations do not and are real-time.

SimPy is a DES package for Python that uses one coroutine for each "process" in the simulation.

The book discusses a taxi simulator prowling the streets, picking up passengers, and dropping them off.

It seems to me that the coroutine/generator function serves as a little taxi state machine and is advanced to its next state by being sent the simulation master time. Remember that a coroutine is a kind of swapping procedure; I give you this, you do something, and you give me something in return.

The most confusing thing is how the value sent into the coroutine ends up in the time variable. Perhaps the initial setting `time = yield ...` hooks it up "permanently."

P. 512 - Not sure of the need to take notes on this, but the example is pretty nice in that it's a great way to show the practical use of coroutines as state machine "actors".

On p. 513 is a Simulator class and its `__init__` constructor setting up the 'events' and 'procs' attributes. (I think this is the proper name; they're "props" or "fields" to me!!), and it very much resembles a C# or any other OO language class constructor.

P. 514 has quite a dense little clip of what I would say is "normal" application code, and it reads much like anything else except that it's Python. Note that we have still not covered arranging and organizing code yet.

The author points out the use of `else` for the outer `while` loop, which is indeed useful for knowing when the main while condition has been broken, as opposed to encountering a break.

### Use Cases

There are generally three kinds of uses for coroutines: pull-style iterators, push-style aggregators, and tasks.

The chapter ends by calling out David Beazley's writings and tutorials on Python coroutines.

Coroutines lie at the heart of `asyncio`, too.

**Chapter Summary Note:** I don't normally dwell on the end-of-chapter blurb, but it feels like coroutines are complex, powerful, and different from anything I've seen in other languages. It mentions that Brett Slatkin's *"Effective Python"* shows Conway's Game of Life implemented using coroutines, and that the code is available on GitHub at github.com/bslatkin/effectivepython.

**Resource Note:** The end of the chapter also calls out good resources for Discrete Event Simulations, which seems like a subdiscipline that hiring managers might think is cool or esoteric.

---

## Chapter 17: Concurrency with Futures

### Library Support

Focuses on the `concurrent.futures` library introduced in Python 3.2, which is also available via the `futures` package for Python 2.5+ on PyPI.

**Language Comparison:** C# has Tasks, and it seems Python uses the term Futures for its object. This idea is the foundation for `concurrent.futures` and also for the `asyncio` package covered in Chapter 18.

### Simple Sequential Downloader

**Code Implementation (p. 526):** Page 526 shows a simple sequential downloader script/program. The most interesting thing is how this exemplifies writing a short executable program and how to invoke main or write an entry point:

```python
if __name__ == '__main__':
    main(download_many)
```

...where `download_many` is a function reference passed into `main`.

**Dependency Note:** It also uses the improved `requests` library by Kenneth Reitz instead of the standard `urllib.request`.

### Futures Overview

As of Python 3.4, there are **two** classes named Future: `concurrent.futures.Future` and `asyncio.Future`. Weirdly, they both serve the same purpose: to be like a Task in C# or Promise in JavaScript.

Unlike a C# Task, a Python Future should not be instantiated by the developer.

They are meant to be created by your concurrency framework, for example `concurrent.futures.Executor` via its `.submit()` function.

Both futures have:
- `.done()` method which returns a boolean
- `.add_done_callback()` method
- `.result()` method (behaves differently depending on whether the Future is complete or not; in concurrent.futures it will block, while the asyncio version is designed to be called via yield from)
- Obviously, if complete, it returns the result, or it throws whatever exception had occurred.

### as_completed() vs executor.map()

P531 is a pretty familiar page of code to any C# developer having kicked off a series of concurrent tasks and waited for them all to complete. The `as_completed()` function is nice and only yields futures as they complete, presumably blocking until one becomes available.

**Bombshell! None of the concurrent implementation demos are actually downloading in parallel!!!**

The speedup of 5x is real, however, there is a "Global Interpreter Lock" (GIL) that only lets one thread run at a time!

The CPython interpreter is not thread-safe!!!

However, during downloads and I/O, the C code releases the GIL and so Python can continue executing on its single thread. So it seems the future concept really represents asynchronous I/O; so on Windows, that's I/O Completion Ports, and on Linux, that's epoll and io_uring.

However, there is `ProcessPoolExecutor` which is like Node `worker_threads` where the worker runs in a dedicated runtime, I think.

Real OS threads are scheduled, but the GIL prevents any Python bytecode from running in parallel; but with `ProcessPoolExecutor`, it's for real concurrent CPU-bound work.

The book shows a 4x speedup from doing SHA256 hashing using 4 workers on a 4-core machine.

Interestingly, it gets a 3.8x speedup using PyPy, which is an alternative to CPython (which also uses a single-threaded GIL!?), and it mentions that other Python runtimes like IronPython **are** truly parallel.

P536 shows using the useful `executor.map()` to fire off general callables:

```python
executor.map(do_work, range(5))
```

Where it'll iterate the range and push the value into the callable, and return the whole batch of futures.

```python
futures = executor.map(..., …)
```

### executor.map vs as_completed

The book hints that `executor.map(...)` always returns futures in the order they were started, meaning it blocks until the last one finishes. `executor.submit()` combined with `futures.as_completed()` is generally much more flexible.

### Synchronization Primitives

**Appendix Material (Pages 539–540):** Discusses several Python scripts using futures, though a lot of it is just noise regarding progress bar libraries.

**Synchronization Primitives (Page 550):** The book mentions that Python has `Thread`, `Lock`, `Semaphore`, and thread-safe queues in the `queue` module.

> *Critique:* This feels odd—are these meant for use with `ProcessPoolExecutor`? The text then suggests using `Process...`, almost implying these synchronization primitives are meant for a *single-threaded* `ThreadPoolExecutor`?!

### Ecosystem & Multiprocessing

**Ecosystem & Multiprocessing (Page 552):**
- There is plenty more Python can do with concurrent programming, but the author chooses to skip it to focus heavily on `asyncio`.
- The end-of-chapter notes highlight **PEP 371** and the `multiprocessing` package (locks, queues, pipes, shared memory), which forms the actual basis for `concurrent.futures.ProcessPoolExecutor`.
- The book also calls out Apache Spark's distributed computing engine and its Python API.

**Author Perspective:** It feels like low-level concurrent programming might be outside the author's primary skill set or interest. He explicitly notes that he likes **Elixir** and **Go**!

---

## Chapter 18: Concurrency with asyncio

### Background & Versions

**Origin:** The standard `"Tulip"` package became the `asyncio` package in Python 3.4. It is strictly incompatible with Python 3.2 and below.

**Backports:** `Trollius` exists as a backport of `asyncio` for Python ≥ 2.6.

**The Reality of Async:** It is crucial to remember that `asyncio` is **not true multithreading**—it is syntax for efficient async I/O management. The book illustrates this early on by comparing a command-line spinner written using `threading.Thread` against one written via `asyncio`.

### Coroutine Mechanics (Python 3.4 Era)

**Rules for Coroutines:** To be compatible with `asyncio`, a coroutine must use `yield from` in its body (not a bare `yield`). It must be driven by a caller invoking it through `yield from` or by passing it to functions like `asyncio.async(...)`, and it must have the `@asyncio.coroutine` decorator applied.

**The Spinner Example:** The `asyncio` spinner loops through the character symbols (`/-\|`) to print them, but crucially yields control using `yield from asyncio.sleep(0.1)` inside a `try/except` block that catches `asyncio.CancelledError`.

**C# / .NET Analogies:**
- `@asyncio.coroutine` ↔ Like `async` in C#.
- `yield from` ↔ Like `await` in C#.
- Every calling function up the chain must use these patterns.

### The Application Entry Point

The top of the call stack in `main` typically handles the event loop lifecycle manually:

```python
import asyncio

def main():
    loop = asyncio.get_event_loop()
    result = loop.run_until_complete(my_func())
    loop.close()
    print('Answer:', result)
```

### Key Gotchas

**Blocking the Loop:** You must use `yield from asyncio.sleep(...)`. Using an unawaited, standard `time.sleep(...)` will completely freeze the event loop and effectively pause your entire application.

**The Decorator:** While `@asyncio.coroutine` isn't strictly mandatory, it is highly recommended because it triggers a warning if the coroutine is Garbage Collected before being yielded from.

**Terminology (Python vs. C#):**
- Python uses the name `Task` similarly to how C# uses `Task` as a "sort of like a Thread" abstraction.
- `asyncio.async(...)` acts very similarly to C#'s `Task.Run`.

**Cancellation Safety:** It is safe to call `.cancel()` on a Task/coroutine and handle the `CancelledError`. Because Python's async runs under a global lock, it can only interrupt execution and raise a cancellation at an explicit `yield` line—never halfway through synchronous execution logic.

### The Future vs. Task Confusion

> **WTF Moment:** Python features both `asyncio.Future` and `concurrent.futures.Future`, and they are **not** compatible with one another.
> **Resolution:** So, what is a `Task`? A `Task` is simply a subclass of `asyncio.Future` specifically designed to wrap a coroutine.

**Extracting Results:** The text reiterates that `asyncio.result()` must be called via `yield from`. However, you rarely need to call `.result()` directly, because you are simply supposed to assignment-unwrap it:

```python
result = yield from my_future
```

### Practical Demos & Architecture

**Entry Point Differences:** The architecture is so close to .NET's async-await framework that it's barely worth writing down. The primary difference is that Python's entry point (`main`) is not natively an async coroutine. You have to bootstrap the entire execution chain manually using `asyncio.async(coro_or_future)` or `BaseEventLoop.create_task(coro)` (introduced in Python 3.4.2+).

**Protocol Support:** As of Python 3.4, `asyncio` natively supports only TCP and UDP. For HTTP async capability, an external library like `aiohttp` is mandatory.

### Script Demo Analysis (Page 570)

The book's script closely mimics standard async/await patterns. The core entry point function looks like this:

```python
def download_many(cc_list):
    loop = asyncio.get_event_loop()
    # Schedule all coroutines
    to_dos = [download_one(cc) for cc in sorted(cc_list)]
    wait_coro = asyncio.wait(to_dos)
    
    # Run loop until everything finishes
    res, _ = loop.run_until_complete(wait_coro)
    loop.close()
    return len(res)
```

**API Observations:**
- `asyncio.wait` behaves exactly like `Task.WaitAll()` in C#/.NET.
- **Event Loop Ownership:** You explicitly fetch the event loop because external frameworks might "own" its runtime lifetime.
- **Naming Quirks:** The variable `res` returns a set of futures (classic terrible naming).
- **The Discard Variable (`_`):** The second variable in the assignment tuple contains a set of *incomplete* futures, which only populates if you pass a `timeout` argument to `asyncio.wait`.
- **Page 575 Caveat:** I'm still skeptical if this is truly asynchronous; no other substantial work can step in when the execution thread is released by an underlying mechanism like an I/O Completion Port (IOCP).

### Shared State & Concurrency Bugs

Okay, so after chatting to the AI, my intuition is correct that asyncio depends upon the program being designed around a loop.

With a loop and several coroutines running concurrently, but not at the exact same time, you can still get bugs around shared values **"suddenly"** changing. This happens when one coroutine yields due to I/O, and some other coroutine makes some forward progress and also changes the value!

The book doesn't make super clear the importance of needing a loop due to its single-threaded nature.

At the time of publication, asyncio does not provide an asynchronous filesystem API.

P578 uses a `yield from` semaphore to limit concurrent HTTP requests to the same server:

```python
async with sem:
    await some_io()
```

The semaphore is used in a `with` block, which is like a `using` statement in C#, remember!

It has `.acquire()` and `.release()` functions, but for some reason, they are not used in the `with` code block.

These pages are essentially the author explaining the code reorganization needed to rewrite his downloader and spinner example around coroutines/async.

P582 discusses using `loop.run_in_executor(...)` to schedule the local filesystem save (synchronous) on the "thread pool," but this makes no sense to me because surely if Python is single-threaded, even with its "pool," then...?!

**Correction:** `loop.run_in_executor(...)`

Not explained in the book is that CPython will drop the GIL (Global Interpreter Lock) just before making the C call to the I/O. So, if this happens on a different "thread," then it'll yield and allow other "threads" to execute in the meantime.

P592 shows a useful TCP socket server, but the code gets the event loop, starts the TCP server with `asyncio.start_server(...)`, and then weirdly calls `loop.run_until_complete()` but *also* then calls `loop.run_forever()`!? Unless the first call drains anything before starting "properly." It is not explained.

The next server demo uses the `aiohttp` library again because it also supports running a little web server.

You pass the event loop into the library and create a web application like so:

```python
app = web.Application(loop=loop)
app.router.add_route('GET', '/', home)
handler = app.make_handler()
server = yield from loop.create_server(handler, address, port)
```

There's more to it **than** the above.

See p596, as **it's** a little odd/weird. Perhaps I am misunderstanding what `loop.run_until_complete(...)` does?

Perhaps these setup functions are themselves coroutines, so they need to be run asynchronously until completion just to start. That requires them to be executed in the event loop, because in Python, all async functions/coroutines must be run in the loop.

### Route Handlers

The 'home' function which handles the route doesn't need to be async/coro.

Obviously when making I/O from a handler which is most of the time you need to use async else what's the point?!

But also, for Python, any long CPU bound work must be artificially split into async chunks else it'll hog the single thread!!

---

## Chapter 19: Dynamic Attributes & Properties

### Attributes & Methods

Data attributes and methods are known as attributes in Python (as opposed to members?) and a function is just a callable attribute.

Oh! It says you can create properties which replace the (field-like) data attribs with accessors.

Weirdly has the opposite principle to e.g. C# where the dot notation does not betray whether it's using storage or computation!!

Python calls `__getattr__` and `__setattr__` which do whatever logic including virtual attributes (maybe like C# dynamic).

P610 shows accessing loaded JSON via indexer type syntax:

```python
feed['schedule']['speakers'][-1]['name']
```

P612 shows small class which implements `__init__` and `__getattr__` and recursively makes and returns this small wrapper class via a static constructor. This allows `feed.schedule.speakers...`

P614 shows a solution to JSON fields/keys which are reserved keywords in Python and appends `_` on the end of them like `feed.kids.jane.class_`

Wow okay so `__init__` gets called the constructor by some devs but actually it is `__new__` is the real constructor! It must return an instance.

**WTF:** `__new__` can return an instance of another class (in which case `__init__` is not called), and so now the 'build' static constructor in his example can just construct/return a new FrozenJSON class "as per normal" but it may not always be a FrozenJSON — sounds cool, probably stupid.

P619 the `__init__` has a cool hack to turn keyword args/dictionary of KV's into object attribs!

P619 - the code is unintelligible due to its deplorable variable names. Gemini does a great job of reprinting the code with clear names.

It's really just recursing through the JSON items using for loops and adding stuff to a "database" provided by a lib called "shelve".

P625 - this section much more clearly shows how to really build object models using Python even if the var names are BAD!

### Properties & Validation

P628 - Using `@property` decorator for read/writeable properties and validation.

Properties are just classes (which are themselves just abstractions over dictionaries!) for example:

```python
class LineItem:
    def ... # define getter and setter
    some_prop = property(getter, setter)
```

Book says that this explicit creation of a property via a class above is sometimes preferred to using the `@property` decorator.

But the presence of a property in a class affects how attributes in instances of the class can be found in a surprising way...

**Fuck! I think I've misunderstood Python syntax all the way through —**

```python
class Class:
    data = 'something'
```

`data` is not an instance field like in C++ but a public static string in C# parlance!!

Setting `my_class_instance.data = 'poo'` will then append a new instance attrib which shadows out the static data attribute!

P633 shows the "surprising" way that trying to set/overwrite a property works — it doesn't allow `obj.prop = 'foo'` which is expected/good to my eyes but perhaps if you're used to the "shadowing statics" behaviour previously explained then you may find this "surprising"?!

P633 the code then shows how overwriting class property destroys the Property and reveals the underlying dict class attribute called prop — it's worth a look at p633 carefully because it really is weird.

P634 shows the opposite where a class has an instance attrib which is then shadowed by sticking a new property on it and then deleting it again, unshadows the property. Wild usage of `del` to delete the prop: `del Class.my_prop`

Essentially `obj.something` doesn't start searching in obj instance attributes but instead from `obj.__class__`, then attributes.

Every Python code unit (modules, functions, classes, methods) can have a "DocString".

When console `help()` or IDEs need to access property documentation, they get it from `__doc__` attribute of the property.

You set the doc string like so:

```python
weight = property(getter, setter, doc="Here!")
```

Or stick it inside the getter func marked with the `@property` decorator.

Example worth reading because the setter decorator is not intuitive; see `@bar.setter`:

```python
class Foo:
    @property
    def bar(self):
        """The bar attribute."""
        return self.__dict__['bar']

    @bar.setter
    def bar(self, value):
        self.__dict__['bar'] = value
```

### Property Factory

Okay so the author returns to the problem of wanting to protect/guard against bad values being set.

For this, he turns to making a "property factory" and in particular a quantity property factory.

**Fascinating:** so the idea is that quantities often should not be zero or negative so we may as well encode this rule once.

The property factory is a function defined in the class (perhaps?!) which defines two further functions inside itself for the getter and setter (with the rules in the setter obvs) and then returns a new property. See ex:

```python
def quantity(storage_name):
    def qty_getter(instance):
        return instance.__dict__[storage_name]
    
    def qty_setter(instance, value):
        if value > 0:
            instance.__dict__[storage_name] = value
        else:
            raise ValueError('value must be > 0')
            
    return property(qty_getter, qty_setter)
```

### Attribute Access Overview

The last part of the chapter from page 640 lays out a bunch of "essential attributes" and "functions for attribute handling".

These are: `__class__`, `__dict__`, `__slots__`; and functions include: `dir(...)`, `getattr(...)`, `hasattr(...)`, `setattr(...)`, and `vars(...)`.

Then special methods for attribute handling: `__delattr__(...)`, `__dir__(...)`, `__getattr__(...)`, `__getattribute__(...)`, `__setattr__(...)`.

Chapter further reading hints at how subclasses might find it tricky to override the getter and setter logic in its superclass.

In "Soapbox" section the author says "I shouldn't have to care whether coconut.price fetches an attrib or performs a computation" — wrong!?!

---

## Chapter 20: Attribute Descriptors

### What Are Descriptors?

**"Descriptors"** seem to be used by ORMs to map between db fields and Python attribs.

A descriptor is a **class** that implements a protocol of `__get__`, `__set__` and `__delete__`.

The `property` class (we just covered?!) implements the full descriptor protocol.

Descriptors are a distinguishing feature of Python! Apparently methods (themselves) and the `classmethod` and `staticmethod` decorators use them.

### Descriptor Example

The author sticks with the quantity property factory and turns it into a Quantity descriptor class, since a descriptor is a 'class which implements a contract'; `__get__`, `__set__`, `__delete__`.

You **use** a descriptor by declaring instances of it as class attributes of another class (!?).

Through to page 655 the author makes a huge meal of explaining it. I think to a beginner coder this might be tough but it's really very simple — it's the same as the quantity factory function from before except the factory func is just a class and I guess, obviously, you therefore need to initialise it with the storage key name and... the instance of the class that owns it (since it's no longer inside that class).

The next 'version' of the descriptor class no longer takes in a key name but creates a "random" key by incrementing a counter and prefixing "Quantity#", though I don't see a problem with passing in a key — perhaps potential bug when a subclass defines another using the same key?

"Owner" (my term) above is misleading — in this pattern "owner" is a recognised term meaning the class that owns it, as opposed to the instance.

P658 recommends updating the `__get__(...)` implementation to return `self` when instance is None so that `MyClass.weight` returns the descriptor when called instead of throwing.

Interestingly the `__set__(...)` implementation for when instance is None is undefined!?

Author points out the apparent effort of all this but does explain that your "Quantity" is highly reusable and should be placed in a good library.

P661+ begins to approach the descriptor classes in a typical OO/frameworky way by composing them and designing an ABC version which calls an abstract validator method.

All of this is simple to follow for anyone who has used an OO language for long enough.

### Descriptor Shadowing & Search Order

P665 Reminds that reading an attribute via an instance returns that instance-defined one, but if an instance attrib doesn't exist then it falls back to a class attrib (weird).

But assigning to a missing instance attribute will define a new one!

P666 begins to explain how this (weird) behaviour can be used to make "overriding" and "nonoverriding" descriptors by defining three new descriptor classes; one with `__get__` and `__set__` and then two more; one with `__set__` and one with `__get__` only.

It then makes a new class which uses all those different types of descriptor.

Interesting comment says "methods are also descriptors."

Its too complicated to write down! Basically, whether your descriptor implements get and set determines how the shadowing will work, obviously.

I have no idea why anyone would want to introduce such mad chaos into the app code.

### Methods Are Descriptors!

P671 - Methods Are Descriptors!

I assume it means that this is how Python does it!? It says that they have a `__get__` so they are **nonoverriding descriptors**, and that means setting `obj.myfunc = 7` shadows it.

Gosh this is all really complicated and seems to demonstrate how Python is a shitty load of abstractions over dictionaries, including methods.

### Descriptor Guidelines

The chapter ends with some useful guidelines for using descriptors:

- Use `property` built-in - KISS
- Remember that even readonly descriptors need a `__set__` else setting it will just add a shadow!
- Pure validation descriptors may work only with `__set__` (which does the checks) because the missing `__get__` will fall back to reading the dict/attrib — I think!?
- Wow, so caching an expensive value can be done with just a `__get__` which I think then adds a namesake attrib to the instance with the result and so it shadows itself!
- Lastly it points out that the descriptors' docstring is what ends up in the docs for any prop using it, which it says is unfortunate but doesn't offer a solution.

---

## Chapter 21: Class Metaprogramming

Chapter opens with a quote from Tim Peters saying that you almost never need metaclasses.

Creating or customizing classes at runtime.

`type()` function can be used to create a class at runtime (no `class` keyword needed).

Class decorators are functions capable of inspecting, changing or entirely replacing the class.

Metaclasses are hard to get right and class decorators can often accomplish the same goal more simply.

The author's own go to demo for metaclasses was rendered moot by class decorators in version 2.6.

I might skip this chapter: "If you are not authoring a framework, you should not be writing metaclasses unless for fun / education".

**Skipped**

End of Book!!

Remember to copy paste all into Claude in my Notes repo and move to GitHub.
