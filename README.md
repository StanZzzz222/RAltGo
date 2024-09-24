## 

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/altv.svg" alt="" width="59"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/plus.png" alt="" width="58"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/rust.png" alt="" width="89"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/equal.png" alt="" width="59"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### Server-side alt:V API for Go (Developed based on [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)))

---

#### **This module has reduced the performance overhead caused by CGO, allowing you to get closer to docking Go performance**

###### Thanks to the contributors of the altv-rust module for allowing me to make this Go module in Rust

###### 

**It is still under development and is not recommended for use in any production environment at the moment.**

<strong>But you can try to use it and raise a related issue if you encounter problems</strong>

**<font color="red">Note: Since the author has other things to do, he can only try to spend as much time as possible to update. You can rest assured that I will keep updating in my free time</font>**

We know many people may be concerned about the performance of this module, so we've added benchmarks to the examples resources. The current code and results are below. If you have any questions, you can try the sample server. Because resource scheduling will only be performed onTick, your performance issues will be alerted by the server-side resourceManager.

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/code.png" alt="" data-align="center">

Measured screenshots:

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/server.png" alt="" data-align="left" width="648">

Some APIs that have been implemented:

```
Object:
1. Blip - Mostly done
2. Player - Mostly done
3. Vehicle - Mostly done
4. Ped - Mostly done
5. Colshape - Mostly done
6. Checkpoint - Mostly done
7. Marker- Mostly done
8. Object - Mostly done
9. Pools - The object pool has implemented all the above objects
10. VisualEntity - To be implemented
11. VirtualEntityGroup - To be implemented
12. VoiceChannel - To be implemented
13. ConnectionInfo - To be implemented

Events:
1. onStart - Done
2. onServerStarted - Done
3. onStop - Done
4. onPlayerConnect - Done
5. onPlayerDisconnect - Done
6. onEnterVehicle - Done
7. onLeaveVehicle - Done
8. onEnterColshape - Done
9. onLeaveColshape - Done
10. onClientEvent - Done
12. onCommandEvent - Done
....
```

<font color="red">Note: Because it is developed based on altv-rust, the problems that occur in altv-rust may also occur in this module!</font>

[I need an example?]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[中文文档](ZH_CH.md)
