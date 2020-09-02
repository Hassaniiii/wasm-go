extern {
  fn go_api(x: i32, y: i32) -> i32;
}

#[no_mangle]
pub extern fn api(x: i32, y: i32) -> i32 {
    unsafe { go_api(x, y) }
}

