Here are your reformatted notes with technical jargon and programming terms enclosed in backticks:

- `cargo` is the build system and package manager
- `cargo` is a command-line executable tool
- Command `cargo new hello_cargo` creates a new project using:
  - `Cargo.toml`
  - `src/main.rs`
  - Also initializes a new Git repo with a `.gitignore` file
- `TOML` stands for Tom's Obvious Minimal Language
- It looks like a Windows `.ini` file but with comments
- `cargo` is the build system and package manager
- `cargo` is a command-line executable tool
- `cargo new hello_cargo` creates a new project using:
  - `Cargo.toml`
  - `src/main.rs`
  - Also initializes a new Git repo with a `.gitignore` file
- `TOML` stands for Tom's Obvious Minimal Language
  - Looks like a Windows `.ini` file but with comments
- Packages are called crates
- Use `cargo build` for more verbose output and chunk files in `/target/debug/hello_cargo` folder
- `cargo run` will build and run
- `cargo check` is like `tsc --noEmit` command, a fast way to check for compilation errors
- `cargo build --release` to build an optimized proper release binary
- Chapter 2
- Rust has a standard set of items it brings into scope of each program, called the Prelude
- Function names ending with an exclamation mark are actually macro calls
- By default, variables are immutable
- `mut` keyword makes a variable mutable (like `let` versus `const` in JavaScript)
- `String::new()` makes a new UTF-8 encoded string
- Can call into library using longer syntax: `std::io::stdin`
- Use `use std::io;`
- The `stdin()` function returns an instance of `std::io::Stdin`
- `io::stdin().read_line(&mut guess)`
  - `&` denotes a reference
  - `mut` denotes mutability of the reference
- `.expect("Failed to read line");`
- The `read_line` method returns a `Result` enum that can be one of many state variants, `Ok` or `Err` in this case
- `expect` is a method on the `Result`
- `expect` is used to catch and throw expected errors
- If `Ok`, `expect` returns the value that `Ok` is holding (in this case, the number of bytes read)
- If you don't call `expect`, the compiler will warn you about unused results
- String interpolation is like C#, except it seems to only take a variable name in the curly brackets
- Must use empty curly brackets for positional placeholders and then supply the expression after the string
- Example: `println!("x = {}, y + 2 = {}", x, y + 2);`
- `Cargo.lock` locks in versions until you specifically update, and this file should be checked into git
- `cargo update` will ignore the lock file and look for latest versions while adhering to semver in `Cargo.toml` file
- Once the dependency is downloaded, just use `rand::Rng;`
- Page 23
- `cargo doc --open` CLI command will build docs for all dependencies
- `match` keyword performs a pattern match and has arms like cases in C#
- `let guess: u32 = guess.trim().parse().expect("Please type a number!");`
- Looks like a variable can be converted in-situ; this is called shadowing
- An `expect` style call can be converted to use `match`:
  ```rust
  let guess: u32 = match guess.trim().parse() {
      Ok(num) => num,
      Err(_) => continue,
  };
  ```
- Chapter 3: Common Programming Concepts
- `const` can be declared in any scope, must be annotated with its type, and use shouting caps
- Using the `let` keyword can shadow an existing variable. Shadows are locally scoped
- Four scalar types: integers, floating-points, booleans, characters
- Integer sizing: 8, 16, 32, 64, and 128 bits. Can use underscore in literals like in JavaScript, e.g., `1_000`
- Overflows will panic in debug mode, wrap in release mode
- Can use `wrapping_*`, `checked_*`, `overflowing_*`, and `saturating_*` type methods to get specific behaviors
- Integer literals can be:
  - `98_222`
  - `0xff` for hex
  - `0o77` for octal
  - `0b1111_0000` for binary
  - `b'A'` for a byte
