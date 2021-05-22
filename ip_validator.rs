use std::net::{Ipv4Addr, Ipv6Addr};
use std::ffi::CStr;
use std::mem;
use std::os::raw::{c_char, c_void};

/* allocate/deallocate taken from:
https://github.com/wasmerio/wasmer-go/blob/75326bf847202945964e09905ecf64bd9d1baeee/wasmer/test/testdata/examples/greet.rs
*/

#[no_mangle]
pub extern fn allocate(size: usize) -> *mut c_void {
    let mut buffer = Vec::with_capacity(size);
    let pointer = buffer.as_mut_ptr();
    mem::forget(buffer);

    pointer as *mut c_void
}

#[no_mangle]
pub extern fn deallocate(pointer: *mut c_void, capacity: usize) {
    unsafe {
        let _ = Vec::from_raw_parts(pointer, 0, capacity);
    }
}

#[no_mangle]
pub extern "C" fn validate_ip(ip: *const c_char) -> i32 {
    let c_ip_str = unsafe { CStr::from_ptr(ip)};
    let ip_str: &str = c_ip_str.to_str().unwrap();
    return match ip_str.parse::<Ipv6Addr>() {
        Ok(_) => 1,
        Err(_) => match ip_str.parse::<Ipv4Addr>() {
            Ok(_) => 1,
            Err(_) => 0
        }
    };
}