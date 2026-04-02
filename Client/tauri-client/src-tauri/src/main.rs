// Prevents additional console window on Windows in release
#![cfg_attr(all(windows, not(debug_assertions)), windows_subsystem = "windows")]

fn main() {
    owncord_client_lib::run()
}
