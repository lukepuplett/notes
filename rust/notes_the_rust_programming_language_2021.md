## Chapter 1 - Getting Started
  
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
  
## Chapter 2 - Programming a Guessing Game

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
  
Page 23

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
  
## Chapter 3 - Common Programming Concepts

- `const` can be declared in any scope, must be annotated with its type, and use shouting caps
  
- Using the `let` keyword can shadow an existing variable. Shadows are locally scoped
  
- Four scalar types: integers, floating-points, booleans, characters
  
- Integer sizing: 8, 16, 32, 64, and 128 bits. Can use underscore in literals like in JavaScript, e.g., `1_000`
  
- **IMPORTANT** - Overflows will panic in debug mode, wrap in release mode. This looks dangerous to me and I'd prefer overflows always crash the app.
  
- Can use `wrapping_*`, `checked_*`, `overflowing_*`, and `saturating_*` type methods to get specific behaviors

- `checked_*` and `overflowing_*` methods are safest as they produce either an `Option` or a `(8, false)` tuple where the boolean `false` indicates no overflow occurred.
  
- Integer literals can be:
  - `98_222`
  - `0xff` for hex
  - `0o77` for octal
  - `0b1111_0000` for binary
  - `b'A'` for a byte using ASCII
    
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
  
- `let y = { let x = 3; x + 1 }` works as long as there's no final semicolon `x + 1`, which would cause it to become a statement and not an expression that returns a value. In this case the code in the `{}` runs and is 'expressed' which causes it to be assigned to `y`.
  
- `fn() -> i32 { }` uses the arrow to denote a return type
  
- `if number < 5 { ... } else { ... }` No brackets around conditions in Rust, and it also uses the `else if` syntax
  
- `let number = if something { 5 } else { 6 };`
  
- `let result = loop { ... break counter * 2; };` The break from the loop can return a value, and I think the semi-colon is correct even though you'd normally need to 'express' the value - odd.
  
- `'some_name: loop { ... break 'some_name; }` In this syntax, the loop has a name or label called `'some_name` so the correct loop can be exited.
  
- `while some_condition { ... }` This is like C#
  
- `for element in ... { ... }` This syntax is for-each style
  
- `for number in (1..4).rev() { ... }` This syntax demonstrates a for-each over a group of numbers where the numbers have been reversed

## Chapter 4 - Understanding Ownership

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
  
- **Learned the hard way** - structs own their field values such that sometimes you'll use the field with code that either you wrote yourself or with functions in the built-in libs that wants to move or consume the field value, and you cannot!! An example is with the `JoinHandle<T>` type which has a `.join(self)` method which takes ownership of the handle. So your own struct method can't performing actions on its own data! e.g. `self.handle.join();` won't compile because the struct (i.e. self) owns it. The workaround is to wrap the JoinHandle in an `Option`. This let's you "swap" it out for `None`, e.g. `let h = self.handle.take();`

```rust
impl Drop for Worker {
    fn drop(&mut self) {
        let thread = self.thread.take();

        match thread {
            Some(handle) => match handle.join() {
                Ok(_) => (),
                Err(_) => println!("error in thread {}", self.id),
            },
            _ => return,
        }
    }
}
```

It might be that this is a pattern for structs which have member fields which could be "null", so to speak. It makes total sense that such fields would be optional and have either Some or None. If you need to transfer ownership and literally consume or "burn up" the value, then its field needs to have None. Amazing really.
  
## Chapter 5 - Using Structs to Structure Related Data

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

## Chapter 6 - Enums and Pattern Matching

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
## Chapter 7 - Managing Growing Projects with Packages, Crates, and Modules

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

- The following is a sample structure:
  ```
  backyard
  |- Cargo.lock
  |- Cargo.toml
  |- src
     |- garden
     |  |- vegetables.rs
     |- garden.rs
     |- main.rs
  ```
- I think `main.rs` would declare `mod garden` and its code would be in `/src/garden.rs`. Then, the garden module would declare `mod vegetables` and its code would be in `/src/garden/vegetables.rs` and so on.

- I think this is a little strange and perhaps why the "old" `mod.rs` way works, too, but presumably that's problematic with source control, seeing that a mod file was updated but not knowing what module it is.

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
  
