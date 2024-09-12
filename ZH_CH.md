## RAltGo

##### 适用于 AltV 服务端的 Go 模块 (基于 [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)) 进行开发)

---

##### **该模块已经尽可能的减少 CGO 造成的性能开销，使得您能更接近原生的 Go 性能**

###### 感谢 altv-rust 模块的贡献者，让我可以利用 Rust 来制作这个 Go 模块

###### 

**它依旧还在进行开发中，暂时不建议用于任何生产环境。**

**但是您可以尝试使用它，如果遇到问题可以提出相关 issue**

*我们已对部分 ALTV API 性能进行测试，结果显示目前该Go模块在 ALTV API 调用方面的速度并不比 C# 和 JS 资源差，甚至比它们俩都快。以下我们仅展示分别用 C# 与 Go模块创建5000辆车，如果您有异议您依然可以尝试比对其他资源或者其他内容：*

C#:

![](C:\Users\29970\Desktop\6C980DC59B12E0995A4D04399B563BF4.png)

Go:

![](C:\Users\29970\Desktop\63291F350C0E5E6BDA5ECA16CED9229C.png)

<font color="red">注：因基于 altv-rust 进行开发，所以 altv-rust 出现的问题该模块也都可能会出现！</font>

[我需要一个案例？]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[English](README.md)