- Two types of floating-point: `f32` and `f64` (default, same speed as `f32`)
- Character literals use single quotes: `'z'` or `'😻'`
- Compound types: tuples and arrays
- Tuple example: `let tup: (i32, f64, u8) = (500, 6.4, 1);`
- Destructuring syntax: `let (x, y, z) = tup;` and `let five_hundred = tup.0;`
- A tuple with no values `()` is called a unit and represents an empty return value
- Arrays are of a fixed type and length: `let a = [1, 2, 3, 4, 5];`
- Arrays are allocated on the stack. Use a `Vec` (vector) from the standard library for elastic sets
- Define: `let a: [i32; 5] = [1, 2, 3, 4, 5];`
- Also: `let a = [3; 5];` has five elements all set to 3
- `let first = a[0];`
- `usize` means `u32` or `u64`, depending on the platform architecture
- Rust uses snake_case for function names
- `let y = { let x = 3; x + 1 }` works as long as there's no final semicolon, which would cause it to become a statement and not an expression that returns a value
- `fn() -> i32 { }` uses the arrow to denote a return type
- `if number < 5 { ... } else { ... }` No brackets around conditions in Rust, and it also uses the `else if` syntax
- `let number = if something { 5 } else { 6 };`
- `let result = loop { ... break counter * 2; };` The break from the loop can return a value
- `'some_name: loop { ... break 'some_name; }` In this syntax, the loop has a name or label called `'some_name`
- `while some_condition { ... }` This is like C#
- `for element in ... { ... }` This syntax is for-each style
- `for number in (1..4).rev() { ... }` This syntax demonstrates a for-each over a group of numbers where the numbers have been reversed
Chapter 4: Understanding Ownership
- Each value has an owner
- There can only be one owner at a time
- When the owner goes out of scope, the value will be dropped
Page 62:
- Rust calls `drop` as soon as a value goes out of scope
- `let s1 = String::from("hello");`
  `let s2 = s1;`
  The assignment in `s2` copies the pointer on the stack, but not the heap object value itself
- `s1` becomes invalid and cannot be used again
- In Rust, this is known as a move
- Types annotated with the `Copy` trait will be copied
- Types with the `Drop` trait cannot also be annotated with `Copy`
- The scalar types all implement `Copy`
- Passing a value to a function will move or copy
Page 70:
- To return a value from a function, just express it without the semicolon
- Rust concept of references, which allow a function to borrow. Example: `fn(s: &String)`
- The opposite of `&` referencing is `*` dereferencing (Chapter 15 and Chapter 8)
- We're not allowed to mutate borrowed references
- Instead, we must make a mutable reference: `&mut s`
- Example: `change(&mut s)`
- Change the method signature to take a mutable reference: `fn change(some_string: &mut String) { ... }`
- Cannot have multiple references to the same value at the same time or same scope
- The scope of a reference ends on its last use, so you can define a mutable reference after that last usage
- Cannot create dangling references where the underlying value has been dropped while a reference is still pointing at it
- A slice is like a reference, has common range syntax: `[7..16]`, `[7..]`, `[..]`
- String slices must begin and end at UTF-8 character boundaries
- The string slice is a first-class type with symbol `&str`
- Benefit: a string slice is bound to its underlying memory
- A string literal (e.g., `"hello"`) is a string slice and `&str`
- Use `&str` in function signatures, e.g., `fn first_word(s: &str) { ... }`
Chapter 5: Using Structs to Structure Related Data
- `struct User { active: bool, username: String, ... }`
- `let user1 = User { active: true, ... }` Note that it uses colons to set values, but uses dot notation to access fields, e.g., `user1.active`
- The entire instance is mutable, and Rust has no access modifiers
- Can use field init shorthand, e.g., `User { username, email, ... }` where the name of the variable/argument matches the name of the field
- Struct update syntax makes a new instance using values from an existing instance: `let user2 = User { active: false, ..user1 }`
- But the values are moved, and `user1` values can no longer be used
- Tuple structs seem to be tuples but with a name, which makes them different types
- Can define unit-like structs, which have no fields. This appears to be like a flag interface or trait in Rust parlance
- The book explains why the `username` field is `String` and not `&str`: it's so that the `User` struct owns the value, as opposed to being lent the data
- A struct can store references, but it needs a lifetime specified
Page 94:
- The book does a detour on printing a struct for debugging via `println!`, which means perhaps implementing a trait called `Debug` by adding an attribute to the struct: `#[derive(Debug)] struct Rectangle { ... }`
- It mentions that the `dbg!` macro will take ownership of a value and print it along with the file and line number to standard error, not standard out
- `dbg!(30 * scale)` will work because `dbg!` hands back the value
I apologize for the omissions. You're right, I should have included more of the code examples you provided. Let me try again with a more comprehensive transcription of your notes:
- Chapter 6: Enums and Pattern Matching
- Similar to C# enums, but with key differences
- Basic enum syntax:
  ```rust
  enum IpAddrKind {
    V4,
    V6
  }
  ```