- Shorten `use` syntax with nested imports:
  ```rust
  use std::{cmp::Ordering, io};
  ```
  Can also place `self` in the curly braces
- Append asterisk to bring everything into scope (globbing):
  ```rust
  use std::collections::*;
  ```

## Chapter 8 - Common Collections

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

## Chapter 9 - Error Handling

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

## Chapter 10 - Generic Types, Traits, and Lifetimes

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

- Every reference in Rust has a lifetime - the scope for which that reference is valid.

- This function will not compile because the compiler cannot determine to which reference the return refers (X or Y), because it is dependent on information at runtime, i.e. the length:

```rust
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() { x } else { y }
}
```

- We can annotate the method using generic syntax to constrain the function so that the return ref lifetime is the same as the two parameters. For example:

```rust
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    // ...
}
```

- We're not changing the actual lifetimes. Instead, we're constraining the function to only accept references that adhere to this relationship.

- We can annotate a struct so it can hold a reference. This constrains the struct lifetime to be at least the same time as its field. I.e., the struct cannot outlive its field(s). Example:

```rust
struct ImportantExcerpt<'a> {
    text: &'a str
}
```

- The Rust compiler has lifetime elision rules where annotations are not needed. The book gives an example, which I think is obvious because the reference passed in forms the basis for the reference returned. It is not always conditional on runtime information.

- For methods or associated functions, lifetime annotation must go on the impl and after the type name where they refer to the lifetimes on fields. But where they relate to the lifetimes on parameters, then they go in the signature itself. For example:

```rust
impl<'a> ImportantExcerpt<'a> {
    // ...
}
```

- Note that due to Rust's elision rules, annotations are only needed on method signatures themselves.

- The special 'static lifetime denotes the reference lives as long as the program itself. For example:

```rust
let s: &'static str = "I have a static life";
```

- Use commas to separate the constraints inside angle brackets (generic syntax). For example:

```rust
fn example<'a, T>(x: &'a str) -> &'a str
where
    T: Display
{
    // ...
}
```

## Chapter 11 - Writing Automated Tests

- Using the `cargo test` command, Rust looks for functions with the attribute `#[test]` and compiles a binary test runner.

- Whenever we make a new library in cargo, it adds a new test module and new test function for us. E.g., `cargo new adder --lib` command.

- The default generated code contains a mod with the attribute `#[cfg(test)]`.

- Within the module's curly brackets is a test function with the test attribute, which makes use of `assert_eq!()`. For example:
  ```rust
  assert_eq!(result, 4);
  ```

- To run the tests, use the `cargo test` command line.

- We can ignore and filter as per many test frameworks can, using the `ignore` attribute.

- Performance benchmarking tests can be written too, but you should Google it.

- You can also write documentation tests.

- Page 220 shows usage of `use super::*;` to import all the functions defined in the outer parent module.

- When a test fails, arguments to the comparison `assert!()` macro will use the `PartialEq` and `Debug` traits.

- Both these traits are derivable, and it's often as simple as sticking them on your structs.

- To check that code correctly panics, use the `should_panic` attribute.

- Rust automatically checks/asserts test functions that return `Result<T, E>`. For example:
  ```rust
  fn some_test() -> Result<(), String>
  ```

- But to check for a particular Error variant: `assert!(value.is_err());`

- To run tests consecutively, use `cargo test -- --test-threads=1`.

- Standard output is not shown when a test passes, only on failure. Must use `cargo test -- --show-output` to see all standard out.

- Use `cargo test fn_name_here` to run a single test. But note that this uses a contains match, so it will also match other tests by name.

- The attribute `#[cfg(test)]` is actually a compiler instruction to include the block because in Rust, the unit tests are included in the same code files as the program code (though integration tests are not).

- Rust's privacy rules let you test private code.

- Integration tests go in a `tests` folder, sibling to the `src` folder.

- Each file in the `tests` directory is a separate crate, so need to use `use` keyword paths to bring them into scope.

- Don't need `#[cfg(test)]`. Cargo runs unit tests, then integration tests, then doc tests, and will halt proceedings on a failure.

- Can run integration tests by name and by file name.

- To add common utility code (e.g., for setting up test data), you must place it in a `mod.rs` file in a sub module so that it is not mistaken for an integration test crate. For example: `/tests/common/mod.rs`.

