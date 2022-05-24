# wingetupd-go
A tiny command line tool using WinGet to update Windows packages (written in Go) 

### Reasons

I recently was curious about the Go programming language (golang). So i started some re-development of my existing [wingetupd](https://github.com/MBODM/wingetupd) project. This version of `wingetupd.exe` (this time written in Go) shall become exactly the same tool as the original `wingetupd.exe` (which was written in C# and .NET 6).

__Another reason was also:__ The original `wingetupd.exe` (as a .NET 6 self-contained application) is around 10-15 MB in size. The `wingetupd.exe` in Go will become more like 1-3 MB in size.

__But, to be honest, the most prominent reason was:__ I recently was also curious about the Rust programming language. So i already started a re-development in Rust (see [wingetupd-rust](https://github.com/MBODM/wingetupd-rust)). After 2 weeks of Rust i came to the conclusion, that Rust isn´t for me. Even when i finally understood Rust´s system of borrowing, ownership, and lifetime, it´s literally „painful to the point of unusability“, as said in this blog post: http://esr.ibiblio.org/?p=7294 I personally, for myself, totally agree to what this guy says. So i stopped my Rust experiment after 2 weeks and now i will have a look, how Go feels, when rewriting the tool.

#### So, here we "go" again. 😁