- Accessing enum variants: `IpAddrKind::V6`

- Rust enums can hold data:
  ```rust
  enum IpAddr {
    V4(String),
    V6(String)
  }
  ```
- Example usage:
  ```rust
  let home = IpAddr::V4(String::from("127.0.0.1"));
  ```

- Enums can hold tuples:
  ```rust
  enum IpAddr {
    V4(u8, u8, u8, u8),
    // ...
  }
  ```

- Enums can contain structs:
  ```rust
  enum IpAddr {
    V4(Ipv4Address),
    // ...
  }
  ```

- Enums can include other enums
- Enums can be quite complex, resembling structs
- Different variants can have different data types
- All variants are of the same enum type, allowing them to be passed as a single parameter to functions
- Methods can be defined for enums using `impl`

- Rust uses `Option<T>` instead of null:
  ```rust
  enum Option<T> {
    None,
    Some(T)
  }
  ```

- `match` expression for pattern matching:
  ```rust
  match coin {
    Coin::Penny => 1,
    Coin::Nickel => 5,
    // ...
  }
  ```
- Match arms can include multi-line code blocks
- Can extract values from enums using binding patterns:
  ```rust
  match coin {
    // ...
    Coin::Quarter(state) => {
      // use the state variable
    }
  }
  ```

- Match arms must cover all possibilities
- Example of matching raw values with a catch-all pattern:
  ```rust
  match dice_roll {
    6 => celebrate(),
    other => move_player(other)
  }
  ```
- Underscore `_` pattern for ignoring values
- Combine `_` with unit `()` for no-op:
  ```rust
  match dice_roll {
    6 => celebrate(),
    _ => ()
  }
  ```

- `if let` syntax for matching a single outcome:
  ```rust
  if let Some(v) = some_value {
    // ...
  }
  ```
- Chapter 7: Managing Growing Projects with Packages, Crates, and Modules

- Package: Can contain multiple binary crates and optionally one library crate
  - Similar to Visual Studio solution
  - Contains a `Cargo.toml` file
  - Used to build, test, and share crates

- Crate: A tree of modules producing a library or executable
  - Smallest amount of code Rust compiler considers at a time
  - Can be a single .rs file compiled by `rustc` command-line utility

- Module: Controls organization, scope, and privacy of paths

- Path: Way of naming items (struct, function, module)

- Crate root: Source file where Rust compiler starts (root module)

- Package structure:
  - `src/main.rs`: Convention for binary crate root
  - `src/lib.rs`: Convention for library crate root
  - Can have both, resulting in two crates

- Module declaration:
  ```rust
  mod name {
    // Inline code
  }
  ```
  Or in separate files:
  - `src/name.rs`
  - `src/name/mod.rs` (old way)

- Privacy:
  - Module code is private by default
  - Use `pub mod` to make module public
  - Use `pub` for items within the module

- `use` keyword: Similar to C#'s `using`, but can import specific types

- Module referencing:
  - Relative path syntax
  - `self` syntax
  - `super` syntax
  - Identifier

- Sub-modules can use code from ancestor modules, but not vice versa

- Rust API Guidelines: https://rust-lang.github.io/api-guidelines/

- Packages with bin and lib:
  - Define module tree in `lib.rs`
  - Keep bin lightweight

- Idiomatic to bring types into scope via full path
  - Use parent module path for naming collisions:
    ```rust
    fmt::Result
    ```

- Aliasing with `as` keyword

- Re-exporting with `pub use`:
  - Allows modules to present different public structure

- Chapter 7 (continued):
  
- Shorten `use` syntax with nested imports:
  ```rust
  use std::{cmp::Ordering, io};
  ```
  Can also place `self` in the curly braces