- Page 241 continued: To use the common module code, we must define the submodule in the test files so the compiler knows to look for it and build it. Then refer to its members using paths as per usual: `common::setup`.

## Chapter 12 - An I/O Project - Building a Command Line Program.

## Chapter 13 - Iterators and Closures.

- Code listing starting on page 274 has some interesting uses of the language:

```rust
fn giveaway(&self, user_preference: Option<ShirtColor>) -> ShirtColor {
    user_preference.unwrap_or_else(|| self.most_stocked())
}
```

This is fascinating for taking in an Option. The double pipe `||` is a closure with no parameters/arguments that's evaluated if the option has None. It's also interesting in that it can use the self value (capture it).

- Page 277 has some cool examples of closure syntax. The compiler can infer information from usage. A function is:

```rust
fn add_1(x: u32) -> u32 {
    x + 1
}
```

And the closure would be:

```rust
let add_1 = |x| x + 1;
```

We can call it like this: `let res = add_1(15);`

- Here's the really interesting part: The first use of a closure without any annotations will fix the types inferred.

```rust
let no_annotations = |x| x;
```

If we call this e.g. by passing a string, which it comes back in this particular example, then this closure will be typed as taking a string and returning a string. So a second call that passes an integer will break compilation.

- Just like functions, closures can capture values either by borrowing mutably or borrowing immutably or taking ownership. (Page 278)

- Rust allows multiple concurrent immutable references, so this one is easy for the compiler to emit. But a mutable borrow (e.g. where the closure adds something to a borrowed vec) prevents other borrows (usages) of the vec between the definition of the closure and its usage/call.

- To force the closure to take ownership, you must use the move keyword. This is useful when passing the closure to e.g. a new thread:

```rust
thread::spawn(move || {
    println!("from thread: {:?}", list)
}).join().unwrap();
```

- Page 280 explains, rather curiously, that the main thread might finish before when we are creating and to drop the value.

- A closure body can move a captured value out of the closure, mutate it, neither move nor mutate it, or capture nothing to begin with.

- The way a closure captures and handles values affects the traits it implements, and traits are how functions and structs specify which kinds of closures they can use.

Thank you for providing the additional transcribed notes. I'll process them, correct any misspoken parts, recognize Rust code and terminology, apply syntax highlighting, and ensure the code is legal and correctly formatted. Here's the processed version:

- Closures will automatically implement one, two, or all three of the following `Fn` traits, which I think are additive:

- `FnOnce` applies to closures that can be called once. A closure that moves values out will only ever implement this trait. Because if it can be only called once, all others will always implement `FnOnce`, else you would never be able to call them.

- `FnMut` applies to closures that don't move, but may mutate and can be called many times.

- `Fn` applies to those which don't move or mutate, and those that don't capture anything. These can be called many times and concurrently. Presumably, these are pure functions.

```rust
pub fn unwrap_or_else<F>(self, f: F) -> T
where
    F: FnOnce() -> T
{
    // ...
}
```

- As in C#, we can pass the name of a function that satisfies the trait and signature constraint instead of writing a lambda/closure. For example:

```rust
unwrap_or_else(Vec::new)
```

Where `Vec::new` is called when the value is None.

Interestingly:

```rust
list.sort_by_key(|r| r.width);
```

Requires an `FnMut`. It doesn't mutate anything but does call the closure many times.

- In Rust, like .NET, iterators are lazy execution.

And like .NET, they're implemented via `Iterator`, a trait that has a type `Item` property and function:

```rust
fn next(&mut self) -> Option<Self::Item>
```

Some new syntax was introduced above: `type Item` and `Self::Item`, which is super interesting as it means implementing this trait means also defining an associated item type (see Chapter 19).

So the trait says: Define your item, then implement the `next` function.

Most interesting is that we need to make the iterator instance mutable because calling `next` changes internal state, but when we use `for`, the loop takes mutable ownership behind the scenes.

```rust
let mut my_iter = some_vec.iter();
```

The values we get from `next` are immutable references. Use `into_iter` to create an iterator that takes ownership and returns owned values, and use `iter_mut` to iterate over mutable references.

The standard library has a bunch of default implementations that call `next` and consume or use up the iterator. These are like LINQ in .NET. Some iterator adapters produce different iterators. For example, `map`, which returns a new iterator with the modified items.

But remember, you must call `collect` to actually execute the iterator.

