## RAltGo

##### Server-side alt:V API for Go (Developed based on [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)))

---

#### **This module has reduced the performance overhead caused by CGO, allowing you to get closer to docking Go performance**

###### Thanks to the contributors of the altv-rust module for allowing me to make this Go module in Rust

###### 

**It is still under development and is not recommended for use in any production environment at the moment.**

<strong>But you can try to use it and raise a related issue if you encounter problems</strong>

We have tested some ALTV API performance and the results show that the current Go module is no worse than C# and JS resources in ALTV API calls, and is even faster than both of them. Below we only show 5000 created using C# and Go modules respectively. If you have any objections, you can still try other resources or other content:

Csharp:

![](C:\Users\29970\Desktop\6C980DC59B12E0995A4D04399B563BF4.png)

Golang:

![](C:\Users\29970\Desktop\63291F350C0E5E6BDA5ECA16CED9229C.png)

<font color="red">Note: Because it is developed based on altv-rust, the problems that occur in altv-rust may also occur in this module!</font>

[I need an example?]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[中文文档](ZH_CH.md)