- Append asterisk to bring everything into scope (globbing):
  ```rust
  use std::collections::*;
  ```

Chapter 8: Common Collections
- All Rust collections: https://doc.rust-lang.org/std/collections/index.html
- Vector (`Vec<T>`), similar to C# List<T>:
  ```rust
  let v: Vec<i32> = Vec::new();
  let v = vec![1, 2, 3];  // macro syntax, defaults to i32
  ```
- Add items with `push()` (requires `mut`)
- Access items:
  - `v[2]` (panics if out of bounds)
  - `v.get(2)` (returns `Option<&T>`)
- Borrow checker gotcha example:
  ```rust
  let mut v = vec![1, 2, 3, 4, 5];
  let first = &v[0];  // immutable borrow
  v.push(6);  // mutable borrow
  println!("The first element is {first}");  // Error
  ```
  This fails because `push` may cause internal reallocation
- Iterating:
  ```rust
  for i in &v { ... }  // Borrowing its iterable
  for i in &mut v { ... }
  ```
- Modifying values while iterating:
  ```rust
  for i in &mut v {
      *i += 50;  // dereferencing (see p. 322)
  }
  ```
- Cannot mutate vector contents from within an iterator due to borrow checker
- Use enums to store different types in a vector:
  ```rust
  enum SpreadsheetCell {
      Int(i32),
      Text(String),
  }
  ```
- For unknown types at compile time, use trait objects (see Chapter 17)

Strings:
- Book is confusing: mentions only one string type (`&str`), but also `String`
- `String` type is a wrapper around `Vec<u8>`
- `to_string()` method is on the `Display` trait
- Creating strings:
  ```rust
  let s = "literal".to_string();
  let s = String::from("literal");
  ```
- Concatenation:
  ```rust
  s1 + &s2  // s1 is moved
  format!("{}{}{}", s1, s2, s3)  // doesn't move
  ```
- Appending:
  ```rust
  s.push_str("bar");
  s.push('b');
  ```
- `+` operator uses `fn add(self, s: &str) -> String`
- Multiple concatenations are unwieldy with `+`, use `format!` macro instead
- Cannot index strings by integer (may break UTF-8):
  ```rust
  let s1 = String::from("hello");
  let h = s1[0];  // This will panic
  ```

- Chapter 8 (continued):

- String handling:
- Avoid using string slices with ranges (`[..]`) as it may cause runtime panic if not on character boundaries.
- Use `chars()` function to iterate over characters, or `bytes()` for byte access (reliable only for ASCII).
- The book demonstrates complexities of grapheme clusters using an obscure Asian script.
- Refers to crates.io for specialized language support.
- Unclear if standard library fully supports basic English, Spanish, and Latin languages.

- HashMap:
- Creation: `let mut scores = HashMap::new();`
- Adding items: `scores.insert(String::from("Italy"), 2);`
- Accessing with default value:
  ```rust
  let score = scores.get(&String::from("Italy")).copied().unwrap_or(0);
  ```
  This uses `copied()` to convert `Option<&i32>` to `Option<i32>`.
- Iteration: `for (key, value) in &scores { ... }`
- Types with the Copy trait are copied into the map, others are moved.
- References can be put in the map but must remain valid for the map's lifetime.
- `insert()` overwrites by default (last writer wins).
- Using `entry()` for conditional insertion:
  ```rust
  scores.entry(String::from("Italy")).or_insert(5);
  ```
- Updating based on old value:
  ```rust
  let score = map.entry(String::from("Turkey")).or_insert(0);
  *score += 1;
  ```
- Default hashing uses SipHash (implements `BuildHasher` trait).

- Chapter 9: Error Handling
- Rust uses `Result<T, E>` instead of exceptions, and `panic!` macro for unrecoverable errors.
- Call stack can be displayed when an environment variable is set.
- Pattern matching on `Result::Ok(T)` or `Result::Err(E)`.
- Further pattern matching on error types:
  ```rust
  match error.kind() {
      ErrorKind::NotFound => { /* etc */ },
      other_error => {
          // Could call panic!(error_message) and use other_error variable
      }
  }
  ```