Here's a function which takes an iterator with an associated item type of `String`:

```rust
fn something(mut args: impl Iterator<Item = String>) {
    // ...
}
```
Thank you for providing the transcript. I'll process it, preserving the bullet points, correcting any misspoken parts, recognizing Rust code and terminology, applying syntax highlighting, and ensuring the code is legal and correctly formatted. Here's the processed version:

## Chapter 14 - More about Cargo and Crates.io

- Cargo.toml file can have sections and config per release profile. For example:

```toml
[profile.dev]
opt-level = 0
```

- Like NuGet, you can publish to crates.io. Invest time in markdown comments. Page 297 shows triple slash doc comments with examples in markdown.

- Running `cargo doc --open` will build and open your browser to your doc page.

- It's common to have examples, panics, errors, and safety sections.

- If you add examples using your function/method and use `assert!`, then `cargo test` will also execute doc comment examples!

- Inside the crate root, use `//!` to apply comments to the crate itself.

- The `//!` comments get applied to any logical code container like crate or module.

- Page 302 shows using `pub use` to re-export stuff from lib.rs, which can be used by closures or they can use the original paths.

- You can also re-export stuff used in your dependencies.

- `cargo login <your API key string here>` to store it in `~/.cargo/credentials`.

- Before publishing, you need a [package] section and `name = "some_string"` key-value pair, as well as a description and a license type, name, version, edition of rust.

- You can then define a Cargo.toml file in a root folder and add a [workspace] section. Then add a JSON-like list of crates called members containing paths to those crates:

```toml
[workspace]
members = ["some_path"]
```

Where "some_path" in this case would be "/some_path/Cargo.toml".

- Important: Cargo doesn't assume crates in a workspace depend on each other. So the dependencies need to be defined as usual via the [dependencies] section. For example:

```toml
[dependencies]
add_one = { path = "../add_one" }
```

- Use `cargo run -p <name>` to specifically run a certain crate by name. "adder" in this example.

- Workspaces have a single root Cargo.lock file to ensure the whole tree uses the same versions.

- Use `cargo install` to install binary crates for tools in the Rust bin folder, likely $HOME/.cargo/bin, which implies it must be in your PATH.

## Chapter 15 - Smart Pointers

- Smart pointers behave like normal pointers/refs but have added data and capabilities. For instance, the reference counting smart pointer.

- In many cases, smart pointers actually own the data they point to.

- Technically, the String and Vec<T> types are smart pointers.

- Smart pointers are usually implemented using structs, but implement the Deref and Drop traits.

- Deref trait lets it behave like a normal reference so you can code against it as either a smart pointer or reference.

- Drop affords custom code to run when it goes out of scope and sounds a bit like a destructor to me.

- You can write your own smart pointers, and many libraries ship their own smart pointers.

- Common ones in the standard Library are `Box<T>`, `Rc<T>`, `Ref<T>`, `RefMut<T>` and `RefCell<T>`. `Box<T>` allocates values on the Heap. `Rc<T>` is a reference counter.`Ref<T>` and `RefMut<T>` are accessed through `RefCell<T>`, which enforces borrowing rules at runtime, not compile time.

- Use a `Box<T>` when a type of a known size at compile time is needed, but you have a value of an unknowable size. Or when you have a large amount of data that you want to transfer ownership and be sure it won't be copied. When you want to own a value and you care only that it implements a particular trait rather than being of a particular type.

- `Box<T>` is especially useful for recursive types, where the type contains itself, so Rust cannot compute its total size of the instance. Page 318 shows using an `enum` in Rust to implement a cons list, which is a Lisp way of doing a linked list.

```rust
enum List {
    Cons(i32, List),
    Nil,
}
```

vs.

```rust
enum List {
    Cons(i32, Box<List>),
    Nil,
}
```

In this version, `List` appears within the nested `Cons`, which is self-referential, and therefore, computing the size would be an infinite loop.

- `Box<T>` implements the `Drop` trait, which runs code to clean up its Heap memory. Implementing `Deref` lets you customize the `*` dereference operator, and by coding it in a certain way, you can use the smart pointer like a regular reference. Functions expecting straight up stack values.

Functions expecting straight up stack values, not references, can have a value dereferenced into them.

```rust
let y = 5;
assert_eq!(y, 5);
```

