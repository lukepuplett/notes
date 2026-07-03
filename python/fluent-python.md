# Fluent Python – Comprehensive Notes

Source: *Fluent Python* by Luciano Ramalho

---

## Chapter 1: Python's Data Model

The Python data model revolves around **dunder (double underscore) functions** that allow your classes to integrate seamlessly with Python's syntax and built-in operations.

### Key Concepts

- **83 dunder functions** exist to implement the protocol (see p.13)
- Examples: `len()`, indexation `[]`, string representation, addition `+`, multiplication `*`
- When you implement `__len__`, calling `len(obj)` works
- Dunder functions enable operator overloading and protocol implementation

### Common Dunder Functions

- `__init__` – Constructor/initializer
- `__repr__` – Official string representation (for developers)
- `__str__` – User-friendly string representation
- `__len__` – Length protocol
- `__getitem__` – Indexation
- `__eq__`, `__lt__`, etc. – Comparison operators
- `__add__`, `__mul__`, etc. – Arithmetic operators
- `__bool__` – Truthiness

---

## Chapter 2: An Array of Sequences

### Container vs. Flat Sequences

**Container sequences** hold references to objects (can mix types):
- `list`, `tuple`, `collections.deque`

**Flat sequences** hold values as copies (fixed type):
- `str`, `bytes`, `bytearray`, `memoryview`, `array.array`

Immutable: `tuple`, `str`, `bytes`

### List Comprehensions & Generator Expressions

```python
dummy = [ord(x) for x in x]  # List comprehension syntax
cartesian = [(color, size) for color in colors for size in sizes]
```

Code in `[]` can span multiple lines for readability.

**Generator expressions** use `()` instead of `[]` and iterate lazily:
```python
array.array('I', (ord(symbol) for symbol in symbols))  # Memory efficient
```

### Unpacking

```python
lat, long = coords  # Unpack tuple into variables
b, a = a, b  # Swap values
send_missile(*coords)  # Unpack into function args
a, b, *rest = range(5)  # Grab operator: (0, 1, [2, 3, 4])
a, *body, c, d = range(5)  # Can appear anywhere: (0, [1, 2], 3, 4)
```

### Named Tuples

```python
City = collections.namedtuple('City', 'name country')  # Space-separated fields
tokyo = City('Tokyo', 'Japan')
print(tokyo.country)  # "Japan"
print(tokyo[0])  # "Tokyo"
```

Members: `City._fields`, `City._make(iterable)`, `city._asdict()`

### Slicing & Mutation

```python
my_list[:3]  # Excludes item at index 3
[start:stop:stride]  # Stride can be negative to reverse

UNIT_PRICE = slice(0, 6)
print(text[UNIT_PRICE])

nums[2:5] = [20, 30]  # Mutate slice contents
del nums[2:5]  # Delete slice
```

Multidimensional: `[i, j]`, `[m:n, k:l]`

### Arrays & Sorting

```python
tic_tac_toe_board = [['_'] * 3 for i in range(3)]
list.sort()  # In-place
sorted(items, key=str.lower, reverse=True)  # New instance
```

Arrays are lean (C-level) and need a type code (e.g., `'b'` for signed char).

### Memoryview & NumPy

