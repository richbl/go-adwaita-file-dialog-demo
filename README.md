## Overview

<p align="center">
<img width="610" height="422" alt="file-dialog-demo-dark" src="https://github.com/user-attachments/assets/de68dabf-85a4-4c10-bcc6-a2af33426169" />
<img width="610" height="422" alt="file-dialog-demo-light" src="https://github.com/user-attachments/assets/1eca0ef7-177e-4874-b199-5dc57d8c553b" />
</p>

**Go Adwaita File Dialog
 Demo** is a simple Go-based demo using GTK4/Adwaita libraries ([GoTK4](https://github.com/diamondburned/gotk4)/[GoTK4-Adwaita](https://github.com/diamondburned/gotk4-adwaita)).

It's presented here as a simple project demo to serve as a functional example of how to implement a file dialog in Go that follows idiomatic best practices of the [Gnome Human Interface Guidelines (HIG)](https://developer.gnome.org/hig/).

For a good example of how this file dialog is used in a more significant project, check out the [**BLE Sync Cycle project**](https://github.com/richbl/go-ble-sync-cycle).

## Installation

A note on installation: since the GTK4/Adwaita packages ([GoTK4](https://github.com/diamondburned/gotk4)/[GoTK4-Adwaita](https://github.com/diamondburned/gotk4-adwaita)) require local compilation (the native libraries are written in C), expect a delay in first-time application execution (~10-15 minutes depending on CPU speed).