We can also use:

```rust
let y = Box::new(42);
let x = *y;
```

Page 323 is interesting, because it defines a generic struct with no fields.

To implement the `Deref` trait:

```rust
impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}
```

The expression `&self.0` simply accesses the first tuple struct value. Interestingly, this code still returns a reference, i.e., `&self`. Behind the scenes, Rust calls `*(y.deref())`.

Note that if the value was returned raw, i.e., not a reference, then it would be moved out of `self`.

Auto-deref coercion kicks in when calling a method where the parameters don't match what we're passing in and reduces the amount of `&` and `*` syntax to type in.

Implementing `Drop` is just:

```rust
fn drop(&mut self) {
    // ...
}
```

And very much like a C# `Dispose` method. You can call `std::mem::drop` manually to force it. The only use case is a lock/semaphore that you might want to exactly control the drop/unlock of.

- `Rc<T>` is for reference counting lifetime. Like in Visual Basic, this is needed in graph object models where many nodes point at each other and only the last node pointing at any node is conceptually owned by the nodes pointing at it.

- Chapter 16 deals with reference counting in multi-threaded situations. 

- If you don't use `Rc<T>`, then you have to use regular references and that will entail using lifetime syntax to constrain the calling code to ensure all nodes' lives depend on each other.

- Use `Rc::clone(&some_other_rc_instance)` to make a new one where this may actually just increment the reference counter internally.  

- Using `Rc<T>` and Rust's immutable references, you can share read-only data.  

- `RefCell<T>` along with `Rc<T>` affords multiple mutable references via the interior mutability pattern.  

- It seems that the `RefCell<T>` encapsulates some unsafe code to allow some bending of the rules though the pattern/type is safe to use by its consumer at runtime, I think.

- Because `RefCell<T>` allows mutable borrows at runtime, you can mutate the value inside even when the `RefCell<T>` is immutable.  

```rust
let x = 5;
let y = &mut x; // illegal because `x` is immutable.
```

- We need a way to build a value where its methods are permitted to mutate its value, but outsiders see it as immutable.  

- Page 339 sets up a complex situation to show the need for `RefCell<T>`. But the situation seems common to me. It defines a trait with a method that updates a value that the trait implementer holds via its `&self` reference. But it cannot because `&self` is immutable and `&mut self` breaks the signature of the trait method.  

Instead, it puts the value inside a `RefCell<T>` and calls `borrow_mut()` to get a mutable reference, which it can update.  

- The `borrow` method returns a `Ref<T>` and `borrow_mut` returns a `RefMut<T>`.  

- Internally, `RefCell<T>` actually counts the mutable and immutable references being held and sticks to the same rules as the Rust compiler, but at runtime. So two calls to `borrow_mut` in the same scope will cause a runtime panic.  

- By putting a `RefCell<T>` inside a `Rc<T>`, you can get a multiple owner writable reference.

- On page 342, there's some crazy long-winded syntax for working with a graph object where the nodes are mutable. For example:

```rust
let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil))); 
```

- As in Visual Basic, it's possible to leak memory using reference counting by creating a self-referential island of objects. See page 343 to 345.  

- Can defend against this by designing types that don't own their values.  

- Rust has a built-in weak reference `Weak<T>`. You can create `Weak<T>` by calling `Rc::downgrade()`, passing in an instance of an `Rc<T>`.

- To use the weak reference, you must ensure it still holds a value by calling its upgrade method which returns an Option<Rc<T>>:

```rust
struct Node {
  parent: RefCell<Weak<Node>>,
  children: RefCell<Vec<Rc<Node>>>,
}

let leaf = Rc::new(Node {
  parent: RefCell::new(Weak::new()),
  children: RefCell::new(vec![]),  
});

*leaf.parent.borrow_mut() = Rc::downgrade(&leaf);
```

The reason this is not what it seems - assigning to a method - is because the `*` operator dereferences the actual value, so we can assign a new value over the top.

## Chapter 16 - Fearless Concurrency

- Rust appears to have threads, message passing, shared and sync and send traits

- We use `thread::spawn(|| {...})` and when the main program thread exits, all spawned threads will be shut down

- Spawn returns a `JoinHandle<T>` and you can call `.join()` to block the current thread

- Remember to use the `move` keyword in the closure:

  ```rust
  thread::spawn(move || {
    // ...
  });
  ```