- Page 168 shows a more concise example using closures and `unwrap_or_else()` method.

- Chapter 9 (continued):
  
- Shortcut: Use `unwrap()` on `Result`, which returns the value or panics:
  ```rust
  File::open("file.txt").unwrap()
  ```
- Custom panic message with `expect()`:
  ```rust
  File::open("file.txt").expect("Failed to open file")
  ```
- Propagate errors with `?` operator:
  ```rust
  fn read_username() -> Result<String, io::Error> {
      let mut file = File::open("username.txt")?;
      // ...
  }
  ```
- `?` operator uses `From` trait to convert error types
- `impl` allows defining methods on types we don't own (like C# extension methods)
- Chaining method calls with `?`:
  ```rust
  File::open("a.txt")?.read_to_string(&mut s)?;
  ```
- `?` can be used with `Option<T>` to propagate `None`
- Example of `Option<T>` with `?`:
  ```rust
  fn last_char_of_first_line(text: &str) -> Option<char> {
      text.lines().next()?.chars().last()
  }
  ```
- Convert between `Result` and `Option` using `ok()` and `ok_or()`
- `main()` can return `Result<(), Box<dyn Error>>`
- `Box<dyn Error>` means any kind of error (see p. 379)
- `unwrap()` and `expect()` are normal in tests

Chapter 10: Generic Types, Traits, and Lifetimes
- Generics similar to C#, including constraints:
  ```rust
  fn largest<T: PartialOrd>(list: &[T]) -> &T
  ```
- Constraint `PartialOrd` is from `std::cmp` and is a trait
- Structs, enums, and methods can be generic
- Implement methods for specific types:
  ```rust
  impl Point<f32> { 
      // methods specific to Point<f32>
  }
  ```
Here's the revised version with hyphens for bullets, corrected sentences, and recognized spoken symbols:

- Page 192. Traits are like interfaces, but with some differences. They are defined using:

```rust
pub trait Summary {
    fn summarize(&self) -> String;
}
```

- With no implementation code, just like an interface. To implement the trait, we must use an Associated function as you would expect, but with special syntax:

```rust
impl Summary for NewsArticle
```

- Key difference to C#: the trait must be brought into scope via use. E.g.:

```rust
use aggregator::{Summary, Tweet};
```

- Other crates that depend on aggregator can implement the trait on their own types, too. But cannot stick external traits on external types. So we can implement our own traits on external types.

- A big difference to C# seems to be that a trait can have an implementation. This is like a cross between a base class and an interface. 

- You can specify a function parameter can only take a reference to a value with a trait by using the impl keyword. For example:

```rust
fn do_something(item: &impl Summary) { ... }
```

- For more complex signatures, can use the fuller trait bound syntax, which uses generic syntax. For example:

```rust
fn do_something<T: Summary>(item: &T) { ... }
```

- Cool feature: trait bounds can combine traits. For example:

```rust
fn do_something(item: &(impl Summary + Display)) { ... }
```

- Another example is:

```rust
fn do_something<T: Summary + Display>(item: &T) { ... }
```

- You can also use a where clause syntax just like C#. For example:

```rust
fn do_something<T, U>(t: &T, u: &U) -> i32
where
    T: Display + Clone,
    U: Clone + Debug,
{ ... }
```

- Functions can also return trait implementers using this syntax:

```rust
-> impl Summary
```

- Thirdly, you cannot return either one implementer type or another. You must return only a single type.

- My own note: In general, you should return the specific type in programming. Generally. But I suspect being able to return a trait is helpful for signatures. Specify accepting a function or closure, though we have not touched upon passing functions as objects yet.

- Can use combined trait bounds (+ sign) syntax to add methods (Associated functions) or just those types that implement both or more traits. This is done using the impl as usual, but with a combining generic T in angle brackets. Syntax, for example:

```rust
impl<T: Display + PartialOrd> { ... }
```

- Remember `SomeTrait for SomeStruct`? Remember using that to implement a trait? Well, adding a trait bound generic syntax to the impl block means we can implement a trait only if the struct implements a different trait. For example:

```rust
impl<T: Display> ToString for T { ... }
```

- Anything with the Display trait also gets the ToString trait.
