`cargo new hello`

creates the new binary package for you, with the Cargo.toml and src dir with a main.rs

`cargo run`

compiles and runs the program

notice it creates ./target/debug ... have VCS ignore all this

to compile without debug symbols

`cargo run --release`