- Rust's channels are like those in Go, with a transmitter and a receiver, for example:

  ```rust 
  let (tx, rx) = std::sync::mpsc::channel();
  ```

- The channel looks like a blocking queue in C# with `.send()` and `.recv()` methods and `try_recv()` which is useful for when the receiving loop has other work it can do. These return `Result<T>` or `Result<Ok<T>, E>`

- Once the value is sent the channel takes ownership or it's copied 

- Can treat the receiver as an infinite iterator, and clone the transmitter

- A Rust mutex is quite safe to use due to all the other Rust rules and paradigms, for example:

  ```rust
  let mut num = mtx.lock().unwrap();
  *num = 10;
  ```

- The call to `.lock()` blocks until the mutex lock lets you set the value and the mutex lock will be dropped when it goes out of scope i.e. when the curly brackets end

- Interestingly, the call to `.lock()` will fail if its current holder panics

- So we need to clone the mutex to send it into a new thread closure. So we declare the mutex inside an `Rc<T>` which allows many references via ref counting, but `Rc<T>` itself isn't thread safe and its counting. So you have to use `Arc<T>` which is atomic, for example:

  ```rust
  let mtx = Arc::new(Mutex::new(0));
  ```

- The `Arc<T>` struct implements a trait called `Send`, which allows it to be sent across thread boundaries


- There are other interesting types in `std::sync::Atomic`. 

- Almost every type has the `Send` marker.

- In `[std::marker]`, the `Send` marker trait indicates that the type can be referenced from multiple threads.

- Any type `T` is `Send` if and only if an immutable reference to `T` is `Send`. Similar to `Send`, primitive types, `Sync`, and types composed of them are also `Sync`.

- Because of the composition rules above, there is no reason to manually implement or mark your own types with either of these traits, unless you're using unsafe Rust code features.

## Chapter 17 - Object-Oriented Programming Features

- Rust is somewhat object-oriented. It has methods and public/private.

- Rust has no inheritance, but it does have traits with default implementations, and we can override them.

- Pages 379 to 384 show how to use traits for polymorphism, it's obvious.

- Pages 384 to 391 show how to implement the state pattern, where each state is represented by a struct type that implements a state trait, such that each state object contains its own logic for what to do to move between states. It's confusing at first, but quite interesting to contrast with using enums.

- The point of the exercise is to show that Rust has better/idiomatic ways to implement this kind of state-changing program. It then describes using types to represent the blog articles in all their various states, which is possibly how I may have coded it in C#, anyway. And it doesn't, to my mind, show anything specific to Rust that isn't possible in any statically typed language.

## Chapter 18: Patterns and Matching

- Skipping over `match` keyword and its arms. The `if let` syntax is also covered again.

- We can combine `else` and `else if let`. And the patterns do not have to relate to each other.

- You can also use `while let`, and the value following `for` in a `for` loop is actually a pattern, often destructuring.

- Function parameters can also be patterns. Example:

```rust
fn example(&(x, y): &(i32, i32), ...) { ... }
```

- Patterns are either refutable or irrefutable (i.e., they cannot fail to match).

- When a pattern is refutable, we often just need an `if let` to branch execution for the non-match situation.

- You can match literal values, too. For example:

```rust
match {
    1 => println!("One"),
    ...
}
```

- When matching into named variables, we've seen examples like `Some(v) => ...`, but we can also match a literal here, too. For example `Some(12307)` => ...`

- We can use `|` OR operator. When matching as well as ranges. `1 = 5` goes to `...`.

- Destructuring a struct's fields gets crazy when it's destructured into another struct. For example:
```rust
let Point { x, y } = some_other;
```

- This has more readable shorthand:
```rust
let Point { y, .. } = some_other;
```
And this means we can destructure literal values in the fields, for example:

```rust
match Point { x, y: 63 } {
    Point { x, y } => println!("Blah, blah, blah. x = {x}"),
}
```

- Page 409 shows destructuring enums, including pulling out values embedded in variants, for example:
```rust
match some_enum {
    SomeEnum::SomeVariant { x, y } => { ... }
}
```
You can even match and destructure nested enums within enums, see page 410. And at the top of page 411, it shows complex destructuring of structs and tuples.

