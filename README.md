# wingetupd-go
A tiny command line tool using WinGet to update Windows packages (written in Go) 

### Reasons

I recently was curious about the Go programming language (Golang). So i started some re-development of my existing [wingetupd](https://github.com/MBODM/wingetupd) project. This version of `wingetupd.exe` (this time written in Go) shall become exactly the same tool as the original `wingetupd.exe` (that was written in C#).

__Another reason was also:__ The original `wingetupd.exe` (as a .NET 6 self-contained application) is around 10-15 MB in size. The `wingetupd.exe` in Go will become more like 1-3 MB in size.

__But, to be honest, the most prominent reason was:__ I recently was curious about the Rust programming language. So i started a redevelopment of my existing wingetupd project in Rust. After 2 weeks of Rust i came to the conclusion, that Rust isn´t for me. Even when i finally understood Rust´s system of borrowing, ownership, and lifetime, it´s literally a pain in the a** to program with.

#### So, here we "go" again. 😁
