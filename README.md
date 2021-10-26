# cgscript

The problem with go currenlty is it's startup time for very small scripts. I use go as a scripting/automation language, but the problem is while my script runs only for 9ms, it compiles for a second!

cgscript will try to run your go script from cache, eliminating the startup problem.

## Usage

```
cgscript ./build.go
```

This line runs contents of build.go from cache or compiles into the cache dir them and then runs. `Temp dir`/cgscript is cache dir by default.

## Features

- [ ] Change the cache dir
- [ ] Changing the run command
- [ ] Changing the compile command