- The `_` and `..` can be combined to match and ignore. The `_` is like a C# discard and can/should be used even on function signatures where your implementation does not use a parameter. In Rust pattern matching, it seems to mean "don't care" or "whatever".

- Using a `_` before a variable indicates that it's okay/valid that the variable is not used, but it still binds a value to it, as opposed to a true discard.

- The `..` syntax (two periods) is the "rest of the fields" or "all the others" and works like a discard `_`. For example, `[first, _, last, ..]` goes to `{ ... }`.

- Match guards are additional conditions on the arms of a `match` which must also be true. For example:
```rust
match num {
    Some(x) if x % 2 == 0 => { ... },
    _ => { ... },
}
```

Unfortunately, Rust will suspend checking for exhaustiveness of the arms. You can combine with the `|` or `||` operator, for example: `4 | 5 | 6`.

- You can even use the `@` symbol to capture a destructured value and also test it. See page 417.

## Chapter 19 - Advanced Features

- To use unsafe code, stick it in an `unsafe` block. Then you can dereference an unsafe pointer, call an unsafe function or method, access or modify a mutable static variable, implement an unsafe trait, or access fields of unions.
  
- The borrow checker still works/applies in `unsafe` code blocks.
  
- The `unsafe` keyword only gives you access to the unsafe features listed above. You'll still get some degree of safety by limiting access to within these blocks, effectively guaranteeing that bugs are inside these `unsafe` blocks, making debugging simpler.
  
- Raw pointers can be mutable or immutable and are written as `*const T` and `*mut T` respectively. The `*` is not a dereference but part of the type name itself.
  
- Raw pointers can ignore the boring rules and have both mutable and immutable pointers or many mutable pointers to the same location.

- Rule pointers are not guaranteed to point to valid memory. 

- Raw pointers are allowed to be null, and they do not implement auto cleanup.

Creating raw pointers from references:
```rust
let mut num = 5;
let r1 = &num as *const i32;
let r2 = &mut num as *mut i32;
```

In the code above, there's no `unsafe` block yet. We can create the raw pointers, which is safe to do, but we cannot dereference them outside of an `unsafe` block.

We can create a raw pointer to an arbitrary location in memory:
```rust
let address = 0x012345usize;
let r = address as *const i32;
```

Here's how we use the raw pointer:
```rust
unsafe {
    println!("r1 is: {}", *r1);
}
```

- You can define unsafe functions by prepending `unsafe` to the function definition, like `unsafe fn danger() { ... }`.

- You can only call the unsafe function from within an `unsafe` block.

- Page 425 shows using unsafe code from inside a safe normal function to abstract it. It has an `unsafe` block, which calls the `to_and_say` functions.

- Rust has a keyword `extern` for creating foreign function interface (FFI).

- The `extern` keyword creates a block within which we can define the signatures of the external functions we want to call from our Rust code.
  ```rust
  extern "C" { ... }
  ```
  This defines a block using the C language ABI (application binary interface).

- We can also use `extern` to create an interface to allow other languages to call Rust. See page 427.
  ```rust
  #[no_mangle]
  pub extern "C" fn cool() { ... }
  ```

- Calling out to external code is inherently unsafe, but incoming calls are safe.

In Rust, global variables are called static variables and they're named in `SCREAMING_SNAKE_CASE`.

- Accessing an immutable static variable is safe. They have the `'static` lifetime.

- Accessing and modifying (reading and writing) mutable statics is safe, but it can introduce data races in multi-threaded programs, so you should use types from the `sync` crate like `Mutex`.

- An `unsafe` trait is any trait where at least one method has some invariant that the compiler cannot verify.
  ```rust
  unsafe trait MySafeTrait { ... }
  ```

- If we need to manually mark a type as `Send` or `Sync`, then Rust can't verify that it upholds those guarantees, so we must use `unsafe` to do those checks.

- A union is like a struct, but only one declared field is used in an instance at a time. They're used to interface with unions in C code.

Page 430. 
- **Advanced traits. Associated types.** This is sort of like generics for traits, except the syntax is different. And for some reason, the book doesn't make this analogy. Actually, it does a bit later.

```rust
trait Iterator {
    type Item;
    fn next(&mut self) -> Option<Self::Item>;
}
```

- The implementer must specify the type.
- `type Item = u32;`
- But the method implementation still uses `Option<Self::Item>`.