`memoryview` provides shared-memory access (like `Span<T>` in C#):
```python
nums = array.array('h', [-2, -1, 0, 1, 2])
memv = memoryview(nums)
memv2 = memv.cast('B')  # Reinterpret bytes
```

**NumPy** (built with C/Fortran):
```python
a = numpy.arange(12)
a.shape = 3, 4  # Reshape in-place
a.transpose()  # Flip rows/cols
```

### Deque & Queues

```python
dq = deque(range(10), maxlen=10)  # Double-ended queue
dq.append() and dq.popleft()  # Atomic, FIFO/stack friendly
```

- `queue.Queue`, `LifoQueue`, `PriorityQueue` – Block when full, thread-safe
- `multiprocessing.Queue`, `JoinableQueue`
- `asyncio` – All queue types for async code
- `heapq` – `heappush()`, `heappop()` on mutable sequences

---

## Chapter 3: Dictionaries & Sets

### Creating Dicts

```python
a = dict(one=1, two=2, three=3)
b = {'one': 1, 'two': 2, 'three': 3}
c = dict(zip(['one', 'two', 'three'], [1, 2, 3]))
d = dict([('two', 2), ('one', 1), ('three', 3)])
a == b == c == d == e  # All equal despite construction order
```

### Dict Comprehensions

```python
DIAL_CODES = [(86, 'China'), ...]
country_code = {country: code for code, country in DIAL_CODES}
{code: country.upper() for country, code in country_code.items() if code < 66}
```

### Hashable Keys

All standard lib maps use `dict` backed by a hash table. Keys must be hashable:
- Must implement `__hash__()`
- Must implement `__eq__()` for equality comparison

### Mapping Methods

- `d.setdefault(idx, []).append(item)` – Set if missing or get if exists
- `d.update(m, [kwargs])` – Uses duck typing to detect iterables
- `defaultdict(list)` – Auto-creates missing keys (only for `d[key]`, not `d.get(key)`)

### Dict Variations

- `OrderedDict` – Preserves insertion order (less needed in 3.7+)
- `ChainMap` – Chains multiple dicts
- `Counter` – Tallies int values per key
- `UserDict` – Pure Python, designed for subclassing
- `MappingProxyType` – Read-only proxy of a mapping

### Set Theory

```python
set(l)  # Unique hashable values
a | b  # Union
a & b  # Intersection
a - b  # Difference
{1, 2, 3}  # Literal syntax (no empty set literal: use set())
frozenset(range(10))  # Immutable set (no literal syntax)
```

Set literals `{1, 2, 3}` are much faster than `set([1, 2, 3])` which creates a list first.

**Warning:** Don't update a dict or set while iterating—you won't reliably iterate all items.

---

## Chapter 4: Text vs. Bytes

### Unicode & Encoding

Python 3 strings are Unicode. Each character is a **code point** (U+0000 to U+10FFFF in hex):
- `str.encode('utf8')` → `bytes`
- Bytes depend on encoding (e.g., `\x41` for 'A' in UTF-8)

### Creating Bytes

```python
bytes.fromhex('31 4B CE A9')
bytes('foo', encoding='utf-8')
bytes([255, 0, 1, 255, ...])
bytes(obj_implements_buffer_protocol)  # bytes, bytearray, memoryview, array.array
```

### Binary Operations

- `memoryview` – Share memory between binary structures
- `struct` module – Binary serialization/deserialization (like .NET BinarySerializer)
- `mmap` module – Memory-map files

### Encodings & Errors

>Python bundles >100 codecs (cp1252, utf-8, cp437, etc.)

Most non-UTF codecs handle only a subset of Unicode. Pass `errors` argument:
```python
errors='ignore'  # Skip unsupported chars
errors='replace'  # Replace with '?'
```

**Motto:** Always specify the encoding!

### Text Normalization

- `str.casefold()` – Like `.lower()` but converts special Unicode chars (µ, ß)
- `unicodedata.normalize()` – Normalize accents and combining chars
- `locale.strxfrm()` – Locale-aware sorting (OS-dependent syntax)
- **PyUCA** lib – Pure Python Unicode Collation Algorithm (by James Tauber)

### Regex & Filenames

- Regex behaves differently on `str` vs `bytes` (bytes is stricter)
- `sys.getfilesystemencoding()` – Convert between str/bytes filenames
- `os.fsencode()`, `os.fsdecode()` – Filename encoding helpers
- `os.listdir(b'.')` → Returns bytes if given bytes

---

## Chapter 5: First-Class Functions

Functions are objects. You can assign them, pass them, and return them.

### Functions as Objects

```python
def factorial(n): ...
fact = factorial  # Assign function to variable
map(factorial, range(11))  # Pass function to map
sorted(fruits, key=len)  # Functions as keys
```

### Lambda & Readability

```python
lambda word: word[::-1]  # Anonymous function (body is expression only)
sorted(fruits, key=lambda word: word[::-1])
```

Prefer **listcomp** or **genexp** for clarity:
```python
[fact(n) for n in range(6)]  # More readable than map()
[factorial(n) for n in range(6) if n % 2]  # Filter built-in
```

### Callable Objects

```python
callable(obj)  # Check if object is callable
```

Classes, functions, methods, and instances with `__call__` are callable.

### Function Introspection

```python
dir(func)  # Lists all dunders
func.__dict__  # Arbitrary values attached to function
func.__code__  # Code object
```

### Parameters: Positional, Keyword, Variadic

```python
def tag(name, *content, cls=None, **attrs):
    # name – positional
    # *content – remaining positional args (tuple)
    # cls – keyword-only (must pass as name=value)
    # **attrs – keyword args (dict)

# Call:
tag('p', 'Hello', id='content', cls='info')
```

### The inspect Module

```python
from inspect import signature
sig = signature(func)
sig.bind(values)  # Try to bind args using Python's rules
```

### Function Annotations (Python 3)

```python
def clip(text: str, max_len: int = 80) -> str:
    ...
```

**No processing!** Annotations are stored in `__annotations__` for IDEs and linters only.

### Functional Programming: operator & functools

```python
operator.mul(a, b)  # Function version of *
itemgetter(1, 0, 4)  # Extract fields: like lambda x: (x[1], x[0], x[4])
attrgetter('name', 'coord.lat')  # Extract attributes
methodcaller('replace', ' ', '-')  # Call method with args

functools.partial(mul, 3)  # Partial application: fix first arg
functools.reduce(func, iterable, [initial])
```

---

## Chapter 6: Design Patterns with First-Class Functions

### Strategy Pattern → Functions

The classic GoF strategy pattern is unnecessary in Python. Replace strategy classes with functions.

### Registration Systems

```python
globals()  # Dict of current global symbol table
promos = [globals()[name] for name in globals()
          if name.endswith('_promo') and name != 'best_promo']
```

Or use `inspect.getmembers(module, inspect.isfunction)`.

---

## Chapter 7: Function Decorators & Closures

### Decorators Basics

A decorator is a callable that takes a function and returns a function (or callable):

```python
@decorate
def target():
    print('running')

# Equivalent to:
target = decorate(target)
```

**Decorators execute at import time** when modules load. Great for registration systems.

### Variable Scope & nonlocal

```python
b = 6
def f2(a):
    print(a)
    print(b)
    b = 9  # Assignment makes b local!
```

Python assumes assignment creates local variables. Use `nonlocal` to rebind captured variables:

```python
def outer():
    count = 0
    def inner():
        nonlocal count
        count += 1
    return inner
```

### Closures

Closures capture referenced variables from enclosing scope. Essential for async callbacks and functional-style code.

### Standard Library Decorators

- `@property` – Replace field with computed property
- `@classmethod` – Receives `cls` instead of `self`
- `@staticmethod` – Regular function attached to class (no `cls`/`self`)
- `@functools.wraps` – Preserve original function metadata in wrapper

### Memoization with lru_cache

```python
@functools.lru_cache()  # Note the () – it can take args
def expensive_func(n):
    ...
```

Recursive Fibonacci: 17.7s → 0.0005s with caching!

### singledispatch (Python 3.4+)

Function overloading based on first argument type:

```python
@singledispatch
def htmlize(obj):
    return f'<pre>{obj!r}</pre>'

@htmlize.register(str)
def _(text):
    return f'<p>{text.replace("<", "&lt;")}</p>'

@htmlize.register(int)
def _(n):
    return f'<pre>{n:b}</pre>'
```

Prefer abstract base classes like `numbers.Integral` over concrete types.

### Parameterized Decorators

Three-level structure:

```python
def myDecorator(arg):  # Factory – takes config
    def actualDecorator(func):  # Decorator – takes function
        def wrapper(*vargs, **kwargs):  # Wrapper – runs function
            # ... use arg ...
            return func(*vargs, **kwargs)
        return wrapper
    return actualDecorator

@myDecorator(X)
def something():
    ...
```

---

## Chapter 8: Object References, Mutability & Recycling

### Variables Are Labels

Variables don't contain objects; they're labels pointing to objects.

```python
a is b  # Same object (identity)
a == b  # Same value (equality)
```

`is` compares identities (using `id()`), `==` is sugar for `__eq__()`.

### Copying

```python
copy.copy(obj)  # Shallow copy
copy.deepcopy(obj)  # Deep copy (recursively)
```

Shallow copies reference the same nested objects. Lists and dicts are reference types.

### Mutable Default Arguments (Pitfall!)

```python
def __init__(self, passengers=[]):  # DON'T!
    self.passengers = passengers
```

The `[]` is evaluated once at definition time. All instances share the same list!

### del & Garbage Collection

`del x` unbinds the label `x` and decrements reference count. Objects are freed when refcount hits zero.

CPython uses **reference counting** (primary GC), plus **generational GC** (cycles).

### Weak References

Weak references don't prevent GC:

```python
wref = weakref.ref(obj)
obj_or_none = wref()  # Returns object or None if GC'd
```

Collections:
- `WeakValueDictionary` – Values are weak; keys removed on GC
- `WeakKeyDictionary` – Keys are weak
- `WeakSet` – Track class instances without keeping them alive

---

## Chapter 9: A Pythonic Object

### String Representation Dunders

```python
__repr__()  # Official representation (for devs)
__str__()   # User-friendly
__format__(format_spec)  # Custom formatting
__bytes__()  # Byte representation
```

### Class vs. Static Methods

```python
@classmethod
def frombytes(cls, octets):
    typecode = chr(octets[0])
    memv = memoryview(octets[1:]).cast(typecode)
    return cls(*memv)  # cls = the class itself

@staticmethod
def regular_func():  # No special first param
    ...
```

`@classmethod` creates alternative constructors. `@staticmethod` is just a regular function on the class.

### Type Coercion

Implement for numeric types:
```python
__int__()   # int() conversion
__float__()  # float() conversion
```

### Private Attributes & Name Mangling

```python
class Dog:
    def __init__(self):
        self.__mood = 'happy'
```

Stored as `_Dog__mood` in `__dict__`. Prevents accidental overwrites in subclasses (not true privacy).

Convention: single underscore `_mood` signals "private" without mangling.

### __slots__

Eliminate per-instance `__dict__` overhead:

```python
class Vector2D:
    __slots__ = ('__x', '__y')
```

10M instances: `__dict__` version ~1.5GB, `__slots__` version ~0.65GB!

**Note:** Use NumPy/SciPy for numeric data, not homemade classes.

---

## Chapter 12: Inheritance

### Multiple Inheritance

Python supports it. Most use with ABCs (interfaces) from `collections.abc`.

### Pitfall: Subclassing Built-ins

C-implemented built-ins (dict, list, str) **bypass overridden special methods** for speed:

```python
class BadDict(dict):
    def __setitem__(self, key, value):
        super().__setitem__(key, [value] * 2)  # Ignored by __init__!

d = BadDict(a=1)
print(d)  # {'a': 1}, not {'a': [1, 1]}
```

**Solution:** Subclass from `collections.UserDict`, `UserList`, or `UserString` instead.

### MRO & __mro__

```python
bool.__mro__  # (bool, int, object)
```

### Mixins

Mixins are lightweight base classes offering reusable functionality without standing alone. Useful for multiple inheritance to avoid deep hierarchies. Do not leave methods abstract.

### Design Principles

Recommend not subclassing from more than one concrete class. ABCs are fine for multiple inheritance. Modern frameworks often favor composition over inheritance.

---

## Chapter 13: Operator Overloading

### Limitations

- Cannot overload for built-in types
- Cannot create new operators (only overload existing ones)
- Cannot overload `is`, `and`, `or`, `not`

### Unary Operators

```python
__neg__()  # Negation (-)
__pos__()  # Positive (+)
__invert__()  # Bitwise NOT (~)
__abs__()  # abs()
```

### Binary & Reflected Operators

```python
__add__(self, other)  # Addition
__radd__(self, other)  # Reversed: other + self
```

Python dispatches left, then right if left returns `NotImplemented`.

Return `NotImplemented` (not `NotImplementedError`) to allow other to try:

```python
def __radd__(self, other):
    return self + other  # Delegate to __add__
```

### Rich Comparisons

```python
__eq__, __ne__
__lt__, __le__
__gt__, __ge__
```

For reverse: `__gt__` reverse is `__lt__`, etc.

### Augmented Assignment

```python
__iadd__(self, other)  # += (must return self)
```

Works "for free" if `__add__` exists, but can optimize in-place if overridden.

### itertools

Useful for combining iterables:
```python
itertools.zip_longest(...)
itertools.takewhile(lambda n: n < 3, itertools.count(1, .5))
```

---

## Chapter 14: Iterables, Iterators & Generators

### Iterable Protocol

```python
class MyIterable:
    def __iter__(self):
        return iter(self.items)  # Returns an iterator
```

Iterables are registered via `collections.abc.Iterable` (duck typing via `__subclasshook__`).

### Iterator Protocol

```python
class MyIterator:
    def __next__(self):
        ...
    def __iter__(self):
        return self  # Iterator returns itself
```

Call `next(iterator)` to advance.

### Generator Functions

Pythonic way to create iterators using `yield`:

```python
def __iter__(self):
    for word in self.words:
        yield word
```

The function body becomes a generator object with `__next__` and `__iter__`. Lazy evaluation.

### Generator Expressions

Lazy version of listcomp:

```python
(match.group() for match in RE.finditer(self.text))
```

### Reducing Functions

- `all(iterable)`, `any(iterable)`
- `max(iterable, [key=], [default=])`
- `min(...)`
- `functools.reduce(func, iterable, [initial])`
- `sum(iterable)` – Use `math.fsum()` for floats

### iter() with Sentinel

```python
iter(callable, sentinel)  # Calls callable until it returns sentinel
```

Use case: read lines from file until blank line.

---

## Chapter 15: Context Managers & Else Blocks

### else Blocks

Not unique to `with`. Useful with `for`, `while`, `try`:

```python
for item in items:
    if not found(item):
        break
else:
    print("All items checked")  # Runs if no break
```

```python
try:
    ...
except SomeError:
    ...
else:
    ...  # Runs if no exception, outside try scope
```

### Context Managers

```python
with open('file.txt') as f:
    ...
```

The `with` statement calls:
- `__enter__()` – Returns the object bound by `as`
- `__exit__(exc_type, exc_value, traceback)` – Cleanup

`exc_*` are `None` if block exited cleanly.

### contextlib

```python
contextlib.closing(obj)  # Make context manager from close() method
contextlib.suppress(ExceptionType)  # Ignore specific exceptions
contextlib.ExitStack  # Stack multiple contexts
```

---

## Chapter 16: Coroutines

### Yield as Receiver

```python
def simple_coroutine():
    print('started!')
    x = yield
    print('received:', x)

my_coro = simple_coroutine()
next(my_coro)  # Prime: 'started!'
my_coro.send(42)  # 'received: 42'
```

Coroutines are bi-directional: send data in, yield data out.

### Priming

Must call `next(c)` or `c.send(None)` before sending data. States: `GEN_CREATED`, `GEN_RUNNING`, `GEN_SUSPENDED`, `GEN_CLOSED`.

### Close & Throw

```python
c.close()  # Injects GeneratorExit
c.throw(ExcType, [exc_value, [traceback]])  # Throw exception
```

Coroutine can catch and handle, advancing to next `yield`.

### yield from

Delegate to sub-generator:

```python
def delegator():
    yield from sub_generator()
```

Allows coroutine composition. Sub-generator can send/receive values through the delegator.

### Use Cases

- Pull-style iterators
- Push-style aggregators
- Tasks (especially with SimPy for discrete event simulation)

Taxi simulator example: each taxi is a coroutine state machine advanced by simulation time.

---

## Chapter 17: Concurrency with Futures

### Futures (Tasks in C#)

```python
from concurrent.futures import Future, Executor

executor = ThreadPoolExecutor()
future = executor.submit(func, arg1, arg2)
result = future.result()  # Block until done
```

Methods:
- `.done()` – Boolean
- `.result()` – Blocks in concurrent.futures
- `.add_done_callback(func)` – Called when done
- `futures.as_completed(futures)` – Yields futures as they complete

### ThreadPoolExecutor vs. ProcessPoolExecutor

**ThreadPoolExecutor:** Limited by **GIL** (Global Interpreter Lock). Only one thread executes Python bytecode at a time. I/O releases the GIL, so network/disk benefit from threads.

**ProcessPoolExecutor:** True parallelism for CPU-bound work. Each worker is a separate process. Overhead is higher.

5x speedup from ThreadPoolExecutor on I/O. 4x speedup from ProcessPoolExecutor on CPU (4-core machine, SHA256 hashing).

### executor.map()

```python
futures = executor.map(do_work, range(5))
```

Returns futures in submission order; blocks until all complete.

---

## Chapter 18: Concurrency with asyncio

### Event Loop

```python
loop = asyncio.get_event_loop()
result = loop.run_until_complete(coro)
loop.close()
```

Entry point manually bootstraps the loop.

### Coroutines with asyncio

```python
@asyncio.coroutine
def my_coro():
    result = yield from asyncio.sleep(0.1)
    ...
```

In Python 3.5+, use native `async`/`await`:

```python
async def my_coro():
    result = await asyncio.sleep(0.1)
    ...
```

### Key Differences from Threads

- **Single-threaded:** Only one coroutine executes at a time
- **Cooperative:** Only yields at explicit `yield`/`await` points
- **No GIL issues:** But shared mutable state can still change "suddenly" between yields

### asyncio.wait() & Tasks

```python
to_dos = [my_coro(cc) for cc in cc_list]
done, pending = await asyncio.wait(to_dos)  # Like Task.WaitAll()
```

`asyncio.wait()` optionally takes `timeout` parameter; `pending` is populated if timeout expires.

### Semaphore for Rate Limiting

```python
sem = asyncio.Semaphore(5)
async with sem:
    await fetch(url)  # Max 5 concurrent fetches
```

### run_in_executor

```python
loop.run_in_executor(None, sync_filesystem_op)
```

Runs synchronous code on thread pool. CPython drops the GIL for the duration.

### TCP/HTTP Servers

```python
loop = asyncio.get_event_loop()
server = await loop.create_server(handler, address, port)
```

External library `aiohttp` for HTTP async support (not in stdlib 3.4).

---

## Chapter 19: Dynamic Attributes & Properties

### __getattr__ & __setattr__

```python
def __getattr__(self, name):
    return self.__dict__.get(name, '<missing>')

def __setattr__(self, name, value):
    self.__dict__[name] = value
```

### FrozenJSON Pattern

Create nested wrapper objects for JSON dict access:

```python
feed['schedule']['speakers'][-1]['name']
# Becomes:
feed.schedule.speakers[-1].name
```

### __new__ vs __init__

```python
__new__(cls, ...)  # Allocates instance (must return instance)
__init__(self, ...)  # Initializes instance
```

`__new__` can return instance of different class; `__init__` not called in that case.

### @property Decorator

```python
@property
def x(self):
    return self.__x

@x.setter
def x(self, value):
    self.__x = value
```

Replaces fields with computed properties seamlessly.

### Property Factory

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

class LineItem:
    weight = quantity('weight')
    price = quantity('price')
```

Encodes validation rules once, reusable across many classes.

### Attribute Access Order

Python searches:
1. Data descriptors (have `__set__`)
2. Instance `__dict__`
3. Non-data descriptors (have `__get__` only)
4. Class `__dict__`
5. Parent classes

### Essential Attributes & Functions

- `__class__`, `__dict__`, `__slots__`
- `dir(obj)`, `getattr()`, `hasattr()`, `setattr()`, `vars()`
- Special methods: `__delattr__`, `__dir__`, `__getattr__`, `__getattribute__`, `__setattr__`

---

## Chapter 20: Attribute Descriptors

### Descriptor Protocol

A descriptor is a class implementing `__get__`, `__set__`, and/or `__delete__`:

```python
class Quantity:
    def __init__(self, name):
        self.name = name
    
    def __get__(self, instance, owner):
        if instance is None:
            return self
        return instance.__dict__[self.name]
    
    def __set__(self, instance, value):
        if value > 0:
            instance.__dict__[self.name] = value
        else:
            raise ValueError('value > 0')
```

Use as class attribute:
```python
class LineItem:
    weight = Quantity('weight')
    price = Quantity('price')
```

### Data vs. Non-Data Descriptors

- **Data descriptor:** Has both `__get__` and `__set__` (or `__delete__`)
  - Takes precedence over instance attributes
- **Non-data descriptor:** Has only `__get__`
  - Instance attributes shadow it

### Methods Are Descriptors

Methods have `__get__`, making them non-data descriptors. This allows `obj.method = 7` to shadow the method with a value.

### Caching with __get__

A descriptor with only `__get__` can cache by setting an instance attribute of the same name, shadowing itself:

```python
def __get__(self, instance, owner):
    if instance is None:
        return self
    value = expensive_computation()
    instance.__dict__[self.name] = value  # Shadow
    return value
```

### Guidelines

1. Prefer built-in `@property` (KISS)
2. Readonly descriptors need `__set__` (raising error) else shadowing occurs
3. Pure validation descriptors may use only `__set__`, leaving read to fallback
4. Descriptor docstrings become the property documentation

---

## Chapter 21: Class Metaprogramming

> "Metaclasses are 99% magic. If you're not authoring a framework, you shouldn't be writing metaclasses." — Tim Peters

**Intentionally skipped.** The chapter covers:
- Creating classes at runtime with `type()`
- Class decorators for inspecting/modifying classes
- Metaclasses for advanced customization

If not authoring a framework, class decorators solve most metaprogramming needs more simply than metaclasses.

---

## Appendix: Key Terminology & Patterns

### Pythonic Principles

- Duck typing: Rely on object behavior, not type
- EAFP (Easier to Ask for Forgiveness than Permission): Try/except preferred over isinstance checks
- Lazy evaluation: Generators and iterators defer computation
- Closures & decorators: Powerful functional programming patterns

### C# / .NET Parallels

| Python | C# / .NET |
|--------|-----------|
| `async def` / `await` | `async Task` / `await` |
| `concurrent.futures.Future` | `Task<T>` / `Promise<T>` |
| `@asyncio.coroutine` + `yield from` | `async` + `await` |
| `@property` | Auto property `{ get; set; }` |
| `__slots__` | Struct (value type) |
| Descriptor | Implicit property getter/setter |
| `nonlocal` | Closure capture |

### Performance Tuning

- Use `array.array` instead of `list` for homogeneous numeric data
- Use `__slots__` for millions of small objects
- Use NumPy for large numeric arrays
- ProcessPoolExecutor for CPU-bound tasks
- ThreadPoolExecutor or `asyncio` for I/O-bound tasks

---

**End of Notes**