- Rust allows operator overloading for specific types. And the associated type and the default generic type features come into play.

- The example on page 433 shows `struct Millimeters(i32)`, which is known as the "newtype pattern".
- All its operators get overloaded.
- The book notes that adding a generic to a trait after it has shipped is possible, because you can set its default. So everything still compiles.

- A type can implement two traits which specify the same method name, and you can indicate which one to call by fully qualifying its path, e.g. `Pilot::method(&arg)`.

- On page 435, it shows that traits can specify non-method associated functions, or what you call "static methods" in C#.
- To disambiguate these associated functions: `<Dog as Animal>::_name()`.

- We can define a trait so its implementer must also implement another trait. See page 437:
```rust
trait MyTrait: OtherTrait { ... }
```

- The "newtype pattern" lets us get around the inability to extend types. The thin wrapper struct is removed during compilation, so there's no runtime cost.

- The book gives a nice example of a `People` newtype that wraps a `HashMap<i32, String>`.

- The book mentions "type aliases", which are used to avoid having to keep track of lengthy and complex type names, e.g. `Box<dyn Fn(i32) -> i32 + Send + 'static>`.

- Rust has a `!` (never) type, which is implied by code that shifts execution to another, like `continue` or `panic!()`.

- Page 445 discusses the special `?` sized syntax for dynamically sized traits, but you didn't get it.

- Passing functions to functions via function pointers:
  - `Fn` is a type, not a trait, unlike closures.
  - `fn example(func: fn(i32) -> i32) { ... }`
  - It recommends using the closure syntax instead, which is useful for interacting with languages like C that have function pointers but not closures.
  - You cannot return an `Fn` function pointer, but you can return a closure in a `Box`.

- Macros:
  - Rust has declarative macros and three kinds of procedural macros:
    - Custom derive macros that define the code added to `#[derive]` for structs and enums.
    - Attribute-like macros that define custom attributes usable on any item.
    - Function-like macros that look like normal calls but operate on the tokens specified as their arguments.
  - Macros are a way of writing code that writes other code. They expand to produce more code than you would have written manually.

- Macros can take a variable number of arguments. Macros are expanded before the compiler understands the meaning of the code. Code that writes code is hard to understand and write. Rust is no panacea for this problem. Macros must be brought into scope or defined lexically prior to calling.

- The most widely used are declarative macros. Declarative macros work a bit like pattern matching where the code passed to the matching syntax is the arguments and the return value is the expanded final code the compiler should emit.

- Page 450 shows a slightly simplified example of the `vec!` macro. The macro is defined using `macro_rules!` followed by its name. It can be annotated with the `#[macro_export]` attribute to make it available to the importers of the current crate. The Rust pattern matching syntax for macros can be found at https://doc.rust-lang.org/reference/macros-by-example.html.

- Writing macros is a highly advanced, yet simple in concept feature, which is like using regex and can be painfully solved with functions.

- Procedural macros don't use pattern matching. These are defined as functions, which take a token stream as input and return a token stream. Page 453 mentions that Rust has no reflection, which might be why macros are useful.

- Page 457 explains how attribute-like macros are used to make attributes that can adorn functions, for example, in popular web frameworks to annotate functions as HTTP URL routes.

- Attribute-like macros are implemented very much the same as custom derive type macros. They have two token stream arguments: one for the arguments on the attribute in square brackets, and another which captures the code of the function.

- Finally, function-like macros combine the declarative style macro rules types, which can be called like functions with procedural using token streams.

- Note: Macros being an ass to write might make them perfect for AI/LLMs.

## Chapter 20: Building a multi-threaded web server.

- There's one more interesting thing to learn from the tutorial/lesson; the `if let` and `while let` syntax creates a scope that is as long as their blocks `{}`. In the lesson, a Mutex is used to access the receiving end of a channel, like a blocking queue, and obtaining the Mutex and dequeuing a channel item is done using this code:

```rust
let job = receiver.lock().unwrap().recv().unwrap();
```

The lock is released when the MutexGuard that is returned drops out of scope, which is at the end of the assignment in this code, but in a `while let Ok(job) = receiver.lock().unwrap().recv() { ... }` would be the end of the block! This would hold the lock for the entire duration of running the job, so no other jobs can be run at the same time!

<end>
